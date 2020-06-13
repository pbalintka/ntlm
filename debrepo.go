package main

import (
    "compress/gzip"
    "fmt"
    "net/http"
    "io"
    "os"
    "strings"
)

func FileDownload(src string, dst string) error {
    resp, err := http.Get(src)
    if err!=nil {
        return err
    }
    defer resp.Body.Close()

    out, err := os.Create(dst)
    if err!=nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, resp.Body)
    return err
}

func Gunzip(filename string) error {
    f, err := os.Open(filename)
    if err!=nil {
        return err
    }
    defer f.Close()

    reader, err := gzip.NewReader(f)
    if err!=nil {
        return err
    }
    defer reader.Close()

    newfilename := strings.TrimSuffix(filename, ".gz")
    writer, err := os.Create(newfilename)
    if err!=nil {
        return err
    }
    defer writer.Close()

    _, err = io.Copy(writer, reader)
    if err!=nil {
        return err
    }

    err = os.Remove(filename)
    return err
}

func main() {
    var relfile string = "http://ftp.hu.debian.org/debian/dists/buster/Release"
    var packsfile string = "http://ftp.hu.debian.org/debian/dists/buster/contrib/binary-amd64/Packages.gz"

    err := FileDownload(relfile, "./tmp/Release")
    fmt.Println(err)

    err = FileDownload(packsfile, "./tmp/Packages.gz")
    fmt.Println(err)

    err = Gunzip("./tmp/Packages.gz")
    fmt.Println(err)
}

