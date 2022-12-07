package system

import (
	"fmt"
	"github.com/inancgumus/screen"
	"github.com/mrtuxa/gosh/color"
	"os"
	"os/exec"
)

func ClearScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

func ExecCommand(command string) (int, error) {
	cmd := exec.Command(command)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return 0, nil
	}
	return fmt.Println(string(stdout))
}

func CheckCommand(command string) (int, error) {
	cmd := exec.Command("which", command)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return 0, nil
	}
	return fmt.Println(string(stdout))
}

func GetHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

func GetUsername() string {
	return os.Getenv("USER")
}

func ShellLine() {
	color.Print(color.Red(), "[")
	color.Print(color.Yellow(), GetUsername())
	color.Print(color.Green(), "@")
	color.Print(color.Blue(), GetHostname())
	color.Print(color.Red(), "]")
	fmt.Print("$ ")
}

func Exit() {
	os.Exit(1)
}

func Command(command string, input string) int {

	if command == input {
		return 0
	}
	if command < input {
		return -1
	}
	return +1
}

func Println(command string) {
	fmt.Println(command)
}

func Print(command string) {
	fmt.Print(command)
}
