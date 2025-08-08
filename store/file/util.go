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
	"strings"

	"github.com/quickfixgo/quickfix"
)

func createFilenamePrefix(s quickfix.SessionID) string {
	sender := []string{s.SenderCompID}
	if s.SenderSubID != "" {
		sender = append(sender, s.SenderSubID)
	}
	if s.SenderLocationID != "" {
		sender = append(sender, s.SenderLocationID)
	}

	target := []string{s.TargetCompID}
	if s.TargetSubID != "" {
		target = append(target, s.TargetSubID)
	}
	if s.TargetLocationID != "" {
		target = append(target, s.TargetLocationID)
	}

	fname := []string{s.BeginString, strings.Join(sender, "_"), strings.Join(target, "_")}
	if s.Qualifier != "" {
		fname = append(fname, s.Qualifier)
	}
	return strings.Join(fname, "-")
}

// closeSyncFile behaves like Sync and Close, except that no error is returned if the file does not exist.
func closeSyncFile(f *os.File) error {
	if f != nil {
		if err := f.Sync(); err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		}
		if err := f.Close(); err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		}
	}
	return nil
}

// removeFile behaves like os.Remove, except that no error is returned if the file does not exist.
func removeFile(fname string) error {
	if err := os.Remove(fname); (err != nil) && !os.IsNotExist(err) {
		return fmt.Errorf("remove %v: %w", fname, err)
	}
	return nil
}

// openOrCreateFile opens a file for reading and writing, creating it if necessary.
func openOrCreateFile(fname string, perm os.FileMode) (f *os.File, err error) {
	if f, err = os.OpenFile(fname, os.O_RDWR, perm); err != nil {
		if f, err = os.OpenFile(fname, os.O_RDWR|os.O_CREATE, perm); err != nil {
			return nil, fmt.Errorf("error opening or creating file: %s: %s", fname, err.Error())
		}
	}
	return f, nil
}
