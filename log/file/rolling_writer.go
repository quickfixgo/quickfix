// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package file

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// rollingWriter wraps a file with rotation support based on size and age.
type rollingWriter struct {
	mu sync.Mutex

	filename    string
	currentFile *os.File
	maxSize     int64 // in bytes, 0 means disabled
	maxBackups  int
	maxAge      int // in days, 0 means disabled
	compress    bool

	currentSize int64
	lastRotate  time.Time
}

// newRollingWriter creates a new rolling writer.
// filename: base filename (e.g., "prefix.event.current.log")
// maxSize: maximum size in megabytes before rotation (0 = disabled)
// maxBackups: maximum number of old log files to keep (0 = keep all)
// maxAge: maximum age in days for old log files (0 = disabled)
// compress: whether to compress rotated files
func newRollingWriter(filename string, maxSize int, maxBackups int, maxAge int, compress bool) (*rollingWriter, error) {
	w := &rollingWriter{
		filename:   filename,
		maxSize:    int64(maxSize) * 1024 * 1024, // convert MB to bytes
		maxBackups: maxBackups,
		maxAge:     maxAge,
		compress:   compress,
		lastRotate: time.Now(),
	}

	if err := w.openCurrentFile(); err != nil {
		return nil, err
	}

	// Get initial file size
	if info, err := w.currentFile.Stat(); err == nil {
		w.currentSize = info.Size()
	}

	return w, nil
}

// openCurrentFile opens the current log file in append mode.
func (w *rollingWriter) openCurrentFile() error {
	dir := filepath.Dir(w.filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	fileFlags := os.O_RDWR | os.O_CREATE | os.O_APPEND
	file, err := os.OpenFile(w.filename, fileFlags, os.ModePerm)
	if err != nil {
		return err
	}

	w.currentFile = file
	return nil
}

// Write writes data to the log file, rotating if necessary.
func (w *rollingWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// Check if rotation is needed before writing
	if err := w.maybeRotate(); err != nil {
		return 0, err
	}

	n, err = w.currentFile.Write(p)
	if err == nil {
		w.currentSize += int64(n)
		// Check again after writing in case we exceeded the limit
		if err = w.maybeRotate(); err != nil {
			return n, err
		}
	}
	return n, err
}

// maybeRotate checks if rotation is needed and performs it.
func (w *rollingWriter) maybeRotate() error {
	now := time.Now()
	needsRotate := false

	// Check size-based rotation
	if w.maxSize > 0 && w.currentSize >= w.maxSize {
		needsRotate = true
	}

	// Check time-based rotation (daily)
	if !needsRotate && w.maxAge > 0 {
		// Rotate if it's a new day
		if now.YearDay() != w.lastRotate.YearDay() || now.Year() != w.lastRotate.Year() {
			needsRotate = true
		}
	}

	if !needsRotate {
		return nil
	}

	return w.rotate()
}

// rotate performs the actual rotation.
func (w *rollingWriter) rotate() error {
	// Close current file
	if w.currentFile != nil {
		if err := w.currentFile.Close(); err != nil {
			return err
		}
	}

	// Generate rotated filename with timestamp
	timestamp := time.Now().Format("20060102-150405")
	ext := filepath.Ext(w.filename)
	base := strings.TrimSuffix(w.filename, ext)
	rotatedFilename := fmt.Sprintf("%s.%s%s", base, timestamp, ext)

	// Rename current file to rotated filename
	if err := os.Rename(w.filename, rotatedFilename); err != nil {
		// If rename fails, try to open current file again
		w.openCurrentFile()
		return err
	}

	// Compress if enabled
	if w.compress {
		if err := w.compressFile(rotatedFilename); err != nil {
			// Log error but continue
			_ = err
		} else {
			// Remove uncompressed file after compression
			if err := os.Remove(rotatedFilename); err != nil {
				// Log error but continue
				_ = err
			}
		}
	}

	// Clean up old files
	if err := w.cleanup(); err != nil {
		// Log error but continue
		_ = err
	}

	// Open new current file
	if err := w.openCurrentFile(); err != nil {
		return err
	}

	// Reset size and update last rotate time
	if info, err := w.currentFile.Stat(); err == nil {
		w.currentSize = info.Size()
	} else {
		w.currentSize = 0
	}
	w.lastRotate = time.Now()

	return nil
}

// compressFile compresses a file using gzip.
func (w *rollingWriter) compressFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	compressedFilename := filename + ".gz"
	compressedFile, err := os.Create(compressedFilename)
	if err != nil {
		return err
	}
	defer compressedFile.Close()

	gzWriter := gzip.NewWriter(compressedFile)
	defer gzWriter.Close()

	_, err = io.Copy(gzWriter, file)
	return err
}

// cleanup removes old log files based on maxBackups and maxAge.
func (w *rollingWriter) cleanup() error {
	dir := filepath.Dir(w.filename)
	base := filepath.Base(w.filename)
	ext := filepath.Ext(base)
	baseName := strings.TrimSuffix(base, ext)

	// Pattern to match rotated files: baseName.YYYYMMDD-HHMMSS.ext or baseName.YYYYMMDD-HHMMSS.ext.gz
	pattern := filepath.Join(dir, baseName+".*"+ext+"*")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	// Filter out current file
	var rotatedFiles []string
	for _, match := range matches {
		if match != w.filename {
			rotatedFiles = append(rotatedFiles, match)
		}
	}

	// Sort by modification time (oldest first)
	sort.Slice(rotatedFiles, func(i, j int) bool {
		infoI, errI := os.Stat(rotatedFiles[i])
		infoJ, errJ := os.Stat(rotatedFiles[j])
		if errI != nil || errJ != nil {
			return false
		}
		return infoI.ModTime().Before(infoJ.ModTime())
	})

	now := time.Now()
	var toDelete []string

	// Apply maxAge filter
	if w.maxAge > 0 {
		for _, file := range rotatedFiles {
			info, err := os.Stat(file)
			if err != nil {
				continue
			}
			age := now.Sub(info.ModTime())
			if age > time.Duration(w.maxAge)*24*time.Hour {
				toDelete = append(toDelete, file)
			}
		}
	}

	// Apply maxBackups filter
	if w.maxBackups > 0 {
		// Remove files that exceed maxBackups (after removing those deleted by maxAge)
		remaining := make([]string, 0)
		for _, file := range rotatedFiles {
			isDeleted := false
			for _, deleted := range toDelete {
				if file == deleted {
					isDeleted = true
					break
				}
			}
			if !isDeleted {
				remaining = append(remaining, file)
			}
		}

		if len(remaining) > w.maxBackups {
			excess := len(remaining) - w.maxBackups
			for i := 0; i < excess; i++ {
				toDelete = append(toDelete, remaining[i])
			}
		}
	}

	// Delete files
	for _, file := range toDelete {
		_ = os.Remove(file)
	}

	return nil
}

// Close closes the current log file.
func (w *rollingWriter) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.currentFile != nil {
		return w.currentFile.Close()
	}
	return nil
}

// Sync syncs the current log file.
func (w *rollingWriter) Sync() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.currentFile != nil {
		return w.currentFile.Sync()
	}
	return nil
}
