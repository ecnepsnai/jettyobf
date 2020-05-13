package jettyobf_test

import (
	"testing"

	"github.com/ecnepsnai/jettyobf"
)

func MustMatch(t *testing.T, expected, result string) {
	if result != expected {
		t.Errorf("Unexpected result.\nExpected: '%s'\nGot:      '%s'\n", expected, result)
	}
}

func TestDeobfuscate(t *testing.T) {
	obf := "OBF:1jn91yte1uvc1z0f1uuu1yt81jk9"
	expected := "hunter2"
	result := jettyobf.Deobfuscate(obf)
	MustMatch(t, expected, result)
}

func TestObfuscate(t *testing.T) {
	plain := "hunter2"
	expected := "OBF:1jn91yte1uvc1z0f1uuu1yt81jk9"
	result := jettyobf.Obfuscate(plain)
	MustMatch(t, expected, result)
}

func TestDeobfuscateUnicode(t *testing.T) {
	obf := "OBF:U12n3U0vk3"
	expected := "ß"
	result := jettyobf.Deobfuscate(obf)
	MustMatch(t, expected, result)
}

func TestObfuscateUnicode(t *testing.T) {
	plain := "ß"
	expected := "OBF:U12n3U0vk3"
	result := jettyobf.Obfuscate(plain)
	MustMatch(t, expected, result)
}
