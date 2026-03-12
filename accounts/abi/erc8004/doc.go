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

// Package erc8004 provides early support for ERC-8004 (Trustless Agents),
// the Ethereum standard for on-chain AI agent identity, reputation, and
// validation registries activating on mainnet March 31, 2026.
//
// ERC-8004 defines three registry contracts:
//   - Identity Registry: each agent is an ERC-721 NFT with a URI to an off-chain AgentCard
//   - Reputation Registry: on-chain feedback scores with EIP-191/ERC-1271 signature verification
//   - Validation Registry: generic hooks for independent task validation (staking, zkML, TEE)
//
// See https://eips.ethereum.org/EIPS/eip-8004
package erc8004
