// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2026, runasexe mcatc-faag@runas.name

package modes

import "fmt"

// TableIA5 maps 6-bit IA-5 character codes (subset defined in ICAO Annex 10 Volume IV, 3.1.2.9.1.2)
// to ASCII characters. Valid index range is [0x00, 0x3F].
var TableIA5 = [0x40]byte{
	0x01: 'A',
	0x02: 'B',
	0x03: 'C',
	0x04: 'D',
	0x05: 'E',
	0x06: 'F',
	0x07: 'G',
	0x08: 'H',
	0x09: 'I',
	0x0a: 'J',
	0x0b: 'K',
	0x0c: 'L',
	0x0d: 'M',
	0x0e: 'N',
	0x0f: 'O',
	0x10: 'P',
	0x11: 'Q',
	0x12: 'R',
	0x13: 'S',
	0x14: 'T',
	0x15: 'U',
	0x16: 'V',
	0x17: 'W',
	0x18: 'X',
	0x19: 'Y',
	0x1a: 'Z',
	0x20: ' ',
	0x30: '0',
	0x31: '1',
	0x32: '2',
	0x33: '3',
	0x34: '4',
	0x35: '5',
	0x36: '6',
	0x37: '7',
	0x38: '8',
	0x39: '9',
}

// TableIA5Reversed maps ASCII byte values [0x00, 0xFF] back to their corresponding 6-bit IA-5 character codes.
var TableIA5Reversed [0x100]uint8

func init() {
	for idx, e := range TableIA5 {
		if idx > 0x3F {
			panic(fmt.Errorf("invalid value for %d, valid range [0, 63]", idx))
		}
		TableIA5Reversed[e] = uint8(idx)
	}
}
