package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "flag"
  "strings"
  "strconv"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("A file name required")
    os.Exit(1)
  }
  data, err := ioutil.ReadFile(os.Args[1])
  if err != nil {
    panic(err)
  }
  var lineNumbers = flag.Bool("n", true, "Line numbers")
  str := strings.TrimRight(string(data), "\n")
  if *lineNumbers {
    lines := strings.Split(str, "\n")
    linesWithNums := make([]string, 0, len(lines))
    var width = 1
    for num := len(lines); num > 10; width++ {
      num /= 10
    }
    lineNum := 1
    for _, line := range(lines) {
      linesWithNums = append(linesWithNums, fmt.Sprintf("%" + strconv.Itoa(width) + "d %s", lineNum, line))
      lineNum++
    }
    str = strings.Join(linesWithNums, "\n")
  }
  fmt.Println(str)
}
