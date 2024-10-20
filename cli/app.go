package main

import (
	"fmt"
	"log"
	"time"

	md "bless.rng/smarthome/device/modbus"
	"github.com/goburrow/modbus"
)

func main() {
	handler := modbus.NewRTUClientHandler("COM3")
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 2
	handler.Timeout = 5 * time.Second

	err := handler.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer handler.Close()

	client := modbus.NewClient(handler)

	d := md.MR6CUV2{SlaveId: 78}
	d.WriteMultipleCoils(client, handler, [6]bool{})
	states := d.ReadCoilsStates(client, handler)

	fmt.Println(states)
}
