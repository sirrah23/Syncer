package main

import (
    "os"
    "bufio"
    "strings"
    "errors"
    "flag"
    "fmt"
    "os/exec"
    "log"
    "sync"
)

type SrcDest struct{
    src, dest string
}

//Read a csv file with source-destination file pairs
func srcDestRead(fname string) ([]SrcDest, []string, []string, error){
    var pairs []SrcDest
    var srcs, dests []string
    file, err := os.Open(fname)
    defer file.Close()
    if err != nil{
        log.Fatalf("Error: %s", err)
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

//Given a list of files, check to see that they all exist
func filesExist(files []string)(bool){
    for _, file := range files{
        if _, err := os.Stat(file); os.IsNotExist(err){
            return false
        }
    }
    return true
}

//Given a list of strings, make sure there are no duplicates in the list
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

//Run the rsync command for a given source and destination directory
func rsyncCmd(src, dest string, wg *sync.WaitGroup){
    if src[len(src)-1] != '/'{
        src = src + string('/')
    }
    if src[len(dest)-1] != '/'{
        dest = dest + string('/')
    }
    cmd := exec.Command("rsync", "-av", src, dest)
    _, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatalf("Could not rsync %s and %s", src, dest)
    }
    log.Println(fmt.Sprintf("%s and %s have been synced", src, dest))
    wg.Done()
}

//Run a set of rsync commands in parallel and wait for them all to finish
func rsyncs(srcdests []SrcDest){
    var wg sync.WaitGroup
    wg.Add(len(srcdests))
    for _, srcdest := range srcdests{
        rsyncCmd(srcdest.src, srcdest.dest, &wg)
    }
    wg.Wait()
}

//Coordinates the syncing between source and destination directories
func syncer(filesList string) error{
    //TODO: Check for overlap between source & dest files
    pairs, srcs, dests, err := srcDestRead(filesList)
    if err != nil {
        return err
    }
    if len(pairs) == 0{
        return errors.New("Input file is empty")
    }
    if !filesExist(append(srcs, dests...)){
        return errors.New("One or more files listed in input file do not exists")
    }
    if !isUnique(dests){
        return errors.New("Source or destination directory has been duplicated")
    }
    rsyncs(pairs)
    return nil
}

func main(){
    filesListPtr := flag.String("files", "", "List of file pairs to rsync")
    flag.Parse()
    if len(*filesListPtr) == 0{
        log.Fatal("No input file provided")
    }
    err := syncer(*filesListPtr)
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
}
