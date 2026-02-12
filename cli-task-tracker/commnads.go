package main
import (
	"fmt"
	"strings"
	"strconv"
	"time"
)

func statusSymbol(status string) string {
    switch status {
    case StatusTodo:
        return "[ ]"
    case StatusInProgress:
        return "[~]"
    case StatusDone:
        return "[X]"
    default:
        return "[?]"
    }
}

func ListTasks() {
    tasks, err := LoadTasks()
    if err != nil {
		fmt.Println("error loading tasks: ", err, " :(")
        return
    }
    if len(tasks) == 0 {
        fmt.Println("chill! no tasks.")
        return
    }
    for _, task := range tasks {
        symbol := statusSymbol(task.Status)
        fmt.Printf("%s %d. %s\n",
			symbol,
            task.ID,
            task.Title,
        )
    }
}

func getNextID(tasks []*Task) int {
    maxID := 0
    for _, t := range tasks {
        if t.ID > maxID {
            maxID = t.ID
        }
    }
    return maxID + 1
}

func AddTask(title []string) {
    tasks, err := LoadTasks()
    if err != nil {
        fmt.Println("error loading tasks: ", err, " :(")
        return
    }
	concatTitle := strings.Join(title, " ")
    newID := getNextID(tasks)
    task := NewTask(newID, concatTitle)
    tasks = append(tasks, task)
    err = SaveTasks(tasks)
    if err != nil {
        fmt.Println("error saving tasks: ", err, " :(")
        return
    }
    fmt.Println("task added, get working!")
}

func Mark(idStr string, status string) {
    id, err := strconv.Atoi(idStr)
    if err != nil {
        fmt.Println("get some help!")
        return
    }
    tasks, err := LoadTasks()
    if err != nil {
        fmt.Println("error loading tasks: ", err, " :(")
        return
    }
    found := false
    for _, task := range tasks {
        if task.ID == id {
            if task.Status == status {
                fmt.Println("already marked!")
                return
            }
            task.Status = status
            task.UpdatedAt = time.Now()
            found = true
            break
        }
    }
    if !found {
        fmt.Println("no such task!")
        return
    }
    err = SaveTasks(tasks)
    if err != nil {
        fmt.Println("error saving tasks: ", err, " :(")
        return
    }
    fmt.Println("marked!")
}

func ClearCompleted() {
    tasks, err := LoadTasks()
    if err != nil {
        fmt.Println("error loading tasks: ", err, " :(")
        return
    }

    if len(tasks) == 0 {
        fmt.Println("chill!, no tasks.")
        return
    }

    var filtered []*Task
    removed := 0

    for _, task := range tasks {
        if task.Status != StatusDone {
            filtered = append(filtered, task)
        } else {
            removed++
        }
    }

    if removed == 0 {
        fmt.Println("nothing to clear! :')")
        return
    }
	err = SaveTasks(filtered)
    if err != nil {
        fmt.Println("error saving tasks: ", err, " :(")
        return
    }
    fmt.Printf("cleared %d completed task(s).\n", removed)
}

