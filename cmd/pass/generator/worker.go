package generator

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	charsetSymbols          = "qwertyuiopasdfghjklzxcvbnm"
	charsetSymbolsUppercase = "QWERTYUIOPASDFGHJKLZXCVBNM"
	charsetNumbers          = "1234567890"
	charsetSpecialSymbols   = "!@#$%^&*()_+{}[]/|?.,<>;:"
)

type Worker struct {
	rand *rand.Rand

	taskCh <-chan *Task
}

func NewWorker(taskCh <-chan *Task) *Worker {
	g := Worker{
		rand:   rand.New(rand.NewSource(time.Now().UnixNano())),
		taskCh: taskCh,
	}

	return &g
}

func (w *Worker) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case task := <-w.taskCh:
			task.result <- w.generate(task.Config)
		}
	}
}

func (w *Worker) generate(config Config) string {
	var b strings.Builder

	fmt.Fprintf(&b, charsetSymbols)

	if config.SymbolsUppercase {
		fmt.Fprintf(&b, charsetSymbolsUppercase)
	}

	if config.Numbers {
		fmt.Fprintf(&b, charsetNumbers)
	}

	if config.SpecialSymbols {
		fmt.Fprintf(&b, charsetSpecialSymbols)
	}

	return w.stringWithCharset(config.Length, b.String())
}

func (w *Worker) stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[w.rand.Intn(len(charset))]
	}
	return string(b)
}
