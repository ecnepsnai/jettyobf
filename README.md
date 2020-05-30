# jettyobf

[![Go Report Card](https://goreportcard.com/badge/github.com/ecnepsnai/jettyobf?style=flat-square)](https://goreportcard.com/report/github.com/ecnepsnai/jettyobf)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/ecnepsnai/jettyobf)
[![Releases](https://img.shields.io/github/release/ecnepsnai/jettyobf/all.svg?style=flat-square)](https://github.com/ecnepsnai/jettyobf/releases)
[![LICENSE](https://img.shields.io/github/license/ecnepsnai/jettyobf.svg?style=flat-square)](https://github.com/ecnepsnai/jettyobf/blob/master/LICENSE)

jettyobf is a go implementation of the [Jetty Password Obfuscation Algorithm](https://www.eclipse.org/jetty/documentation/current/configuring-security-secure-passwords.html).

Obfuscated passwords have no security benifits. Obfuscated passwords are easily reverted to their plain-text form.
**THIS IS NOT A REPLACEMENT FOR PROPER PASSWORD HASHING OR ENCRYPTION.**

# Usage

## Obfuscate

```go
input := "hunter2"
result := jettyobf.Obfuscate(input)
fmt.Println(result) // will print 'OBF:1jn91yte1uvc1z0f1uuu1yt81jk9'
```

## Deobfuscate

```go
input := "OBF:1jn91yte1uvc1z0f1uuu1yt81jk9"
result := jettyobf.Deobfuscate(input)
fmt.Println(result) // will print 'hunter2'
```
