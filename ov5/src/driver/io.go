package driver

/*
#cgo LDFLAGS: -lcomedi -lm
#include "io.h"
typedef int h
*/
import "C"

func io_init() int {
  return int(C.io_init());
}

func io_set_bit(int channel){
  C.io_set_bit(C.h(channel))
}

func io_clear_bit(int channel){
  C.io_clear_bit(C.h(int channel))
}

func io_write_analog(int channel, int value) {
  C.io_write_analog(C.int(channel), C.int(value))
}

func io_read_bit(int channel) {
  C.io_read_bit(C.h(channel))
}

func io_read_analog(int channel) {
  C.io_read_analog(C.int(channel))
}
