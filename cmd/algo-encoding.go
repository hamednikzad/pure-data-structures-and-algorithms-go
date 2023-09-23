package main

import (
	"fmt"
	"github.com/hamednikzad/pure-data-structures-and-algorithms-go/pkg/algo/encoding/rle_encoder"
)

func rle() {
	testCode := "aaaaaaaaaaaaaaaaaaaaaasdadddfdf fdf"
	code := rle_encoder.Encode(testCode)
	decode := rle_encoder.Decode(code)

	fmt.Println("Run-length encoding")
	fmt.Printf("Original: %s\n", testCode)
	fmt.Printf("Encoded:  %s\n", code)
	fmt.Printf("Decoded:  %s\n", decode)
}
