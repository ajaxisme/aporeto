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
        fmt.Printf("URL error: \"%s\", Exiting...\n", url)
        os.Exit(1)
	}	
	
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        // Read error
        fmt.Printf("URL Content Read Error: \"%s\". Exiting...\n", url);
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
    // Write frequency table to output file

    f, err := os.Create(filename)
	if err != nil {
        fmt.Printf("Can't create file: \"%s\", Exiting...\n", filename)
        os.Exit(1)
    }

    defer f.Close()

    write_content := fmt.Sprintf("url: %s\n", url)
    for word, word_count := range word_counts {
        s := fmt.Sprintf("\t%s %d\n", word, word_count)
        write_content = write_content + s
    }

    n, err := f.WriteString(write_content)
    if err != nil {
        fmt.Printf("Unable to write to file: \"%s\", Exiting...\n", filename)
        os.Exit(1)
    }

    _ = n // Not using

}

func (urls *stringslice) String() string {
    return fmt.Sprintf("url1,url2,...\t")
}

func (urls *stringslice) Set(value string) error {
   *urls = strings.Split(value, ",")

   return nil
}


func main() {

    //default_url := "https://raw.githubusercontent.com/aporeto-inc/internship2016/master/samples/problem2/uniquified_file.txt"
    flag.Usage = func() {
        // Modifying usage function for better readability
        fmt.Printf("Usage: %s -urls=<comma-separated-one-or-more-urls>\n\n", os.Args[0])
        flag.PrintDefaults()
    }

    var urls stringslice
    flag.Var(&urls, "urls", "Comma-separated list of urls")
    flag.Parse()

    if flag.NFlag() == 0 {
        fmt.Printf("Invalid command: No arguments provided\n")
        flag.Usage()
        os.Exit(1)
    }

    for i, url := range urls {
        content := readURL(url)
        filename := fmt.Sprintf("url%d.txt",i+1)
        write_to_file(count_words(get_words(content)), filename, url)	
    }
}
