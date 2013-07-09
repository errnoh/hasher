// Copyright 2013 errnoh. All rights reserved.
// Use of this source code is governed by
// MIT License that can be found in the LICENSE file.

// Hasher passes stdin to multiple hash functions and outputs checksum from each in hex and base64 format.
package main

import (
	"bytes"
	"code.google.com/p/go.crypto/md4"
	"code.google.com/p/go.crypto/ripemd160"
	"code.google.com/p/go.crypto/sha3"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"io"
	"os"
	"text/tabwriter"
)

var algorithms = []struct {
	name string
	hash hash.Hash
}{
	{"adler32", adler32.New()},
	{"crc32", crc32.NewIEEE()},
	{"crc64", crc64.New(crc64.MakeTable(crc64.ISO))},
	{"keccak224", sha3.NewKeccak224()},
	{"keccak256", sha3.NewKeccak256()},
	{"md4", md4.New()},
	{"md5", md5.New()},
	{"ripemd160", ripemd160.New()},
	{"sha1", sha1.New()},
	{"sha256", sha256.New()},
}

func main() {
	var (
		writer = tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)
		input  = new(bytes.Buffer)
	)
	io.Copy(input, os.Stdin)

	fmt.Fprintf(writer, "hash\thex\tbase64\n")
	for _, a := range algorithms {
		hexVal, base64Val := calculate(a.hash, input.Bytes())
		fmt.Fprintf(writer, "%s\t%s\t%s\n", a.name, hexVal, base64Val)
	}
	writer.Flush()
}

func calculate(h hash.Hash, input []byte) (hexVal, base64Val string) {
	h.Write(input)
	input = h.Sum(nil)
	hexVal = hex.EncodeToString(input)
	base64Val = base64.StdEncoding.EncodeToString(input)
	return
}
