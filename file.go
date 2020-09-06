package main

import "custom"

func main() {
    custom.MakeProxy("http://localhost:8080", ":9090")
}