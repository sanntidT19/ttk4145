package main

import (
	."Driver"
	"fmt"
	."time"
	list "container/list"
)

type Elevator struct {
	stopList *list.List
	speed int
	direction int
	currentFloor int
}

func main() {
	e := InitElevator()
	
	

      
	for  {
		e.UpdateList()
		Sleep(4*Millisecond)
		
		if ElevGetStopSignal() == 1 {
			break
		}
	}

	
}

func InitElevator() *Elevator {
		
	if ElevInit() == 0 {
		fmt.Println("I am Error")
	}

	e := new(Elevator)	
	e.stopList = list.New()
	e.direction = 0
	
	if ElevGetFloorSensorSignal() == -1 {
		// goes down to nearest floor if between twoo floors
		e.speed = ElevSetSpeed(-300,0);
		for ElevGetFloorSensorSignal() == -1 {
			Sleep(10*Millisecond)
		}
	}
	
	e.speed = ElevSetSpeed(0,-300);	
	e.currentFloor = ElevGetFloorSensorSignal()
	
	return e
}

func (e *Elevator) GetNextFloor() int {
	if e.stopList.Front() != nil {
		return e.stopList.Front().Value.(int)
	} else {
		return -1
	}
}

func (e *Elevator) UpdateList() {
	
	l := e.stopList
	
	for floor := 0; floor < N_FLOORS; floor++ {
		for button:= 0; button < N_BUTTONS; button++ {
			if ElevGetButtonSignal(button, floor) == 1 {
						
				if l.Len() == 0 {
					// insert at front if list is empty
					_ = l.PushFront(floor)
					break
				}
				
				if Contains(l,floor) {
					break
				}
				 
				if e.direction == 1 {
					// insert before smallest element larger than floor
					for e := l.Front(); e != nil; e = e.Next() {
						if e.Value.(int) > floor {
							_ = l.InsertBefore(floor,e)
							break
						}
					}
				} else if e.direction == -1 {
					// insert before largest element smaller than floor
					for e := l.Front(); e != nil; e = e.Next() {
						if e.Value.(int) < floor {
							_ = l.InsertBefore(floor,e)
							break
						}
					}
				}
			} 
		}
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













