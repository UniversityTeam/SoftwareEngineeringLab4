package engine

import (
	"fmt"
	"strings"
)

type IHandler interface {
	Post(cmd Command)
}

type Handler struct {
	eventLoop *EventLoop
}

type Command interface {
	Execute(handler IHandler)
}

type PrintCmd struct {
	Msg string
}

func (pc PrintCmd) Execute(handler IHandler) {
	fmt.Println(pc.Msg)
}

type DeleteCmd struct {
	Str    string
	Symbol byte
}

func (dc DeleteCmd) Execute(handler IHandler) {
	r := strings.ReplaceAll(string(dc.Str), string(dc.Symbol), "")
	var printCmd Command = &PrintCmd{r}
	handler.Post(printCmd)
}

type EventLoop struct {
	queue []Command
}

func (eventLoop *EventLoop) Post(cmd Command) {
	eventLoop.queue = append(eventLoop.queue, cmd)
}

func (eventLoop *EventLoop) Start() {
	eventLoop.queue = make([]Command, 0)
}

func (eventLoop *EventLoop) AwaitFinish() {
	for len(eventLoop.queue) > 0 {
		cmd := eventLoop.queue[0]
		eventLoop.queue = eventLoop.queue[1:]
		cmd.Execute(eventLoop)
	}
}

func Parse(str string) Command {
	array := strings.Fields(str)
	if len(array) == 0 {
		return &PrintCmd{Msg: "Line cannot be empty!!!"}
	}
	if array[0] == "delete" {
		if len(array) == 3 {
			byteArray := []byte(array[2])
			if len(byteArray) == 1 {
				return &DeleteCmd{Str: array[1], Symbol: byteArray[2]}
			}
			return &PrintCmd{Msg: "Must be 1 char you want to remove!!!"}
		}
		return &PrintCmd{Msg: "Must be 3 arguments: command, string which you wanna change and char!!!"}
	} else if array[0] == "print" {
		if len(array) == 2 {
			return &PrintCmd{Msg: array[1]}
		}
		return &PrintCmd{Msg: "Must be 2 arguments: command and string which you wanna print!!!"}
	}
	return &PrintCmd{Msg: "Wrong command name. Exists only delete and print commands!!!"}
}
