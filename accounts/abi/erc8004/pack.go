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
)

// PackRegisterAgent ABI-encodes a call to Identity Registry registerAgent(string).
func PackRegisterAgent(agentURI string) ([]byte, error) {
	return IdentityRegistry.Pack("registerAgent", agentURI)
}

// PackAgentURI ABI-encodes a call to Identity Registry agentURI(uint256).
func PackAgentURI(tokenID *big.Int) ([]byte, error) {
	return IdentityRegistry.Pack("agentURI", tokenID)
}

// PackOwnerOf ABI-encodes a call to Identity Registry ownerOf(uint256).
func PackOwnerOf(tokenID *big.Int) ([]byte, error) {
	return IdentityRegistry.Pack("ownerOf", tokenID)
}

// PackGiveFeedback ABI-encodes a call to Reputation Registry
// giveFeedback(uint256,uint8,string,bytes).
func PackGiveFeedback(agentID *big.Int, score uint8, comment string, signature []byte) ([]byte, error) {
	return ReputationRegistry.Pack("giveFeedback", agentID, score, comment, signature)
}

// PackGetReputationSummary ABI-encodes a call to Reputation Registry
// getReputationSummary(uint256).
func PackGetReputationSummary(agentID *big.Int) ([]byte, error) {
	return ReputationRegistry.Pack("getReputationSummary", agentID)
}

// PackRequestValidation ABI-encodes a call to Validation Registry
// requestValidation(uint256,bytes32,uint8).
func PackRequestValidation(agentID *big.Int, taskHash [32]byte, validationType uint8) ([]byte, error) {
	return ValidationRegistry.Pack("requestValidation", agentID, taskHash, validationType)
}

// PackSubmitValidation ABI-encodes a call to Validation Registry
// submitValidation(uint256,bool,bytes).
func PackSubmitValidation(requestID *big.Int, result bool, proof []byte) ([]byte, error) {
	return ValidationRegistry.Pack("submitValidation", requestID, result, proof)
}
