package generator

import (
	"context"
)

type Generator struct {
	workers []*Worker

	taskCh chan *Task
}

func NewGenerator(ctx context.Context, workersCount int) *Generator {
	g := Generator{
		taskCh: make(chan *Task, 10000),
	}

	workers := make([]*Worker, workersCount)
	for i := 0; i < workersCount; i++ {
		w := NewWorker(g.taskCh)
		go w.Start(ctx)

		workers[i] = w
	}

	g.workers = workers

	return &g
}

func (g *Generator) Generate(task *Task) {
	g.taskCh <- task
}
