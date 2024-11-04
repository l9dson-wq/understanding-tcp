package main

import(
    "fmt"
    "net"
    "log"
    "io"
)

func main() {
    // Define the network type and the port
    li, err := net.Listen("tcp", ":8090")
    if err != nil {
        log.Fatalln(err)
    }
    defer li.Close()

    fmt.Println("Listening....")

    for {
        conn, err := li.Accept()
        if err != nil {
            log.Println(err)
            continue
        }

        io.WriteString(conn, "\nHello from TCP server\n")
        fmt.Fprintln(conn, "How is your day going?")
        fmt.Fprintf(conn, "well, I hope!")
        
        conn.Close()
    }
}
