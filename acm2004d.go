package acm2004d

import (
	"fmt"
	"time"

	"github.com/d2r2/go-i2c"
)

type LCD struct {
	I2C *i2c.I2C
}

func InitLcd(addr uint8, bus int) (*LCD, error) {
	i2c, err := i2c.NewI2C(addr, bus)
	if err != nil {
		i2c.Close()
		return nil, err
	}
	// Func Set
	time.Sleep(20 * time.Millisecond)
	_, err = i2c.WriteBytes([]byte{0x00, 0x38})
	if err != nil {
		i2c.Close()
		return nil, err
	}
	// Clear Display
	time.Sleep(20 * time.Millisecond)
	_, err = i2c.WriteBytes([]byte{0x00, 0x01})
	if err != nil {
		i2c.Close()
		return nil, err
	}
	// Return Home
	time.Sleep(20 * time.Millisecond)
	_, err = i2c.WriteBytes([]byte{0x00, 0x02})
	if err != nil {
		i2c.Close()
		return nil, err
	}
	// Display On
	time.Sleep(20 * time.Millisecond)
	_, err = i2c.WriteBytes([]byte{0x00, 0x0C})
	if err != nil {
		i2c.Close()
		return nil, err
	}
	time.Sleep(20 * time.Millisecond)
	res := &LCD{
		I2C: i2c,
	}
	return res, nil
}

func (LCD *LCD) Write(line int, data [20]byte) error {
	var err error
	// Set lines
	switch line {
	case 1:
		_, err = LCD.I2C.WriteBytes([]byte{0x80, 0x00})
	case 2:
		_, err = LCD.I2C.WriteBytes([]byte{0x80, 0x40})
	case 3:
		_, err = LCD.I2C.WriteBytes([]byte{0x80, 0x14})
	case 4:
		_, err = LCD.I2C.WriteBytes([]byte{0x80, 0x54})
	default:
		return fmt.Errorf("Error: Undefined Line Number")
	}
	if err != nil {
		return err
	}
	for i := 0; i < 20; i++ {
		_, err = LCD.I2C.WriteBytes([]byte{0x40, data[i]})
		if err != nil {
			return err
		}
	}
	return nil
}

func (LCD *LCD) Close() error {
	return LCD.I2C.Close()
}
