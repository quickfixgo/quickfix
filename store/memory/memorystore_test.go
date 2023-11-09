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

package memory

import (
	"testing"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/internal/testsuite"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// MemoryStoreTestSuite runs all tests in the MessageStoreTestSuite against the MemoryStore implementation.
type MemoryStoreTestSuite struct {
	testsuite.StoreTestSuite
}

func (suite *MemoryStoreTestSuite) SetupTest() {
	var err error
	suite.MsgStore, err = quickfix.NewMemoryStoreFactory().Create(quickfix.SessionID{})
	require.Nil(suite.T(), err)
}

func TestMemoryStoreTestSuite(t *testing.T) {
	suite.Run(t, new(MemoryStoreTestSuite))
}
