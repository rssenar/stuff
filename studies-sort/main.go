package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func WordCount(rdr io.Reader) map[string]int {
	counts := map[string]int{}
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		// word := scanner.Text()
		word := strings.ToLower(scanner.Text())
		counts[word]++
	}
	return counts
}

func main() {
	srcFile, err := os.Open("moby.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer srcFile.Close()

	counts := WordCount(srcFile)
	fmt.Println("Number of whales:", counts["whale"])
	fmt.Println("Items in Map:", len(counts))

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range counts {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Key > ss[j].Key
	})

	for _, str := range ss {
		fmt.Printf("%s, %d\n", str.Key, str.Value)
	}
}
