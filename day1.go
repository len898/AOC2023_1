package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    
    total := 0

    for scanner.Scan() {
        var first string
        var last string
        var sb strings.Builder
        fmt.Printf("line: %s\n", scanner.Text())
        for i, c := range scanner.Text(){
            fmt.Println(i, "=>", string(c))
            if _, err := strconv.Atoi(string(c)); err == nil{
               if first == "" {
                    first = string(c)
                    last = string(c)
                    fmt.Println("First is %s", first)
               }else{
                    last = string(c)
                    fmt.Println("Last is %s", last)
               } 
            }
        }
        sb.WriteString(first)
        sb.WriteString(last)
        
        concatValue, _ := strconv.Atoi(sb.String())
        total += concatValue
    }

    if err := scanner.Err(); err != nil{
        log.Fatal(err)
    }

    fmt.Println("Total is %d", total)
}

