package cli

import (
	"os"
	"log"
	"fmt"
	"strconv"
)

type Args struct {
	Port uint16
}

func get_default_args() Args {
	return Args{
		Port: 8080,
	}
}

func HandleArguments() Args {
	args := get_default_args()
	var pointer = 0
	cli_args := os.Args[1:]
	for pointer < len(cli_args) {
		switch cli_args[pointer] {
		case "--port":
			if pointer++; pointer < len(cli_args) {
				converted, err := strconv.Atoi(cli_args[pointer])
				if err != nil || converted < 0 || converted > 65535 {
					log.Fatal("Port is not an uint16")
				}
				args.Port = uint16(converted)
				pointer++
			}
		case "--help":
			fmt.Println("Usage: collatz-go [OPTIONS]")
			fmt.Println("Options:")
			fmt.Println("\t--port [uint16] port of the server")
			fmt.Println("\t--help          display help message")
			os.Exit(0)
		default:
			log.Fatalf("Unknown option %v", cli_args[pointer])
		}
	}
	return args
}
