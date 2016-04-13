package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"regexp"
)

func readURL(url string) string {
	// Read the URL and return body contents

	resp, err := http.Get(url)
	if err != nil {
		// URL Error 
		os.Exit(1)
	}	
	
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Read error
		os.Exit(1)
	}	

	return string(body)
}

func get_words (contents string) []string{
	// make a list of all distinct words

	words := regexp.MustCompile("\\w+")
	return words.FindAllString(contents, -1)
}

func count_words (words []string) map[string]int {
	// Count the frequency

	word_counts := make(map[string]int)
	for _, word := range words {
		word_counts[word]++
	}	
	return word_counts
}

func write_to_file (word_counts map[string]int, filename string) {

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for word, word_count := range word_counts {
		if _, err = f.WriteString("%v %v\n", word, word_count); err != nil {
			panic(err)
		}
		//fmt.Printf("%v %v\n", word, word_count)
	}
}

func main() {
	url := "https://raw.githubusercontent.com/aporeto-inc/internship2016/master/samples/problem2/uniquified_file.txt"

	content := readURL(url)
	write_to_file(count_words(get_words(content)), "output")	
}
