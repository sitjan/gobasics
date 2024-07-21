package main

import (
	"fmt"
)

type car struct {
	gas_pedal      uint16
	brake_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

func main() {
	a_car := car{gas_pedal: 2215, brake_pedal: 0, steering_wheel: 1022, top_speed_kmh: 29.21}

	fmt.Println(a_car.gas_pedal)
}
