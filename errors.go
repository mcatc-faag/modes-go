// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2026, runasexe mcatc-faag@runas.name

package modes

import "errors"

var (
	// ErrModeSInvalidSymbol indicates that an input character is not supported by the Mode S IA-5 encoding subset.
	ErrModeSInvalidSymbol = errors.New("modes: invalid symbol")

	// ErrDataSizeExceeds indicates that the input string length exceeds the maximum allowed 8 characters for Mode S callsign encoding.
	ErrDataSizeExceeds = errors.New("modes: the data size exceeds the maximum possible")
)
