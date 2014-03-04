package Driver


import (
	."Matrix"
	"math"
	"C"
	"log"
)

const N_FLOORS = 4
const N_BUTTONS = 3

const (
	BUTTON_CALL_UP int = iota
	BUTTON_CALL_DOWN int = iota
	BUTTON_COMMAND int = iota
)

var lamp_channel_matrix Matrix
var button_channel_matrix Matrix

func ElevInit() int {

	lamp_channel_matrix = NewMatrix(4,3)
	button_channel_matrix = NewMatrix(4,3)
	initMatrices()
	
	if IOInit() == 0{
		//feil
		return 0
	}
	
	for i := 0; i < N_FLOORS; i++{
		if ( i != 0){
			ElevSetButtonLamp(BUTTON_CALL_DOWN, i, 0)
		}
		if (i != N_FLOORS-1){
			ElevSetButtonLamp(BUTTON_CALL_UP, i, 0)
		}
	
		ElevSetButtonLamp(BUTTON_COMMAND, i, 0)
	}
	
	ElevSetStopLamp(0)
	ElevSetDoorOpenLamp(0)
	ElevSetFloorIndicator(0)
	
	return 1
}

func ElevSetSpeed(speed int, lastSpeed int) int { //returns last speed
	if speed > 0 {
		IOClearBit(MOTORDIR)
	} else if (speed < 0){
		IOSetBit(MOTORDIR)
	} else if (lastSpeed < 0){
		IOClearBit(MOTORDIR)
	} else if (lastSpeed > 0){
		IOSetBit(MOTORDIR)
	}
	
	IOWriteAnalog(MOTOR, 2048 + 4*int(math.Abs(float64(speed))))
	
	return speed
}
  
func ElevGetFloorSensorSignal() int {
	if (IOReadBit(SENSOR1) == 1){
		return 0
	} else if (IOReadBit(SENSOR2) == 1) {
		return 1
	} else if (IOReadBit(SENSOR3) == 1) {
		return 2
	} else if (IOReadBit(SENSOR4) == 1) {
		return 3
	} else {
		return -1
	}
}

func ElevGetButtonSignal(button int,floor int) int {
	assert(floor >= 0)
	assert(floor < N_FLOORS)
//	assert(!(button == BUTTON_CALL_UP && floor == N_FLOORS-1))
//	assert(!(button == BUTTON_CALL_DOWN && floor == 0))
	assert( button == BUTTON_CALL_UP || button == BUTTON_CALL_DOWN || button == BUTTON_COMMAND)

	if (button == BUTTON_CALL_UP && floor == N_FLOORS-1) || (button == BUTTON_CALL_DOWN && floor == 0) {
		return 0
	}

	if IOReadBit( button_channel_matrix.Get(floor, button)) == 1 {
		return 1
	} else {
		return 0
	}
}

func ElevGetStopSignal() int {
	return IOReadBit(STOP)
}

func ElevGetObstructionSignal() int {
	return IOReadBit(OBSTRUCTION)
}

func ElevSetFloorIndicator(floor int) {
	assert(floor >= 0)
	assert(floor < N_FLOORS)
	
	if (floor & 0x02) != 0 {
		IOSetBit(FLOOR_IND1)
	} else {
		IOClearBit(FLOOR_IND1)
	}
	
	if (floor & 0x01) != 0 {
		IOSetBit(FLOOR_IND2)
	}
}

func ElevSetButtonLamp(button int, floor int, value int) {
	assert(floor >= 0)
	assert(floor < N_FLOORS)
	assert(!(button == BUTTON_CALL_UP && floor == N_FLOORS-1))
	assert(!(button == BUTTON_CALL_DOWN && floor == 0))
	assert( button == BUTTON_CALL_UP || button == BUTTON_CALL_DOWN || button == BUTTON_COMMAND)
	
    	if (value == 1){
      	IOSetBit(lamp_channel_matrix.Get(floor,button));
    	} else {
      	IOClearBit(lamp_channel_matrix.Get(floor,button)); 
	}
}
func ElevSetStopLamp(value int) {
	if value != 0 {
		IOSetBit(LIGHT_STOP)
	} else {
		IOClearBit(LIGHT_STOP)
	}
}

func ElevSetDoorOpenLamp(value int) {
	if value != 0 {
		IOSetBit(DOOR_OPEN)
	} else {
		IOClearBit(DOOR_OPEN)
	}
}

func assert(t bool) {
	if !t {
		log.Panic("assertion failed!")
	}
}

func initMatrices() {
	lamp_channel_matrix.Set(0,0,LIGHT_UP1)
	lamp_channel_matrix.Set(0,1,LIGHT_DOWN1)
	lamp_channel_matrix.Set(0,1,LIGHT_COMMAND1)
	
	lamp_channel_matrix.Set(1,0,LIGHT_UP2)
	lamp_channel_matrix.Set(1,1,LIGHT_DOWN2)
	lamp_channel_matrix.Set(1,2, LIGHT_COMMAND2)
	
	lamp_channel_matrix.Set(2,0,LIGHT_UP3)
	lamp_channel_matrix.Set(2,1,LIGHT_DOWN3)
	lamp_channel_matrix.Set(2,2, LIGHT_COMMAND3)
	
	lamp_channel_matrix.Set(3,0,LIGHT_UP4)
	lamp_channel_matrix.Set(3,1,LIGHT_DOWN4)
	lamp_channel_matrix.Set(3,2,LIGHT_COMMAND4)
	
	
	button_channel_matrix.Set(0,0,FLOOR_UP1)
	button_channel_matrix.Set(0,1,FLOOR_DOWN1)
	button_channel_matrix.Set(0,1,FLOOR_COMMAND1)
	
	button_channel_matrix.Set(1,0,FLOOR_UP2)
	button_channel_matrix.Set(1,1,FLOOR_DOWN2)
	button_channel_matrix.Set(1,2, FLOOR_COMMAND2)
	
	button_channel_matrix.Set(2,0,FLOOR_UP3)
	button_channel_matrix.Set(2,1,FLOOR_DOWN3)
	button_channel_matrix.Set(2,2, FLOOR_COMMAND3)
	
	button_channel_matrix.Set(3,0,FLOOR_UP4)
	button_channel_matrix.Set(3,1,FLOOR_DOWN4)
	button_channel_matrix.Set(3,2,FLOOR_COMMAND4)
}
