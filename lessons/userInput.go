package lessons

import (
	"bufio"
	"fmt"
	"os"
	"researching-go/pkg/logger"
	"strings"
)

func UserInputExample() {
	standardInput()
}

func standardInput() {
	scanner := bufio.NewScanner(os.Stdin) // bufio - I/O, buffering; os - features of operating system; Stdin - input;
	commands := map[string]string{
		"add":    "you want to add something",
		"delete": "do you want to delete choose items?",
		"update": "your items is update",
		"help":   "list of access commands: add, delete, update, help",
		"exit":   "program is stopped, goodbye!",
	}
	for {
		fmt.Print("type to command :")

		if ok := scanner.Scan(); !ok { // if there is any error
			logger.Ptc("Error is occurred")
			return
		}

		text := scanner.Text() // some text typed user
		logger.Ptc("user command: ", text)

		splitString := strings.Fields(text)

		command := splitString[0]
		var items string
		for index := 1; index < len(splitString); index++ {
			items += splitString[index]
			if index != len(splitString)-1 {
				items += " "
			}
		}

		if items == "" && command != "exit" {
			logger.Ptc("you dont pointed items, try again!")
			return
		}
		_, isExistCommand := commands[command]
		if !isExistCommand {
			logger.Ptc("Unknown command, try again, goodbye!")
			return
		}
		if command == "exit" {
			logger.Ptc(commands[command])
			return
		}

		logger.Ptc("your command: ", command, ".", commands[command], ": ", items, " items")
	}
}
