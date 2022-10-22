package main

import (
	b64 "encoding/base64"
	"fmt"
	"time"
)

func main() {
	// what is this nonsense?! This message is just strange!
	magicalMessage := "SGVsbG8gdGhlcmUhIFlvdSBhcmUgd2VsbCBvbiB5b3VyIHdheSB0byBiZWNvbWluZy" +
		"BhIEdvbGFuZyB3aXphcmQhIEtlZXAgaXQgdXAsIHRoZXJlIGFyZSByZWFsbHkgY29vbCB0aGluZ3MgYWhlYWQ" +
		"hCgpUaGVyZSBpcyBhIGxvdCBnb2luZyBvbiBpbiB0aGlzIGNvZGUsIGJ1dCBkb24ndCB3b3JyeSwgaXQgd2lsbCBhbGwgbWFrZSBzZW5zZSBzb29uLg=="

	// we have to say a spell to learn what it says! VERTO VORTO!!
	message, _ := b64.StdEncoding.DecodeString(magicalMessage)

	// now, let's see if we can understand the message...
	for _, letter := range message {
		time.Sleep(40 * time.Millisecond)
		fmt.Print(string(letter))
	}
}
