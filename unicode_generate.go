// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The code in this file was taken from golang.org/x/text/internal/export/unicode.

// +build generate

// Unicode table generator.
// Data read from the web.

package main

import (
	"flag"
	"fmt"

	"golang.org/x/text/unicode/rangetable"
)

func main() {
	flag.Parse()
	setupOutput()
	loadChars() // always needed
	printCategories()
	printSizes()
	flushOutput()
}

var output *gen_CodeWriter

func setupOutput() {
	output = gen_NewCodeWriter()
}

func flushOutput() {
	output.WriteGoFile("tables.go", "password")
}

func printf(format string, args ...interface{}) {
	fmt.Fprintf(output, format, args...)
}

func print(args ...interface{}) {
	fmt.Fprint(output, args...)
}

func println(args ...interface{}) {
	fmt.Fprintln(output, args...)
}

type Char struct {
	codePoint rune // if zero, this index is not a valid code point.
	category  string
}

const MaxChar = 0x10FFFF

var chars = make([]Char, MaxChar+1)

func loadChars() {
	ucd_Parse(gen_OpenUCDFile("UnicodeData.txt"), func(p *ucd_Parser) {
		codePoint := p.Rune(0)
		chars[codePoint] = Char{
			codePoint: codePoint,
			category:  p.String(ucd_GeneralCategory),
		}
	})
}

func printCategories() {
	println("import \"unicode\"\n\n")

	println("// unicodeVersion is the Unicode edition from which the tables are derived.")
	printf("const unicodeVersion = %q\n\n", gen_UnicodeVersion())

	// TODO(tmthrgd): Review these ranges.

	dumpRange("rangeTableASCII", func(code rune) bool {
		return code >= 0x20 && code <= 0x7e
	})

	dumpRange("allowedRangeTable", func(code rune) bool {
		if code <= 0x7e { // Special case ASCII.
			return code >= 0x20
		}

		c := chars[code]
		switch c.category {
		case "":
			return false
		case "Lu", "Ll", "Lt", "Lo":
			return true
		case "Sm", "Sc", "So":
			return true
		}
		switch c.category[0] {
		case 'N':
			return true
		case 'P':
			return true
		default:
			return false
		}
	})
}

type Op func(code rune) bool

func dumpRange(name string, inCategory Op) {
	runes := []rune{}
	for i := range chars {
		r := rune(i)
		if inCategory(r) {
			runes = append(runes, r)
		}
	}
	printRangeTable(name, runes)
}

func printRangeTable(name string, runes []rune) {
	rt := rangetable.New(runes...)
	printf("var %s = &unicode.RangeTable{\n", name)
	println("\tR16: []unicode.Range16{")
	for _, r := range rt.R16 {
		printf("\t\t{Lo: %#04x, Hi: %#04x, Stride: %d},\n", r.Lo, r.Hi, r.Stride)
		range16Count++
	}
	println("\t},")
	if len(rt.R32) > 0 {
		println("\tR32: []unicode.Range32{")
		for _, r := range rt.R32 {
			printf("\t\t{Lo: %#x, Hi: %#x, Stride: %d},\n", r.Lo, r.Hi, r.Stride)
			range32Count++
		}
		println("\t},")
	}
	if rt.LatinOffset > 0 {
		printf("\tLatinOffset: %d,\n", rt.LatinOffset)
	}
	printf("}\n\n")
}

var range16Count = 0  // Number of entries in the 16-bit range tables.
var range32Count = 0  // Number of entries in the 32-bit range tables.
var foldPairCount = 0 // Number of fold pairs in the exception tables.

func printSizes() {
	println()
	printf("// Range entries: %d 16-bit, %d 32-bit, %d total.\n", range16Count, range32Count, range16Count+range32Count)
	range16Bytes := range16Count * 3 * 2
	range32Bytes := range32Count * 3 * 4
	printf("// Range bytes: %d 16-bit, %d 32-bit, %d total.\n", range16Bytes, range32Bytes, range16Bytes+range32Bytes)
}
