package videosource

import (
	"fmt"
	"time"
)

type VS struct {
	id   string
	done chan bool
}

func NewVS(id string) *VS {
	return &VS{id: id, done: make(chan bool)}
}

func (vs *VS) Process(sourceURL string) {
	go func() {
		for {
			select {
			case <-vs.done:
				{
					return
				}
			default:
				{
					// read image
					// process image
					// send results
					fmt.Printf("source %s is processing\n", sourceURL)
					time.Sleep(time.Second)
				}
			}
		}
	}()

}

func (vs *VS) EndProcess() {
	vs.done <- true
}
