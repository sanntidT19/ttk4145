package Driver

/*
#cgo LDFLAGS: -lcomedi -lm -std=c99
#include "io.h"

*/
import "C"

func IOInit() int {
  return int(C.io_init());
}

func IOSetBit(channel int){
  C.io_set_bit(C.int(channel))
}

func IOClearBit(channel int){
  C.io_clear_bit(C.int(channel))
}

func IOWriteAnalog(channel int, value int) {
  C.io_write_analog(C.int(channel), C.int(value))
}

func IOReadBit(channel int) int {
  return int(C.io_read_bit(C.int(channel)))
}

func IOReadAnalog(channel int) int {
  return int(C.io_read_analog(C.int(channel)))
}
