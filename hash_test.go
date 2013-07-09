package main

import (
	"testing"
)

func TestHash(t *testing.T) {
	expectedHex := map[string]string{
		"adler32":   "05e80204",
		"crc32":     "a4b40d18",
		"md4":       "62445769fb62f35ca36acffde98ae3a9",
		"md5":       "7813258ef8c6b632dde8cc80f6bda62f",
		"ripemd160": "58ff8220275e6592716d40fbe8fc19283304dc99",
		"sha1":      "8abf15bef376e0e21f1f9e9c3d74483d5018f3d5",
		"sha256":    "9cca0703342e24806a9f64e08c053dca7f2cd90f10529af8ea872afb0a0c77d4",
	}

	expectedBase64 := map[string]string{
		"md4":       "YkRXafti81yjas/96YrjqQ==",
		"md5":       "eBMljvjGtjLd6MyA9r2mLw==",
		"ripemd160": "WP+CICdeZZJxbUD76PwZKDME3Jk=",
		"sha1":      "ir8VvvN24OIfH56cPXRIPVAY89U=",
		"sha256":    "nMoHAzQuJIBqn2TgjAU9yn8s2Q8QUpr46ocq+woMd9Q=",
	}

	for _, a := range algorithms {
		hexVal, base64Val := calculate(a.hash, []byte("bacon"))

		if expectedVal, ok := expectedHex[a.name]; ok {
			if hexVal != expectedVal {
				t.Logf("   Hex: %s: expected '%s', got '%s'", a.name, expectedVal, hexVal)
				t.Fail()
			}
		}

		if expectedVal, ok := expectedBase64[a.name]; ok {
			if base64Val != expectedVal {
				t.Logf("Base64: %s: expected '%s', got '%s'", a.name, expectedVal, hexVal)
				t.Fail()
			}
		}
	}
}
