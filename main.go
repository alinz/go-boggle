package main

import (
	"bufio"
	"fmt"
	"github.com/alinz/go-boggle/boggle"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	graph := boggle.Graph{}

	//I used anonymouse function to read and construct a graph
	func() {
		nodesFile, err := os.Open("nodes.txt")
		if err != nil {
			log.Fatal(err)
		}

		defer nodesFile.Close()

		nodesFileBuffer := bufio.NewReader(nodesFile)

		firstLine := true
		for {
			line, isPrefix, err := nodesFileBuffer.ReadLine()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
			}

			if isPrefix {
				log.Fatal("Error: line is too long, WTF?!", nodesFile.Name())
			}

			stringLine := string(line)

			if stringLine == "" {
				continue
			}

			labels := strings.Split(stringLine, " ")

			if firstLine {
				for _, label := range labels {
					graph.AddVertex(label)
				}
				firstLine = false
			} else {
				for i, label := range labels {
					//skip the first one because this is a node that we are connection others to
					if i == 0 {
						continue
					}
					graph.AddEdge(labels[0], label)
				}
			}
		}
	}()

	boggle.GenrateAllPossibleWords(&graph, func(str string) {
		//Put your discitionay search here. str is a word found by system based on graph
		fmt.Println(str)
	})
}
