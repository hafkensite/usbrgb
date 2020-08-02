package main

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gousb"
)

type input string

var colormapping = map[string]string{
	"black":   "000000",
	"silver":  "c0c0c0",
	"gray":    "808080",
	"white":   "ffffff",
	"maroon":  "800000",
	"red":     "ff0000",
	"purple":  "800080",
	"fuchsia": "ff00ff",
	"green":   "008000",
	"lime":    "00ff00",
	"olive":   "808000",
	"yellow":  "ffff00",
	"navy":    "000080",
	"blue":    "0000ff",
	"teal":    "008080",
	"aqua":    "00ffff"}

func getRGB(i input) ([]byte, error) {
	rgb := make([]byte, 3)
	x, ok := colormapping[string(i)]
	if ok {
		len, err := hex.Decode(rgb, []byte(x))
		if len != 3 {
			return nil, fmt.Errorf("Color code not 3 bytes: %v", x)
		}
		return rgb, err
	}

	len, err := hex.Decode(rgb, []byte(string(i)))
	if len != 3 {
		return nil, fmt.Errorf("Color code not 3 bytes: %v", i)
	}
	return rgb, err
}

func readNativeMessage() (*input, error) {
	lenbytes := make([]byte, 4)
	n, err := os.Stdin.Read(lenbytes)
	if err != nil {
		return nil, err
	}
	if n != 4 {
		return nil, fmt.Errorf("Did not read 4 bytes as len, got %d", n)
	}
	len := binary.LittleEndian.Uint32(lenbytes)

	msgbytes := make([]byte, len)
	n, err = os.Stdin.Read(msgbytes)
	if err != nil {
		return nil, err
	}
	if uint32(n) != len {
		return nil, fmt.Errorf("Did not read %d bytes as %d", n, len)
	}

	var i input
	err = json.Unmarshal(msgbytes, &i)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func setColor(device *gousb.Device, color []byte) error {
	var err error
	for x := 0; x < 100; x++ {
		_, err := device.Control(gousb.ControlVendor|gousb.ControlDevice|gousb.ControlOut, 0, uint16(color[1])*256+uint16(color[0]), uint16(color[2]), []byte{})
		if err == nil {
			// log.Println("device.Control")
			// log.Print(err)
			break
		}
		time.Sleep(time.Millisecond * 1)
	}
	if err != nil {
		return err
	}

	for x := 0; x < 100; x++ {
		_, err := device.Control(gousb.ControlVendor|gousb.ControlDevice|gousb.ControlOut, 1, uint16(color[1])*256+uint16(color[0]), uint16(color[2]), []byte{})
		if err == nil {
			// log.Println("device.Control")
			// log.Print(err)
			break
		}
		time.Sleep(time.Millisecond * 1)
	}
	return err
}

func main() {

	f, err := os.OpenFile("/tmp/testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(f)
	ctx := gousb.NewContext()
	defer ctx.Close()

	device, err := ctx.OpenDeviceWithVIDPID(0x16c0, 0x05df)
	if err != nil {
		// f.WriteString(err.Error())
		log.Fatal(err)
		return
	}
	if device == nil {
		log.Fatal("Device not found")
		return
	}

	for i := 0; i < 255; i += 8 {
		rgb := make([]byte, 3)
		rgb[0] = byte(i)
		rgb[1] = 0
		rgb[2] = 0
		err = setColor(device, rgb)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		time.Sleep(time.Millisecond * 10)
	}

	for i := 255; i >= 0; i -= 8 {
		rgb := make([]byte, 3)
		rgb[0] = byte(i)
		rgb[1] = 0
		rgb[2] = 0
		err = setColor(device, rgb)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		time.Sleep(time.Millisecond * 10)
	}

	rgb := make([]byte, 3)
	rgb[0] = 0
	rgb[1] = 0
	rgb[2] = 0
	err = setColor(device, rgb)

	for true {
		i, err := readNativeMessage()
		if err != nil {
			log.Print(err.Error())
			return
		}

		rgb, err := getRGB(*i)
		if err != nil {
			log.Print(err.Error())
			return
		}

		err = setColor(device, rgb)
		if err != nil {
			log.Print(err.Error())
			return
		}
	}
}

