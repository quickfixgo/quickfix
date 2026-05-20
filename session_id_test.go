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

package quickfix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionID_String(t *testing.T) {
	var testCases = []struct {
		sessionID      SessionID
		expectedString string
	}{
		{SessionID{BeginString: "FIX.4.2", SenderCompID: "SND", TargetCompID: "TAR"}, "FIX.4.2:SND->TAR"},
		{SessionID{BeginString: "FIX.4.2", SenderCompID: "SND", TargetCompID: "TAR", Qualifier: "BLAH"}, "FIX.4.2:SND->TAR:BLAH"},
		{SessionID{BeginString: "FIX.4.2", SenderCompID: "SND", SenderSubID: "SSUB", SenderLocationID: "SLOC",
			TargetCompID: "TAR", TargetSubID: "TSUB", TargetLocationID: "TLOC",
			Qualifier: "BLAH"}, "FIX.4.2:SND/SSUB/SLOC->TAR/TSUB/TLOC:BLAH"},
		{SessionID{BeginString: "FIX.4.2", SenderCompID: "SND", SenderLocationID: "SLOC",
			TargetCompID: "TAR", TargetSubID: "TSUB", TargetLocationID: "TLOC",
		}, "FIX.4.2:SND/SLOC->TAR/TSUB/TLOC"},
	}

	for _, tc := range testCases {
		actual := tc.sessionID.String()
		assert.Equal(t, tc.expectedString, actual)
	}
}
