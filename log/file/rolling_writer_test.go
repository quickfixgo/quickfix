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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestRollingWriter_SizeBasedRotation(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), fmt.Sprintf("TestRolling-%d", os.Getpid()))
	defer os.RemoveAll(tmpDir)

	filename := filepath.Join(tmpDir, "test.log")
	// Use very small max size (1MB) for testing
	writer, err := newRollingWriter(filename, 1, 3, 0, false)
	if err != nil {
		t.Fatal("Failed to create rolling writer", err)
	}
	defer writer.Close()

	// Write data that exceeds 1MB
	data := strings.Repeat("x", 1024*1024+1)
	_, err = writer.Write([]byte(data))
	if err != nil {
		t.Fatal("Failed to write", err)
	}

	// Sync to ensure rotation happens
	writer.Sync()

	// Small delay to ensure file operations complete
	time.Sleep(100 * time.Millisecond)

	// Check if rotation occurred
	matches, err := filepath.Glob(filepath.Join(tmpDir, "test.*.log"))
	if err != nil {
		t.Fatal("Failed to glob", err)
	}

	// Should have at least one rotated file
	if len(matches) == 0 {
		t.Error("Expected rotated file, but found none")
	}
}

func TestRollingWriter_MaxBackups(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), fmt.Sprintf("TestRolling-%d", os.Getpid()))
	defer os.RemoveAll(tmpDir)

	filename := filepath.Join(tmpDir, "test.log")
	// Use very small max size and limit backups to 2
	writer, err := newRollingWriter(filename, 1, 2, 0, false)
	if err != nil {
		t.Fatal("Failed to create rolling writer", err)
	}
	defer writer.Close()

	// Trigger multiple rotations
	for i := 0; i < 5; i++ {
		data := strings.Repeat("x", 1024*1024+1)
		_, err = writer.Write([]byte(data))
		if err != nil {
			t.Fatal("Failed to write", err)
		}
		time.Sleep(10 * time.Millisecond) // Small delay to ensure different timestamps
	}

	// Check rotated files
	matches, err := filepath.Glob(filepath.Join(tmpDir, "test.*.log"))
	if err != nil {
		t.Fatal("Failed to glob", err)
	}

	// Should have at most 2 rotated files (maxBackups=2)
	if len(matches) > 2 {
		t.Errorf("Expected at most 2 rotated files, but found %d", len(matches))
	}
}

func TestRollingWriter_Compression(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), fmt.Sprintf("TestRolling-%d", os.Getpid()))
	defer os.RemoveAll(tmpDir)

	filename := filepath.Join(tmpDir, "test.log")
	// Enable compression
	writer, err := newRollingWriter(filename, 1, 1, 0, true)
	if err != nil {
		t.Fatal("Failed to create rolling writer", err)
	}
	defer writer.Close()

	// Trigger rotation
	data := strings.Repeat("x", 1024*1024+1)
	_, err = writer.Write([]byte(data))
	if err != nil {
		t.Fatal("Failed to write", err)
	}

	// Sync to ensure rotation happens
	writer.Sync()

	// Small delay to ensure file operations complete
	time.Sleep(100 * time.Millisecond)

	// Check if compressed file exists
	matches, err := filepath.Glob(filepath.Join(tmpDir, "test.*.log.gz"))
	if err != nil {
		t.Fatal("Failed to glob", err)
	}

	// Should have at least one compressed file
	if len(matches) == 0 {
		t.Error("Expected compressed file, but found none")
	}
}

func TestRollingWriter_NoRotationWhenDisabled(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), fmt.Sprintf("TestRolling-%d", os.Getpid()))
	defer os.RemoveAll(tmpDir)

	filename := filepath.Join(tmpDir, "test.log")
	// All rolling options disabled
	writer, err := newRollingWriter(filename, 0, 0, 0, false)
	if err != nil {
		t.Fatal("Failed to create rolling writer", err)
	}
	defer writer.Close()

	// Write large amount of data
	data := strings.Repeat("x", 10*1024*1024)
	_, err = writer.Write([]byte(data))
	if err != nil {
		t.Fatal("Failed to write", err)
	}

	// Check that no rotation occurred
	matches, err := filepath.Glob(filepath.Join(tmpDir, "test.*.log"))
	if err != nil {
		t.Fatal("Failed to glob", err)
	}

	// Should have no rotated files
	if len(matches) > 0 {
		t.Error("Expected no rotated files when rotation is disabled")
	}
}
