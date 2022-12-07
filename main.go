package main

import (
	"bufio"
	"github.com/fatih/color"
	"github.com/mrtuxa/gosh/system"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		system.ShellLine()
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		if system.Command("btw", input) == 0 {
			system.Println("I Use arch btw")
		} else if system.Command("exit", input) == 0 {
			color.Green("Goodbye")
			system.Exit()
		} else if system.Command("clear", input) == 0 {
			system.ClearScreen()
		} else {
			_, err := system.ExecCommand(input)
			if err != nil {
				return
			}
			/*
				if system.CheckCommand(input) == "" {
					fmt.Println("gosh: command not found:" + input)
				}
			*/
		}
	}
}
