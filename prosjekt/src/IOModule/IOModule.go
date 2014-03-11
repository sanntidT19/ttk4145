package IOModule

import (
."bufio"
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

func InitIOModule(i chan string, o chan string) *IOModule {
    m := new(IOModule)
    m.in = i
    m.out = o
    m.reader = NewReader(os.Stdin)
    
    return m
}

func (m *IOModule) ConsoleIn() {
    
    
    input,_ := m.reader.ReadString('\n')
    
    m.out <- strings.TrimRight(input,"\n")
}

func (m *IOModule) ConsoleOut() {
    var str string
    for {
        str = <- m.in
        fmt.Println("[CONSOLE] ", str)
    }

}
