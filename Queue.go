package main

import "fmt"

type Queue struct {
	data []string
}

// push method : Enqueue

func (q *Queue) Enqueue(name string) {
	q.data = append(q.data, name)
}

// pop method : Dequeue :

func (q *Queue) Dequeue() string {
	todelete := q.data[0]
	q.data = q.data[1:]
	return todelete
}

func (q Queue) get(index int) string {
	if q.isEmpty() || index > len(q.data) || index < 0 {
		return ""
	}
	return q.data[index]
}

// is empty method:

func (q Queue) isEmpty() bool {
	return len(q.data) == 0
}

// display method :

func (q Queue) ShowQueue() {
	i := 0
	for i != len(q.data) {
		fmt.Printf(" :=:> %s", q.data[i])
		i++
	}
}
