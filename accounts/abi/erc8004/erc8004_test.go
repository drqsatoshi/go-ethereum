// Copyright 2026 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package erc8004

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// TestRegistryABIParsing verifies that all three registry ABIs parse without error
// and contain the expected methods, events, and errors.
func TestRegistryABIParsing(t *testing.T) {
	// Identity Registry
	if _, ok := IdentityRegistry.Methods["registerAgent"]; !ok {
		t.Fatal("Identity Registry missing registerAgent method")
	}
	if _, ok := IdentityRegistry.Methods["agentURI"]; !ok {
		t.Fatal("Identity Registry missing agentURI method")
	}
	if _, ok := IdentityRegistry.Methods["ownerOf"]; !ok {
		t.Fatal("Identity Registry missing ownerOf method")
	}
	if _, ok := IdentityRegistry.Events["AgentRegistered"]; !ok {
		t.Fatal("Identity Registry missing AgentRegistered event")
	}
	if _, ok := IdentityRegistry.Errors["Unauthorized"]; !ok {
		t.Fatal("Identity Registry missing Unauthorized error")
	}

	// Reputation Registry
	if _, ok := ReputationRegistry.Methods["giveFeedback"]; !ok {
		t.Fatal("Reputation Registry missing giveFeedback method")
	}
	if _, ok := ReputationRegistry.Methods["getReputationSummary"]; !ok {
		t.Fatal("Reputation Registry missing getReputationSummary method")
	}
	if _, ok := ReputationRegistry.Events["FeedbackGiven"]; !ok {
		t.Fatal("Reputation Registry missing FeedbackGiven event")
	}
	if _, ok := ReputationRegistry.Errors["InvalidSignature"]; !ok {
		t.Fatal("Reputation Registry missing InvalidSignature error")
	}

	// Validation Registry
	if _, ok := ValidationRegistry.Methods["requestValidation"]; !ok {
		t.Fatal("Validation Registry missing requestValidation method")
	}
	if _, ok := ValidationRegistry.Methods["submitValidation"]; !ok {
		t.Fatal("Validation Registry missing submitValidation method")
	}
	if _, ok := ValidationRegistry.Events["ValidationRequested"]; !ok {
		t.Fatal("Validation Registry missing ValidationRequested event")
	}
	if _, ok := ValidationRegistry.Events["ValidationSubmitted"]; !ok {
		t.Fatal("Validation Registry missing ValidationSubmitted event")
	}
	if _, ok := ValidationRegistry.Errors["ValidationNotFound"]; !ok {
		t.Fatal("Validation Registry missing ValidationNotFound error")
	}
}

// TestPackRegisterAgent verifies ABI encoding of a registerAgent call.
func TestPackRegisterAgent(t *testing.T) {
	data, err := PackRegisterAgent("https://example.com/agent.json")
	if err != nil {
		t.Fatalf("PackRegisterAgent failed: %v", err)
	}
	// 4-byte selector + ABI-encoded string
	if len(data) < 4 {
		t.Fatalf("packed data too short: %d bytes", len(data))
	}
	// Verify the 4-byte method selector matches
	method := IdentityRegistry.Methods["registerAgent"]
	for i := 0; i < 4; i++ {
		if data[i] != method.ID[i] {
			t.Fatalf("selector mismatch at byte %d: got %#x, want %#x", i, data[i], method.ID[i])
		}
	}
}

// TestPackGiveFeedback verifies ABI encoding of a giveFeedback call.
func TestPackGiveFeedback(t *testing.T) {
	agentID := big.NewInt(42)
	sig := []byte{0x01, 0x02, 0x03}
	data, err := PackGiveFeedback(agentID, 5, "great work", sig)
	if err != nil {
		t.Fatalf("PackGiveFeedback failed: %v", err)
	}
	if len(data) < 4 {
		t.Fatalf("packed data too short: %d bytes", len(data))
	}
	method := ReputationRegistry.Methods["giveFeedback"]
	for i := 0; i < 4; i++ {
		if data[i] != method.ID[i] {
			t.Fatalf("selector mismatch at byte %d: got %#x, want %#x", i, data[i], method.ID[i])
		}
	}
}

// TestPackRequestValidation verifies ABI encoding of a requestValidation call.
func TestPackRequestValidation(t *testing.T) {
	agentID := big.NewInt(1)
	taskHash := [32]byte{0xab, 0xcd}
	data, err := PackRequestValidation(agentID, taskHash, uint8(ValidationZKML))
	if err != nil {
		t.Fatalf("PackRequestValidation failed: %v", err)
	}
	if len(data) < 4 {
		t.Fatalf("packed data too short: %d bytes", len(data))
	}
	method := ValidationRegistry.Methods["requestValidation"]
	for i := 0; i < 4; i++ {
		if data[i] != method.ID[i] {
			t.Fatalf("selector mismatch at byte %d: got %#x, want %#x", i, data[i], method.ID[i])
		}
	}
}

// TestPackSubmitValidation verifies ABI encoding of a submitValidation call.
func TestPackSubmitValidation(t *testing.T) {
	reqID := big.NewInt(99)
	proof := []byte{0xde, 0xad, 0xbe, 0xef}
	data, err := PackSubmitValidation(reqID, true, proof)
	if err != nil {
		t.Fatalf("PackSubmitValidation failed: %v", err)
	}
	if len(data) < 4 {
		t.Fatalf("packed data too short: %d bytes", len(data))
	}
	method := ValidationRegistry.Methods["submitValidation"]
	for i := 0; i < 4; i++ {
		if data[i] != method.ID[i] {
			t.Fatalf("selector mismatch at byte %d: got %#x, want %#x", i, data[i], method.ID[i])
		}
	}
}

// TestUnpackReputationSummary verifies decoding of getReputationSummary return data.
func TestUnpackReputationSummary(t *testing.T) {
	// Encode known values, then decode and compare.
	totalScore := big.NewInt(450)
	feedbackCount := big.NewInt(10)

	// Pack the output values using the method's output arguments.
	method := ReputationRegistry.Methods["getReputationSummary"]
	packed, err := method.Outputs.Pack(totalScore, feedbackCount)
	if err != nil {
		t.Fatalf("failed to pack outputs: %v", err)
	}

	summary, err := UnpackReputationSummary(packed)
	if err != nil {
		t.Fatalf("UnpackReputationSummary failed: %v", err)
	}
	if summary.TotalScore.Cmp(totalScore) != 0 {
		t.Fatalf("totalScore mismatch: got %s, want %s", summary.TotalScore, totalScore)
	}
	if summary.FeedbackCount.Cmp(feedbackCount) != 0 {
		t.Fatalf("feedbackCount mismatch: got %s, want %s", summary.FeedbackCount, feedbackCount)
	}
}

// TestUnpackRegisterAgentReturn verifies decoding of registerAgent return data.
func TestUnpackRegisterAgentReturn(t *testing.T) {
	expected := big.NewInt(7)
	method := IdentityRegistry.Methods["registerAgent"]
	packed, err := method.Outputs.Pack(expected)
	if err != nil {
		t.Fatalf("failed to pack output: %v", err)
	}
	tokenID, err := UnpackRegisterAgentReturn(packed)
	if err != nil {
		t.Fatalf("UnpackRegisterAgentReturn failed: %v", err)
	}
	if tokenID.Cmp(expected) != 0 {
		t.Fatalf("tokenID mismatch: got %s, want %s", tokenID, expected)
	}
}

// TestUnpackAgentURIReturn verifies decoding of agentURI return data.
func TestUnpackAgentURIReturn(t *testing.T) {
	expected := "https://example.com/agent.json"
	method := IdentityRegistry.Methods["agentURI"]
	packed, err := method.Outputs.Pack(expected)
	if err != nil {
		t.Fatalf("failed to pack output: %v", err)
	}
	uri, err := UnpackAgentURIReturn(packed)
	if err != nil {
		t.Fatalf("UnpackAgentURIReturn failed: %v", err)
	}
	if uri != expected {
		t.Fatalf("URI mismatch: got %q, want %q", uri, expected)
	}
}

// TestUnpackOwnerOfReturn verifies decoding of ownerOf return data.
func TestUnpackOwnerOfReturn(t *testing.T) {
	expected := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	method := IdentityRegistry.Methods["ownerOf"]
	packed, err := method.Outputs.Pack(expected)
	if err != nil {
		t.Fatalf("failed to pack output: %v", err)
	}
	owner, err := UnpackOwnerOfReturn(packed)
	if err != nil {
		t.Fatalf("UnpackOwnerOfReturn failed: %v", err)
	}
	if owner != expected {
		t.Fatalf("owner mismatch: got %s, want %s", owner.Hex(), expected.Hex())
	}
}

// TestUnpackAgentRegisteredEvent verifies decoding of AgentRegistered event data.
func TestUnpackAgentRegisteredEvent(t *testing.T) {
	expected := "https://example.com/my-agent"
	event := IdentityRegistry.Events["AgentRegistered"]
	// Pack the non-indexed inputs (only agentURI).
	packed, err := event.Inputs.NonIndexed().Pack(expected)
	if err != nil {
		t.Fatalf("failed to pack event data: %v", err)
	}
	uri, err := UnpackAgentRegisteredEvent(packed)
	if err != nil {
		t.Fatalf("UnpackAgentRegisteredEvent failed: %v", err)
	}
	if uri != expected {
		t.Fatalf("agentURI mismatch: got %q, want %q", uri, expected)
	}
}

// TestUnpackFeedbackGivenEvent verifies decoding of FeedbackGiven event data.
func TestUnpackFeedbackGivenEvent(t *testing.T) {
	event := ReputationRegistry.Events["FeedbackGiven"]
	packed, err := event.Inputs.NonIndexed().Pack(uint8(4), "nice agent")
	if err != nil {
		t.Fatalf("failed to pack event data: %v", err)
	}
	score, comment, err := UnpackFeedbackGivenEvent(packed)
	if err != nil {
		t.Fatalf("UnpackFeedbackGivenEvent failed: %v", err)
	}
	if score != 4 {
		t.Fatalf("score mismatch: got %d, want 4", score)
	}
	if comment != "nice agent" {
		t.Fatalf("comment mismatch: got %q, want %q", comment, "nice agent")
	}
}

// TestUnpackValidationSubmittedEvent verifies decoding of ValidationSubmitted event data.
func TestUnpackValidationSubmittedEvent(t *testing.T) {
	event := ValidationRegistry.Events["ValidationSubmitted"]
	packed, err := event.Inputs.NonIndexed().Pack(true)
	if err != nil {
		t.Fatalf("failed to pack event data: %v", err)
	}
	result, err := UnpackValidationSubmittedEvent(packed)
	if err != nil {
		t.Fatalf("UnpackValidationSubmittedEvent failed: %v", err)
	}
	if !result {
		t.Fatal("expected result=true, got false")
	}
}

// TestUnpackRequestValidationReturn verifies decoding of requestValidation return data.
func TestUnpackRequestValidationReturn(t *testing.T) {
	expected := big.NewInt(123)
	method := ValidationRegistry.Methods["requestValidation"]
	packed, err := method.Outputs.Pack(expected)
	if err != nil {
		t.Fatalf("failed to pack output: %v", err)
	}
	reqID, err := UnpackRequestValidationReturn(packed)
	if err != nil {
		t.Fatalf("UnpackRequestValidationReturn failed: %v", err)
	}
	if reqID.Cmp(expected) != 0 {
		t.Fatalf("requestID mismatch: got %s, want %s", reqID, expected)
	}
}

// TestValidationTypeConstants verifies the validation type enum values.
func TestValidationTypeConstants(t *testing.T) {
	if ValidationStaking != 0 {
		t.Fatalf("ValidationStaking: got %d, want 0", ValidationStaking)
	}
	if ValidationZKML != 1 {
		t.Fatalf("ValidationZKML: got %d, want 1", ValidationZKML)
	}
	if ValidationTEE != 2 {
		t.Fatalf("ValidationTEE: got %d, want 2", ValidationTEE)
	}
}
