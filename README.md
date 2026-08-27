# modes-go — Mode S Aircraft Identification Encoder & Decoder for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/mcatc-faag/modes-go.svg)](https://pkg.go.dev/github.com/mcatc-faag/modes-go)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg)](LICENSE)

`modes-go` provides functions for packing and unpacking Mode S aircraft callsigns between standard 8-character string representations and 6-byte (48-bit) packed binary formats.

The encoding complies with ICAO Annex 10 Volume IV (Section 3.1.2.9.1.2) and EUROCONTROL ASTERIX Category 048 (Item I048/240).

## Overview

In aviation surveillance systems (Mode S secondary radar, ADS-B, ASTERIX I048/240), aircraft callsigns consist of up to 8 characters. Allowed characters include uppercase letters A–Z, digits 0–9, and trailing space characters for padding.

Each character is encoded using a 6-bit subset of IA-5 (International Alphabet No. 5). Eight 6-bit characters are packed into 6 bytes (48 bits).

## Documentation

- [Russian (Русский)](README.ru.md)

## Bit Layout

The 48-bit buffer (`[6]byte`) stores 8 characters ($C_1$ to $C_8$) across 6 bytes ($B_0$ to $B_5$):

| Byte | Bit 7 | Bit 6 | Bit 5 | Bit 4 | Bit 3 | Bit 2 | Bit 1 | Bit 0 |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| Byte 0 ($B_0$) | $C_1[5]$ | $C_1[4]$ | $C_1[3]$ | $C_1[2]$ | $C_1[1]$ | $C_1[0]$ | $C_2[5]$ | $C_2[4]$ |
| Byte 1 ($B_1$) | $C_2[3]$ | $C_2[2]$ | $C_2[1]$ | $C_2[0]$ | $C_3[5]$ | $C_3[4]$ | $C_3[3]$ | $C_3[2]$ |
| Byte 2 ($B_2$) | $C_3[1]$ | $C_3[0]$ | $C_4[5]$ | $C_4[4]$ | $C_4[3]$ | $C_4[2]$ | $C_4[1]$ | $C_4[0]$ |
| Byte 3 ($B_3$) | $C_5[5]$ | $C_5[4]$ | $C_5[3]$ | $C_5[2]$ | $C_5[1]$ | $C_5[0]$ | $C_6[5]$ | $C_6[4]$ |
| Byte 4 ($B_4$) | $C_6[3]$ | $C_6[2]$ | $C_6[1]$ | $C_6[0]$ | $C_7[5]$ | $C_7[4]$ | $C_7[3]$ | $C_7[2]$ |
| Byte 5 ($B_5$) | $C_7[1]$ | $C_7[0]$ | $C_8[5]$ | $C_8[4]$ | $C_8[3]$ | $C_8[2]$ | $C_8[1]$ | $C_8[0]$ |

## Installation

```bash
go get github.com/mcatc-faag/modes-go
```

## API Reference

### Types and Variables

- `type ModeS [6]byte`: Represents a 6-byte packed Mode S aircraft callsign.
- `var Zero = ModeS{}`: Mode S array initialized to zero bytes.
- `var Space = ModeS{...}`: Mode S array pre-filled with IA-5 space padding.
- `var TableIA5 = [0x40]byte{...}`: Lookup array mapping 6-bit IA-5 values to ASCII bytes.
- `var TableIA5Reversed = [0x100]uint8{...}`: Lookup array mapping ASCII bytes to 6-bit IA-5 values.

### Methods

- `func (s ModeS) String() string`: Decodes the 6-byte Mode S buffer into an 8-character ASCII string.
- `func (s ModeS) SpaceTerminatedString() string`: Decodes the Mode S buffer and removes trailing spaces.

### Functions

- `func MakeModeSZero(s string) (ModeS, error)`: Encodes a callsign string into a 6-byte Mode S buffer, padding remaining characters with zero.
- `func MakeModeSSpace(s string) (ModeS, error)`: Encodes a callsign string into a 6-byte Mode S buffer, padding remaining characters with spaces.

### Errors

- `ErrModeSInvalidSymbol`: Returned when an input string contains characters outside the IA-5 subset.
- `ErrDataSizeExceeds`: Returned when string length exceeds 8 characters.

## Usage Examples

### Encoding a Callsign

```go
package main

import (
	"fmt"
	"log"

	"github.com/mcatc-faag/modes-go"
)

func main() {
	buf, err := modes.MakeModeSSpace("BCS4907")
	if err != nil {
		log.Fatalf("Encoding error: %v", err)
	}

	fmt.Printf("Raw bytes: %X\n", buf)
	// Output: Raw bytes: 0834F4E70DE0
}
```

### Decoding a Callsign

```go
package main

import (
	"fmt"

	"github.com/mcatc-faag/modes-go"
)

func main() {
	raw := modes.ModeS{0x08, 0x34, 0xf4, 0xe7, 0x0d, 0xe0}

	fmt.Printf("String: %q\n", raw.String())
	// Output: String: "BCS4907 "

	fmt.Printf("Clean:  %q\n", raw.SpaceTerminatedString())
	// Output: Clean:  "BCS4907"
}
```

### Error Handling

```go
package main

import (
	"errors"
	"fmt"

	"github.com/mcatc-faag/modes-go"
)

func main() {
	_, err := modes.MakeModeSSpace("CALLSIGN_TOO_LONG")
	if errors.Is(err, modes.ErrDataSizeExceeds) {
		fmt.Println("Error: callsign exceeds 8 characters")
	}

	_, err = modes.MakeModeSSpace("A@B")
	if errors.Is(err, modes.ErrModeSInvalidSymbol) {
		fmt.Println("Error: invalid character in callsign")
	}
}
```

## Testing

Run tests with:

```bash
go test -v ./...
```

## Legal Notices

This repository is distributed under the Mozilla Public License, version 2.0 (MPL-2.0). See [LICENSE](LICENSE).

For legal, safety, certification, and operational notices, refer to the following documents:

- [Legal Notice (English)](docs/legal/LEGAL_NOTICE.en.md)
- [Юридические уведомления (Русский)](docs/legal/LEGAL_NOTICE.ru.md)

## License

Copyright (c) 2026 runasexe (mcatc-faag@runas.name). Distributed under the [Mozilla Public License 2.0](LICENSE).
