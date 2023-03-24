package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		filename := scanner.Text()
		fileh, err := os.Open(filename)
		if err != nil {
			log.Fatalf("error opening file %s: %s", filename, err)
		}
		defer fileh.Close()
		var buffer [4]byte
		var sbuffer [4]byte
		n := 0
		n, err = io.ReadFull(fileh, buffer[:])
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				// Partial read, pad the buffer with zeros
				for i := n; i < len(buffer); i++ {
					buffer[i] = 0
				}
			} else {
				log.Fatalf("error reading file %s: %s", filename, err)
			}
		}
		for i, b := range buffer {
			if unicode.IsPrint(rune(b)) {
				sbuffer[i] = buffer[i]
			} else {
				sbuffer[i] = '.'
			}
		}
		fileinfo, err := fileh.Stat()
		if err != nil {
			log.Fatalf("error getting file info for %s: %s", filename, err)
		}
		dir, file := filepath.Split(filename)
		dir = filepath.Clean(dir)
		ext := filepath.Ext(file)
		fmt.Printf("%s,%s,%s,%d,%x,%s\n", dir, file[:len(file)-len(ext)], ext, fileinfo.Size(), buffer, sbuffer)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning input: %s", err)
	}
}
