package IOModule

import (
"bufio"
"fmt"
"strings"
"os"
)

type IOModule struct {
    
    // channel for incoming messages
    in chan string
    
    // channel for outcoming messages
    out chan string
    
    // console reader
    reader *Reader
    
}

func InitIOModule(i chan string, o chan string) IOModule {
    var mod IOModule
    mod.in = i
    mod.out = o
    mod.reader = bufio.NewReader(os.Stdin)
    
    return mod
}

func (IOModule) ConsoleIn() {
    
    Print(prompt)
    input,_ := reader.ReadString('\n')
    
    return strings.TrimRight(input,"\n")
}

func (IOModule) ConsoleOut() {
    var str string
    for {
        str = <- in
        Println("[CONSOLE] ", str)
    }

}
