package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/xerrors"

	"passwordhelper/pkg/passwordhelper"
)

func main() {
	flag.Parse()
	flags := flag.Args()
	if len(flags) < 2 {
		log.Fatal("You need to provide a password and at least one character you require.")
	}
	password := flags[0]
	indexes := flags[1:]
	var i []int
	for _, v := range indexes {
		a, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("index %s cannot be converted to int", v)
		}
		i = append(i, a)
	}
	s := &passwordhelper.Service{}
	res, err := s.CharsAt(password, i...)
	if err != nil {
		switch {
		case xerrors.Is(err, passwordhelper.ErrEmptyPass):
			log.Fatalf("You provided an empty password :(")
		case xerrors.Is(err, passwordhelper.ErrInvalidIndex):
			log.Fatalf("You provided an invalid index for a character. :(")
		default:
			log.Fatal("I failed for an unknown reason...sorry!")
		}
	}
	fmt.Println(*res)
	os.Exit(0)
}
