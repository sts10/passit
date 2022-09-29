package passit

import (
	"encoding/binary"
	"fmt"
	"io"
	"unicode"
)

const maxUint32 = (1 << 32) - 1

func readBytes(r io.Reader, buf []byte) (int, error) {
	n, err := io.ReadFull(r, buf)
	if err != nil {
		return n, fmt.Errorf("passit: failed to read entropy: %w", err)
	}
	return n, nil
}

func readUint32(r io.Reader) (uint32, error) {
	var buf [4]byte
	_, err := readBytes(r, buf[:])
	return binary.LittleEndian.Uint32(buf[:]), err
}

func readUint32n(r io.Reader, n uint32) (uint32, error) {
	// This is based on golang.org/x/exp/rand:
	// https://github.com/golang/exp/blob/ec7cb31e5a562f5e9e31b300128d2f530f55d127/rand/rand.go#L91-L109.

	// If n is 1, meaning the result will always be 0, avoid reading anything
	// from r and immediately return 0.
	if n == 1 {
		return 0, nil
	}

	v, err := readUint32(r)
	if err != nil {
		return 0, err
	}

	if n&(n-1) == 0 { // n is power of two, can mask
		if n == 0 {
			panic("passit: invalid argument to readUint32n")
		}
		return v & (n - 1), nil
	}

	// If n does not divide v, to avoid bias we must not use
	// a v that is within maxUint32%n of the top of the range.
	if v > maxUint32-n { // Fast check.
		ceiling := maxUint32 - maxUint32%n
		for v >= ceiling {
			v, err = readUint32(r)
			if err != nil {
				return 0, err
			}
		}
	}

	return v % n, nil
}

func countTableRunes(tab *unicode.RangeTable) int {
	var c int
	for _, r16 := range tab.R16 {
		c += int((r16.Hi-r16.Lo)/r16.Stride) + 1
	}
	for _, r32 := range tab.R32 {
		c += int((r32.Hi-r32.Lo)/r32.Stride) + 1
	}

	return c
}

func readRune(r io.Reader, tab *unicode.RangeTable, count int) (rune, error) {
	if int(uint32(count)) != count {
		panic("passit: unicode.RangeTable is too large")
	}

	v, err := readUint32n(r, uint32(count))
	if err != nil {
		return 0, err
	}

	for _, r16 := range tab.R16 {
		size := uint32((r16.Hi-r16.Lo)/r16.Stride) + 1
		if v < size {
			return rune(r16.Lo + uint16(v)*r16.Stride), nil
		}
		v -= size
	}

	for _, r32 := range tab.R32 {
		size := (r32.Hi-r32.Lo)/r32.Stride + 1
		if v < size {
			return rune(r32.Lo + v*r32.Stride), nil
		}
		v -= size
	}

	panic("passit: internal error: unicode.RangeTable did not contain rune")
}
