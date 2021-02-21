package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"Gosh/alias"
)

var ALIAS = alias.NewAlias()

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

func parseArgs(input string) []string {
	if strings.Contains(input, "=") {
		return strings.SplitN(input, " ", 2)
	}
	return strings.Split(input, " ")
}

func execInput(input string) error {
	input = os.ExpandEnv(input)
	args := parseArgs(input)

	if args[0] == "cd" {
		err := os.Chdir(args[1])

		return err
	}

	if args[0] == "export" {
		kv := strings.Split(args[1], "=")

		err := os.Setenv(kv[0], kv[1])

		return err
	}

	if args[0] == "unset" {
		key := args[1]
		err := os.Unsetenv(key)

		return err
	}

	if args[0] == "alias" {
		kv := strings.Split(args[1], "=")
		ALIAS.SetAlias(kv[0], kv[1])
		return nil
	}

	aliasCmd, err := ALIAS.ExpendAlias(args[0])

	if err == nil {
		args = strings.Split(strings.ReplaceAll(aliasCmd, "\"", ""), " ")
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	return err
}

func showPrompt() {
	user, _ := user.Current()
	hostName, _ := os.Hostname()
	wd, _ := os.Getwd()

	fmt.Printf("%s@%s %s > ", user.Username, hostName, wd)
}
