package jettyobf_test

import (
	"fmt"

	"github.com/ecnepsnai/jettyobf"
)

func ExampleDeobfuscate() {
	input := "OBF:1jn91yte1uvc1z0f1uuu1yt81jk9"
	result := jettyobf.Deobfuscate(input)
	fmt.Println(result) // will print 'hunter2'
}

func ExampleObfuscate() {
	input := "hunter2"
	result := jettyobf.Obfuscate(input)
	fmt.Println(result) // will print 'OBF:1jn91yte1uvc1z0f1uuu1yt81jk9'
}
