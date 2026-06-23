package main

import (
	"fmt"
	"miniq/internal/db"
	"os"
	"strings"
)

func main() {
	// check if arg was provided
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		printHelp()
		return
	}
	command := os.Args[1]

	switch command {
	case "help":
		printHelp()
	case "version":
		// do code
		fmt.Println("miniq v0.1.0")
	case "hello":
		fmt.Println("hello from miniq")
	case "echo":
		handleEcho(os.Args[2:])
	case "init":
		handleInit()
	default:
		fmt.Println("unknown command:", command)
		printHelp()
	}
}

// prints the help menu
func printHelp() {
	fmt.Println()
	fmt.Println("miniq - a tiny job queue")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println(" miniq <command>")
	fmt.Println("Commands:")
	fmt.Println(" help - Show this help message")
	fmt.Println(" version - Show the version")
	fmt.Println(" hello - print a test message")
	fmt.Println(" echo - print a message")
	fmt.Println()

}

// handles the formatting of the echo command string
func handleEcho(args []string) {
	message := strings.Join(args, " ")
	fmt.Println(message)
}

// function to initialise thne sqlite database
func handleInit() {
	// open, or create, the sql database file for the jobs
	database, err := db.Open("miniq.db")
	if err != nil {
		fmt.Println("failed to open database:", err)
		return
	}
	defer database.Close() // schedule close database connection
	// run migration to create actual jobs table
	if err := db.Migrate(database); err != nil {
		fmt.Println("failed to migrate database:", err)
		return
	}
	// print success
	fmt.Println("database initialised")

}
