package main

import (
    "fmt"
    "net"
    "log"
    "bufio"
)

func main() {
    // start listening
    li, err := net.Listen("tcp", ":8090")
    if err != nil {
        log.Fatalln(err)
    }

    defer li.Close()

    fmt.Println("Listening and reading...")

    for {
        conn, err := li.Accept()
        if err != nil {
            fmt.Println(err)
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
        fmt.Fprintf(conn, "I've heard you said: %s\n", ln)
    }

    defer conn.Close()

    // we never get here
    // we have an open stream connection
    // how does the above reader know it's done?
    fmt.Println("CODE GETTING HERE")
}
