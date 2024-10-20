package modbus

import (
	"log"

	"github.com/goburrow/modbus"
)

const (
	ON  = 0xFF00
	OFF = 0x0000
)

type MR6CUV2 struct {
	SlaveId byte
}

func (r MR6CUV2) ReadCoilsStates(client modbus.Client, handler *modbus.RTUClientHandler) [6]bool {
	handler.SlaveId = r.SlaveId
	coils, err := client.ReadCoils(0, 6)
	if err != nil {
		handler.Logger.Println(err)
	}

	dst := make([]bool, 0)
	for _, v := range coils {
		for i := 7; i >= 2; i-- {
			move := uint(7 - i)
			dst = append(dst, ((int((v >> move) & 1)) == 1))
		}
	}
	return [6]bool(dst)
}

func (r MR6CUV2) WriteSingleCoil(coilIndex byte, client modbus.Client, handler *modbus.RTUClientHandler, state bool) {
	handler.SlaveId = r.SlaveId
	value := OFF
	if state {
		value = ON
	}
	_, err := client.WriteSingleCoil(uint16(coilIndex), uint16(value))
	if err != nil {
		handler.Logger.Println(err)
	}
}

func (r MR6CUV2) WriteMultipleCoils(client modbus.Client, handler *modbus.RTUClientHandler, values [6]bool) {
	handler.SlaveId = r.SlaveId

	var value byte = 0
	for i, v := range values {
		if v {
			value = value | (1 << (0 + i))
		}
	}

	_, err := client.WriteMultipleCoils(0, 6, []byte{value})
	if err != nil {
		log.Fatal(err)
	}
}
