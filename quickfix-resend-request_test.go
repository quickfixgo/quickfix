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

import "bytes"

// TestFIXMsgInResendRequestFIXInputValidation checks FIX input constraints:
// required integer fields, BeginSeqNo >= 1, and EndSeqNo >= 0.
// BeginSeqNo <= EndSeqNo follows the ordered finite range definition; EndSeqNo=0
// is open-ended in FIX 4.2. See FIX Session Layer sections 4.1, 4.8.2, 9.3
// and FIX Session Testcases section 4.5.13 for Reject reason categories:
// https://www.fixtrading.org/wp-content/uploads/download-manager-files/FIX_Session_Layer_June_2020.pdf
// https://www.fixtrading.org/wp-content/uploads/download-manager-files/FIX_Session_Testcases_June_2020.pdf
// The dictionary validator is unset so the handler must enforce these constraints.
// Normal finite/open-ended requests and replay are covered by existing session
// and acceptance tests. End clamping and rejecting unsent ranges are engine
// policy; an originally reversed finite range is an invalid input.
func (s *InSessionTestSuite) TestFIXMsgInResendRequestFIXInputValidation() {
	tests := []struct {
		name           string
		beginSeqNo     string // An empty string omits the required field.
		endSeqNo       string
		disablePersist bool
		rejectReason   int
		refTagID       Tag
	}{
		{name: "missing_begin", beginSeqNo: "", endSeqNo: "10", rejectReason: rejectReasonRequiredTagMissing, refTagID: tagBeginSeqNo},
		{name: "missing_end", beginSeqNo: "7", endSeqNo: "", rejectReason: rejectReasonRequiredTagMissing, refTagID: tagEndSeqNo},
		{name: "non_integer_begin", beginSeqNo: "abc", endSeqNo: "10", rejectReason: rejectReasonIncorrectDataFormatForValue, refTagID: tagBeginSeqNo},
		{name: "non_integer_end", beginSeqNo: "7", endSeqNo: "abc", rejectReason: rejectReasonIncorrectDataFormatForValue, refTagID: tagEndSeqNo},
		{name: "zero_begin_open_ended", beginSeqNo: "0", endSeqNo: "0", rejectReason: rejectReasonValueIsIncorrect, refTagID: tagBeginSeqNo},
		{name: "negative_begin", beginSeqNo: "-1", endSeqNo: "10", rejectReason: rejectReasonValueIsIncorrect, refTagID: tagBeginSeqNo},
		{name: "negative_end", beginSeqNo: "7", endSeqNo: "-1", rejectReason: rejectReasonValueIsIncorrect, refTagID: tagEndSeqNo},
		{name: "reversed_range", beginSeqNo: "10", endSeqNo: "7", rejectReason: rejectReasonValueIsIncorrect, refTagID: tagEndSeqNo},
		{name: "reversed_range_with_999999_end", beginSeqNo: "1000000", endSeqNo: "999999", rejectReason: rejectReasonValueIsIncorrect, refTagID: tagEndSeqNo},
		// One representative invalid request must not bypass validation when persistence is disabled.
		{name: "zero_begin_no_persist", beginSeqNo: "0", endSeqNo: "0", disablePersist: true, rejectReason: rejectReasonValueIsIncorrect, refTagID: tagBeginSeqNo},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			s.SetupTest()
			s.session.DisableMessagePersist = test.disablePersist
			s.Require().Nil(s.session.store.SetNextSenderMsgSeqNum(11))
			if !test.disablePersist {
				// A replayable message makes accidental replay observable.
				msg := NewMessage()
				s.session.fillDefaultHeader(msg, nil)
				msg.Header.SetField(tagMsgType, FIXString("D"))
				msg.Header.SetField(tagMsgSeqNum, FIXInt(7))
				s.Require().Nil(s.session.store.SaveMessage(7, msg.build()))
			}

			request := s.ResendRequest(1)
			request.Body.Remove(tagBeginSeqNo)
			request.Body.Remove(tagEndSeqNo)
			if test.beginSeqNo != "" {
				request.Body.SetField(tagBeginSeqNo, FIXString(test.beginSeqNo))
			}
			if test.endSeqNo != "" {
				request.Body.SetField(tagEndSeqNo, FIXString(test.endSeqNo))
			}
			s.MockApp.On("FromAdmin").Return(nil).Maybe()
			s.MockApp.On("ToAdmin").Maybe()
			s.MockApp.On("ToApp").Return(nil).Maybe()

			s.fixMsgIn(s.session, request)

			s.MockApp.AssertNumberOfCalls(s.T(), "ToApp", 0)
			s.State(inSession{})
			s.NextSenderMsgSeqNum(12) // Only the Reject consumes a sender sequence number.
			s.NextTargetMsgSeqNum(2)
			s.Require().Len(s.Receiver.sendChannel, 1, "send only a Reject, without replay or GapFill")
			msgBytes, ok := s.Receiver.LastMessage()
			s.Require().True(ok)
			response := NewMessage()
			s.Require().Nil(ParseMessage(response, bytes.NewBuffer(msgBytes)))
			s.MessageType(string(msgTypeReject), response)
			s.FieldEquals(tagMsgSeqNum, 11, response.Header)
			s.FieldEquals(tagRefSeqNum, 1, response.Body)
			// These optional diagnostic fields use the engine's existing Reject encoding.
			s.FieldEquals(tagSessionRejectReason, test.rejectReason, response.Body)
			s.FieldEquals(tagRefTagID, int(test.refTagID), response.Body)
			s.FieldEquals(tagRefMsgType, string(msgTypeResendRequest), response.Body)
		})
	}
}

// TestFIXMsgInResendRequestFIXInputValidationValidRange checks valid FIX 4.2 ranges.
// Equal bounds request one message, EndSeqNo=0 is open-ended,
// and EndSeqNo=999999 is finite (FIX Session Layer 4.8.2).
func (s *InSessionTestSuite) TestFIXMsgInResendRequestFIXInputValidationValidRange() {
	tests := []struct {
		name             string
		beginSeqNo       int
		endSeqNo         int
		nextSenderSeqNum int
		expectedSeqNums  []int
	}{
		{name: "single_message", beginSeqNo: 7, endSeqNo: 7, nextSenderSeqNum: 9, expectedSeqNums: []int{7}},
		{name: "open_ended", beginSeqNo: 7, endSeqNo: 0, nextSenderSeqNum: 9, expectedSeqNums: []int{7, 8}},
		{name: "last_sent_message", beginSeqNo: 8, endSeqNo: 0, nextSenderSeqNum: 9, expectedSeqNums: []int{8}},
		{name: "999999_is_finite", beginSeqNo: 999999, endSeqNo: 999999, nextSenderSeqNum: 1000001, expectedSeqNums: []int{999999}},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			s.SetupTest()
			s.Require().Nil(s.session.store.SetNextSenderMsgSeqNum(test.nextSenderSeqNum))
			for seqNum := test.beginSeqNo - 1; seqNum < test.nextSenderSeqNum; seqNum++ {
				msg := NewMessage()
				s.session.fillDefaultHeader(msg, nil)
				msg.Header.SetField(tagMsgType, FIXString("D"))
				msg.Header.SetField(tagMsgSeqNum, FIXInt(seqNum))
				s.Require().Nil(s.session.store.SaveMessage(seqNum, msg.build()))
			}

			request := s.ResendRequest(test.beginSeqNo)
			request.Body.SetField(tagEndSeqNo, FIXInt(test.endSeqNo))
			s.MockApp.On("FromAdmin").Return(nil).Once()
			s.MockApp.On("ToApp").Return(nil).Times(len(test.expectedSeqNums))
			s.MockApp.On("ToAdmin").Maybe()

			s.fixMsgIn(s.session, request)

			s.MockApp.AssertExpectations(s.T())
			s.MockApp.AssertNumberOfCalls(s.T(), "ToAdmin", 0)
			s.State(inSession{})
			s.NextSenderMsgSeqNum(test.nextSenderSeqNum)
			s.NextTargetMsgSeqNum(2)
			s.Require().Len(s.Receiver.sendChannel, len(test.expectedSeqNums), "resend only the requested messages, without Reject or GapFill")
			for _, seqNum := range test.expectedSeqNums {
				msgBytes, ok := s.Receiver.LastMessage()
				s.Require().True(ok)
				response := NewMessage()
				s.Require().Nil(ParseMessage(response, bytes.NewBuffer(msgBytes)))
				s.MessageType("D", response)
				s.FieldEquals(tagMsgSeqNum, seqNum, response.Header)
				s.FieldEquals(tagPossDupFlag, true, response.Header)
			}
		})
	}
}

// TestFIXMsgInResendRequestPolicyUnsentRange checks the engine policy of rejecting
// requests starting at an unsent sequence number, including when persistence is
// disabled. This response is not mandated by FIX.
func (s *InSessionTestSuite) TestFIXMsgInResendRequestPolicyUnsentRange() {
	tests := []struct {
		name           string
		beginSeqNo     int
		endSeqNo       int
		disablePersist bool
	}{
		{name: "begin_at_next_sender_open_ended", beginSeqNo: 11, endSeqNo: 0},
		{name: "begin_above_next_sender_finite", beginSeqNo: 12, endSeqNo: 20},
		{name: "begin_at_next_sender_no_persist", beginSeqNo: 11, endSeqNo: 0, disablePersist: true},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			s.SetupTest()
			s.session.DisableMessagePersist = test.disablePersist
			s.Require().Nil(s.session.store.SetNextSenderMsgSeqNum(11))
			request := s.ResendRequest(test.beginSeqNo)
			request.Body.SetField(tagEndSeqNo, FIXInt(test.endSeqNo))
			s.MockApp.On("FromAdmin").Return(nil).Once()
			s.MockApp.On("ToAdmin").Once()

			s.fixMsgIn(s.session, request)

			s.MockApp.AssertExpectations(s.T())
			s.MockApp.AssertNumberOfCalls(s.T(), "ToApp", 0)
			s.State(inSession{})
			s.NextSenderMsgSeqNum(12)
			s.NextTargetMsgSeqNum(2)
			s.Require().Len(s.Receiver.sendChannel, 1, "send only a Reject, without replay or GapFill")
			msgBytes, ok := s.Receiver.LastMessage()
			s.Require().True(ok)
			response := NewMessage()
			s.Require().Nil(ParseMessage(response, bytes.NewBuffer(msgBytes)))
			s.MessageType(string(msgTypeReject), response)
			s.FieldEquals(tagMsgSeqNum, 11, response.Header)
			s.FieldEquals(tagRefSeqNum, 1, response.Body)
			s.FieldEquals(tagSessionRejectReason, rejectReasonValueIsIncorrect, response.Body)
			s.FieldEquals(tagRefTagID, int(tagBeginSeqNo), response.Body)
			s.FieldEquals(tagRefMsgType, string(msgTypeResendRequest), response.Body)
		})
	}
}

// TestFIXMsgInResendRequestRejectSequence checks that a Reject does not advance
// the inbound sequence for an out-of-order request or an already processed duplicate.
func (s *InSessionTestSuite) TestFIXMsgInResendRequestRejectSequence() {
	tests := []struct {
		name       string
		nextTarget int
		requestSeq int
		beginSeqNo int
		endSeqNo   int
		refTagID   Tag
		duplicate  bool
	}{
		{name: "unsent_range_high", nextTarget: 1, requestSeq: 5, beginSeqNo: 11, endSeqNo: 20, refTagID: tagBeginSeqNo},
		{name: "unsent_range_duplicate", nextTarget: 3, requestSeq: 1, beginSeqNo: 11, endSeqNo: 20, refTagID: tagBeginSeqNo, duplicate: true},
		{name: "reversed_range_high", nextTarget: 1, requestSeq: 5, beginSeqNo: 10, endSeqNo: 7, refTagID: tagEndSeqNo},
		{name: "reversed_range_duplicate", nextTarget: 3, requestSeq: 1, beginSeqNo: 10, endSeqNo: 7, refTagID: tagEndSeqNo, duplicate: true},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			s.SetupTest()
			s.Require().Nil(s.session.store.SetNextSenderMsgSeqNum(11))
			s.Require().Nil(s.session.store.SetNextTargetMsgSeqNum(test.nextTarget))
			request := s.ResendRequest(test.beginSeqNo)
			request.Header.SetField(tagMsgSeqNum, FIXInt(test.requestSeq))
			request.Body.SetField(tagEndSeqNo, FIXInt(test.endSeqNo))
			if test.duplicate {
				request.Header.SetField(tagPossDupFlag, FIXBoolean(true))
				sendingTime, err := request.Header.GetString(tagSendingTime)
				s.Require().Nil(err)
				request.Header.SetField(tagOrigSendingTime, FIXString(sendingTime))
			}
			s.MockApp.On("FromAdmin").Return(nil).Once()
			s.MockApp.On("ToAdmin").Once()

			s.fixMsgIn(s.session, request)

			s.MockApp.AssertExpectations(s.T())
			s.MockApp.AssertNumberOfCalls(s.T(), "ToApp", 0)
			s.State(inSession{})
			s.NextSenderMsgSeqNum(12)
			s.NextTargetMsgSeqNum(test.nextTarget)
			s.Require().Len(s.Receiver.sendChannel, 1, "send only a Reject, without replay or GapFill")
			msgBytes, ok := s.Receiver.LastMessage()
			s.Require().True(ok)
			response := NewMessage()
			s.Require().Nil(ParseMessage(response, bytes.NewBuffer(msgBytes)))
			s.MessageType(string(msgTypeReject), response)
			s.FieldEquals(tagMsgSeqNum, 11, response.Header)
			s.FieldEquals(tagRefSeqNum, test.requestSeq, response.Body)
			s.FieldEquals(tagSessionRejectReason, rejectReasonValueIsIncorrect, response.Body)
			s.FieldEquals(tagRefTagID, int(test.refTagID), response.Body)
			s.FieldEquals(tagRefMsgType, string(msgTypeResendRequest), response.Body)
		})
	}
}
