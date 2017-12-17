package main

import (
    "os"
    "bufio"
    "strings"
    "errors"
)

type SrcDest struct{
    src, dest string
}

func srcDestRead(fname string) ([]SrcDest, []string, []string, error){
    var pairs []SrcDest
    var srcs, dests []string
    file, err := os.Open(fname)
    defer file.Close()
    if err != nil{
        panic(err)
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        lineText := strings.Split(scanner.Text(), ",")
        if len(lineText) > 2 {
            return nil, nil, nil, errors.New("Lines must contain source-destination pairs")
        }
        srcs = append(srcs, lineText[0])
        dests = append(dests, lineText[1])
        pairs = append(pairs, SrcDest{lineText[0], lineText[1]})
    }
    return pairs, srcs, dests,  nil
}

func filesExist(files []string)(bool){
    for _, file := range files{
        if _, err := os.Stat(file); os.IsNotExist(err){
            return false
        }
    }
    return true
}

func isUnique(list []string) bool{
    m := make(map[string]bool)
    for _, item := range list{
        if _, prs := m[item]; prs{
            return false
        } else {
            m[item] = true
        }
    }
    return true
}

func main(){
    _, srcs, dests, err := srcDestRead("./test.csv")
    if err != nil {
        panic(err)
    }
    if ! filesExist(append(srcs, dests...)){
        panic("One or more files listed in input file do not exists")
    }
}
