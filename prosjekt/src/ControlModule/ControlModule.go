package ControlModule

import (
	."DriverModule"
	."fmt"
	."time"
	."NetworkModule"
	."IOModule"
	list "container/list"
	."strings"
)

/*
Forslag:
Alle beskjedane er av typen
"<id>:<kommandotype>:<arg0>:<arg1>: ... :<argN>"
der <kommandotype> for eskempel kan vere "ButtonPush" eller "Message"

Då kan vi bruke SplitN for å skaffe relevante dela av beskjeden
*/

/* Den gamle testfunksjonen
func ControlModule(){
		e := InitElevator()
		//network := InitNetworkModule()
		//io := InitIOModule()

		go e.Update()
		go e.RunDMC()
		go e.Buttons()
}
*/


type Control struct {
	//own elevator
	e *Elevator	
	
	//status
	master bool
	
	//storage
	buttonPushes [][]bool
	elevators list.List
	isAlive []bool
	handledFloors [][][]bool //handledFLoors[floor][button][elevator_id]
	
	//modules
	network NetworkModule
	io IOModule
	
}

func ControlModule() {
	var control Control
	control.e := InitElevator()
	control.network := InitNetworkModule()
	control.io := InitIOModule()
	control.elevators := list.New()
	control.e.id = -1
	
	
	go control.e.Update()
	go control.e.RunDMC()
	go control.e.Buttons()
	
	control.Backup()
	control.Master()
	
}

func (c *Control) Master(){

	if c.e.id == -1{ //ingen boss
		c.e.id = 0	
		control.isAlive[e.id] = true
	}
	
	masterMessage := "Boss Alive"
	
	c.elevators.PushFront(e)	
	
	
	for {
		network.in <- masterMessage
		
		
	}
	
	
	
}

func Backup(){
	return 
}

func (c *Control) ComputeMessages() {
	
	
	for {
		msg := <-c.network.in
		
		substrings := SplitN(msg, ":")
		
		if substrings[2] == "ButtonPush") { //new button push
			floor := substrings[3]
			button := substrings[4]
			c.buttonPushes[floor][button] = true //specify new button push			
		} else if substrings[2] == "Livsteikn" { //livsteikn
			
		}
	
	}
}

func (c *Control) ComputeMap() {
	for floor := 0; floor < N_FLOORS; floor++ {
		for button:= 0; button < N_BUTTONS; button++ {
			if (button == BUTTON_CALL_UP && floor == N_FLOORS-1) || (button == BUTTON_CALL_DOWN && floor == 0){
				continue
			}
			if c.buttonPushes[floor][button] && !handledFloors[floor][button]{
				AddFloor(floor, button)
			}
		}
	}
}

func (c *Control) AddFloor(floor, button int){
	
}
