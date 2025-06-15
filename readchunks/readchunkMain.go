package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Println("Read chunks of text")
	myReadLine("test.txt")

}

func myReadLine(fileName string) {
	/**
		Regex used in this function
		Detect first line of the block

	**/
	frame_begin := regexp.MustCompile("This is a file part\\s+\\d+")
	var file_line string
	var string_accumulator []string
	file, err := os.Open(fileName)
	var first_rec bool = true
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close() // Ensure the file is closed

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		file_line = scanner.Text()
		if frame_begin.MatchString(file_line) && first_rec {
			fmt.Printf("Start block found: %s\n", file_line)
			string_accumulator = append(string_accumulator, string(file_line))
			first_rec = false
		} else if frame_begin.MatchString(file_line) && !first_rec { // A record was accumulated
			//fmt.Printf("Record found : %q\n", string_accumulator) // Print or process the record
			//printMyRecordLines(string_accumulator)
			processRecoreLevel1(string_accumulator)
			string_accumulator = nil                                           // Empty the accumultor for next record
			string_accumulator = append(string_accumulator, string(file_line)) // New record
		} else {
			string_accumulator = append(string_accumulator, string(file_line))
		}
	}
	//fmt.Printf("Record found : %s\n", string_accumulator) // Last record
	printMyRecordLines(string_accumulator)

}

func printMyRecordLines(recordList []string) {
	for _, rec_line := range recordList {
		fmt.Printf("Record line: %s\n", rec_line)
	}
}

func processRecoreLevel1(toplevelRec []string) {
	for _, recLine := range toplevelRec {
		// Start processing lines here for sub-levels
		fmt.Printf("Process line: %s\n", recLine)
	}
}
