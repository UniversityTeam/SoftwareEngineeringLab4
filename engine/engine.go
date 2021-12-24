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