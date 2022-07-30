package concurrency_pattern

type T = interface{}

type ThreadPool interface {
	Run()
	AddTask(task func())
}
type threadPool struct {
	workerNumber int
	processTasks chan func()
}

func NewThreadPool(workerNumber int) ThreadPool {
	t := &threadPool{workerNumber: workerNumber, processTasks: make(chan func())}
	return t
}

func (tp *threadPool) Run() {
	for i := 0; i < tp.workerNumber; i++ {
		go func(work int) {
			for processTask := range tp.processTasks {
				processTask()
			}
		}(i + 1)
	}
}
func (tp *threadPool) AddTask(function func()) {
	tp.processTasks <- function
}
