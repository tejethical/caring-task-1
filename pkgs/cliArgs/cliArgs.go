package cliArgs

import (
	"log"
	"os"
)

type Args struct {
	Role     string
	Port     string
	FilePath string
}

func Get() *Args {
	if len(os.Args) < 3 {
		log.Fatal("Insufficient number of arguments")
	}
	args := os.Args[1:]
	if args[0] == "worker" {
		return &Args{Role: args[0], Port: args[1], FilePath: ""}
	}
	return &Args{Role: args[0], FilePath: args[1], Port: ""}
}
