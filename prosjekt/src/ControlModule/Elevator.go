package ControlModule

import (
	."DriverModule"
	."fmt"
	."time"
	list "container/list"
)

type Elevator struct {
	id int
	stopList *list.List
	speed int
	direction int
	currentFloor int
	location int
}

func InitElevator() *Elevator {
	// initializes elevator object	
	if ElevInit() == 0 {
		Println("I am Error")
	}

	e := new(Elevator)	
	e.stopList = list.New()
	e.direction = 0
	//go e.PrintStatus()
	if ElevGetFloorSensorSignal() == -1 {
		// goes down to nearest floor if between two floors
		e.speed = ElevSetSpeed(-300,0);
		for ElevGetFloorSensorSignal() == -1 {}
	}
	
	e.speed = ElevSetSpeed(0,-300);	
	e.currentFloor = ElevGetFloorSensorSignal()
	
	return e
}

func (e *Elevator) Update() {
	for {
		e.location = ElevGetFloorSensorSignal()

		if e.location != -1 {
			e.currentFloor = e.location
		}
		Sleep(Millisecond)
	}
}


func (e *Elevator) Buttons() {
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
		    if button == BUTTON_COMMAND {
		    	e.UpdateList(button, floor)
		    	//other buttons are handled by control
		    }
		}
		//Println(button,floor)
		
		if ElevGetStopSignal() == 1 {
			return
		}
		Sleep(Millisecond)
	}
	
	
}

func (e *Elevator) GetNextFloor() int {
	// returns next floor in stopList
	if e.stopList.Front() != nil {
		return e.stopList.Front().Value.(int)
	} else {
		return -1
	}
}

func (e *Elevator) HandleButtonPress(oldButton int, oldFloor int) (int,int) {
	// check for first button press
	for floor := 0; floor < N_FLOORS; floor++ {
		for button:= 0; button < N_BUTTONS; button++ {
			if (button == BUTTON_CALL_UP && floor == N_FLOORS-1) || (button == BUTTON_CALL_DOWN && floor == 0){
				continue
			}
			if ElevGetButtonSignal(button, floor) == 1 && !(ElevGetButtonLamp(button, floor)){
			    Println("Button", button, "at floor", floor, "was pressed")
                return button, floor			
			} 
		}
	}
	return -1,-1
}

func (e *Elevator) RunDMC() {
    for {
        destination := e.GetNextFloor()
        
        if  e.location != -1 {
           
            ElevSetFloorIndicator(e.currentFloor)
        }
            
        
        
        
        if destination == -1 {
        	e.direction = 0
        } else if e.currentFloor > destination {
            e.direction = -1
            e.speed = ElevSetSpeed(-300, e.speed)
        } else if e.currentFloor < destination {
            e.direction = 1
            e.speed = ElevSetSpeed(300, e.speed)
        } else if ElevGetStopSignal() == 1 {
        	e.speed = ElevSetSpeed(0, e.speed)
        	break
        } else {
           
            e.speed = ElevSetSpeed(0, e.speed)
            e.stopList.Remove(e.stopList.Front())
            
            if e.stopList.Len() == 0  || ElevGetButtonLamp(0, destination) != ElevGetButtonLamp(1, destination){
            	ElevSetButtonLamp(1, destination, 0)
            	ElevSetButtonLamp(0, destination, 0)
            } else if destination < e.GetNextFloor() {
            	ElevSetButtonLamp(0, destination, 0)
            } else if destination > e.GetNextFloor() {
            	ElevSetButtonLamp(1, destination, 0)
            } else  {
            	ElevSetButtonLamp(1, destination, 0)
            	ElevSetButtonLamp(0, destination, 0)
            } 
           
           ElevSetButtonLamp(2, destination, 0)
           Sleep(Second)            
        }
        Sleep(Millisecond)
    }
}

func (e *Elevator) UpdateList(button, floor int) int {
	// inserts the floor to an appropriate position in the list
	// returns 1 if floor was added and 0 orherwise
	
	l := e.stopList
							
	if l.Len() == 0 || e.direction == 0{
		// insert at front if list is empty
		_ = l.PushFront(floor)
		return 1
	}
	
	if Contains(l,floor) {
	    // do nothing if list already contains floor
	    Println("Contains")
		return 0
	}
	inserted := false 
	if e.direction == 1 {
		// insert before smallest element larger than floor
		for el := l.Front(); el != nil; el = el.Next() {
			if el.Value.(int) > floor && floor > e.currentFloor{
				if (button == BUTTON_CALL_DOWN) {
						continue
				}//nedoverknappen
				_ = l.InsertBefore(floor,el)
				Println("opp")
				inserted = true
				break
			} 
		}
	} else if e.direction == -1 {
		// insert before largest element smaller than floor
		for el := l.Front(); el != nil; el = el.Next() {
			if el.Value.(int) < floor && floor < e.currentFloor{
				if button == BUTTON_CALL_UP {
					continue
				}//oppoverknappen
				_ = l.InsertBefore(floor,el)
				inserted = true
				break
			}
		}
	} 
	if inserted {
		return 1
	} else {
		_ = l.PushBack(floor)
		return 0
	}

}

func Contains(l *list.List, val int) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == val {
			return true
		}
	}
	return false
}

func (e *Elevator) PrintStatus() {
    for {
        Println("Elevator stats")
        Println("Speed:\t", e.speed)
        Println("Direction:\t",e.direction)
        Println("Current Floor:\t", e.currentFloor)
        Println("Number of floors in queue:\t", e.stopList.Len())
        Sleep(500*Millisecond)
    }
}
