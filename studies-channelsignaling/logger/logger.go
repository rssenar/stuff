package logger

import (
	"fmt"
	"io"
	"sync"
)

type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

func New(w io.Writer, cap int) *Logger {
	log := Logger{
		ch: make(chan string, cap),
	}
	log.wg.Add(1)
	go func() {
		for v := range log.ch {
			fmt.Fprintln(w, v)
		}
		log.wg.Done()
	}()
	return &log
}

func (l *Logger) Stop() {
	close(l.ch)
	l.wg.Wait()
}

func (l *Logger) Println(s string) {
	select {
	case l.ch <- s:
	default:
		fmt.Println("Logging Error")
	}
}
