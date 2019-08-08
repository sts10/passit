// Code generated by running "go generate". DO NOT EDIT.

package password

import "unicode"

// unicodeVersion is the Unicode edition from which the tables are derived.
const unicodeVersion = "10.0.0"

var rangeTableASCII = &unicode.RangeTable{
	R16: []unicode.Range16{
		{Lo: 0x0020, Hi: 0x007e, Stride: 1},
	},
	LatinOffset: 1,
}

var allowedRangeTable = &unicode.RangeTable{
	R16: []unicode.Range16{
		{Lo: 0x0020, Hi: 0x007e, Stride: 1},
		{Lo: 0x00a1, Hi: 0x00a7, Stride: 1},
		{Lo: 0x00a9, Hi: 0x00ac, Stride: 1},
		{Lo: 0x00ae, Hi: 0x00b0, Stride: 2},
		{Lo: 0x00b1, Hi: 0x00b3, Stride: 1},
		{Lo: 0x00b5, Hi: 0x00b7, Stride: 1},
		{Lo: 0x00b9, Hi: 0x02af, Stride: 1},
		{Lo: 0x0370, Hi: 0x0373, Stride: 1},
		{Lo: 0x0376, Hi: 0x0377, Stride: 1},
		{Lo: 0x037b, Hi: 0x037f, Stride: 1},
		{Lo: 0x0386, Hi: 0x038a, Stride: 1},
		{Lo: 0x038c, Hi: 0x038e, Stride: 2},
		{Lo: 0x038f, Hi: 0x03a1, Stride: 1},
		{Lo: 0x03a3, Hi: 0x0482, Stride: 1},
		{Lo: 0x048a, Hi: 0x052f, Stride: 1},
		{Lo: 0x0531, Hi: 0x0556, Stride: 1},
		{Lo: 0x055a, Hi: 0x055f, Stride: 1},
		{Lo: 0x0561, Hi: 0x0587, Stride: 1},
		{Lo: 0x0589, Hi: 0x058a, Stride: 1},
		{Lo: 0x058d, Hi: 0x058f, Stride: 1},
		{Lo: 0x05be, Hi: 0x05c0, Stride: 2},
		{Lo: 0x05c3, Hi: 0x05c6, Stride: 3},
		{Lo: 0x05d0, Hi: 0x05ea, Stride: 1},
		{Lo: 0x05f0, Hi: 0x05f4, Stride: 1},
		{Lo: 0x0606, Hi: 0x060f, Stride: 1},
		{Lo: 0x061b, Hi: 0x061e, Stride: 3},
		{Lo: 0x061f, Hi: 0x063f, Stride: 1},
		{Lo: 0x0641, Hi: 0x064a, Stride: 1},
		{Lo: 0x0660, Hi: 0x066f, Stride: 1},
		{Lo: 0x0671, Hi: 0x06d5, Stride: 1},
		{Lo: 0x06de, Hi: 0x06e9, Stride: 11},
		{Lo: 0x06ee, Hi: 0x070d, Stride: 1},
		{Lo: 0x0710, Hi: 0x0712, Stride: 2},
		{Lo: 0x0713, Hi: 0x072f, Stride: 1},
		{Lo: 0x074d, Hi: 0x07a5, Stride: 1},
		{Lo: 0x07b1, Hi: 0x07c0, Stride: 15},
		{Lo: 0x07c1, Hi: 0x07ea, Stride: 1},
		{Lo: 0x07f6, Hi: 0x07f9, Stride: 1},
		{Lo: 0x0800, Hi: 0x0815, Stride: 1},
		{Lo: 0x0830, Hi: 0x083e, Stride: 1},
		{Lo: 0x0840, Hi: 0x0858, Stride: 1},
		{Lo: 0x085e, Hi: 0x0860, Stride: 2},
		{Lo: 0x0861, Hi: 0x086a, Stride: 1},
		{Lo: 0x08a0, Hi: 0x08b4, Stride: 1},
		{Lo: 0x08b6, Hi: 0x08bd, Stride: 1},
		{Lo: 0x0904, Hi: 0x0939, Stride: 1},
		{Lo: 0x093d, Hi: 0x0950, Stride: 19},
		{Lo: 0x0958, Hi: 0x0961, Stride: 1},
		{Lo: 0x0964, Hi: 0x0970, Stride: 1},
		{Lo: 0x0972, Hi: 0x0980, Stride: 1},
		{Lo: 0x0985, Hi: 0x098c, Stride: 1},
		{Lo: 0x098f, Hi: 0x0990, Stride: 1},
		{Lo: 0x0993, Hi: 0x09a8, Stride: 1},
		{Lo: 0x09aa, Hi: 0x09b0, Stride: 1},
		{Lo: 0x09b2, Hi: 0x09b6, Stride: 4},
		{Lo: 0x09b7, Hi: 0x09b9, Stride: 1},
		{Lo: 0x09bd, Hi: 0x09ce, Stride: 17},
		{Lo: 0x09dc, Hi: 0x09dd, Stride: 1},
		{Lo: 0x09df, Hi: 0x09e1, Stride: 1},
		{Lo: 0x09e6, Hi: 0x09fd, Stride: 1},
		{Lo: 0x0a05, Hi: 0x0a0a, Stride: 1},
		{Lo: 0x0a0f, Hi: 0x0a10, Stride: 1},
		{Lo: 0x0a13, Hi: 0x0a28, Stride: 1},
		{Lo: 0x0a2a, Hi: 0x0a30, Stride: 1},
		{Lo: 0x0a32, Hi: 0x0a33, Stride: 1},
		{Lo: 0x0a35, Hi: 0x0a36, Stride: 1},
		{Lo: 0x0a38, Hi: 0x0a39, Stride: 1},
		{Lo: 0x0a59, Hi: 0x0a5c, Stride: 1},
		{Lo: 0x0a5e, Hi: 0x0a66, Stride: 8},
		{Lo: 0x0a67, Hi: 0x0a6f, Stride: 1},
		{Lo: 0x0a72, Hi: 0x0a74, Stride: 1},
		{Lo: 0x0a85, Hi: 0x0a8d, Stride: 1},
		{Lo: 0x0a8f, Hi: 0x0a91, Stride: 1},
		{Lo: 0x0a93, Hi: 0x0aa8, Stride: 1},
		{Lo: 0x0aaa, Hi: 0x0ab0, Stride: 1},
		{Lo: 0x0ab2, Hi: 0x0ab3, Stride: 1},
		{Lo: 0x0ab5, Hi: 0x0ab9, Stride: 1},
		{Lo: 0x0abd, Hi: 0x0ad0, Stride: 19},
		{Lo: 0x0ae0, Hi: 0x0ae1, Stride: 1},
		{Lo: 0x0ae6, Hi: 0x0af1, Stride: 1},
		{Lo: 0x0af9, Hi: 0x0b05, Stride: 12},
		{Lo: 0x0b06, Hi: 0x0b0c, Stride: 1},
		{Lo: 0x0b0f, Hi: 0x0b10, Stride: 1},
		{Lo: 0x0b13, Hi: 0x0b28, Stride: 1},
		{Lo: 0x0b2a, Hi: 0x0b30, Stride: 1},
		{Lo: 0x0b32, Hi: 0x0b33, Stride: 1},
		{Lo: 0x0b35, Hi: 0x0b39, Stride: 1},
		{Lo: 0x0b3d, Hi: 0x0b5c, Stride: 31},
		{Lo: 0x0b5d, Hi: 0x0b5f, Stride: 2},
		{Lo: 0x0b60, Hi: 0x0b61, Stride: 1},
		{Lo: 0x0b66, Hi: 0x0b77, Stride: 1},
		{Lo: 0x0b83, Hi: 0x0b85, Stride: 2},
		{Lo: 0x0b86, Hi: 0x0b8a, Stride: 1},
		{Lo: 0x0b8e, Hi: 0x0b90, Stride: 1},
		{Lo: 0x0b92, Hi: 0x0b95, Stride: 1},
		{Lo: 0x0b99, Hi: 0x0b9a, Stride: 1},
		{Lo: 0x0b9c, Hi: 0x0b9e, Stride: 2},
		{Lo: 0x0b9f, Hi: 0x0ba3, Stride: 4},
		{Lo: 0x0ba4, Hi: 0x0ba8, Stride: 4},
		{Lo: 0x0ba9, Hi: 0x0baa, Stride: 1},
		{Lo: 0x0bae, Hi: 0x0bb9, Stride: 1},
		{Lo: 0x0bd0, Hi: 0x0be6, Stride: 22},
		{Lo: 0x0be7, Hi: 0x0bfa, Stride: 1},
		{Lo: 0x0c05, Hi: 0x0c0c, Stride: 1},
		{Lo: 0x0c0e, Hi: 0x0c10, Stride: 1},
		{Lo: 0x0c12, Hi: 0x0c28, Stride: 1},
		{Lo: 0x0c2a, Hi: 0x0c39, Stride: 1},
		{Lo: 0x0c3d, Hi: 0x0c58, Stride: 27},
		{Lo: 0x0c59, Hi: 0x0c5a, Stride: 1},
		{Lo: 0x0c60, Hi: 0x0c61, Stride: 1},
		{Lo: 0x0c66, Hi: 0x0c6f, Stride: 1},
		{Lo: 0x0c78, Hi: 0x0c80, Stride: 1},
		{Lo: 0x0c85, Hi: 0x0c8c, Stride: 1},
		{Lo: 0x0c8e, Hi: 0x0c90, Stride: 1},
		{Lo: 0x0c92, Hi: 0x0ca8, Stride: 1},
		{Lo: 0x0caa, Hi: 0x0cb3, Stride: 1},
		{Lo: 0x0cb5, Hi: 0x0cb9, Stride: 1},
		{Lo: 0x0cbd, Hi: 0x0cde, Stride: 33},
		{Lo: 0x0ce0, Hi: 0x0ce1, Stride: 1},
		{Lo: 0x0ce6, Hi: 0x0cef, Stride: 1},
		{Lo: 0x0cf1, Hi: 0x0cf2, Stride: 1},
		{Lo: 0x0d05, Hi: 0x0d0c, Stride: 1},
		{Lo: 0x0d0e, Hi: 0x0d10, Stride: 1},
		{Lo: 0x0d12, Hi: 0x0d3a, Stride: 1},
		{Lo: 0x0d3d, Hi: 0x0d4e, Stride: 17},
		{Lo: 0x0d4f, Hi: 0x0d54, Stride: 5},
		{Lo: 0x0d55, Hi: 0x0d56, Stride: 1},
		{Lo: 0x0d58, Hi: 0x0d61, Stride: 1},
		{Lo: 0x0d66, Hi: 0x0d7f, Stride: 1},
		{Lo: 0x0d85, Hi: 0x0d96, Stride: 1},
		{Lo: 0x0d9a, Hi: 0x0db1, Stride: 1},
		{Lo: 0x0db3, Hi: 0x0dbb, Stride: 1},
		{Lo: 0x0dbd, Hi: 0x0dc0, Stride: 3},
		{Lo: 0x0dc1, Hi: 0x0dc6, Stride: 1},
		{Lo: 0x0de6, Hi: 0x0def, Stride: 1},
		{Lo: 0x0df4, Hi: 0x0e01, Stride: 13},
		{Lo: 0x0e02, Hi: 0x0e30, Stride: 1},
		{Lo: 0x0e32, Hi: 0x0e33, Stride: 1},
		{Lo: 0x0e3f, Hi: 0x0e45, Stride: 1},
		{Lo: 0x0e4f, Hi: 0x0e5b, Stride: 1},
		{Lo: 0x0e81, Hi: 0x0e82, Stride: 1},
		{Lo: 0x0e84, Hi: 0x0e87, Stride: 3},
		{Lo: 0x0e88, Hi: 0x0e8a, Stride: 2},
		{Lo: 0x0e8d, Hi: 0x0e94, Stride: 7},
		{Lo: 0x0e95, Hi: 0x0e97, Stride: 1},
		{Lo: 0x0e99, Hi: 0x0e9f, Stride: 1},
		{Lo: 0x0ea1, Hi: 0x0ea3, Stride: 1},
		{Lo: 0x0ea5, Hi: 0x0ea7, Stride: 2},
		{Lo: 0x0eaa, Hi: 0x0eab, Stride: 1},
		{Lo: 0x0ead, Hi: 0x0eb0, Stride: 1},
		{Lo: 0x0eb2, Hi: 0x0eb3, Stride: 1},
		{Lo: 0x0ebd, Hi: 0x0ec0, Stride: 3},
		{Lo: 0x0ec1, Hi: 0x0ec4, Stride: 1},
		{Lo: 0x0ed0, Hi: 0x0ed9, Stride: 1},
		{Lo: 0x0edc, Hi: 0x0edf, Stride: 1},
		{Lo: 0x0f00, Hi: 0x0f17, Stride: 1},
		{Lo: 0x0f1a, Hi: 0x0f34, Stride: 1},
		{Lo: 0x0f36, Hi: 0x0f3a, Stride: 2},
		{Lo: 0x0f3b, Hi: 0x0f3d, Stride: 1},
		{Lo: 0x0f40, Hi: 0x0f47, Stride: 1},
		{Lo: 0x0f49, Hi: 0x0f6c, Stride: 1},
		{Lo: 0x0f85, Hi: 0x0f88, Stride: 3},
		{Lo: 0x0f89, Hi: 0x0f8c, Stride: 1},
		{Lo: 0x0fbe, Hi: 0x0fc5, Stride: 1},
		{Lo: 0x0fc7, Hi: 0x0fcc, Stride: 1},
		{Lo: 0x0fce, Hi: 0x0fda, Stride: 1},
		{Lo: 0x1000, Hi: 0x102a, Stride: 1},
		{Lo: 0x103f, Hi: 0x1055, Stride: 1},
		{Lo: 0x105a, Hi: 0x105d, Stride: 1},
		{Lo: 0x1061, Hi: 0x1065, Stride: 4},
		{Lo: 0x1066, Hi: 0x106e, Stride: 8},
		{Lo: 0x106f, Hi: 0x1070, Stride: 1},
		{Lo: 0x1075, Hi: 0x1081, Stride: 1},
		{Lo: 0x108e, Hi: 0x1090, Stride: 2},
		{Lo: 0x1091, Hi: 0x1099, Stride: 1},
		{Lo: 0x109e, Hi: 0x10c5, Stride: 1},
		{Lo: 0x10c7, Hi: 0x10cd, Stride: 6},
		{Lo: 0x10d0, Hi: 0x10fb, Stride: 1},
		{Lo: 0x10fd, Hi: 0x1248, Stride: 1},
		{Lo: 0x124a, Hi: 0x124d, Stride: 1},
		{Lo: 0x1250, Hi: 0x1256, Stride: 1},
		{Lo: 0x1258, Hi: 0x125a, Stride: 2},
		{Lo: 0x125b, Hi: 0x125d, Stride: 1},
		{Lo: 0x1260, Hi: 0x1288, Stride: 1},
		{Lo: 0x128a, Hi: 0x128d, Stride: 1},
		{Lo: 0x1290, Hi: 0x12b0, Stride: 1},
		{Lo: 0x12b2, Hi: 0x12b5, Stride: 1},
		{Lo: 0x12b8, Hi: 0x12be, Stride: 1},
		{Lo: 0x12c0, Hi: 0x12c2, Stride: 2},
		{Lo: 0x12c3, Hi: 0x12c5, Stride: 1},
		{Lo: 0x12c8, Hi: 0x12d6, Stride: 1},
		{Lo: 0x12d8, Hi: 0x1310, Stride: 1},
		{Lo: 0x1312, Hi: 0x1315, Stride: 1},
		{Lo: 0x1318, Hi: 0x135a, Stride: 1},
		{Lo: 0x1360, Hi: 0x137c, Stride: 1},
		{Lo: 0x1380, Hi: 0x1399, Stride: 1},
		{Lo: 0x13a0, Hi: 0x13f5, Stride: 1},
		{Lo: 0x13f8, Hi: 0x13fd, Stride: 1},
		{Lo: 0x1400, Hi: 0x167f, Stride: 1},
		{Lo: 0x1681, Hi: 0x169c, Stride: 1},
		{Lo: 0x16a0, Hi: 0x16f8, Stride: 1},
		{Lo: 0x1700, Hi: 0x170c, Stride: 1},
		{Lo: 0x170e, Hi: 0x1711, Stride: 1},
		{Lo: 0x1720, Hi: 0x1731, Stride: 1},
		{Lo: 0x1735, Hi: 0x1736, Stride: 1},
		{Lo: 0x1740, Hi: 0x1751, Stride: 1},
		{Lo: 0x1760, Hi: 0x176c, Stride: 1},
		{Lo: 0x176e, Hi: 0x1770, Stride: 1},
		{Lo: 0x1780, Hi: 0x17b3, Stride: 1},
		{Lo: 0x17d4, Hi: 0x17d6, Stride: 1},
		{Lo: 0x17d8, Hi: 0x17dc, Stride: 1},
		{Lo: 0x17e0, Hi: 0x17e9, Stride: 1},
		{Lo: 0x17f0, Hi: 0x17f9, Stride: 1},
		{Lo: 0x1800, Hi: 0x180a, Stride: 1},
		{Lo: 0x1810, Hi: 0x1819, Stride: 1},
		{Lo: 0x1820, Hi: 0x1842, Stride: 1},
		{Lo: 0x1844, Hi: 0x1877, Stride: 1},
		{Lo: 0x1880, Hi: 0x1884, Stride: 1},
		{Lo: 0x1887, Hi: 0x18a8, Stride: 1},
		{Lo: 0x18aa, Hi: 0x18b0, Stride: 6},
		{Lo: 0x18b1, Hi: 0x18f5, Stride: 1},
		{Lo: 0x1900, Hi: 0x191e, Stride: 1},
		{Lo: 0x1940, Hi: 0x1944, Stride: 4},
		{Lo: 0x1945, Hi: 0x196d, Stride: 1},
		{Lo: 0x1970, Hi: 0x1974, Stride: 1},
		{Lo: 0x1980, Hi: 0x19ab, Stride: 1},
		{Lo: 0x19b0, Hi: 0x19c9, Stride: 1},
		{Lo: 0x19d0, Hi: 0x19da, Stride: 1},
		{Lo: 0x19de, Hi: 0x1a16, Stride: 1},
		{Lo: 0x1a1e, Hi: 0x1a54, Stride: 1},
		{Lo: 0x1a80, Hi: 0x1a89, Stride: 1},
		{Lo: 0x1a90, Hi: 0x1a99, Stride: 1},
		{Lo: 0x1aa0, Hi: 0x1aa6, Stride: 1},
		{Lo: 0x1aa8, Hi: 0x1aad, Stride: 1},
		{Lo: 0x1b05, Hi: 0x1b33, Stride: 1},
		{Lo: 0x1b45, Hi: 0x1b4b, Stride: 1},
		{Lo: 0x1b50, Hi: 0x1b6a, Stride: 1},
		{Lo: 0x1b74, Hi: 0x1b7c, Stride: 1},
		{Lo: 0x1b83, Hi: 0x1ba0, Stride: 1},
		{Lo: 0x1bae, Hi: 0x1be5, Stride: 1},
		{Lo: 0x1bfc, Hi: 0x1c23, Stride: 1},
		{Lo: 0x1c3b, Hi: 0x1c49, Stride: 1},
		{Lo: 0x1c4d, Hi: 0x1c77, Stride: 1},
		{Lo: 0x1c7e, Hi: 0x1c88, Stride: 1},
		{Lo: 0x1cc0, Hi: 0x1cc7, Stride: 1},
		{Lo: 0x1cd3, Hi: 0x1ce9, Stride: 22},
		{Lo: 0x1cea, Hi: 0x1cec, Stride: 1},
		{Lo: 0x1cee, Hi: 0x1cf1, Stride: 1},
		{Lo: 0x1cf5, Hi: 0x1cf6, Stride: 1},
		{Lo: 0x1d00, Hi: 0x1d2b, Stride: 1},
		{Lo: 0x1d6b, Hi: 0x1d77, Stride: 1},
		{Lo: 0x1d79, Hi: 0x1d9a, Stride: 1},
		{Lo: 0x1e00, Hi: 0x1f15, Stride: 1},
		{Lo: 0x1f18, Hi: 0x1f1d, Stride: 1},
		{Lo: 0x1f20, Hi: 0x1f45, Stride: 1},
		{Lo: 0x1f48, Hi: 0x1f4d, Stride: 1},
		{Lo: 0x1f50, Hi: 0x1f57, Stride: 1},
		{Lo: 0x1f59, Hi: 0x1f5f, Stride: 2},
		{Lo: 0x1f60, Hi: 0x1f7d, Stride: 1},
		{Lo: 0x1f80, Hi: 0x1fb4, Stride: 1},
		{Lo: 0x1fb6, Hi: 0x1fbc, Stride: 1},
		{Lo: 0x1fbe, Hi: 0x1fc2, Stride: 4},
		{Lo: 0x1fc3, Hi: 0x1fc4, Stride: 1},
		{Lo: 0x1fc6, Hi: 0x1fcc, Stride: 1},
		{Lo: 0x1fd0, Hi: 0x1fd3, Stride: 1},
		{Lo: 0x1fd6, Hi: 0x1fdb, Stride: 1},
		{Lo: 0x1fe0, Hi: 0x1fec, Stride: 1},
		{Lo: 0x1ff2, Hi: 0x1ff4, Stride: 1},
		{Lo: 0x1ff6, Hi: 0x1ffc, Stride: 1},
		{Lo: 0x2010, Hi: 0x2027, Stride: 1},
		{Lo: 0x2030, Hi: 0x205e, Stride: 1},
		{Lo: 0x2070, Hi: 0x2074, Stride: 4},
		{Lo: 0x2075, Hi: 0x207e, Stride: 1},
		{Lo: 0x2080, Hi: 0x208e, Stride: 1},
		{Lo: 0x20a0, Hi: 0x20bf, Stride: 1},
		{Lo: 0x2100, Hi: 0x218b, Stride: 1},
		{Lo: 0x2190, Hi: 0x2426, Stride: 1},
		{Lo: 0x2440, Hi: 0x244a, Stride: 1},
		{Lo: 0x2460, Hi: 0x2b73, Stride: 1},
		{Lo: 0x2b76, Hi: 0x2b95, Stride: 1},
		{Lo: 0x2b98, Hi: 0x2bb9, Stride: 1},
		{Lo: 0x2bbd, Hi: 0x2bc8, Stride: 1},
		{Lo: 0x2bca, Hi: 0x2bd2, Stride: 1},
		{Lo: 0x2bec, Hi: 0x2bef, Stride: 1},
		{Lo: 0x2c00, Hi: 0x2c2e, Stride: 1},
		{Lo: 0x2c30, Hi: 0x2c5e, Stride: 1},
		{Lo: 0x2c60, Hi: 0x2c7b, Stride: 1},
		{Lo: 0x2c7e, Hi: 0x2cee, Stride: 1},
		{Lo: 0x2cf2, Hi: 0x2cf3, Stride: 1},
		{Lo: 0x2cf9, Hi: 0x2d25, Stride: 1},
		{Lo: 0x2d27, Hi: 0x2d2d, Stride: 6},
		{Lo: 0x2d30, Hi: 0x2d67, Stride: 1},
		{Lo: 0x2d70, Hi: 0x2d80, Stride: 16},
		{Lo: 0x2d81, Hi: 0x2d96, Stride: 1},
		{Lo: 0x2da0, Hi: 0x2da6, Stride: 1},
		{Lo: 0x2da8, Hi: 0x2dae, Stride: 1},
		{Lo: 0x2db0, Hi: 0x2db6, Stride: 1},
		{Lo: 0x2db8, Hi: 0x2dbe, Stride: 1},
		{Lo: 0x2dc0, Hi: 0x2dc6, Stride: 1},
		{Lo: 0x2dc8, Hi: 0x2dce, Stride: 1},
		{Lo: 0x2dd0, Hi: 0x2dd6, Stride: 1},
		{Lo: 0x2dd8, Hi: 0x2dde, Stride: 1},
		{Lo: 0x2e00, Hi: 0x2e2e, Stride: 1},
		{Lo: 0x2e30, Hi: 0x2e49, Stride: 1},
		{Lo: 0x2e80, Hi: 0x2e99, Stride: 1},
		{Lo: 0x2e9b, Hi: 0x2ef3, Stride: 1},
		{Lo: 0x2f00, Hi: 0x2fd5, Stride: 1},
		{Lo: 0x2ff0, Hi: 0x2ffb, Stride: 1},
		{Lo: 0x3001, Hi: 0x3004, Stride: 1},
		{Lo: 0x3006, Hi: 0x3029, Stride: 1},
		{Lo: 0x3030, Hi: 0x3036, Stride: 6},
		{Lo: 0x3037, Hi: 0x303a, Stride: 1},
		{Lo: 0x303c, Hi: 0x303f, Stride: 1},
		{Lo: 0x3041, Hi: 0x3096, Stride: 1},
		{Lo: 0x309f, Hi: 0x30fb, Stride: 1},
		{Lo: 0x30ff, Hi: 0x3105, Stride: 6},
		{Lo: 0x3106, Hi: 0x312e, Stride: 1},
		{Lo: 0x3131, Hi: 0x318e, Stride: 1},
		{Lo: 0x3190, Hi: 0x31ba, Stride: 1},
		{Lo: 0x31c0, Hi: 0x31e3, Stride: 1},
		{Lo: 0x31f0, Hi: 0x321e, Stride: 1},
		{Lo: 0x3220, Hi: 0x32fe, Stride: 1},
		{Lo: 0x3300, Hi: 0x4db5, Stride: 1},
		{Lo: 0x4dc0, Hi: 0x9fea, Stride: 1},
		{Lo: 0xa000, Hi: 0xa014, Stride: 1},
		{Lo: 0xa016, Hi: 0xa48c, Stride: 1},
		{Lo: 0xa490, Hi: 0xa4c6, Stride: 1},
		{Lo: 0xa4d0, Hi: 0xa4f7, Stride: 1},
		{Lo: 0xa4fe, Hi: 0xa60b, Stride: 1},
		{Lo: 0xa60d, Hi: 0xa62b, Stride: 1},
		{Lo: 0xa640, Hi: 0xa66e, Stride: 1},
		{Lo: 0xa673, Hi: 0xa67e, Stride: 11},
		{Lo: 0xa680, Hi: 0xa69b, Stride: 1},
		{Lo: 0xa6a0, Hi: 0xa6ef, Stride: 1},
		{Lo: 0xa6f2, Hi: 0xa6f7, Stride: 1},
		{Lo: 0xa722, Hi: 0xa76f, Stride: 1},
		{Lo: 0xa771, Hi: 0xa787, Stride: 1},
		{Lo: 0xa78b, Hi: 0xa7ae, Stride: 1},
		{Lo: 0xa7b0, Hi: 0xa7b7, Stride: 1},
		{Lo: 0xa7f7, Hi: 0xa7fa, Stride: 3},
		{Lo: 0xa7fb, Hi: 0xa801, Stride: 1},
		{Lo: 0xa803, Hi: 0xa805, Stride: 1},
		{Lo: 0xa807, Hi: 0xa80a, Stride: 1},
		{Lo: 0xa80c, Hi: 0xa822, Stride: 1},
		{Lo: 0xa828, Hi: 0xa82b, Stride: 1},
		{Lo: 0xa830, Hi: 0xa839, Stride: 1},
		{Lo: 0xa840, Hi: 0xa877, Stride: 1},
		{Lo: 0xa882, Hi: 0xa8b3, Stride: 1},
		{Lo: 0xa8ce, Hi: 0xa8d9, Stride: 1},
		{Lo: 0xa8f2, Hi: 0xa8fd, Stride: 1},
		{Lo: 0xa900, Hi: 0xa925, Stride: 1},
		{Lo: 0xa92e, Hi: 0xa946, Stride: 1},
		{Lo: 0xa95f, Hi: 0xa97c, Stride: 1},
		{Lo: 0xa984, Hi: 0xa9b2, Stride: 1},
		{Lo: 0xa9c1, Hi: 0xa9cd, Stride: 1},
		{Lo: 0xa9d0, Hi: 0xa9d9, Stride: 1},
		{Lo: 0xa9de, Hi: 0xa9e4, Stride: 1},
		{Lo: 0xa9e7, Hi: 0xa9fe, Stride: 1},
		{Lo: 0xaa00, Hi: 0xaa28, Stride: 1},
		{Lo: 0xaa40, Hi: 0xaa42, Stride: 1},
		{Lo: 0xaa44, Hi: 0xaa4b, Stride: 1},
		{Lo: 0xaa50, Hi: 0xaa59, Stride: 1},
		{Lo: 0xaa5c, Hi: 0xaa6f, Stride: 1},
		{Lo: 0xaa71, Hi: 0xaa7a, Stride: 1},
		{Lo: 0xaa7e, Hi: 0xaaaf, Stride: 1},
		{Lo: 0xaab1, Hi: 0xaab5, Stride: 4},
		{Lo: 0xaab6, Hi: 0xaab9, Stride: 3},
		{Lo: 0xaaba, Hi: 0xaabd, Stride: 1},
		{Lo: 0xaac0, Hi: 0xaac2, Stride: 2},
		{Lo: 0xaadb, Hi: 0xaadc, Stride: 1},
		{Lo: 0xaade, Hi: 0xaaea, Stride: 1},
		{Lo: 0xaaf0, Hi: 0xaaf2, Stride: 1},
		{Lo: 0xab01, Hi: 0xab06, Stride: 1},
		{Lo: 0xab09, Hi: 0xab0e, Stride: 1},
		{Lo: 0xab11, Hi: 0xab16, Stride: 1},
		{Lo: 0xab20, Hi: 0xab26, Stride: 1},
		{Lo: 0xab28, Hi: 0xab2e, Stride: 1},
		{Lo: 0xab30, Hi: 0xab5a, Stride: 1},
		{Lo: 0xab60, Hi: 0xab65, Stride: 1},
		{Lo: 0xab70, Hi: 0xabe2, Stride: 1},
		{Lo: 0xabeb, Hi: 0xabf0, Stride: 5},
		{Lo: 0xabf1, Hi: 0xabf9, Stride: 1},
		{Lo: 0xac00, Hi: 0xd7a3, Stride: 1},
		{Lo: 0xd7b0, Hi: 0xd7c6, Stride: 1},
		{Lo: 0xd7cb, Hi: 0xd7fb, Stride: 1},
		{Lo: 0xf900, Hi: 0xfa6d, Stride: 1},
		{Lo: 0xfa70, Hi: 0xfad9, Stride: 1},
		{Lo: 0xfb00, Hi: 0xfb06, Stride: 1},
		{Lo: 0xfb13, Hi: 0xfb17, Stride: 1},
		{Lo: 0xfb1d, Hi: 0xfb1f, Stride: 2},
		{Lo: 0xfb20, Hi: 0xfb36, Stride: 1},
		{Lo: 0xfb38, Hi: 0xfb3c, Stride: 1},
		{Lo: 0xfb3e, Hi: 0xfb40, Stride: 2},
		{Lo: 0xfb41, Hi: 0xfb43, Stride: 2},
		{Lo: 0xfb44, Hi: 0xfb46, Stride: 2},
		{Lo: 0xfb47, Hi: 0xfbb1, Stride: 1},
		{Lo: 0xfbd3, Hi: 0xfd3f, Stride: 1},
		{Lo: 0xfd50, Hi: 0xfd8f, Stride: 1},
		{Lo: 0xfd92, Hi: 0xfdc7, Stride: 1},
		{Lo: 0xfdf0, Hi: 0xfdfd, Stride: 1},
		{Lo: 0xfe10, Hi: 0xfe19, Stride: 1},
		{Lo: 0xfe30, Hi: 0xfe52, Stride: 1},
		{Lo: 0xfe54, Hi: 0xfe66, Stride: 1},
		{Lo: 0xfe68, Hi: 0xfe6b, Stride: 1},
		{Lo: 0xfe70, Hi: 0xfe74, Stride: 1},
		{Lo: 0xfe76, Hi: 0xfefc, Stride: 1},
		{Lo: 0xff01, Hi: 0xff3d, Stride: 1},
		{Lo: 0xff3f, Hi: 0xff41, Stride: 2},
		{Lo: 0xff42, Hi: 0xff6f, Stride: 1},
		{Lo: 0xff71, Hi: 0xff9d, Stride: 1},
		{Lo: 0xffa0, Hi: 0xffbe, Stride: 1},
		{Lo: 0xffc2, Hi: 0xffc7, Stride: 1},
		{Lo: 0xffca, Hi: 0xffcf, Stride: 1},
		{Lo: 0xffd2, Hi: 0xffd7, Stride: 1},
		{Lo: 0xffda, Hi: 0xffdc, Stride: 1},
		{Lo: 0xffe0, Hi: 0xffe2, Stride: 1},
		{Lo: 0xffe4, Hi: 0xffe6, Stride: 1},
		{Lo: 0xffe8, Hi: 0xffee, Stride: 1},
		{Lo: 0xfffc, Hi: 0xfffd, Stride: 1},
	},
	R32: []unicode.Range32{
		{Lo: 0x10000, Hi: 0x1000b, Stride: 1},
		{Lo: 0x1000d, Hi: 0x10026, Stride: 1},
		{Lo: 0x10028, Hi: 0x1003a, Stride: 1},
		{Lo: 0x1003c, Hi: 0x1003d, Stride: 1},
		{Lo: 0x1003f, Hi: 0x1004d, Stride: 1},
		{Lo: 0x10050, Hi: 0x1005d, Stride: 1},
		{Lo: 0x10080, Hi: 0x100fa, Stride: 1},
		{Lo: 0x10100, Hi: 0x10102, Stride: 1},
		{Lo: 0x10107, Hi: 0x10133, Stride: 1},
		{Lo: 0x10137, Hi: 0x1018e, Stride: 1},
		{Lo: 0x10190, Hi: 0x1019b, Stride: 1},
		{Lo: 0x101a0, Hi: 0x101d0, Stride: 48},
		{Lo: 0x101d1, Hi: 0x101fc, Stride: 1},
		{Lo: 0x10280, Hi: 0x1029c, Stride: 1},
		{Lo: 0x102a0, Hi: 0x102d0, Stride: 1},
		{Lo: 0x102e1, Hi: 0x102fb, Stride: 1},
		{Lo: 0x10300, Hi: 0x10323, Stride: 1},
		{Lo: 0x1032d, Hi: 0x1034a, Stride: 1},
		{Lo: 0x10350, Hi: 0x10375, Stride: 1},
		{Lo: 0x10380, Hi: 0x1039d, Stride: 1},
		{Lo: 0x1039f, Hi: 0x103c3, Stride: 1},
		{Lo: 0x103c8, Hi: 0x103d5, Stride: 1},
		{Lo: 0x10400, Hi: 0x1049d, Stride: 1},
		{Lo: 0x104a0, Hi: 0x104a9, Stride: 1},
		{Lo: 0x104b0, Hi: 0x104d3, Stride: 1},
		{Lo: 0x104d8, Hi: 0x104fb, Stride: 1},
		{Lo: 0x10500, Hi: 0x10527, Stride: 1},
		{Lo: 0x10530, Hi: 0x10563, Stride: 1},
		{Lo: 0x1056f, Hi: 0x10600, Stride: 145},
		{Lo: 0x10601, Hi: 0x10736, Stride: 1},
		{Lo: 0x10740, Hi: 0x10755, Stride: 1},
		{Lo: 0x10760, Hi: 0x10767, Stride: 1},
		{Lo: 0x10800, Hi: 0x10805, Stride: 1},
		{Lo: 0x10808, Hi: 0x1080a, Stride: 2},
		{Lo: 0x1080b, Hi: 0x10835, Stride: 1},
		{Lo: 0x10837, Hi: 0x10838, Stride: 1},
		{Lo: 0x1083c, Hi: 0x1083f, Stride: 3},
		{Lo: 0x10840, Hi: 0x10855, Stride: 1},
		{Lo: 0x10857, Hi: 0x1089e, Stride: 1},
		{Lo: 0x108a7, Hi: 0x108af, Stride: 1},
		{Lo: 0x108e0, Hi: 0x108f2, Stride: 1},
		{Lo: 0x108f4, Hi: 0x108f5, Stride: 1},
		{Lo: 0x108fb, Hi: 0x1091b, Stride: 1},
		{Lo: 0x1091f, Hi: 0x10939, Stride: 1},
		{Lo: 0x1093f, Hi: 0x10980, Stride: 65},
		{Lo: 0x10981, Hi: 0x109b7, Stride: 1},
		{Lo: 0x109bc, Hi: 0x109cf, Stride: 1},
		{Lo: 0x109d2, Hi: 0x10a00, Stride: 1},
		{Lo: 0x10a10, Hi: 0x10a13, Stride: 1},
		{Lo: 0x10a15, Hi: 0x10a17, Stride: 1},
		{Lo: 0x10a19, Hi: 0x10a33, Stride: 1},
		{Lo: 0x10a40, Hi: 0x10a47, Stride: 1},
		{Lo: 0x10a50, Hi: 0x10a58, Stride: 1},
		{Lo: 0x10a60, Hi: 0x10a9f, Stride: 1},
		{Lo: 0x10ac0, Hi: 0x10ae4, Stride: 1},
		{Lo: 0x10aeb, Hi: 0x10af6, Stride: 1},
		{Lo: 0x10b00, Hi: 0x10b35, Stride: 1},
		{Lo: 0x10b39, Hi: 0x10b55, Stride: 1},
		{Lo: 0x10b58, Hi: 0x10b72, Stride: 1},
		{Lo: 0x10b78, Hi: 0x10b91, Stride: 1},
		{Lo: 0x10b99, Hi: 0x10b9c, Stride: 1},
		{Lo: 0x10ba9, Hi: 0x10baf, Stride: 1},
		{Lo: 0x10c00, Hi: 0x10c48, Stride: 1},
		{Lo: 0x10c80, Hi: 0x10cb2, Stride: 1},
		{Lo: 0x10cc0, Hi: 0x10cf2, Stride: 1},
		{Lo: 0x10cfa, Hi: 0x10cff, Stride: 1},
		{Lo: 0x10e60, Hi: 0x10e7e, Stride: 1},
		{Lo: 0x11003, Hi: 0x11037, Stride: 1},
		{Lo: 0x11047, Hi: 0x1104d, Stride: 1},
		{Lo: 0x11052, Hi: 0x1106f, Stride: 1},
		{Lo: 0x11083, Hi: 0x110af, Stride: 1},
		{Lo: 0x110bb, Hi: 0x110bc, Stride: 1},
		{Lo: 0x110be, Hi: 0x110c1, Stride: 1},
		{Lo: 0x110d0, Hi: 0x110e8, Stride: 1},
		{Lo: 0x110f0, Hi: 0x110f9, Stride: 1},
		{Lo: 0x11103, Hi: 0x11126, Stride: 1},
		{Lo: 0x11136, Hi: 0x11143, Stride: 1},
		{Lo: 0x11150, Hi: 0x11172, Stride: 1},
		{Lo: 0x11174, Hi: 0x11176, Stride: 1},
		{Lo: 0x11183, Hi: 0x111b2, Stride: 1},
		{Lo: 0x111c1, Hi: 0x111c9, Stride: 1},
		{Lo: 0x111cd, Hi: 0x111d0, Stride: 3},
		{Lo: 0x111d1, Hi: 0x111df, Stride: 1},
		{Lo: 0x111e1, Hi: 0x111f4, Stride: 1},
		{Lo: 0x11200, Hi: 0x11211, Stride: 1},
		{Lo: 0x11213, Hi: 0x1122b, Stride: 1},
		{Lo: 0x11238, Hi: 0x1123d, Stride: 1},
		{Lo: 0x11280, Hi: 0x11286, Stride: 1},
		{Lo: 0x11288, Hi: 0x1128a, Stride: 2},
		{Lo: 0x1128b, Hi: 0x1128d, Stride: 1},
		{Lo: 0x1128f, Hi: 0x1129d, Stride: 1},
		{Lo: 0x1129f, Hi: 0x112a9, Stride: 1},
		{Lo: 0x112b0, Hi: 0x112de, Stride: 1},
		{Lo: 0x112f0, Hi: 0x112f9, Stride: 1},
		{Lo: 0x11305, Hi: 0x1130c, Stride: 1},
		{Lo: 0x1130f, Hi: 0x11310, Stride: 1},
		{Lo: 0x11313, Hi: 0x11328, Stride: 1},
		{Lo: 0x1132a, Hi: 0x11330, Stride: 1},
		{Lo: 0x11332, Hi: 0x11333, Stride: 1},
		{Lo: 0x11335, Hi: 0x11339, Stride: 1},
		{Lo: 0x1133d, Hi: 0x11350, Stride: 19},
		{Lo: 0x1135d, Hi: 0x11361, Stride: 1},
		{Lo: 0x11400, Hi: 0x11434, Stride: 1},
		{Lo: 0x11447, Hi: 0x11459, Stride: 1},
		{Lo: 0x1145b, Hi: 0x1145d, Stride: 2},
		{Lo: 0x11480, Hi: 0x114af, Stride: 1},
		{Lo: 0x114c4, Hi: 0x114c7, Stride: 1},
		{Lo: 0x114d0, Hi: 0x114d9, Stride: 1},
		{Lo: 0x11580, Hi: 0x115ae, Stride: 1},
		{Lo: 0x115c1, Hi: 0x115db, Stride: 1},
		{Lo: 0x11600, Hi: 0x1162f, Stride: 1},
		{Lo: 0x11641, Hi: 0x11644, Stride: 1},
		{Lo: 0x11650, Hi: 0x11659, Stride: 1},
		{Lo: 0x11660, Hi: 0x1166c, Stride: 1},
		{Lo: 0x11680, Hi: 0x116aa, Stride: 1},
		{Lo: 0x116c0, Hi: 0x116c9, Stride: 1},
		{Lo: 0x11700, Hi: 0x11719, Stride: 1},
		{Lo: 0x11730, Hi: 0x1173f, Stride: 1},
		{Lo: 0x118a0, Hi: 0x118f2, Stride: 1},
		{Lo: 0x118ff, Hi: 0x11a00, Stride: 257},
		{Lo: 0x11a0b, Hi: 0x11a32, Stride: 1},
		{Lo: 0x11a3a, Hi: 0x11a3f, Stride: 5},
		{Lo: 0x11a40, Hi: 0x11a46, Stride: 1},
		{Lo: 0x11a50, Hi: 0x11a5c, Stride: 12},
		{Lo: 0x11a5d, Hi: 0x11a83, Stride: 1},
		{Lo: 0x11a86, Hi: 0x11a89, Stride: 1},
		{Lo: 0x11a9a, Hi: 0x11a9c, Stride: 1},
		{Lo: 0x11a9e, Hi: 0x11aa2, Stride: 1},
		{Lo: 0x11ac0, Hi: 0x11af8, Stride: 1},
		{Lo: 0x11c00, Hi: 0x11c08, Stride: 1},
		{Lo: 0x11c0a, Hi: 0x11c2e, Stride: 1},
		{Lo: 0x11c40, Hi: 0x11c45, Stride: 1},
		{Lo: 0x11c50, Hi: 0x11c6c, Stride: 1},
		{Lo: 0x11c70, Hi: 0x11c8f, Stride: 1},
		{Lo: 0x11d00, Hi: 0x11d06, Stride: 1},
		{Lo: 0x11d08, Hi: 0x11d09, Stride: 1},
		{Lo: 0x11d0b, Hi: 0x11d30, Stride: 1},
		{Lo: 0x11d46, Hi: 0x11d50, Stride: 10},
		{Lo: 0x11d51, Hi: 0x11d59, Stride: 1},
		{Lo: 0x12000, Hi: 0x12399, Stride: 1},
		{Lo: 0x12400, Hi: 0x1246e, Stride: 1},
		{Lo: 0x12470, Hi: 0x12474, Stride: 1},
		{Lo: 0x12480, Hi: 0x12543, Stride: 1},
		{Lo: 0x13000, Hi: 0x1342e, Stride: 1},
		{Lo: 0x14400, Hi: 0x14646, Stride: 1},
		{Lo: 0x16800, Hi: 0x16a38, Stride: 1},
		{Lo: 0x16a40, Hi: 0x16a5e, Stride: 1},
		{Lo: 0x16a60, Hi: 0x16a69, Stride: 1},
		{Lo: 0x16a6e, Hi: 0x16a6f, Stride: 1},
		{Lo: 0x16ad0, Hi: 0x16aed, Stride: 1},
		{Lo: 0x16af5, Hi: 0x16b00, Stride: 11},
		{Lo: 0x16b01, Hi: 0x16b2f, Stride: 1},
		{Lo: 0x16b37, Hi: 0x16b3f, Stride: 1},
		{Lo: 0x16b44, Hi: 0x16b45, Stride: 1},
		{Lo: 0x16b50, Hi: 0x16b59, Stride: 1},
		{Lo: 0x16b5b, Hi: 0x16b61, Stride: 1},
		{Lo: 0x16b63, Hi: 0x16b77, Stride: 1},
		{Lo: 0x16b7d, Hi: 0x16b8f, Stride: 1},
		{Lo: 0x16f00, Hi: 0x16f44, Stride: 1},
		{Lo: 0x16f50, Hi: 0x17000, Stride: 176},
		{Lo: 0x17001, Hi: 0x187ec, Stride: 1},
		{Lo: 0x18800, Hi: 0x18af2, Stride: 1},
		{Lo: 0x1b000, Hi: 0x1b11e, Stride: 1},
		{Lo: 0x1b170, Hi: 0x1b2fb, Stride: 1},
		{Lo: 0x1bc00, Hi: 0x1bc6a, Stride: 1},
		{Lo: 0x1bc70, Hi: 0x1bc7c, Stride: 1},
		{Lo: 0x1bc80, Hi: 0x1bc88, Stride: 1},
		{Lo: 0x1bc90, Hi: 0x1bc99, Stride: 1},
		{Lo: 0x1bc9c, Hi: 0x1bc9f, Stride: 3},
		{Lo: 0x1d000, Hi: 0x1d0f5, Stride: 1},
		{Lo: 0x1d100, Hi: 0x1d126, Stride: 1},
		{Lo: 0x1d129, Hi: 0x1d164, Stride: 1},
		{Lo: 0x1d16a, Hi: 0x1d16c, Stride: 1},
		{Lo: 0x1d183, Hi: 0x1d184, Stride: 1},
		{Lo: 0x1d18c, Hi: 0x1d1a9, Stride: 1},
		{Lo: 0x1d1ae, Hi: 0x1d1e8, Stride: 1},
		{Lo: 0x1d200, Hi: 0x1d241, Stride: 1},
		{Lo: 0x1d245, Hi: 0x1d300, Stride: 187},
		{Lo: 0x1d301, Hi: 0x1d356, Stride: 1},
		{Lo: 0x1d360, Hi: 0x1d371, Stride: 1},
		{Lo: 0x1d400, Hi: 0x1d454, Stride: 1},
		{Lo: 0x1d456, Hi: 0x1d49c, Stride: 1},
		{Lo: 0x1d49e, Hi: 0x1d49f, Stride: 1},
		{Lo: 0x1d4a2, Hi: 0x1d4a5, Stride: 3},
		{Lo: 0x1d4a6, Hi: 0x1d4a9, Stride: 3},
		{Lo: 0x1d4aa, Hi: 0x1d4ac, Stride: 1},
		{Lo: 0x1d4ae, Hi: 0x1d4b9, Stride: 1},
		{Lo: 0x1d4bb, Hi: 0x1d4bd, Stride: 2},
		{Lo: 0x1d4be, Hi: 0x1d4c3, Stride: 1},
		{Lo: 0x1d4c5, Hi: 0x1d505, Stride: 1},
		{Lo: 0x1d507, Hi: 0x1d50a, Stride: 1},
		{Lo: 0x1d50d, Hi: 0x1d514, Stride: 1},
		{Lo: 0x1d516, Hi: 0x1d51c, Stride: 1},
		{Lo: 0x1d51e, Hi: 0x1d539, Stride: 1},
		{Lo: 0x1d53b, Hi: 0x1d53e, Stride: 1},
		{Lo: 0x1d540, Hi: 0x1d544, Stride: 1},
		{Lo: 0x1d546, Hi: 0x1d54a, Stride: 4},
		{Lo: 0x1d54b, Hi: 0x1d550, Stride: 1},
		{Lo: 0x1d552, Hi: 0x1d6a5, Stride: 1},
		{Lo: 0x1d6a8, Hi: 0x1d7cb, Stride: 1},
		{Lo: 0x1d7ce, Hi: 0x1d9ff, Stride: 1},
		{Lo: 0x1da37, Hi: 0x1da3a, Stride: 1},
		{Lo: 0x1da6d, Hi: 0x1da74, Stride: 1},
		{Lo: 0x1da76, Hi: 0x1da83, Stride: 1},
		{Lo: 0x1da85, Hi: 0x1da8b, Stride: 1},
		{Lo: 0x1e800, Hi: 0x1e8c4, Stride: 1},
		{Lo: 0x1e8c7, Hi: 0x1e8cf, Stride: 1},
		{Lo: 0x1e900, Hi: 0x1e943, Stride: 1},
		{Lo: 0x1e950, Hi: 0x1e959, Stride: 1},
		{Lo: 0x1e95e, Hi: 0x1e95f, Stride: 1},
		{Lo: 0x1ee00, Hi: 0x1ee03, Stride: 1},
		{Lo: 0x1ee05, Hi: 0x1ee1f, Stride: 1},
		{Lo: 0x1ee21, Hi: 0x1ee22, Stride: 1},
		{Lo: 0x1ee24, Hi: 0x1ee27, Stride: 3},
		{Lo: 0x1ee29, Hi: 0x1ee32, Stride: 1},
		{Lo: 0x1ee34, Hi: 0x1ee37, Stride: 1},
		{Lo: 0x1ee39, Hi: 0x1ee3b, Stride: 2},
		{Lo: 0x1ee42, Hi: 0x1ee47, Stride: 5},
		{Lo: 0x1ee49, Hi: 0x1ee4d, Stride: 2},
		{Lo: 0x1ee4e, Hi: 0x1ee4f, Stride: 1},
		{Lo: 0x1ee51, Hi: 0x1ee52, Stride: 1},
		{Lo: 0x1ee54, Hi: 0x1ee57, Stride: 3},
		{Lo: 0x1ee59, Hi: 0x1ee61, Stride: 2},
		{Lo: 0x1ee62, Hi: 0x1ee64, Stride: 2},
		{Lo: 0x1ee67, Hi: 0x1ee6a, Stride: 1},
		{Lo: 0x1ee6c, Hi: 0x1ee72, Stride: 1},
		{Lo: 0x1ee74, Hi: 0x1ee77, Stride: 1},
		{Lo: 0x1ee79, Hi: 0x1ee7c, Stride: 1},
		{Lo: 0x1ee7e, Hi: 0x1ee80, Stride: 2},
		{Lo: 0x1ee81, Hi: 0x1ee89, Stride: 1},
		{Lo: 0x1ee8b, Hi: 0x1ee9b, Stride: 1},
		{Lo: 0x1eea1, Hi: 0x1eea3, Stride: 1},
		{Lo: 0x1eea5, Hi: 0x1eea9, Stride: 1},
		{Lo: 0x1eeab, Hi: 0x1eebb, Stride: 1},
		{Lo: 0x1eef0, Hi: 0x1eef1, Stride: 1},
		{Lo: 0x1f000, Hi: 0x1f02b, Stride: 1},
		{Lo: 0x1f030, Hi: 0x1f093, Stride: 1},
		{Lo: 0x1f0a0, Hi: 0x1f0ae, Stride: 1},
		{Lo: 0x1f0b1, Hi: 0x1f0bf, Stride: 1},
		{Lo: 0x1f0c1, Hi: 0x1f0cf, Stride: 1},
		{Lo: 0x1f0d1, Hi: 0x1f0f5, Stride: 1},
		{Lo: 0x1f100, Hi: 0x1f10c, Stride: 1},
		{Lo: 0x1f110, Hi: 0x1f12e, Stride: 1},
		{Lo: 0x1f130, Hi: 0x1f16b, Stride: 1},
		{Lo: 0x1f170, Hi: 0x1f1ac, Stride: 1},
		{Lo: 0x1f1e6, Hi: 0x1f202, Stride: 1},
		{Lo: 0x1f210, Hi: 0x1f23b, Stride: 1},
		{Lo: 0x1f240, Hi: 0x1f248, Stride: 1},
		{Lo: 0x1f250, Hi: 0x1f251, Stride: 1},
		{Lo: 0x1f260, Hi: 0x1f265, Stride: 1},
		{Lo: 0x1f300, Hi: 0x1f3fa, Stride: 1},
		{Lo: 0x1f400, Hi: 0x1f6d4, Stride: 1},
		{Lo: 0x1f6e0, Hi: 0x1f6ec, Stride: 1},
		{Lo: 0x1f6f0, Hi: 0x1f6f8, Stride: 1},
		{Lo: 0x1f700, Hi: 0x1f773, Stride: 1},
		{Lo: 0x1f780, Hi: 0x1f7d4, Stride: 1},
		{Lo: 0x1f800, Hi: 0x1f80b, Stride: 1},
		{Lo: 0x1f810, Hi: 0x1f847, Stride: 1},
		{Lo: 0x1f850, Hi: 0x1f859, Stride: 1},
		{Lo: 0x1f860, Hi: 0x1f887, Stride: 1},
		{Lo: 0x1f890, Hi: 0x1f8ad, Stride: 1},
		{Lo: 0x1f900, Hi: 0x1f90b, Stride: 1},
		{Lo: 0x1f910, Hi: 0x1f93e, Stride: 1},
		{Lo: 0x1f940, Hi: 0x1f94c, Stride: 1},
		{Lo: 0x1f950, Hi: 0x1f96b, Stride: 1},
		{Lo: 0x1f980, Hi: 0x1f997, Stride: 1},
		{Lo: 0x1f9c0, Hi: 0x1f9d0, Stride: 16},
		{Lo: 0x1f9d1, Hi: 0x1f9e6, Stride: 1},
		{Lo: 0x20000, Hi: 0x2a6d6, Stride: 1},
		{Lo: 0x2a700, Hi: 0x2b734, Stride: 1},
		{Lo: 0x2b740, Hi: 0x2b81d, Stride: 1},
		{Lo: 0x2b820, Hi: 0x2cea1, Stride: 1},
		{Lo: 0x2ceb0, Hi: 0x2ebe0, Stride: 1},
		{Lo: 0x2f800, Hi: 0x2fa1d, Stride: 1},
	},
	LatinOffset: 6,
}

// Range entries: 420 16-bit, 274 32-bit, 694 total.
// Range bytes: 2520 16-bit, 3288 32-bit, 5808 total.
