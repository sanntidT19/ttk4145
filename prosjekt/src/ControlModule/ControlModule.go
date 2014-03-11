package ControlModule

import (
	."DriverModule"
	."NetworkModule"
	."IOModule"
	list "util"
	."strings"
	"strconv"
	"fmt"
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
	elevators *list.LinkedList
	handledFloors [N_FLOORS][3][256]bool //handledFLoors[floor][button][elevator_id]
	
	//modules
	network *NetworkModule
	io *IOModule
	
	//channels
	toNet chan string
	fromNet chan string
	toIO chan string
	fromIO chan string
	
}

func ControlModule() {

	
	
	control := new(Control)
	control.toNet = make(chan string, 10)
	control.fromNet = make(chan string, 10)
	control.toIO = make(chan string, 10)
	control.fromIO = make(chan string, 10)
	control.e = InitElevator()
	control.network = InitNetworkModule("161.129.241.161", ":20011", control.toNet, control.fromNet)
	control.io = InitIOModule(control.toIO, control.fromIO)
	control.elevators = list.New()
	control.e.id = -1
	
	//elevator procedures
	go control.e.Update()
	go control.e.RunDMC()
	go control.e.Buttons()
	
	//io procedures
	go control.io.ConsoleOut()
	
	//network procedures
	go control.network.Broadcast()
	go control.network.Listen()
	
	control.Backup()
	control.Master()
	
}

func (c *Control) Master(){

	if c.e.id == -1{ //ingen boss
		c.e.id = 0	
	}
	
	masterMessage := "Boss Alive"
	fmt.Println(masterMessage)
	c.elevators.PushFront(c.e)	
	
	
	for {
		//c.toNet <- masterMessage
		c.toIO <-<-c.e.out
		//fmt.Println(msg)
		//c.toIO<-msg
		
	}
	
	
	
}

func (c *Control) Backup(){
	return 
}

func (c *Control) ComputeMessages() {
	
	
	for {
		msg := <-c.fromNet
		
		substrings := SplitN(msg, ":", -1)
		
		if substrings[2] == "ButtonPush" { //new button push
			floor,_ := strconv.Atoi(substrings[3])
			button,_ := strconv.Atoi(substrings[4])
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
			/*if c.buttonPushes[floor][button] && !handledFloors[floor][button]{
				c.AddFloor(floor, button)
			}*/
		}
	}
}

func (c *Control) AddFloor(floor, button int){
	
}
