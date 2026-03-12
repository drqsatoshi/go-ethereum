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

	"github.com/ethereum/go-ethereum/common"
)

// ValidationType enumerates the supported validation mechanisms.
type ValidationType uint8

const (
	ValidationStaking ValidationType = iota
	ValidationZKML
	ValidationTEE
)

// AgentRegistration mirrors the data stored per agent in the Identity Registry.
type AgentRegistration struct {
	TokenID  *big.Int
	Owner    common.Address
	AgentURI string
}

// Feedback represents a single reputation feedback entry.
type Feedback struct {
	AgentID   *big.Int
	Reviewer  common.Address
	Score     uint8
	Comment   string
	Signature []byte
}

// ReputationSummary holds the aggregated reputation data for an agent.
type ReputationSummary struct {
	TotalScore    *big.Int
	FeedbackCount *big.Int
}

// ValidationRequest represents a request for task validation.
type ValidationRequest struct {
	RequestID      *big.Int
	AgentID        *big.Int
	TaskHash       [32]byte
	ValidationType uint8
}

// ValidationResult represents the outcome of a validation.
type ValidationResult struct {
	RequestID *big.Int
	Validator common.Address
	Result    bool
	Proof     []byte
}
