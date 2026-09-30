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
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/quickfixgo/quickfix/datadictionary"
)

// buildNewOrderSingle builds a minimal NewOrderSingle (35=D) message programmatically.
func buildNewOrderSingle() *Message {
	msg := NewMessage()
	msg.Header.SetString(tagBeginString, "FIX.4.2")
	msg.Header.SetString(tagMsgType, "D")
	msg.Header.SetString(tagSenderCompID, "TW")
	msg.Header.SetString(tagTargetCompID, "ISLD")
	msg.Header.SetInt(tagMsgSeqNum, 2)
	msg.Header.SetString(Tag(52), "20140515-19:49:56.659") // SendingTime
	msg.Body.SetString(Tag(11), "100")                     // ClOrdID
	msg.Body.SetString(Tag(21), "1")                       // HandlInst
	msg.Body.SetString(Tag(40), "1")                       // OrdType
	msg.Body.SetString(Tag(54), "1")                       // Side
	msg.Body.SetString(Tag(55), "TSLA")                    // Symbol
	msg.Body.SetString(Tag(60), "00010101-00:00:00.000")   // TransactTime
	return msg
}

// helper: load a data dictionary from the spec directory.
func mustLoadDD(t *testing.T, path string) *datadictionary.DataDictionary {
	t.Helper()
	dd, err := datadictionary.Parse(path)
	require.NoError(t, err)
	return dd
}

// ----------------------------------------------------------------------------
// ToJSON – no data dictionary (numeric keys)
// ----------------------------------------------------------------------------

func TestToJSON_NoDictionary(t *testing.T) {
	msg := buildNewOrderSingle()

	jsonBytes, err := msg.ToJSON(nil)
	require.NoError(t, err)

	var parsed map[string]map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(jsonBytes, &parsed))

	// Top-level sections must be present.
	assert.Contains(t, parsed, "Header")
	assert.Contains(t, parsed, "Body")
	assert.Contains(t, parsed, "Trailer")

	// Field values are stored as numeric keys when no DD is available.
	headerStr := func(key string) string {
		var s string
		require.NoError(t, json.Unmarshal(parsed["Header"][key], &s))
		return s
	}
	assert.Equal(t, "FIX.4.2", headerStr("8"))
	assert.Equal(t, "D", headerStr("35"))
	assert.Equal(t, "TW", headerStr("49"))
	assert.Equal(t, "ISLD", headerStr("56"))

	bodyStr := func(key string) string {
		var s string
		require.NoError(t, json.Unmarshal(parsed["Body"][key], &s))
		return s
	}
	assert.Equal(t, "TSLA", bodyStr("55"))
	assert.Equal(t, "1", bodyStr("54"))

	// CheckSum must NOT appear in the Trailer section.
	_, hasCheckSum := parsed["Trailer"]["10"]
	assert.False(t, hasCheckSum, "CheckSum should be excluded from JSON output")
}

// ----------------------------------------------------------------------------
// ToJSON – with data dictionary (named keys)
// ----------------------------------------------------------------------------

func TestToJSON_WithDictionary(t *testing.T) {
	dd := mustLoadDD(t, "spec/FIX42.xml")

	msg := buildNewOrderSingle()

	jsonBytes, err := msg.ToJSON(dd)
	require.NoError(t, err)

	var parsed map[string]map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(jsonBytes, &parsed))

	strVal := func(section, key string) string {
		t.Helper()
		v, ok := parsed[section][key]
		require.True(t, ok, "section %q key %q not found in JSON", section, key)
		var s string
		require.NoError(t, json.Unmarshal(v, &s))
		return s
	}

	// Named keys from the data dictionary.
	assert.Equal(t, "FIX.4.2", strVal("Header", "BeginString"))
	assert.Equal(t, "D", strVal("Header", "MsgType"))
	assert.Equal(t, "TW", strVal("Header", "SenderCompID"))
	assert.Equal(t, "ISLD", strVal("Header", "TargetCompID"))
	assert.Equal(t, "TSLA", strVal("Body", "Symbol"))
	assert.Equal(t, "1", strVal("Body", "Side"))

	// CheckSum must NOT appear even with a DD.
	_, hasCheckSum := parsed["Trailer"]["CheckSum"]
	assert.False(t, hasCheckSum, "CheckSum should be excluded from JSON output")
}

// ----------------------------------------------------------------------------
// Round-trip: ToJSON → FromJSON (no data dictionary)
// ----------------------------------------------------------------------------

func TestFromJSON_NoDictionary_RoundTrip(t *testing.T) {
	orig := buildNewOrderSingle()

	jsonBytes, err := orig.ToJSON(nil)
	require.NoError(t, err)

	restored := NewMessage()
	require.NoError(t, restored.FromJSON(jsonBytes, nil))

	// Header fields.
	bs, err := restored.Header.GetString(tagBeginString)
	require.NoError(t, err)
	assert.Equal(t, "FIX.4.2", bs)

	mt, err := restored.Header.GetString(tagMsgType)
	require.NoError(t, err)
	assert.Equal(t, "D", mt)

	// Body fields.
	sym, err := restored.Body.GetString(Tag(55))
	require.NoError(t, err)
	assert.Equal(t, "TSLA", sym)

	side, err := restored.Body.GetString(Tag(54))
	require.NoError(t, err)
	assert.Equal(t, "1", side)
}

// ----------------------------------------------------------------------------
// Round-trip: ToJSON → FromJSON (with data dictionary, named keys)
// ----------------------------------------------------------------------------

func TestFromJSON_WithDictionary_RoundTrip(t *testing.T) {
	dd := mustLoadDD(t, "spec/FIX42.xml")

	orig := buildNewOrderSingle()

	jsonBytes, err := orig.ToJSON(dd)
	require.NoError(t, err)

	// The JSON now has named keys; FromJSON should resolve them back to tags.
	restored := NewMessage()
	require.NoError(t, restored.FromJSON(jsonBytes, dd))

	sym, err := restored.Body.GetString(Tag(55))
	require.NoError(t, err)
	assert.Equal(t, "TSLA", sym)

	senderID, err := restored.Header.GetString(tagSenderCompID)
	require.NoError(t, err)
	assert.Equal(t, "TW", senderID)
}

// ----------------------------------------------------------------------------
// Repeating groups
// ----------------------------------------------------------------------------

func TestToJSON_RepeatingGroups(t *testing.T) {
	dd := mustLoadDD(t, "spec/FIX44.xml")

	// ExecutionReport (35=8) with two NoContraBrokers (tag 382) instances.
	// Build the message programmatically.
	msg := NewMessage()
	msg.Header.SetString(tagBeginString, "FIX.4.4")
	msg.Header.SetString(tagMsgType, "8")
	msg.Header.SetString(tagSenderCompID, "SENDER")
	msg.Header.SetString(tagTargetCompID, "TARGET")
	msg.Header.SetInt(tagMsgSeqNum, 1)
	msg.Body.SetString(Tag(37), "ORD001")  // OrderID
	msg.Body.SetString(Tag(17), "EXEC001") // ExecID
	msg.Body.SetString(Tag(150), "0")      // ExecType
	msg.Body.SetString(Tag(39), "0")       // OrdStatus
	msg.Body.SetString(Tag(55), "AAPL")    // Symbol
	msg.Body.SetString(Tag(54), "1")       // Side
	msg.Body.SetString(Tag(151), "100")    // LeavesQty
	msg.Body.SetString(Tag(14), "0")       // CumQty
	msg.Body.SetString(Tag(6), "0")        // AvgPx

	// NoContraBrokers (tag 382) group – use FIX44 template from DD.
	msgDef, ok := dd.Messages["8"]
	require.True(t, ok, "ExecutionReport not found in FIX44 DD")
	grpFieldDef, ok := msgDef.Fields[382]
	require.True(t, ok, "NoContraBrokers (382) not in ExecutionReport")
	require.True(t, grpFieldDef.IsGroup(), "tag 382 should be a group")

	template := buildGroupTemplate(grpFieldDef)
	rg := NewRepeatingGroup(382, template)

	g1 := rg.Add()
	g1.SetString(Tag(375), "BRKR1") // ContraBroker
	g1.SetString(Tag(337), "ACC1")  // ContraTrader

	g2 := rg.Add()
	g2.SetString(Tag(375), "BRKR2")
	g2.SetString(Tag(337), "ACC2")

	msg.Body.SetGroup(rg)

	jsonBytes, err := msg.ToJSON(dd)
	require.NoError(t, err)

	var parsed map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(jsonBytes, &parsed))

	var body map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(parsed["Body"], &body))

	// NoContraBrokers should be serialized as a JSON array.
	rawGrp, ok := body["NoContraBrokers"]
	require.True(t, ok, "NoContraBrokers not found in Body JSON")

	var grpArray []map[string]json.RawMessage
	require.NoError(t, json.Unmarshal(rawGrp, &grpArray))
	require.Len(t, grpArray, 2)

	strFromRaw := func(raw json.RawMessage) string {
		var s string
		require.NoError(t, json.Unmarshal(raw, &s))
		return s
	}

	assert.Equal(t, "BRKR1", strFromRaw(grpArray[0]["ContraBroker"]))
	assert.Equal(t, "ACC1", strFromRaw(grpArray[0]["ContraTrader"]))
	assert.Equal(t, "BRKR2", strFromRaw(grpArray[1]["ContraBroker"]))
	assert.Equal(t, "ACC2", strFromRaw(grpArray[1]["ContraTrader"]))
}

func TestFromJSON_RepeatingGroups_RoundTrip(t *testing.T) {
	dd := mustLoadDD(t, "spec/FIX44.xml")

	// Build the source JSON directly.
	srcJSON := []byte(`{
		"Header": {
			"BeginString": "FIX.4.4",
			"MsgType":     "8",
			"SenderCompID":"SENDER",
			"TargetCompID":"TARGET",
			"MsgSeqNum":   "1"
		},
		"Body": {
			"OrderID":   "ORD001",
			"ExecID":    "EXEC001",
			"ExecType":  "0",
			"OrdStatus": "0",
			"Symbol":    "AAPL",
			"Side":      "1",
			"LeavesQty": "100",
			"CumQty":    "0",
			"AvgPx":     "0",
			"NoContraBrokers": [
				{"ContraBroker": "BRKR1", "ContraTrader": "ACC1"},
				{"ContraBroker": "BRKR2", "ContraTrader": "ACC2"}
			]
		},
		"Trailer": {}
	}`)

	msg := NewMessage()
	require.NoError(t, msg.FromJSON(srcJSON, dd))

	// Verify scalar fields.
	sym, err := msg.Body.GetString(Tag(55))
	require.NoError(t, err)
	assert.Equal(t, "AAPL", sym)

	// Verify repeating group.
	msgDef := dd.Messages["8"]
	grpFieldDef := msgDef.Fields[382]
	template := buildGroupTemplate(grpFieldDef)
	rg := NewRepeatingGroup(382, template)

	require.NoError(t, msg.Body.GetGroup(rg))
	require.Equal(t, 2, rg.Len())

	broker1, err := rg.Get(0).GetString(Tag(375))
	require.NoError(t, err)
	assert.Equal(t, "BRKR1", broker1)

	broker2, err := rg.Get(1).GetString(Tag(375))
	require.NoError(t, err)
	assert.Equal(t, "BRKR2", broker2)
}

// ----------------------------------------------------------------------------
// FromJSON error handling
// ----------------------------------------------------------------------------

func TestFromJSON_InvalidJSON(t *testing.T) {
	msg := NewMessage()
	err := msg.FromJSON([]byte(`not json`), nil)
	assert.Error(t, err)
}
