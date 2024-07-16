package main

import (
	"fmt"
	"strings"
	"time"
)

func NewTask(description string, dueDate time.Time) *Task {
    return &Task{
        Description: description,
        DueDate: dueDate,
        CompletedAt: time.Time{},
    }
}

func (t *Task) MarkComplete() {
    t.CompletedAt = time.Now().Local()
}

func (t *Task) IsComplete() bool {
    return !t.CompletedAt.IsZero()
}

func (t *Task) IsDue() bool {
    pin := t.CompletedAt

    if !t.IsComplete() {
        pin = time.Now()
    }
    
    return !t.DueDate.IsZero() && t.DueDate.Before(pin)
}

func NewManager(ts ...*Task) TaskManager {
    return TaskManager{Tasks: ts}
}

func (tm *TaskManager) Add(t *Task) {
    tm.Tasks = append(tm.Tasks, t)
}

func (tm *TaskManager) Remove(index int) {
    if index < len(tm.Tasks) {
        left := tm.Tasks[0:index]
        right := tm.Tasks[index+1:]
        tm.Tasks = append(left, right...)
    }
}

func (tm *TaskManager) RemoveAny(predicate func(*Task) bool) *Task {
    for i, v := range tm.Tasks {
        if predicate(v) {
            t := tm.Tasks[i]
            tm.Remove(i)
            return t
        }
    }

    return nil
}

func (tm *TaskManager) PrintAllTasks(r Renderer[*Task, string], sep string) {
    tasks := make([]string, 0, len(tm.Tasks))

    for _, t := range tm.Tasks {
        tasks = append(tasks, r(t))
    }

    final := strings.Join(tasks, sep)
    fmt.Println(final)
}

func (tm *TaskManager) ClearComplete() {
    for i := len(tm.Tasks)-1; i >= 0; i-- {
        t := tm.Tasks[i]
        if t.IsComplete() {
            tm.Remove(i)
        }
    }
}

func RenderTask(t *Task) string {
    b := strings.Builder{}
    layout := "02/01/2006 15:04 MST"

    if t.IsComplete() {
        b.Write([]byte("[x] "))
    } else {
        b.Write([]byte("[ ] "))
    }
    
    b.Write([]byte(fmt.Sprintf("%s\n    Due by %s", t.Description, t.DueDate.Format(layout))))

    if t.IsComplete() {
        b.Write([]byte(fmt.Sprintf("\n    Completed at %s", t.CompletedAt.Format(layout))))
    }
    
    if t.IsDue() {
        b.Write([]byte(" (past due!)"))
    }
    
    return b.String()
}

func (s inMemoryStorage[T]) Get(key string) (z T, _ bool) {
    if v, ok := s.storage[key]; ok {
        return v, true
    }
    return z, false
}

func (s inMemoryStorage[T]) Put(key string, value T) bool {
    s.storage[key] = value
    return true
}

func NewInMemoryStorage[T any]() Storage[T] {
    return inMemoryStorage[T]{
        storage: make(map[string]T),
    }
}