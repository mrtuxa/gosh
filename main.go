package main

import (
	"bufio"
	"fmt"
	"gosh/color"
	"gosh/system"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		system.ShellLine()
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		if args[0] == "cd" {
			// change the working directory
			if len(args) == 1 {
				homeDir, _ := os.UserHomeDir()
				err := os.Chdir(homeDir)
				if err != nil {
					return
				}
			} else {
				err := os.Chdir(args[1])
				if err != nil {
					return
				}
			}
		} else if args[0] == "exit" {
			color.Println(color.Green(), "Goodbye :D")
			os.Exit(0)
		} else {
			cmd := exec.Command(args[0], args[1:]...)
			output, err := cmd.Output()
			if err != nil {
				if exitError, ok := err.(*exec.ExitError); ok {
					if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
						if status.ExitStatus() == 127 {
							fmt.Println("Error: Command not found.")
						}
					}
				}
			} else {
				fmt.Print(string(output))
			}
		}
	}
}
