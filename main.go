package main

import (
	"fmt"
	"github.com/pkg/errors"
	DB "http/db"
	"os"
	"strings"
)

func echo(args []string) error {
	if len(args) < 2 {
		return errors.New("no message to echo")
	}
	_, err := fmt.Println(strings.Join(args[1:], " "))
	return err
}

func main() {
	DB.PrintUser()
	if err := echo(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}
