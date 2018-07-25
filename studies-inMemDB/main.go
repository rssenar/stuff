package main

import (
	"fmt"
)

type post struct {
	id      int
	content string
	author  string
}

var postByID map[int]*post
var postsByAuthor map[string][]*post

func store(post post) {
	postByID[post.id] = &post
	postsByAuthor[post.author] = append(postsByAuthor[post.author], &post)
}

func main() {
	postByID = make(map[int]*post)
	postsByAuthor = make(map[string][]*post)

	post1 := post{
		id:      1,
		content: "Hello World!",
		author:  "Sau Sheong",
	}
	post2 := post{
		id:      2,
		content: "Bonjour Monde!",
		author:  "Pierre",
	}
	post3 := post{
		id:      3,
		content: "Hola Mundo!",
		author:  "Pedro",
	}
	post4 := post{
		id:      4,
		content: "Greetings Earthlings!",
		author:  "Sau Sheong",
	}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(postByID[1])
	fmt.Println(postByID[2])

	for _, post := range postsByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}
	for _, post := range postsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}
