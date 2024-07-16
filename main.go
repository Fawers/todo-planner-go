package main

import (
	"fmt"
	"time"
)

func MustParse(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}

func main() {
	key := "tasks"
	sep := "\n=================================\n"
	store := NewInMemoryStorage[[]*Task]()
	tasks := []*Task{
		NewTask("Finish this program", MustParse("02/01/2006 @ 15:04Z0700", "14/07/2024 @ 21:30-0300")),
		NewTask("Mark this complete", MustParse("02/01/2006 @ 15:04Z0700", "14/07/2024 @ 21:45-0300")),
		NewTask("Forget this todo", MustParse("02/01/2006 @ 15:04Z0700", "14/07/2024 @ 22:02-0300")),
	}
	tasks[1].CompletedAt = MustParse("20060102T15:04:05Z0700", "20240714T21:34:00-0300")
	if store.Put(key, tasks) {
		fmt.Println("yay put tasks")
	}
	tasks, _ = store.Get(key)
	mg := NewManager(tasks...)
	mg.Tasks[0].MarkComplete()
	mg.PrintAllTasks(RenderTask, sep)
	mg.ClearComplete()
	fmt.Println("And...\n")
	mg.PrintAllTasks(RenderTask, sep)
}
