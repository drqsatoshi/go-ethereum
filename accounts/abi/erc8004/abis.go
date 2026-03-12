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

// IdentityRegistryABI is the ABI JSON for the ERC-8004 Identity Registry.
// Agents are minted as ERC-721 NFTs with a URI pointing to an off-chain AgentCard.
const IdentityRegistryABI = `[
  {
    "type": "function",
    "name": "registerAgent",
    "inputs": [
      { "name": "agentURI", "type": "string" }
    ],
    "outputs": [
      { "name": "tokenId", "type": "uint256" }
    ],
    "stateMutability": "nonpayable"
  },
  {
    "type": "function",
    "name": "agentURI",
    "inputs": [
      { "name": "tokenId", "type": "uint256" }
    ],
    "outputs": [
      { "name": "", "type": "string" }
    ],
    "stateMutability": "view"
  },
  {
    "type": "function",
    "name": "ownerOf",
    "inputs": [
      { "name": "tokenId", "type": "uint256" }
    ],
    "outputs": [
      { "name": "", "type": "address" }
    ],
    "stateMutability": "view"
  },
  {
    "type": "event",
    "name": "AgentRegistered",
    "inputs": [
      { "name": "tokenId", "type": "uint256", "indexed": true },
      { "name": "owner", "type": "address", "indexed": true },
      { "name": "agentURI", "type": "string", "indexed": false }
    ]
  },
  {
    "type": "error",
    "name": "Unauthorized",
    "inputs": [
      { "name": "caller", "type": "address" }
    ]
  }
]`

// ReputationRegistryABI is the ABI JSON for the ERC-8004 Reputation Registry.
// Provides on-chain feedback with EIP-191/ERC-1271 signature verification.
const ReputationRegistryABI = `[
  {
    "type": "function",
    "name": "giveFeedback",
    "inputs": [
      { "name": "agentId", "type": "uint256" },
      { "name": "score", "type": "uint8" },
      { "name": "comment", "type": "string" },
      { "name": "signature", "type": "bytes" }
    ],
    "outputs": [],
    "stateMutability": "nonpayable"
  },
  {
    "type": "function",
    "name": "getReputationSummary",
    "inputs": [
      { "name": "agentId", "type": "uint256" }
    ],
    "outputs": [
      { "name": "totalScore", "type": "uint256" },
      { "name": "feedbackCount", "type": "uint256" }
    ],
    "stateMutability": "view"
  },
  {
    "type": "event",
    "name": "FeedbackGiven",
    "inputs": [
      { "name": "agentId", "type": "uint256", "indexed": true },
      { "name": "reviewer", "type": "address", "indexed": true },
      { "name": "score", "type": "uint8", "indexed": false },
      { "name": "comment", "type": "string", "indexed": false }
    ]
  },
  {
    "type": "error",
    "name": "InvalidSignature",
    "inputs": [
      { "name": "signer", "type": "address" }
    ]
  }
]`

// ValidationRegistryABI is the ABI JSON for the ERC-8004 Validation Registry.
// Provides generic hooks for independent task validation (staking, zkML, TEE).
const ValidationRegistryABI = `[
  {
    "type": "function",
    "name": "requestValidation",
    "inputs": [
      { "name": "agentId", "type": "uint256" },
      { "name": "taskHash", "type": "bytes32" },
      { "name": "validationType", "type": "uint8" }
    ],
    "outputs": [
      { "name": "requestId", "type": "uint256" }
    ],
    "stateMutability": "nonpayable"
  },
  {
    "type": "function",
    "name": "submitValidation",
    "inputs": [
      { "name": "requestId", "type": "uint256" },
      { "name": "result", "type": "bool" },
      { "name": "proof", "type": "bytes" }
    ],
    "outputs": [],
    "stateMutability": "nonpayable"
  },
  {
    "type": "event",
    "name": "ValidationRequested",
    "inputs": [
      { "name": "requestId", "type": "uint256", "indexed": true },
      { "name": "agentId", "type": "uint256", "indexed": true },
      { "name": "taskHash", "type": "bytes32", "indexed": false },
      { "name": "validationType", "type": "uint8", "indexed": false }
    ]
  },
  {
    "type": "event",
    "name": "ValidationSubmitted",
    "inputs": [
      { "name": "requestId", "type": "uint256", "indexed": true },
      { "name": "validator", "type": "address", "indexed": true },
      { "name": "result", "type": "bool", "indexed": false }
    ]
  },
  {
    "type": "error",
    "name": "ValidationNotFound",
    "inputs": [
      { "name": "requestId", "type": "uint256" }
    ]
  }
]`
