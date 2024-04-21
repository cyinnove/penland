package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"net/http"
	"strings"
	"time"

)

// Define a struct to match the JSON structure
type Writeups struct {
	Data []struct {
		Links    []struct {
			Title string `json:"Title"`
			Link  string `json:"Link"`
		} `json:"Links"`
		Programs []string `json:"Programs"`
		Bugs     []string `json:"Bugs"`
	} `json:"data"`
}

// Result struct for the desired output JSON format
type Result struct {
	URL  string   `json:"url"`
	Targets []string `json:"target"`
	Bugs []string `json:"Bugs"`
}

type Results struct {
	Result []Result `json:"result"`
}

// Flags to capture user input
var (
	title    = flag.String("title", "", "Filter by the title of the writeup")
	link     = flag.String("link", "", "Filter by the link of the writeup")
	programs = flag.String("programs", "", "Filter by the program involved in the writeup")
	bugs     = flag.String("bugs", "", "Filter by the bug types mentioned in the writeup")
	output   = flag.String("output", "", "Specify the output file name")
)
func main() {
	flag.Parse() // Parse the provided flags

	client := http.Client{
		Timeout: 1 * time.Minute,
	}

	// Fetch the JSON data from the URL
	req, err := http.NewRequest("GET", "https://pentester.land/writeups.json", nil)
	if err != nil {
		fmt.Printf("Failed to fetch data: %s\n", err)
		return
	}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	// Read and parse the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read data: %s\n", err)
		return
	}

	var writeups Writeups
	if err = json.Unmarshal(body, &writeups); err != nil {
		fmt.Printf("Failed to parse JSON: %s\n", err)
		return
	}

	var results Results
	for _, data := range writeups.Data {
		match := true
		if *title != "" && !strings.Contains(strings.ToLower(data.Links[0].Title), strings.ToLower(*title)) {
			match = false
		}
		if *link != "" && !strings.Contains(strings.ToLower(data.Links[0].Link), strings.ToLower(*link)) {
			match = false
		}
		if *programs != "" {
			programMatch := false
			for _, program := range data.Programs {
				if strings.Contains(strings.ToLower(program), strings.ToLower(*programs)) {
					programMatch = true
					break
				}
			}
			if !programMatch {
				match = false
			}
		}
		if *bugs != "" {
			bugsMatch := false
			for _, bug := range data.Bugs {
				if strings.Contains(strings.ToLower(bug), strings.ToLower(*bugs)) {
					bugsMatch = true
					break
				}
			}
			if !bugsMatch {
				match = false
			}
		}
		if match {
			for _, link := range data.Links {
				results.Result = append(results.Result, Result{
					URL:     link.Link,
					Targets: data.Programs,
					Bugs:    data.Bugs,
				})
			}
		}
	}

	if *output == "" {
		// Print results to the console if output is not provided or default
		for _, result := range results.Result {
			fmt.Printf("URL: %s\n", result.URL)
			fmt.Printf("Programs: %s\n", strings.Join(result.Targets, ", "))
			fmt.Printf("Bugs: %s\n", strings.Join(result.Bugs, ", "))
			fmt.Println("---")
		}
	} else {
		// Write the results to a JSON file if output is provided
		file, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Printf("Failed to marshal results: %s\n", err)
			return
		}
		if err = os.WriteFile(*output, file, 0644); err != nil {
			fmt.Printf("Failed to write to file: %s\n", err)
			return
		}
		fmt.Printf("Results written to '%s'\n", *output)
	}
}
