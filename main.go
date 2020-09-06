package main

import "custom"

func main() {
    custom.MakeProxy("http://localhost:8080", ":9090")
    // To see this in action check https://droplet.cf:9090
}
