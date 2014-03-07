package main

import (
    ."ControlModule"
    
    //."IOModule"
    //."NetworkModule"
)

func main() {
    go ControlModule()
    i := make(chan int,1)
    
    <-i
    
}
