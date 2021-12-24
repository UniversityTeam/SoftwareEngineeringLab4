package main

import (
	"bufio"
	"os"
    	"flag"
	"main/engine"
)

func main() {
	eLoop := new(engine.EventLoop)
	eLoop.Start()
	flag.Parse()
	inputFile := "filename"

	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		emptyFile := true
		for scanner.Scan() {
			emptyFile = false
			commandLine := scanner.Text()
			cmd := engine.Parse(commandLine)
			eLoop.Post(cmd)
		}
		if emptyFile {
			cmd := &engine.PrintCmd{Msg: "File is empty!!!"}
			eLoop.Post(cmd)
		}
	}
	eLoop.AwaitFinish()
}