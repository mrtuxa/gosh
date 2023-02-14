package system

import (
	"fmt"
	"gosh/color"
	"os"
	"strings"
)

func GetUsername() string {
	return os.Getenv("USER")
}

func GetHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

func ShellLine() {
	color.Print(color.Green(), "╭── ")
	color.Print(color.Yellow(), GetUsername())
	color.Print(color.Green(), "@")
	color.Print(color.Blue(), GetHostname())
	color.Print(color.Magenta(), " \uF07C  "+GetCurrentPath())
	color.Print(color.Red(), "]\n")
	fmt.Print(color.Green() + "╰─  " + color.Reset())
}

func GetCurrentPath() string {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
		} else {
			if strings.HasPrefix(currentDir, homeDir) {
				currentDir = "~" + currentDir[len(homeDir):]
			}
		}
	}
	return currentDir
}
