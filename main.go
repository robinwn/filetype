package main

import (
	"bufio"
	"encoding/csv"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {

	writer := csv.NewWriter(os.Stdout)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		filename := scanner.Text()
		fileh, err := os.Open(filename)
		if err != nil {
			log.Fatalf("error opening file %s: %s", filename, err)
		}
		var buffer [4]byte
		var sbuffer [4]byte
		n, err := io.ReadFull(fileh, buffer[:])
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
			if b >= 32 && b <= 126 {      // only want to show ASCII printable
				sbuffer[i] = buffer[i]
			} else {
				sbuffer[i] = '.'
			}
		}
		fileinfo, err := fileh.Stat()
		if err != nil {
			log.Fatalf("error getting file info for %s: %s", filename, err)
		}
		err = fileh.Close() // close the file explicitly
		if err != nil {
			log.Fatalf("error closing file %s: %s", filename, err)
		}
		dir, file := filepath.Split(filename)
		dir = filepath.Clean(dir)
		ext := filepath.Ext(file)

		cerr := writer.Write([]string{dir, file[:len(file)-len(ext)], ext, strconv.FormatInt(fileinfo.Size(), 10), hex.EncodeToString(buffer[:]), string(sbuffer[:])})
		if cerr != nil {
			panic(cerr)
		}

	}
	writer.Flush()
	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning input: %s", err)
	}
}
