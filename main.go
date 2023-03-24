package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		filename := scanner.Text()
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("error opening file %s: %s", filename, err)
		}
		defer file.Close()
		var buffer [4]byte
		var sbuffer [4]byte
		_, err = io.ReadFull(file, buffer[:])
		if err != nil {
			log.Fatalf("error reading file %s: %s", filename, err)
		}
		for i, b := range buffer {
			if unicode.IsPrint(rune(b)) {
				sbuffer[i] = buffer[i]
			} else {
				sbuffer[i] = '.'
			}
		}
		fmt.Printf("%s,%x,%s\n", filename, buffer, sbuffer)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning input: %s", err)
	}
}
