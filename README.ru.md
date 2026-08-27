# modes-go — кодирование и декодирование идентификатора воздушных судов Mode S

[![Go Reference](https://pkg.go.dev/badge/github.com/mcatc-faag/modes-go.svg)](https://pkg.go.dev/github.com/mcatc-faag/modes-go)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg)](LICENSE)


[![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg)](LICENSE)

Библиотека `modes-go` предоставляет функции для преобразования позывных воздушных судов между 8-символьной строкой и упакованным 6-байтовым (48-битным) бинарным форматом Mode S.

Реализация соответствует спецификациям ICAO Annex 10 Volume IV (раздел 3.1.2.9.1.2) и EUROCONTROL ASTERIX Category 048 (Item I048/240).

## Описание

В радиолокационных системах наблюдения УВД (вторичная радиолокация Mode S, ADS-B, ASTERIX I048/240) идентификатор воздушного судна состоит из 8 символов. Допустимы заглавные латинские буквы A–Z, цифры 0–9 и концевые пробелы.

Каждый символ кодируется 6-битным кодом подмножества IA-5 (International Alphabet No. 5). Восемь 6-битных символов упаковываются в 6 байт (48 бит).

## Документация

- [English (Английский)](README.md)

## Битовая структура

Массив из 6 байт (`[6]byte`) содержит 8 символов ($C_1$ - $C_8$) в 6 байтах ($B_0$ – $B_5$):

| Байт | Бит 7 | Бит 6 | Бит 5 | Бит 4 | Бит 3 | Бит 2 | Бит 1 | Бит 0 |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| Байт 0 ($B_0$) | $C_1[5]$ | $C_1[4]$ | $C_1[3]$ | $C_1[2]$ | $C_1[1]$ | $C_1[0]$ | $C_2[5]$ | $C_2[4]$ |
| Байт 1 ($B_1$) | $C_2[3]$ | $C_2[2]$ | $C_2[1]$ | $C_2[0]$ | $C_3[5]$ | $C_3[4]$ | $C_3[3]$ | $C_3[2]$ |
| Байт 2 ($B_2$) | $C_3[1]$ | $C_3[0]$ | $C_4[5]$ | $C_4[4]$ | $C_4[3]$ | $C_4[2]$ | $C_4[1]$ | $C_4[0]$ |
| Байт 3 ($B_3$) | $C_5[5]$ | $C_5[4]$ | $C_5[3]$ | $C_5[2]$ | $C_5[1]$ | $C_5[0]$ | $C_6[5]$ | $C_6[4]$ |
| Байт 4 ($B_4$) | $C_6[3]$ | $C_6[2]$ | $C_6[1]$ | $C_6[0]$ | $C_7[5]$ | $C_7[4]$ | $C_7[3]$ | $C_7[2]$ |
| Байт 5 ($B_5$) | $C_7[1]$ | $C_7[0]$ | $C_8[5]$ | $C_8[4]$ | $C_8[3]$ | $C_8[2]$ | $C_8[1]$ | $C_8[0]$ |

## Установка

```bash
go get github.com/mcatc-faag/modes-go
```

## Справочник API

### Типы и переменные

- `type ModeS [6]byte`: представляет упакованный 6-байтовый позывной Mode S.
- `var Zero = ModeS{}`: нулевой массив Mode S.
- `var Space = ModeS{...}`: массив Mode S, заполненный пробелами IA-5.
- `var TableIA5 = [0x40]byte{...}`: таблица соответствия 6-битных кодов IA-5 и символов ASCII.
- `var TableIA5Reversed = [0x100]uint8{...}`: таблица соответствия символов ASCII и кодов IA-5.

### Методы

- `func (s ModeS) String() string`: декодирует буфер Mode S в 8-символьную ASCII-строку.
- `func (s ModeS) SpaceTerminatedString() string`: декодирует буфер Mode S и удаляет концевые пробелы.

### Функции

- `func MakeModeSZero(s string) (ModeS, error)`: кодирует позывной в 6-байтовый буфер Mode S с дополнением неиспользованных символов нулями.
- `func MakeModeSSpace(s string) (ModeS, error)`: кодирует позывной в 6-байтовый буфер Mode S с дополнением неиспользованных символов пробелами.

### Ошибки

- `ErrModeSInvalidSymbol`: возвращается при наличии недопустимых символов вне подмножества IA-5.
- `ErrDataSizeExceeds`: возвращается, если длина строки превышает 8 символов.

## Примеры использования

### Кодирование позывного

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
		log.Fatalf("Ошибка кодирования: %v", err)
	}

	fmt.Printf("Байты: %X\n", buf)
	// Вывод: Байты: 0834F4E70DE0
}
```

### Декодирование позывного

```go
package main

import (
	"fmt"

	"github.com/mcatc-faag/modes-go"
)

func main() {
	raw := modes.ModeS{0x08, 0x34, 0xf4, 0xe7, 0x0d, 0xe0}

	fmt.Printf("Строка:  %q\n", raw.String())
	// Вывод: Строка:  "BCS4907 "

	fmt.Printf("Очищено: %q\n", raw.SpaceTerminatedString())
	// Вывод: Очищено: "BCS4907"
}
```

### Обработка ошибок

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
		fmt.Println("Ошибка: длина позывного не должна превышать 8 символов")
	}

	_, err = modes.MakeModeSSpace("A@B")
	if errors.Is(err, modes.ErrModeSInvalidSymbol) {
		fmt.Println("Ошибка: недопустимый символ в позывном")
	}
}
```

## Тестирование

Запуск тестов:

```bash
go test -v ./...
```

## Юридические уведомления

Репозиторий распространяется на условиях Mozilla Public License версии 2.0 (MPL-2.0). См. [LICENSE](LICENSE).

Юридические уведомления, эксплуатационные ограничения и применимые условия представлены в следующих документах:

- [Legal Notice (English)](docs/legal/LEGAL_NOTICE.en.md)
- [Юридические уведомления (Русский)](docs/legal/LEGAL_NOTICE.ru.md)

## Лицензия

Copyright (c) 2026 runasexe (mcatc-faag@runas.name). Распространяется на условиях [Mozilla Public License 2.0](LICENSE).
