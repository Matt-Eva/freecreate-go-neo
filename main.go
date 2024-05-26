package main

import (
	"freecreate/internal/api/routes"
	"fmt"
	"os"
)

func run() error {
	if err := routes.CreateRoutes(); err !=nil{
		return err
	}
	return nil
}

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
