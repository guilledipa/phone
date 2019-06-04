// https://github.com/gophercises/phone
package main

import (
	"bytes"
	"unicode"
)

func normalize(phone string) string {
	var buf bytes.Buffer
	for _, ch := range phone {
		if unicode.IsNumber(ch) {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}

func main() {

}
