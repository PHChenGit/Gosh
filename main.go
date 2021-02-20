package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	for {
		showPrompt()

		input, _ := stdin.ReadString('\n')
		input = strings.TrimSpace(input)

		err := execInput(input)

		if err != nil {
			log.Println(err)
		}
	}
}

func execInput(input string) error {
	args := strings.Split(input, " ")

	if args[0] == "cd" {
		err := os.Chdir(args[1])

		return err
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	return err
}

func showPrompt() {
	user, _ := user.Current()
	hostName, _ := os.Hostname()
	wd, _ := os.Getwd()

	fmt.Printf("%s@%s %s > ", user.Username, hostName, wd)
}
