package main

import "time"

type Completable interface {
    MarkComplete()
    IsComplete() bool
}

type Renderer[In any, Out any] func(In) Out

type Task struct {
    Description string
    DueDate, CompletedAt time.Time
}

type TaskManager struct {
    Tasks []*Task
}

type Storage[T any] interface {
    Get(string) (T, bool)
    Put(string, T) bool
}

type inMemoryStorage[T any] struct {
    storage map[string]T
}
