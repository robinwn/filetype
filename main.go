package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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
		_, err = io.ReadFull(file, buffer[:])
		if err != nil {
			log.Fatalf("error reading file %s: %s", filename, err)
		}
        fmt.Printf("%s,%x,%s\n", filename, buffer, buffer)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning input: %s", err)
	}
}
