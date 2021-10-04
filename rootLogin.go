package main

import (
	"fmt"
	"os/user"
)

func main() {
	passwd, _ := user.Current()
    if passwd.Username == "root" { 
        fmt.Println("Root is logged in!")
    }
}
