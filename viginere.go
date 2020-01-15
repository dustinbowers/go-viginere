package main

import (
	"fmt"
	"os"
	"strings"
)

var grid = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var grid_len = len(grid)
var ord_start = grid[0]

type Operation int

const (
	ENCRYPT Operation = iota
	DECRYPT           = iota
)

func process(key, phrase string, method Operation) string {
	output := ""

	key = strings.ToUpper(key)
	key = strings.Replace(key, " ", "", -1)
	phrase = strings.ToUpper(phrase)
	phrase = strings.Replace(phrase, " ", "", -1)

	direction := 1
	if method == DECRYPT {
		direction = -1
	}
	key_len := len(key)
	for pos, c := range phrase {
		char_index := int(byte(c) - ord_start)
		key_index := pos % key_len
		key_ord := int(key[key_index] - ord_start)
		grid_offset := (char_index + (key_ord * direction)) % grid_len
		if grid_offset < 0 {
			grid_offset += grid_len
		}
		code := grid[grid_offset]
		output += string(code)
	}

	return output
}

func main() {
	fmt.Println(os.Args)

	if len(os.Args) != 3 {
		fmt.Println("Usage: viginere <key> <message>")
	}

	key := os.Args[1]
	phrase := os.Args[2]
	fmt.Println("Key:", key)
	fmt.Println("Phrase:", phrase)

	enc := process(key, phrase, ENCRYPT)
	fmt.Println("Encrypted:", enc)
	decrypted := process(key, enc, DECRYPT)
	fmt.Println("Derypted:", decrypted)
}
