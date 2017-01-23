package mrmiddle

import "time"

// write byte data to designated line
func (mm *MrMiddle) writeByte(y int, v byte) (err error) {
	addr, gpio := y2AddrAndGpio(y - 1)

	err = mm.e.I2cStart(addr)

	if checkError(err) {
		return
	}

	if gpio == GPIOA {
		v = byte2Row(v).reversed().toByte()
	}

	data := []byte{byte(gpio), v}
	err = mm.e.I2cWrite(addr, data)

	if checkError(err) {
		return
	}

	return
}

func (mm *MrMiddle) writeAllLow() (err error) {
	for y := 0; y < 8; y++ {
		err = mm.writeByte(y, 0x00)

		if checkError(err) {
			return
		}
	}

	return
}

// DriveCoil drives coils as given pole direction
func (mm *MrMiddle) DriveCoil(pd Pole) (err error) {
	switch pd {
	case N:
		err = mm.e.DigitalWrite(IN1, 1)

		if checkError(err) {
			return
		}

		err = mm.e.DigitalWrite(IN2, 0)

		if checkError(err) {
			return
		}
	case S:
		err = mm.e.DigitalWrite(IN1, 0)

		if checkError(err) {
			return
		}

		err = mm.e.DigitalWrite(IN2, 1)

		if checkError(err) {
			return
		}
	}

	return
}

// ReleaseCoil releases coils
func (mm *MrMiddle) ReleaseCoil() (err error) {
	err = mm.e.DigitalWrite(IN1, 0)

	if checkError(err) {
		return
	}

	err = mm.e.DigitalWrite(IN2, 0)

	if checkError(err) {
		return
	}

	return
}

// HighWhile make (x, y) to High while ms[msec]
func (mm *MrMiddle) highWhile(x int, y int, ms time.Duration, pd Pole) (err error) {
	bits := byte(0x01 << uint(x-1))

	err = mm.writeByte(y, bits)

	if checkError(err) {
		return
	}

	err = mm.DriveCoil(pd)

	if checkError(err) {
		return
	}

	time.Sleep(ms)

	err = mm.ReleaseCoil()

	if checkError(err) {
		return
	}

	err = mm.writeByte(y, 0x00)

	if checkError(err) {
		return
	}

	return
}

// Flip flips a stone at (x, y)
func (mm *MrMiddle) Flip(x int, y int, pd Pole) (err error) {
	err = mm.highWhile(x, y, FLIPTIME, pd)

	if checkError(err) {
		return wrapError(err)
	}

	return
}
