package driver


import (
	"Matrix"
	"math"
	"C"
)

const N_FLOORS = 4

type elev_button_type_t int
const (
	BUTTON_CALL_UP elev_button_type_t = iota
	BUTTON_CALL_DOWN elev_button_type_t = iota
	BUTTON_COMMAND elev_button_type_t = iota
)

var lamp_channel_matrix Matrix
var button_channel_matrix Matrix
lamp_channel_matrix = NewMatrix(4,3)
button_channel_matrix = NewMatrix(4,3)

//TODO: lampematrise

func elev_init() int {
	if io_init() == 0{
		//feil
		return 0
	}
	initMatrices()
	for i := 0; i < N_FLOORS; i++{
		if ( i != 0){
			elev_set_button_lamp(BUTTON_CALL_DOWN, i, 0)
		}
		if (i != N_FLOORS-1){
			elev_set_button_lamp(BUTTON_CALL_UP, i, 0)
		}
	
		elev_set_button_lamp(BUTTON_COMMAND, i, 0)
	}
	
	elev_set_stop_lamp(0)
	elev_set_door_open_lamp(0)
	elev_set_floor_indicator(0)
	
	return 1
}

func elev_set_speed(speed int, last_speed int) int { //returns last speed
	if speed > 0 {
		io_clear_bit(MOTORDIR)
	} else if (speed < 0){
		io_set_bit(MOTORDIR)
	} else if (last_speed < 0){
		io_clear_bit(MOTORDIR)
	} else if (last_speed > 0){
		io_set_bit(MOTORDIR)
	}
	
	io_write_analog(MOTOR, 2048 + 4*int(math.Abs(float64(speed))))
	
	return speed
}
  
func elev_get_floor_sensor_signal() int {
	if (io_read_bit(SENSOR1) == 1){
		return 0
	} else if (io_read_bit(SENSOR2) == 1) {
		return 1
	} else if (io_read_bit(SENSOR3) == 1) {
		return 2
	} else if (io_read_bit(SENSOR4) == 1) {
		return 3
	} else {
		return -1
	}
}

func elev_get_button_signal(button elev_button_type_t,floor int) {
	assert(floor >= 0)
	assert(floor < N_FLOORS)
	assert(!(button == BUTTON_CALL_UP && floor == N_FLOORS-1))
	assert(!(button == BUTTON_CALL_DOWN && floor == 0))
	assert( button == BUTTON_CALL_UP || button == BUTTON_CALL_DOWN || button == BUTTON_COMMAND)

	if io_read_bit( button_channel_matrix.Get(floor, int(button)))   == 1 {
		return 1
	} else {
		return 0
	}
}

func elev_get_stop_signal() int {
	return io_read_bit(STOP)
}

func elev_get_obstruction_signal() int {
	return io_read_bit(OBSTRUCTION)
}

func elev_set_floor_indicator(floor int) {
	assert(floor >= 0)
	assert(floor < N_FLOORS)
	
	if floor & 0x02 {
		io_set_bit(FLOOR_IND1)
	} else {
		io_clear_bit(FLOOR_IND1)
	}
	
	if floor & 0x01 {
		io_set_bit(FLOOR_IND2)
	}
}

func elev_set_button_lamp(button elev_button_type_t, floor int, value int) {
	assert(floor >= 0)
	assert(floor < N_FLOORS)
	assert(!(button == BUTTON_CALL_UP && floor == N_FLOORS-1))
	assert(!(button == BUTTON_CALL_DOWN && floor == 0))
	assert( button == BUTTON_CALL_UP || button == BUTTON_CALL_DOWN || button == BUTTON_COMMAND)
	
	
	
}
func elev_set_stop_lamp(value int) {
	if value != 0 {
		io_set_bit(LIGHT_STOP)
	} else {
		io_clear_bit(LIGHT_STOP)
	}
}

func elev_set_door_open_lamp(int value) {
	if value != 0 {
		io_set_bit(DOOR_OPEN)
	} else {
		io_clear_bit(DOOR_OPEN)
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



/*
static const int lamp_channel_matrix[N_FLOORS][N_BUTTONS] = {
    {LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
    {LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
    {LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
    {LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
};
  
  

static const int button_channel_matrix[N_FLOORS][N_BUTTONS] = {
    {FLOOR_UP1, FLOOR_DOWN1, FLOOR_COMMAND1},
    {FLOOR_UP2, FLOOR_DOWN2, FLOOR_COMMAND2},
    {FLOOR_UP3, FLOOR_DOWN3, FLOOR_COMMAND3},
    {FLOOR_UP4, FLOOR_DOWN4, FLOOR_COMMAND4},
};

*/