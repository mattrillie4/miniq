package main

import (
	"flag"
	"fmt"
	"miniq/internal/db"
	"miniq/internal/jobs"
	"os"
	"strings"
)

func main() {
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
		fmt.Println("miniq v0.1.0")
	case "push":
		handlePush(os.Args[2:]) // push job onto queue
	case "list":
		handleList() // display job list
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

// handles adding a job to the queue
func handlePush(args []string) {
	// create new flag set for push
	fs := flag.NewFlagSet("push", flag.ExitOnError)

	jobType := fs.String("type", "", "job type")
	payload := fs.String("payload", "", "job payload")

	fs.Parse(args)

	if *jobType == "" {
		fmt.Println("missing required -type")
		return
	}
	if *payload == "" {
		fmt.Println("missing required -payload")
		return
	}

	database, err := db.Open("miniq.db")
	if err != nil {
		fmt.Println("failed to open database:", err)
		return
	}
	defer database.Close()

	if err := jobs.Create(database, *jobType, *payload); err != nil {
		fmt.Println("failed to push job:", err)
		return
	}
	fmt.Println("job pushed")
}

func handleList() {

}

// handles the formatting of the echo command string
func handleEcho(args []string) {
	fs := flag.NewFlagSet("echo", flag.ExitOnError)
	upper := fs.Bool("upper", false, "print message in uppercase")

	fs.Parse(args)
	message := strings.Join(fs.Args(), " ")
	if *upper {
		message = strings.ToUpper(message)
	}
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
