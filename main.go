package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)


func countChar(fileName string) int64 {
    content, err := os.ReadFile(fileName)
    if err != nil {
        log.Fatalf("Failed to open file: %v", err)
    }

    charCount := len(string(content))

    return int64(charCount)
}

func countLines(fileName string) int64 {
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatalf("Faild to open file: %v", err)
    }
    defer file.Close()

    // membaca file baris per baris
    scanner := bufio.NewScanner(file)
    lineCount := 0

    for scanner.Scan() {
        lineCount += 1
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error reading file: %v", err)
    }

    return int64(lineCount)
}

func countWords(fileName string) int64 {
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatalf("Failed to open file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    wordCount := 0

    for scanner.Scan() {
        line := scanner.Text()
        words := bufio.NewScanner(strings.NewReader(line))
        words.Split(bufio.ScanWords)

        for words.Scan(){
            wordCount += 1
        }
    }

    if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

    return int64(wordCount)
}

func countCharMultiByte(fileName string) int64 {
    content, err := os.ReadFile(fileName)
    if err != nil {
        log.Fatalf("Failed to open file: %v", err)
    }

    charCount := utf8.RuneCountInString(string(content))
    return int64(charCount)
}



func main()  {
    if len(os.Args) == 3 {
        fileName := os.Args[2]

        switch os.Args[1] {
        case "-c":
            total := countChar(fileName)
            fmt.Println(total, fileName)
        case "-l":
            total := countLines(fileName)
            fmt.Println(total, fileName)
        case "-w":
            total := countWords(fileName)
            fmt.Println(total, fileName)
        case "-m":
            total := countCharMultiByte(fileName)
            fmt.Println(total, fileName)
 

        default:
            fmt.Println("Unknown argument")
        }
    }

    if len(os.Args) == 2 {
        fileName := os.Args[1]
        lines := countLines(fileName)
        words := countWords(fileName)
        characters := countChar(fileName)

        fmt.Println(lines, words, characters, fileName)
    }
}
