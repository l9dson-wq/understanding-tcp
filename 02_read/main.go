package main

import (
    "fmt"
    "net"
    "bufio"
    "log"
)

func main() {
    // defining the listener.
    li, err := net.Listen("tcp", ":8090")
    if err != nil {
        log.Fatalln(err)
    } 

    defer li.Close()

    fmt.Println("Listening and reading...")

    for {
        conn, err := li.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        go handle(conn)
    }
}

func handle(conn net.Conn) {
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        ln := scanner.Text()
        fmt.Println(ln)
    }
    defer conn.Close()

    // Code never goes here
    fmt.Println("CODE GETTING HERE")
}
