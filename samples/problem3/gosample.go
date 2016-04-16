package main

import (
    "flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
    "strings"
)

type stringslice []string

func readURL(url string) string {
	// Read the URL and return body contents

	resp, err := http.Get(url)
	if err != nil {
		// URL Error 
        panic(err)
        os.Exit(1)
	}	
	
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Read error
        panic(err)
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

func write_to_file (word_counts map[string]int, filename string, url string) {

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

    write_content := fmt.Sprintf("url: %s\n", url)
	for word, word_count := range word_counts {
        s := fmt.Sprintf("\t%s %d\n", word, word_count)
        write_content = write_content + s
	}

    n, err := f.WriteString(write_content)
    if err != nil {
        panic(err)
    }

    _ = n

}

func (urls *stringslice) String() string {
    return fmt.Sprintf("%s", *urls)
}

func (urls *stringslice) Set(value string) error {
   *urls = strings.Split(value, ",")

   return nil
}


func main() {

    //default_url := "https://raw.githubusercontent.com/aporeto-inc/internship2016/master/samples/problem2/uniquified_file.txt"

    var urls stringslice
    flag.Var(&urls, "urls", "where value is comma-separated list of urls")
    flag.Parse()

    if flag.NFlag() == 0 {
        flag.PrintDefaults()
        os.Exit(1)
    }

    for i, url := range urls {
	    content := readURL(url)
        filename := fmt.Sprintf("url%d.txt",i+1)
	    write_to_file(count_words(get_words(content)), filename, url)	
    }
}
