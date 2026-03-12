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
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	// IdentityRegistry is the pre-parsed ABI for the ERC-8004 Identity Registry.
	IdentityRegistry abi.ABI

	// ReputationRegistry is the pre-parsed ABI for the ERC-8004 Reputation Registry.
	ReputationRegistry abi.ABI

	// ValidationRegistry is the pre-parsed ABI for the ERC-8004 Validation Registry.
	ValidationRegistry abi.ABI
)

func init() {
	var err error

	IdentityRegistry, err = abi.JSON(strings.NewReader(IdentityRegistryABI))
	if err != nil {
		panic("erc8004: failed to parse Identity Registry ABI: " + err.Error())
	}
	ReputationRegistry, err = abi.JSON(strings.NewReader(ReputationRegistryABI))
	if err != nil {
		panic("erc8004: failed to parse Reputation Registry ABI: " + err.Error())
	}
	ValidationRegistry, err = abi.JSON(strings.NewReader(ValidationRegistryABI))
	if err != nil {
		panic("erc8004: failed to parse Validation Registry ABI: " + err.Error())
	}
}
