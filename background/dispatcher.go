package background

import "fmt"

type Dispatcher struct {
	*Queue
	WorkerPool chan chan Message
	NumWorkers int
}

func NewDispatcher(queue *Queue, numWorkers int) *Dispatcher {
	return &Dispatcher{queue, make(chan chan Message), numWorkers}
}

func (d *Dispatcher) Run() {
	go func() {
		for i := 1; i <= d.NumWorkers; i++ {
			fmt.Println("Starting Worker", i)
			worker := NewWorker(i)
			worker.Run(d.WorkerPool)
		}

		for {
			select {
			case message := <-d.Queue.Messages:
				go func(message Message) {
					fmt.Printf("Recieved message %s\n", message.Type)

					availableWorkerMessageChan := <-d.WorkerPool

					fmt.Printf("Dispatching message %s\n", message.Type)
					availableWorkerMessageChan <- message
				}(message)
			}
		}
	}()
}
