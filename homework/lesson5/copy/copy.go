package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	objectFrom := flag.String("from", "", "copy from")
	objectTo := flag.String("to", "", "copy to")

	flag.Parse()

	Copy(*objectFrom, *objectTo)

}

func Copy(objectFrom, objectTo string) error {
	fileInfo, err := os.Stat(objectFrom)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if fileInfo.IsDir() {
		log.Fatal("It is directory")
	}

	in, err := os.Open(objectFrom)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(objectTo)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
