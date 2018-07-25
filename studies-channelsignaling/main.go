package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
	"utilities/channelsignaling/logger"
)

type device struct {
	problem bool
}

func (d *device) Write(p []byte) (n int, err error) {
	for d.problem {
		time.Sleep(time.Millisecond)
	}
	fmt.Print(string(p))
	return len(p), nil
}

func main() {
	const grs = 10

	var d device
	log := logger.New(&d, grs)
	// l.SetOutput(&d)

	for i := 0; i < grs; i++ {
		go func(id int) {
			for {
				log.Println(fmt.Sprintf("%d: Log data", id))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		<-sigChan
		d.problem = !d.problem
	}
}
