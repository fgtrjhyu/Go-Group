package main

import (
  "fmt"
  "sync"
  "time"
  "math/rand"
  "./group"
)

type Foo int

func (i Foo) Equiv(e group.Element) bool {
  return i == e
}

func print(g group.Group) {
  for r := g.First; r != nil; r = r.Next {
    for m := r.First; m != nil; m = m.Next {
      if m == r.First {
        fmt.Printf(">%v", m.Value)
      } else {
        fmt.Printf(",%v", m.Value)
      }
    }
    fmt.Println()
  }
}

func usageint() {
  fmt.Println("Begin int")
  g := group.Group{}
  for x := 1; x < 30; x++ {
    g.Add(Foo(rand.Intn(5)))
  }
  print(g)
  fmt.Println("End int")
}

type Bar struct {
  number int
  text string
}

func (b Bar) Equiv(e group.Element) bool {
  return b.text == e.(Bar).text
}

func chooseText(n int) string {
  switch n % 4 {
  default: panic("bug")
  case 0: return "foo"
  case 1: return "bar"
  case 2: return "baz"
  case 3: return "qux"
  }
}

func usagestruct() {
  fmt.Println("Begin struct")
  var wg sync.WaitGroup
  wg.Add(2)
  c := make(chan int)
  d := make(chan Bar)
  go func() {
    defer func() { close(c); wg.Done() }()
    for x := 1; x < 30; x++ {
      c<- rand.Intn(10)
    }
  }()
  go func() {
    defer func() { close(d); wg.Done() }()
    for x, ok := <-c; ok; x, ok = <-c {
      d<- Bar{number: x, text: chooseText(x)}
    }
  }()
  g := group.Group{}
  for x, ok := <-d; ok; x, ok = <-d {
    g.Add(x)
  }
  print(g)
  fmt.Println("End struct")
}

func main() {
  rand.Seed(time.Now().Unix())
  usageint()
  usagestruct()
}
