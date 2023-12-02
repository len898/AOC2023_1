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
            }else{
                
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
    part2()
}

func part2(){
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()
    
    m := make(map[string]string)
    m["one"] = "1"
    m["two"] = "2"
    m["three"] = "3"
    m["four"] = "4"
    m["five"] = "5"
    m["six"] = "6"
    m["seven"] = "7"
    m["eight"] = "8"
    m["nine"] = "9"

    scanner := bufio.NewScanner(f)
    
    total := 0

    for scanner.Scan(){
        var first string
        var last string
        var sb strings.Builder
        line := scanner.Text()
        for idx, runeValue := range line{
            if _, err := strconv.Atoi(string(runeValue)); err == nil{
                if first == ""{
                    first = string(runeValue)
                    last = string(runeValue)
                }else{
                    last = string(runeValue)
                }
            }else{
                remainingLength := len(line) - idx
                var substr string

                if(remainingLength >= 3){
                    substr = line[idx:idx+3]
                    v1 := m[substr]
                    if(v1 != ""){
                        if first == ""{
                            first = m[substr]
                            last = m[substr]
                        }else{
                            last = m[substr]
                        }
                    }
                }
                if(remainingLength >= 4){
                    substr = line[idx:idx+4]
                    v1 := m[substr]
                    if(v1 != ""){
                        if first == ""{
                            first = m[substr]
                            last = m[substr]
                        }else{
                            last = m[substr]
                        }
                    }
                }
                if(remainingLength >= 5){
                    substr = line[idx:idx+5]
                    v1 := m[substr]
                    if(v1 != ""){
                        if first == ""{
                            first = m[substr]
                            last = m[substr]
                        }else{
                            last = m[substr]
                        }
                    }
                }
            }
         }
            sb.WriteString(first)
            sb.WriteString(last)
            concatValue, _ := strconv.Atoi(sb.String())
            total += concatValue

    }
    fmt.Printf("Total is %d\n", total)
}


