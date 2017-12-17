package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "errors"
)

type SrcDest struct{
    src, dest string
}

func srcDestRead(fname string) ([]SrcDest, error){
    var pairs []SrcDest
    file, err := os.Open(fname)
    defer file.Close()
    if err != nil{
        panic(err)
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        lineText := strings.Split(scanner.Text(), ",")
        if len(lineText) > 2 {
            return nil, errors.New("Lines must contain source-destination pairs")
        }
        pairs = append(pairs, SrcDest{lineText[0], lineText[1]})
    }
    return pairs, nil
}

func main(){
    fmt.Println(srcDestRead("./test.csv"))
}
