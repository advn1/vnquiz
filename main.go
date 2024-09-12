package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"github.com/fatih/color"
)


func main() {
	correct_answers := 0
	answer_c := make(chan string)

	custom_csv_path := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit := flag.Int64("limit", 5, "the time limit for the quiz in seconds")


	// Override Usage Helper Text
	flag.Usage = func() {
		usage_color := color.New(color.FgYellow)
		usage_color.Printf("Usage of %s:\n", os.Args[0])

		title_color := color.New(color.FgCyan)
		type_color := color.New(color.FgCyan, color.Bold)
		desc_color := color.New(color.FgMagenta)
		default_color := color.New(color.FgBlue)

		title_color.Printf("-csv")
		type_color.Printf(" -string\n")
		desc_color.Printf("\tpath to csv file in 'question, answer' format'")
		default_color.Printf(" (default \""+*custom_csv_path+"\")\n")

		title_color.Printf("-limit")
		type_color.Printf(" -int\n")
		desc_color.Printf("\tthe time limit for the quiz in seconds")
		default_color.Printf(" (default %d)\n", *limit)
	}	
	flag.Parse()

	// CSV Parsing

	var file *os.File
	var err error

	file, err = os.Open(*custom_csv_path)

	if err != nil {
		log.Fatal(err)
		return
	}
	
	defer file.Close()

	// CSV Reading

	csvReader := csv.NewReader(file)
	parsed, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
		return
	}

	buffer := bufio.NewReader(os.Stdin)

Quiz: 
	for idx, record := range parsed {
		timer := time.NewTimer(time.Duration(*limit) * time.Second)

		question_color := color.New(color.FgCyan, color.Bold)
		question_color.Printf("Question #%d: %s = ", idx+1, record[0])

		go readBuffer(buffer, answer_c)

		select {
		case <-timer.C:
			timeout_color := color.New(color.FgYellow)
			timeout_color.Printf("\nTime Out\n")
			break Quiz
		case answer := <-answer_c:
			if (answer == record[1]) {
				answer_color := color.New(color.FgGreen)
				answer_color.Printf("\nCORRECT\n")
				correct_answers++
			} else {
				incorrect_color := color.New(color.FgRed)
				incorrect_color.Printf("\nINCORRECT\n")
			}
		}
		fmt.Println("--------------------------------------------------")
	}
	results_color := color.New(color.FgMagenta, color.Bold)
	results_color.Printf("Results: %d out of %d", correct_answers, len(parsed))
}

func readBuffer(buffer *bufio.Reader, answer_c chan string) {
	reader, _ := buffer.ReadString('\n')
	reader = strings.TrimSpace(reader[:len(reader)-2])
	answer_c<-reader
}