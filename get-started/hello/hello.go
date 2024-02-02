package main

import (
  "fmt"
  "log"
  "utiiz/greetings"
)

func main() {
  log.SetPrefix("greetings: ")
  log.SetFlags(0)

  names := []string {"John", "Juliana", "Frank"}
  messages, err := greetings.Hellos(names)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(messages)
}
