// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2026, runasexe mcatc-faag@runas.name

// Package modes implements encoding and decoding for Mode S aircraft identification (callsign).
//
// Aircraft identification in Mode S data links is transmitted as an 8-character callsign
// packed into 6 bytes (48 bits) using a 6-bit subset of IA-5 (ICAO Annex 10, Vol IV, 3.1.2.9.1.2).
package modes

import (
	"fmt"
)

// ModeS represents a 6-byte (48-bit) encoded Mode S aircraft identification callsign.
//
// The callsign consists of up to eight characters (A–Z, 0–9, space padding)
// packed into 6 bytes (8 x 6 bits = 48 bits).
// Reference: ICAO Annex 10 Vol IV / ASTERIX Item I048/240.
type ModeS [6]byte

var (
	// Zero represents a 6-byte Mode S value with all bytes initialized to zero.
	Zero = ModeS{}

	// Space represents a 6-byte Mode S value pre-filled with space padding characters.
	Space = ModeS{0b10000010, 0b00001000, 0b00100000, 0b10000010, 0b00001000, 0b00100000}
)

// String decodes the 6-byte Mode S representation into an 8-character string using the IA-5 character mapping.
// It returns an empty string if the first character code resolves to zero/unrecognized.
func (s ModeS) String() string {
	var buf [8]byte

	ch1 := s[0] & 0xFC >> 2
	buf[0] = TableIA5[ch1]

	if buf[0] == 0 {
		return ""
	}

	ch2 := s[0]&0x03<<4 + s[1]&0xF0>>4
	buf[1] = TableIA5[ch2]

	if buf[1] == 0 {
		return string(buf[:1])
	}

	ch3 := s[1]&0x0F<<2 + s[2]&0xC0>>6
	buf[2] = TableIA5[ch3]

	if buf[2] == 0 {
		return string(buf[:2])
	}

	ch4 := s[2] & 0x3F
	buf[3] = TableIA5[ch4]

	if buf[3] == 0 {
		return string(buf[:3])
	}

	ch5 := s[3] & 0xFC >> 2
	buf[4] = TableIA5[ch5]

	if buf[4] == 0 {
		return string(buf[:4])
	}

	ch6 := s[3]&0x03<<4 + s[4]&0xF0>>4
	buf[5] = TableIA5[ch6]

	if buf[5] == 0 {
		return string(buf[:5])
	}

	ch7 := s[4]&0x0F<<2 + s[5]&0xC0>>6
	buf[6] = TableIA5[ch7]

	if buf[6] == 0 {
		return string(buf[:6])
	}

	ch8 := s[5] & 0x3F
	buf[7] = TableIA5[ch8]

	if buf[7] == 0 {
		return string(buf[:7])
	}

	return string(buf[:8])
}

// SpaceTerminatedString decodes the 6-byte Mode S representation into a string and trims any trailing space padding characters.
func (s ModeS) SpaceTerminatedString() string {
	val := s.String()
	for len(val) > 0 && val[len(val)-1] == ' ' {
		val = val[:len(val)-1]
	}
	return val
}

// MakeModeSZero encodes an aircraft callsign string (up to 8 characters) into a 6-byte Mode S representation.
// Unused trailing character positions are zero-padded.
func MakeModeSZero(s string) (buf ModeS, _ error) {
	length := len(s)
	if length == 0 || s[0] == 0 {
		return buf, nil
	}
	if length > 8 {
		return buf, fmt.Errorf(`%w: %d`, ErrDataSizeExceeds, length)
	}

	if val := TableIA5Reversed[s[0]]; val == 0 {
		return buf, fmt.Errorf(`%w: %v`, ErrModeSInvalidSymbol, s[0])
	} else {
		buf[0] = val << 2 // 0b--AAAAAA << 2 => [0] 0bAAAAAA--
	}

	if length == 1 || s[1] == 0 {
		return buf, nil
	}

	if val := TableIA5Reversed[s[1]]; val == 0 {
		return buf, fmt.Errorf(`%s: %v`, ErrModeSInvalidSymbol, s[1])
	} else {
		buf[0] |= (val & 0x30) >> 4 // 0b--BB---- >> 4 => [0] 0b------BB
		buf[1] = (val & 0x0F) << 4  // 0b----BBBB << 4 => [1] 0bBBBB----
	}

	if length == 2 || s[2] == 0 {
		return buf, nil
	}

	if val := TableIA5Reversed[s[2]]; val == 0 {
		return buf, fmt.Errorf(`%s: %v`, ErrModeSInvalidSymbol, s[2])
	} else {
		buf[1] |= (val & 0x3C) >> 2 // 0b--CCCC-- >> 2 => [1] 0b----CCCC
		buf[2] = (val & 0x03) << 6  // 0b------CC << 6 => [2] 0bCC------
	}

	if length == 3 || s[3] == 0 {
		return buf, nil
	}

	if val := TableIA5Reversed[s[3]]; val == 0 {
		return buf, fmt.Errorf(`%s: %v`, ErrModeSInvalidSymbol, s[3])
	} else {
		buf[2] |= val & 0x3F // 0b--DDDDDD << 0 => [3] 0b--DDDDDD
	}

	if length == 4 || s[4] == 0 {
		return buf, nil
	}

	if val := TableIA5Reversed[s[4]]; val == 0 {
		return buf, fmt.Errorf(`%s: %v`, ErrModeSInvalidSymbol, s[4])
	} else {
		buf[3] = val << 2 // 0b--AAAAAA << 2 => [3] 0bAAAAAA--
	}

	if length == 5 || s[5] == 0 {
		return buf, nil
	}

	if val := TableIA5Reversed[s[5]]; val == 0 {
		return buf, fmt.Errorf(`%s: %v`, ErrModeSInvalidSymbol, s[5])
	} else {
		buf[3] |= (val & 0x30) >> 4 // 0b--BB---- >> 4 => [3] 0b------BB
		buf[4] = (val & 0x0F) << 4  // 0b----BBBB << 4 => [4] 0bBBBB----
	}

	if length == 6 || s[6] == 0 {
		return buf, nil
	}

	if val := TableIA5Reversed[s[6]]; val == 0 {
		return buf, fmt.Errorf(`%s: %v`, ErrModeSInvalidSymbol, s[6])
	} else {
		buf[4] |= (val & 0x3C) >> 2 // 0b--CCCC-- >> 2 => [4] 0b----CCCC
		buf[5] = (val & 0x03) << 6  // 0b------CC << 6 => [5] 0bCC------
	}

	if length == 7 || s[7] == 0 {
		return buf, nil
	}

	if val := TableIA5Reversed[s[7]]; val == 0 {
		return buf, fmt.Errorf(`%s: %v`, ErrModeSInvalidSymbol, s[7])
	} else {
		buf[5] |= val & 0x3F // 0b--DDDDDD << 0 => [5] 0b--DDDDDD
	}

	return buf, nil
}

// MakeModeSSpace encodes an aircraft callsign string (up to 8 characters) into a 6-byte Mode S representation.
// Unused trailing character positions are pre-filled with space padding characters.
func MakeModeSSpace(s string) (buf ModeS, _ error) {
	buf = Space

	length := len(s)
	if length == 0 || s[0] == 0 {
		return buf, nil
	}
	if length > 8 {
		return buf, fmt.Errorf(`%w: %d`, ErrDataSizeExceeds, length)
	}

	if val := TableIA5Reversed[s[0]]; val == 0 {
		return buf, fmt.Errorf(`%w: %v`, ErrModeSInvalidSymbol, s[0])
	} else {
		buf[0] = (val << 2) | 0b00000010 // 0b--AAAAAA << 2 => [0] 0bAAAAAA--
	}

	if length == 1 || s[1] == 0 {
		return buf, nil
	}

	if val := TableIA5Reversed[s[1]]; val == 0 {
		return buf, fmt.Errorf(`%s: %v`, ErrModeSInvalidSymbol, s[1])
	} else {
		buf[0] = buf[0]&0b11111100 | ((val & 0x30) >> 4) // 0b--BB---- >> 4 => [0] 0b------BB
		buf[1] = ((val & 0x0F) << 4) | 0b00001000        // 0b----BBBB << 4 => [1] 0bBBBB----
	}

	if length == 2 || s[2] == 0 {
		return buf, nil
	}

	if val := TableIA5Reversed[s[2]]; val == 0 {
		return buf, fmt.Errorf(`%s: %v`, ErrModeSInvalidSymbol, s[2])
	} else {
		buf[1] = buf[1]&0b11110000 | ((val & 0x3C) >> 2) // 0b--CCCC-- >> 2 => [1] 0b----CCCC
		buf[2] = (val&0x03)<<6 | 0b00100000              // 0b------CC << 6 => [2] 0bCC------
	}

	if length == 3 || s[3] == 0 {
		return buf, nil
	}

	if val := TableIA5Reversed[s[3]]; val == 0 {
		return buf, fmt.Errorf(`%s: %v`, ErrModeSInvalidSymbol, s[3])
	} else {
		buf[2] = buf[2]&0b11000000 | val&0x3F // 0b--DDDDDD << 0 => [3] 0b--DDDDDD
	}

	if length == 4 || s[4] == 0 {
		return buf, nil
	}

	if val := TableIA5Reversed[s[4]]; val == 0 {
		return buf, fmt.Errorf(`%s: %v`, ErrModeSInvalidSymbol, s[4])
	} else {
		buf[3] = (val << 2) | 0b00000010 // 0b--AAAAAA << 2 => [3] 0bAAAAAA--
	}

	if length == 5 || s[5] == 0 {
		return buf, nil
	}

	if val := TableIA5Reversed[s[5]]; val == 0 {
		return buf, fmt.Errorf(`%s: %v`, ErrModeSInvalidSymbol, s[5])
	} else {
		buf[3] = buf[3]&0b11111100 | ((val & 0x30) >> 4) // 0b--BB---- >> 4 => [3] 0b------BB
		buf[4] = ((val & 0x0F) << 4) | 0b00001000        // 0b----BBBB << 4 => [4] 0bBBBB----
	}

	if length == 6 || s[6] == 0 {
		return buf, nil
	}

	if val := TableIA5Reversed[s[6]]; val == 0 {
		return buf, fmt.Errorf(`%s: %v`, ErrModeSInvalidSymbol, s[6])
	} else {
		buf[4] = buf[4]&0b11110000 | ((val & 0x3C) >> 2) // 0b--CCCC-- >> 2 => [4] 0b----CCCC
		buf[5] = (val&0x03)<<6 | 0b00100000              // 0b------CC << 6 => [5] 0bCC------
	}

	if length == 7 || s[7] == 0 {
		return buf, nil
	}

	if val := TableIA5Reversed[s[7]]; val == 0 {
		return buf, fmt.Errorf(`%s: %v`, ErrModeSInvalidSymbol, s[7])
	} else {
		buf[5] = buf[5]&0b11000000 | val&0x3F // 0b--DDDDDD << 0 => [5] 0b--DDDDDD
	}

	return buf, nil
}
