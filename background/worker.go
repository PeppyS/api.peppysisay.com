package background

import "fmt"

type Worker struct {
	ID          int
	MessageChan chan Message // Message for worker to process
	QuitChan    chan bool    // Used to signal worker to stop
}

func NewWorker(ID int) *Worker {
	return &Worker{
		ID:          ID,
		MessageChan: make(chan Message),
		QuitChan:    make(chan bool)}
}

func (w *Worker) Run(workerPool chan chan Message) {
	go func() {
		for {
			fmt.Printf("Worker %d adding self to worker queue\n", w.ID)
			workerPool <- w.MessageChan

			fmt.Printf("Worker %d waiting for message...\n", w.ID)
			select {
			case message := <-w.MessageChan:
				fmt.Printf("Worker %d: Received message %s\n", w.ID, message.Type)
			case <-w.QuitChan:
				fmt.Printf("Worker %d: Stopping", w.ID)
			}
		}
	}()
}
