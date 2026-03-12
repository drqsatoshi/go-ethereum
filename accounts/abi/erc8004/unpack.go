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
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// UnpackAgentRegisteredEvent decodes the non-indexed data from an
// AgentRegistered(uint256,address,string) event log.
// Indexed fields (tokenId, owner) must be extracted from the log topics separately.
func UnpackAgentRegisteredEvent(data []byte) (string, error) {
	out, err := IdentityRegistry.Unpack("AgentRegistered", data)
	if err != nil {
		return "", fmt.Errorf("erc8004: unpack AgentRegistered: %w", err)
	}
	agentURI, ok := out[0].(string)
	if !ok {
		return "", fmt.Errorf("erc8004: AgentRegistered: expected string, got %T", out[0])
	}
	return agentURI, nil
}

// UnpackReputationSummary decodes the return value of
// getReputationSummary(uint256) → (uint256 totalScore, uint256 feedbackCount).
func UnpackReputationSummary(data []byte) (*ReputationSummary, error) {
	out, err := ReputationRegistry.Unpack("getReputationSummary", data)
	if err != nil {
		return nil, fmt.Errorf("erc8004: unpack getReputationSummary: %w", err)
	}
	if len(out) != 2 {
		return nil, fmt.Errorf("erc8004: getReputationSummary: expected 2 outputs, got %d", len(out))
	}
	totalScore, ok := out[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("erc8004: getReputationSummary: expected *big.Int for totalScore, got %T", out[0])
	}
	feedbackCount, ok := out[1].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("erc8004: getReputationSummary: expected *big.Int for feedbackCount, got %T", out[1])
	}
	return &ReputationSummary{
		TotalScore:    totalScore,
		FeedbackCount: feedbackCount,
	}, nil
}

// UnpackFeedbackGivenEvent decodes the non-indexed data from a
// FeedbackGiven(uint256,address,uint8,string) event log.
// Indexed fields (agentId, reviewer) must be extracted from the log topics separately.
func UnpackFeedbackGivenEvent(data []byte) (uint8, string, error) {
	out, err := ReputationRegistry.Unpack("FeedbackGiven", data)
	if err != nil {
		return 0, "", fmt.Errorf("erc8004: unpack FeedbackGiven: %w", err)
	}
	if len(out) != 2 {
		return 0, "", fmt.Errorf("erc8004: FeedbackGiven: expected 2 outputs, got %d", len(out))
	}
	score, ok := out[0].(uint8)
	if !ok {
		return 0, "", fmt.Errorf("erc8004: FeedbackGiven: expected uint8 for score, got %T", out[0])
	}
	comment, ok := out[1].(string)
	if !ok {
		return 0, "", fmt.Errorf("erc8004: FeedbackGiven: expected string for comment, got %T", out[1])
	}
	return score, comment, nil
}

// UnpackValidationRequestedEvent decodes the non-indexed data from a
// ValidationRequested(uint256,uint256,bytes32,uint8) event log.
// Indexed fields (requestId, agentId) must be extracted from the log topics separately.
func UnpackValidationRequestedEvent(data []byte) ([32]byte, uint8, error) {
	out, err := ValidationRegistry.Unpack("ValidationRequested", data)
	if err != nil {
		return [32]byte{}, 0, fmt.Errorf("erc8004: unpack ValidationRequested: %w", err)
	}
	if len(out) != 2 {
		return [32]byte{}, 0, fmt.Errorf("erc8004: ValidationRequested: expected 2 outputs, got %d", len(out))
	}
	taskHash, ok := out[0].([32]byte)
	if !ok {
		return [32]byte{}, 0, fmt.Errorf("erc8004: ValidationRequested: expected [32]byte for taskHash, got %T", out[0])
	}
	validationType, ok := out[1].(uint8)
	if !ok {
		return [32]byte{}, 0, fmt.Errorf("erc8004: ValidationRequested: expected uint8 for validationType, got %T", out[1])
	}
	return taskHash, validationType, nil
}

// UnpackValidationSubmittedEvent decodes the non-indexed data from a
// ValidationSubmitted(uint256,address,bool) event log.
// Indexed fields (requestId, validator) must be extracted from the log topics separately.
func UnpackValidationSubmittedEvent(data []byte) (bool, error) {
	out, err := ValidationRegistry.Unpack("ValidationSubmitted", data)
	if err != nil {
		return false, fmt.Errorf("erc8004: unpack ValidationSubmitted: %w", err)
	}
	if len(out) != 1 {
		return false, fmt.Errorf("erc8004: ValidationSubmitted: expected 1 output, got %d", len(out))
	}
	result, ok := out[0].(bool)
	if !ok {
		return false, fmt.Errorf("erc8004: ValidationSubmitted: expected bool for result, got %T", out[0])
	}
	return result, nil
}

// UnpackRegisterAgentReturn decodes the return value of registerAgent(string) → uint256.
func UnpackRegisterAgentReturn(data []byte) (*big.Int, error) {
	out, err := IdentityRegistry.Unpack("registerAgent", data)
	if err != nil {
		return nil, fmt.Errorf("erc8004: unpack registerAgent return: %w", err)
	}
	if len(out) != 1 {
		return nil, fmt.Errorf("erc8004: registerAgent: expected 1 output, got %d", len(out))
	}
	tokenID, ok := out[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("erc8004: registerAgent: expected *big.Int, got %T", out[0])
	}
	return tokenID, nil
}

// UnpackAgentURIReturn decodes the return value of agentURI(uint256) → string.
func UnpackAgentURIReturn(data []byte) (string, error) {
	out, err := IdentityRegistry.Unpack("agentURI", data)
	if err != nil {
		return "", fmt.Errorf("erc8004: unpack agentURI return: %w", err)
	}
	if len(out) != 1 {
		return "", fmt.Errorf("erc8004: agentURI: expected 1 output, got %d", len(out))
	}
	uri, ok := out[0].(string)
	if !ok {
		return "", fmt.Errorf("erc8004: agentURI: expected string, got %T", out[0])
	}
	return uri, nil
}

// UnpackOwnerOfReturn decodes the return value of ownerOf(uint256) → address.
func UnpackOwnerOfReturn(data []byte) (common.Address, error) {
	out, err := IdentityRegistry.Unpack("ownerOf", data)
	if err != nil {
		return common.Address{}, fmt.Errorf("erc8004: unpack ownerOf return: %w", err)
	}
	if len(out) != 1 {
		return common.Address{}, fmt.Errorf("erc8004: ownerOf: expected 1 output, got %d", len(out))
	}
	owner, ok := out[0].(common.Address)
	if !ok {
		return common.Address{}, fmt.Errorf("erc8004: ownerOf: expected common.Address, got %T", out[0])
	}
	return owner, nil
}

// UnpackRequestValidationReturn decodes the return value of
// requestValidation(uint256,bytes32,uint8) → uint256.
func UnpackRequestValidationReturn(data []byte) (*big.Int, error) {
	out, err := ValidationRegistry.Unpack("requestValidation", data)
	if err != nil {
		return nil, fmt.Errorf("erc8004: unpack requestValidation return: %w", err)
	}
	if len(out) != 1 {
		return nil, fmt.Errorf("erc8004: requestValidation: expected 1 output, got %d", len(out))
	}
	requestID, ok := out[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("erc8004: requestValidation: expected *big.Int, got %T", out[0])
	}
	return requestID, nil
}
