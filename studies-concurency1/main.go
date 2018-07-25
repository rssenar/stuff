package main

import (
	"fmt"
	"math/rand"
	"time"
)

type message struct {
	str  string
	wait chan bool
}

func main() {
	c := fanIn(boring("Richard"), boring("Gertie"), boring("Gabie"), boring("Riley"))
	for msg := range c {
		fmt.Println(msg.str)
		msg.wait <- true
	}
	// for i := 0; i < 6; i++ {
	// msg1 := <-c
	// fmt.Println(msg1.str)
	// msg2 := <-c
	// fmt.Println(msg2.str)
	// msg3 := <-c
	// fmt.Println(msg3.str)
	// msg4 := <-c
	// fmt.Println(msg4.str)
	// msg.wait <- true
	// msg.wait <- true
	// msg.wait <- true
	// msg.wait <- true
	// }
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan message {
	c := make(chan message)
	wait := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- message{fmt.Sprintf("%s: %d", msg, i), wait}
			doSomeWork()
			<-wait
		}
	}()
	return c
}

func fanIn(input1, input2, input3, input4 <-chan message) <-chan message {
	c := make(chan message)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			case s := <-input3:
				c <- s
			case s := <-input4:
				c <- s
			}
		}
	}()
	return c
}

func doSomeWork() {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )
// func main() {
// 	rand.Seed(time.Now().UnixNano()) // Try changing this number!
// 	answers := []string{
// 		"It is certain",
// 		"It is decidedly so",
// 		"Without a doubt",
// 		"Yes definitely",
// 		"You may rely on it",
// 		"As I see it yes",
// 		"Most likely",
// 		"Outlook good",
// 		"Yes",
// 		"Signs point to yes",
// 		"Reply hazy try again",
// 		"Ask again later",
// 		"Better not tell you now",
// 		"Cannot predict now",
// 		"Concentrate and ask again",
// 		"Don't count on it",
// 		"My reply is no",
// 		"My sources say no",
// 		"Outlook not so good",
// 		"Very doubtful",
// 	}
// 	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
// }
