package main

import (
	"io"
	"log"

	"github.com/kuentra-official/wal"
)

func main() {
	walOpts := wal.DefaultOptions
	walOpts.DirPath = "tmp"
	kwal, err := wal.Open(walOpts)
	if err != nil {
		panic(err)
	}

	pos1, err := kwal.Write([]byte("first chunk"))
	if err != nil {
		panic(err)
	}
	rval1, err := kwal.Read(pos1)
	if err != nil {
		panic(err)
	}
	log.Println("Is this the expected value? Expect: [first chunk] return Value : [", string(rval1), "]")
	_, err = kwal.Write([]byte("second chunk"))
	if err != nil {
		panic(err)
	}
	_, err = kwal.Write([]byte("third chunk"))
	if err != nil {
		panic(err)
	}
	reader := kwal.NewReader()
	for {
		rv, pos, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		log.Println("[pos] : ", pos, " [return value] : ", string(rv))
		//must print first, second, third chunk
	}
}
