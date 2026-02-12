package main
import (
	"fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        printUsage()
        return
    }
    command := os.Args[1]
    switch command {
		case "list":
			ListTasks()
		case "add":
			if len(os.Args) < 3 {
				fmt.Println("you need not track doing nothing!")
				return
			}
			AddTask(os.Args[2:])
		case "done":
			if len(os.Args) < 3 {
				fmt.Println("woohoo! you did nothing.")
				return
			}
			Mark(os.Args[2], StatusDone)
		case "progress":
			if len(os.Args) < 3 {
				fmt.Println("Please provide task ID.")
				return
			}
			Mark(os.Args[2], StatusInProgress)

		case "clear":
			ClearCompleted()
		default:
			fmt.Println(command, " ,uh?!")
			printUsage()
    }
}

func printUsage() {
    fmt.Println("usage:")
    fmt.Println("  tracker add \"task\"")
    fmt.Println("  tracker list")
    fmt.Println("  tracker done <id>")
    fmt.Println("  tracker progress <id>")
    fmt.Println("  tracker clear")
}
