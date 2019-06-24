package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gobuffalo/here"
)

func main() {
	pwd, _ := os.Getwd()

	args := os.Args[1:]
	if len(args) == 0 {
		args = append(args, pwd)
	}

	fn := here.Dir
	switch args[0] {
	case "pkg":
		fn = here.Package
		args = args[1:]
		if len(args) == 0 {
			log.Fatalf("you must pass at least one package name")
		}
	}

	for _, a := range args {
		i, err := fn(a)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(os.Stdout, i.String())
	}
}
