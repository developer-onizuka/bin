package main

import (
	"os"
	"log"
	"io"
)

func main() {
	if len(os.Args) == 1 {
		do_cat(os.Stdin)
	} else {
		for i:=1; i<len(os.Args); i++ {
			fd, err := os.Open(os.Args[i])
			if err != nil {
				log.Fatal(err)
			}
			defer fd.Close()
			do_cat(fd)
		}
	}
}

func do_cat(f *os.File) {
	for {
		n, err := io.Copy(os.Stdout, f)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
	}
}
