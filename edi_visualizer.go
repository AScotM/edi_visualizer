package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	segmentDelay = 200 * time.Millisecond
	headerColor  = "\033[1;36m" // bright cyan
	dataColor    = "\033[0;37m" // light gray
	resetColor   = "\033[0m"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: edi_viewer <edi-file>")
		return
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	defer file.Close()

	fmt.Println(headerColor + "EDI Visualizer v0.1 - Terminal Link Open\n" + resetColor)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		showSegment(line)
		time.Sleep(segmentDelay)
	}

	fmt.Println(headerColor + "\nTransmission Complete. Press any key to exit." + resetColor)
	fmt.Scanln()
}

func showSegment(line string) {
	if len(line) > 0 {
		segment := line[:3]
		fmt.Print(headerColor + segment + resetColor)
		fmt.Println(dataColor + line[3:] + resetColor)
	}
}

