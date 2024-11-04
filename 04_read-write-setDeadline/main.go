package main

import (
    "fmt"
    "log"
    "net"
    "bufio"
    "time"
)

func main() {
    // start listening
    li, err := net.Listen("tcp", ":8090")
    if err != nil {
        log.Fatalln(err)
    }
    defer li.Close()

    fmt.Println("Listening, reading and writing...")

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
    // setting deadline
    err := conn.SetDeadline(time.Now().Add(10 * time.Second))
    if err != nil {
        log.Fatalln("CONN TIMEOUT")
    }

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        ln := scanner.Text()
        fmt.Println(ln)
        fmt.Fprintf(conn, "I heard you say: %s\n", ln)
    }

    defer conn.Close()

    // now we get here
    // the connection will time out
    // that breaks us out of the scanner loop 
    fmt.Println("***CODE GO HERE***")
} 
