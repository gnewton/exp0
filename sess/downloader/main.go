package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "io"
    "strings"
    "bytes"
)

//  curl http://localhost:8081/upload -F "fileupload=@journal.pcbi.1005413.pdf" -vvv

func main() {
     log.SetFlags(log.LstdFlags | log.Lshortfile)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

    var Buf bytes.Buffer
    // in your case file would be fileupload
    file, header, err := r.FormFile("fileupload")
    if err != nil {
        log.Println(err)
	return
    }
    defer file.Close()
    name := strings.Split(header.Filename, ".")
    fmt.Printf("File name %s\n", name[0])
    // Copy the file data to my buffer
    io.Copy(&Buf, file)
    // do something with the contents...
    // I normally have a struct defined and unmarshal into a struct, but this will
    // work as an example
    contents := string(Buf.Bytes())
    fmt.Println(contents)
    // I reset the buffer in case I want to use it again
    // reduces memory allocations in more intense projects
    Buf.Reset()
    // do something else
    // etc write header

    })
    
    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
