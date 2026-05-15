package main

import (
	"fmt"
	"strings" // I learned it from go docs
	"program/commands"
)


func main() {

	program:
		for {
			var (
				cmd string
			)

			fmt.Print("[SYSTEM]: ")
			fmt.Scan(&cmd)
			if cmd == "exit" {
				fmt.Println("goodbye!")
				break program
			}
			commands.Cmd(strings.ToLower(cmd))
		}

}
