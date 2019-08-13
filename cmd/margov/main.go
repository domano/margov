package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

func main() {
	textBytes, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	chain := NewChain()
	chain.Parse(string(textBytes))

	println(chain.Generate())

}

type Chain struct {
	index map[Link][]Entry
	lineStarts map[Link]int
}

func NewChain() Chain {
	return Chain{index: map[Link][]Entry{}, lineStarts: map[Link]int{}}
}

func (c Chain) Generate() string {
	var result string

	beginning, i := c.chooseBeginning()
	result = fmt.Sprintf("%s %s ", beginning.Key1, beginning.Key2)

	result = c.generate(beginning, i, result)

	return result
}

func (c Chain) generate(beginning Link, i int, result string) string {
	const length = 30
	key1, key2 := beginning.Key1, beginning.Key2
	for i = 0; i < length; i++ {
		entry := c.chooseByProbabilty(key1, key2)
		result = fmt.Sprintf("%s%s ", result, entry.Word)
		key1 = key2
		key2 = entry.Word
	}
	return result
}

func (c Chain) chooseBeginning() (Link, int) {
	var beginning Link
	i := rand.Intn(len(c.lineStarts))
	for k := range c.lineStarts {
		if i == 0 {
			beginning = k
		}
		i--
	}
	return beginning, i
}

func (c Chain) chooseByProbabilty(key1 string, key2 string) (Entry) {
	entries := c.index[Link{Key1: key1, Key2: key2}]
	var probablitySum int
	for _, entry := range entries {
		probablitySum += entry.Count
	}
	index := rand.Intn(probablitySum)+1
	var sum int
	var i int
	for i = 0; sum < index; i++ {
		sum = sum + entries[i].Count
	}
	return entries[i-1]
}

func (c Chain) Parse(s string) {
	cleanedWords := c.cleanWords(s)
	if len(cleanedWords) < 3 { // we need at least enough words to form a single chain link
		return
	}
	for i := 2; i < len(cleanedWords); i++ { // start at a point where we have enough words for our first keys
		link := Link{
			Key1: cleanedWords[i-2],
			Key2: cleanedWords[i-1],
		}
		c.increaseProbability(link, cleanedWords[i])
		c.parseBeginnings(cleanedWords, i)
	}
}

func (c Chain) cleanWords(s string) []string {
	words := strings.Split(s, " ")
	var cleanedWords []string
	for i := range words {
		if words[i] == "" {
			continue
		}
		cleanedWords = append(cleanedWords, words[i])
	}
	return cleanedWords
}

func (c Chain) parseBeginnings(words []string, i int) {
	firstWord := words[i-2]
	if strings.ContainsAny(firstWord[len(firstWord)-1:],".?!") {
		startLink := Link{Key1: words[i-1], Key2: words[i]}
		count, exists := c.lineStarts[startLink]
		if !exists {
			c.lineStarts[startLink] = 1
		}
		count++
		c.lineStarts[startLink] = count
	}
}

func (c Chain) increaseProbability(link Link, word string) {
	entries := c.index[link]
	for i := range entries {
		if entries[i].Word == word {
			entries[i].Count++
			return
		}
	}
	c.index[link] = append(entries, Entry{Word: word, Count: 1})
}

type Link struct {
	Key1, Key2 string
}

type Entry struct {
	Word  string
	Count int
}
