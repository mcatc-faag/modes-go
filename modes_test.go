// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0
//
// Copyright (c) 2026, runasexe mcatc-faag@runas.name

package modes_test

import (
	"testing"

	"github.com/mcatc-faag/modes-go"
	"github.com/stretchr/testify/require"
)

func TestModeSIdWrite(t *testing.T) {
	require.Equal(t, "", (modes.ModeS{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}).String())
	require.Equal(t, "A", (modes.ModeS{0b00000100, 0x00, 0x00, 0x00, 0x00, 0x00}).String())
	require.Equal(t, "AA", (modes.ModeS{0b00000100, 0b00010000, 0x00, 0x00, 0x00, 0x00}).String())
	require.Equal(t, "AAA", (modes.ModeS{0b00000100, 0b00010000, 0b01000000, 0x00, 0x00, 0x00}).String())
	require.Equal(t, "AAAA", (modes.ModeS{0b00000100, 0b00010000, 0b01000001, 0x00, 0x00, 0x00}).String())
	require.Equal(t, "AAAAA", (modes.ModeS{0b00000100, 0b00010000, 0b01000001, 0b00000100, 0x00, 0x00}).String())
	require.Equal(t, "AAAAAA", (modes.ModeS{0b00000100, 0b00010000, 0b01000001, 0b00000100, 0b00010000, 0x00}).String())
	require.Equal(t, "AAAAAAA", (modes.ModeS{0b00000100, 0b00010000, 0b01000001, 0b00000100, 0b00010000, 0b01000000}).String())
	require.Equal(t, "AAAAAAAA", (modes.ModeS{0b00000100, 0b00010000, 0b01000001, 0b00000100, 0b00010000, 0b01000001}).String())

	require.Equal(t, "TEST", (modes.ModeS{0b01010000, 0b01010100, 0b11010100, 0x00, 0x00, 0x00}).String())
	require.Equal(t, "TESTTEST", (modes.ModeS{0b01010000, 0b01010100, 0b11010100, 0b01010000, 0b01010100, 0b11010100}).String())

	require.Equal(t, "9999", (modes.ModeS{0b11100111, 0b10011110, 0b01111001, 0x00, 0x00, 0x00}).String())
	require.Equal(t, "99999999", (modes.ModeS{0b11100111, 0b10011110, 0b01111001, 0b11100111, 0b10011110, 0b01111001}).String())

	require.Equal(t, "ZZZZ", (modes.ModeS{0b01101001, 0b10100110, 0b10011010, 0x00, 0x00, 0x00}).String())
	require.Equal(t, "ZZZZZZZZ", (modes.ModeS{0b01101001, 0b10100110, 0b10011010, 0b01101001, 0b10100110, 0b10011010}).String())

	require.Equal(t, "BCS4907 ", (modes.ModeS{0x08, 0x34, 0xf4, 0xe7, 0x0d, 0xe0}).String())
}

func TestMakeModeSZero(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("")
		require.NoError(t, err)
		require.Equal(t, modes.Zero, buf)
	})
	t.Run("A", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("A")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000100, 0b00000000, 0b00000000, 0b00000000, 0b00000000, 0b00000000}, buf)
	})
	t.Run("AB", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("AB")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000100, 0b00100000, 0b00000000, 0b00000000, 0b00000000, 0b00000000}, buf)
	})
	t.Run("ABC", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("ABC")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000100, 0b00100000, 0b11000000, 0b00000000, 0b00000000, 0b00000000}, buf)
	})
	t.Run("ABCD", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("ABCD")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000100, 0b00100000, 0b11000100, 0b00000000, 0b00000000, 0b00000000}, buf)
	})
	t.Run("ABCDE", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("ABCDE")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000100, 0b00100000, 0b11000100, 0b00010100, 0b00000000, 0b00000000}, buf)
	})
	t.Run("ABCDEF", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("ABCDEF")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000100, 0b00100000, 0b11000100, 0b00010100, 0b01100000, 0b00000000}, buf)
	})
	t.Run("ABCDEFG", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("ABCDEFG")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000100, 0b00100000, 0b11000100, 0b00010100, 0b01100001, 0b11000000}, buf)
	})
	t.Run("ABCDEFGH", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("ABCDEFGH")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000100, 0b00100000, 0b11000100, 0b00010100, 0b01100001, 0b11001000}, buf)
	})
}
func TestModeSSpace(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		buf, err := modes.MakeModeSSpace("")
		require.NoError(t, err)
		require.Equal(t, modes.Space, buf)
	})
	t.Run("A", func(t *testing.T) {
		buf, err := modes.MakeModeSSpace("A")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000110, 0b00001000, 0b00100000, 0b10000010, 0b00001000, 0b00100000}, buf)
	})
	t.Run("AB", func(t *testing.T) {
		buf, err := modes.MakeModeSSpace("AB")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000100, 0b00101000, 0b00100000, 0b10000010, 0b00001000, 0b00100000}, buf)
	})
	t.Run("ABC", func(t *testing.T) {
		buf, err := modes.MakeModeSSpace("ABC")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000100, 0b00100000, 0b11100000, 0b10000010, 0b00001000, 0b00100000}, buf)
	})
	t.Run("ABCD", func(t *testing.T) {
		buf, err := modes.MakeModeSSpace("ABCD")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000100, 0b00100000, 0b11000100, 0b10000010, 0b00001000, 0b00100000}, buf)
	})
	t.Run("ABCDE", func(t *testing.T) {
		buf, err := modes.MakeModeSSpace("ABCDE")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000100, 0b00100000, 0b11000100, 0b00010110, 0b00001000, 0b00100000}, buf)
	})
	t.Run("ABCDEF", func(t *testing.T) {
		buf, err := modes.MakeModeSSpace("ABCDEF")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000100, 0b00100000, 0b11000100, 0b00010100, 0b01101000, 0b00100000}, buf)
	})
	t.Run("ABCDEFG", func(t *testing.T) {
		buf, err := modes.MakeModeSSpace("ABCDEFG")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000100, 0b00100000, 0b11000100, 0b00010100, 0b01100001, 0b11100000}, buf)
	})
	t.Run("ABCDEFGH", func(t *testing.T) {
		buf, err := modes.MakeModeSSpace("ABCDEFGH")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b00000100, 0b00100000, 0b11000100, 0b00010100, 0b01100001, 0b11001000}, buf)
	})
}

func TestModeSIdWriteTo(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, buf)
	})

	t.Run("space", func(t *testing.T) {
		buf, err := modes.MakeModeSSpace("BCS4907")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0x08, 0x34, 0xf4, 0xe7, 0x0d, 0xe0}, buf)
	})

	t.Run("TEST", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("TEST")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b01010000, 0b01010100, 0b11010100, 0x00, 0x00, 0x00}, buf)
	})

	t.Run("TESTTEST", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("TESTTEST")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b01010000, 0b01010100, 0b11010100, 0b01010000, 0b01010100, 0b11010100}, buf)
	})

	t.Run("9999", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("9999")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b11100111, 0b10011110, 0b01111001, 0x00, 0x00, 0x00}, buf)
	})

	t.Run("99999999", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("99999999")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b11100111, 0b10011110, 0b01111001, 0b11100111, 0b10011110, 0b01111001}, buf)
	})

	t.Run("ZZZZ", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("ZZZZ")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b01101001, 0b10100110, 0b10011010, 0x00, 0x00, 0x00}, buf)
	})

	t.Run("ZZZZZZZZ", func(t *testing.T) {
		buf, err := modes.MakeModeSZero("ZZZZZZZZ")
		require.NoError(t, err)
		require.Equal(t, modes.ModeS{0b01101001, 0b10100110, 0b10011010, 0b01101001, 0b10100110, 0b10011010}, buf)
	})

	t.Run("invalid symbol", func(t *testing.T) {
		_, err := modes.MakeModeSZero("T\x80EST")
		require.Error(t, err)
	})
}
