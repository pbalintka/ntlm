package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

var statusfile string = "/var/lib/dpkg/status"

type Item struct {
    Package string `json:"package"`
    Version string `json:"version"`
    Depends string `json:"depends"`
    Description string `json:"description"`
}

var Items []*Item


func Packages() []*Item {
    var item *Item

    f, err := os.Open(statusfile)
    if err!=nil {
        log.Println(err)
        return nil
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        if line=="" {
            Items = append(Items, item)
            continue
        }

        splitted := strings.SplitN(line, ":", 2)
        switch splitted[0] {
        case "Package":
            item = new(Item)
            item.Package = strings.TrimSpace(splitted[1])
        case "Version":
            item.Version = strings.TrimSpace(splitted[1])
        case "Depends":
            item.Depends = strings.TrimSpace(splitted[1])
        case "Description":
            item.Description = strings.TrimSpace(splitted[1])
        }
    }

    return Items
}

func main() {
    packs := Packages()
    for i:=0; i<len(packs); i++ {
        fmt.Println(packs[i])
    }
}

