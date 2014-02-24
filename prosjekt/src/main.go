package main

import (
."ControlModule"
."DriverModule"
."IOModule"
."NetworkModule"
)

func main() {
NetworkModule :=InitNetwork()

go NetworkModule.start()

}
