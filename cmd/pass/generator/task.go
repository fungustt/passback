package generator

type (
	Task struct {
		Config Config
		result chan string
	}

	Config struct {
		Length           int
		SymbolsUppercase bool
		Numbers          bool
		SpecialSymbols   bool
	}
)

func NewTask(config Config) Task {
	return Task{
		Config: config,
		result: make(chan string, 1),
	}
}

func (t *Task) Done() string {
	return <-t.result
}
