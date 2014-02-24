package driver

func elev_init() int {
  if !io_init() {
    //feil
    return 0
  }
  
  for ( i := 0; i < N_FLOORS; i++){
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

func elev_set_Speed(int speed, int lastSpeed) int { //returns last speed
  if (speed > 0){
    io_clear_bit(MOTORDIR)
  } else if (speed < 0){
    io_set_bit(MOTORDIR)
  } else if (last_speed < 0){
    io_clear_bit(MOTORDIR)
  } else if (last_speed > 0){
    io_set_bit(MOTORDIR)
  }
  
  io_write_analog(MOTOR, 2048 + 4*abs(speed))
  
  return speed
}
  
  
func elev_get_floor_sensor_signal() int {
  if (io_read_bit(SENSOR1)){
    return 0
  } else if (io_read_bit(SENSOR2)) {
    return 1
  } else if (io_read_bit(SENSOR3)) {
    return 2
  } else if (io_read_bit(SENSOR4)) {
    reurn 3
  } else {
    return -1
  }
}

func elev_get_button_signal(
  
