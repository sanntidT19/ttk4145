package main

import (
."ControlModule"
."DriverModule"
."IOModule"
."NetworkModule"
)

func main() {
go InitNetworkModule(LocalIP, port, netIn, netOut)
// NEI!!!

}
