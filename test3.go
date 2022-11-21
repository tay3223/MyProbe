package main

import (
    "fmt"
    "net"
)

func main() {
    _, err := net.Dial("tcp", "baidu.com:80")
    if err != nil {
        return
    }
    fmt.Printf("Connection successful！")
}
