package ControlModule

import (
	."DriverModule"
	//."fmt"
	."time"
	."NetworkModule"
	."IOModule"
	list "container/list"
	map "container/map"
)

func ControlModule() {
	e := InitElevator()
	network := InitNetworkModule()
	io := InitIOModule()
	
	go e.Update()
	go e.RunDMC()
	go e.Buttons()
	
	
	
	
	
    /*
    button := -1
    floor := -1
    var nbutton, nfloor int
      
	for  {
		nbutton,nfloor = e.HandleButtonPress(button,floor)
		if !(nbutton == -1 && nfloor == -1){
		    button = nbutton
		    floor = nfloor
		    ElevSetButtonLamp(button, floor , 1)
		    //put floor in queue
		    e.UpdateList(floor)
		}
		//Println(button,floor)
		
		if ElevGetStopSignal() == 1 {
			return
		}
	}
*/
	
}

func Master(e *Elevator, network *NetworkModule, io *IOModule){
	masterMessage := "Boss Alive"
	elevators := list.New()
	elevators.PushFront(e)
	
	
	for {
		network.in <- masterMessage
		
	}
	
	
	
}

func Backup(){

}

