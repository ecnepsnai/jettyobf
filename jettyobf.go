/*
Package jettyobf is a go implementation of the Jetty Password Obfuscation algorithm
https://www.eclipse.org/jetty/documentation/current/configuring-security-secure-passwords.html

Note: obfuscated passwords have no security benifits. Obfuscated passwords are easily reverted to their plain-text form.
THIS IS NOT A REPLACEMENT FOR PROPER PASSWORD HASHING OR ENCRYPTION.
*/
package jettyobf

import (
	"strconv"
	"strings"
	"unsafe"
)

const obfPrefix = "OBF:"

// Deobfuscate will try to translate the provided obfuscated string into its plain-text representation.
// Returns an empty string if the input cannot be deobfuscated.
func Deobfuscate(in string) string {
	s := in
	if strings.HasPrefix(in, obfPrefix) {
		s = in[4:]
	}

	b := []byte{}

	i := 0
	nibbles := []string{}
	for i < len(s) {
		var nibble string
		// Nibbles that start with U are unicode characters
		if s[i] == 'U' {
			nibble = s[i : i+5]
			i++
		} else {
			nibble = s[i : i+4]
		}
		nibbles = append(nibbles, nibble)
		i += 4
	}

	for _, nibble := range nibbles {
		// Nibbles that start with U are unicode characters
		if nibble[0] == 'U' {
			x := nibble[1:]
			i0, err := strconv.ParseInt(x, 36, 64)
			if err != nil {
				return ""
			}
			bx := byte(i0 >> 8)
			b = append(b, bx)
		} else {
			x := nibble
			i0, err := strconv.ParseInt(x, 36, 64)
			if err != nil {
				return ""
			}
			i1 := (i0 / 256)
			i2 := (i0 % 256)
			bx := byte((i1 + i2 - 254) / 2)
			b = append(b, bx)
		}
	}

	return string(b)
}

// Obfuscate will obfuscate the provided string.
func Obfuscate(in string) string {
	buf := obfPrefix

	// Java stores 'byte' objects as a signed integer (for some stupid reason)
	// so just cast the input string (which is always UTF-8) as signed int's (java bytes)
	input := *(*[]int8)(unsafe.Pointer(&in))

	for i, b1 := range input {
		v := len(input) - (i + 1)
		b2 := input[v]

		b1i := int64(b1)
		b2i := int64(b2)

		if b1 < 0 || b2 < 0 {
			i0 := (0xff&b1i)*256 + (0xff & b2i)
			x := strconv.FormatInt(i0, 36)

			// Pad the nibble with U000
			// indicating it's a unicode character
			length := len(x)
			if length == 4 {
				x = "U" + x
			} else if length == 3 {
				x = "U0" + x
			} else if length == 2 {
				x = "U00" + x
			} else if length == 1 {
				x = "U000" + x
			}

			buf += x
		} else {
			i1 := 127 + b1i + b2i
			i2 := 127 + b1i - b2i
			i0 := i1*256 + i2
			x := strconv.FormatInt(i0, 36)
			buf += x
		}
	}

	return buf
}
