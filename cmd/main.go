package main

import (
	"crypto/rand"
	"fmt"

	"github.com/JPD-12/sks"
)

func main() {
	// Updated call: sks.NewKey now requires 5 args; passing nil for the last parameter
	key, err := sks.NewKey("testlabel", "testtag", false, false, nil)
	if err != nil {
		fmt.Println("NewKey failed:", err)
		return
	}
	fmt.Println("Key created:", key.Label(), key.Tag())

	digest := make([]byte, 32)
	_, err = rand.Read(digest)
	if err != nil {
		fmt.Println("Failed to generate random data:", err)
		return
	}

	fmt.Println("Test completed - key created and digest generated")
}
