# usbrgb

Attiny85 based usb pcb that provides two rgb led's, to be used as indicators.

# Hardware

The pcb design is based upon that of the [Digispark USB development board](http://digistump.com/products/1). For which they have released the [Eagle](https://s3.amazonaws.com/digistump-resources/files/3a5187f5_digispark_sources.zip) files.

The Digispark board uses an attiny84 microprocessor, which (with the help of some special software), can be programmed over USB once the bootloader is flashed once.

The width of the original Digispark pcb is slightly too wide to fit in some usb ports, when they are located at an edge. A fairly large part of the pcb is occupied by a voltage regulator. As the usbrgb PCB is to be used in a computer, there is no need for a 5Volt regulator.

Removing the regulator, and a few other non-critical components allowed the board to be narrower. I chose for a 'bigger' package for the passive components, as I want to solder it to manually.

## Layout

[! PCB](./eagle-pcb/top-pcb.png)

[! Schematic](./eagle-pcb/schematic.png)

## Components

| Qty | Value      | Package | Parts  | Description                     |
| --- | ---------- | ------- | ------ | ------------------------------- |
| 1   | ATTINY85   | SOP8    | IC1    | Atmel ATTINY85-20SU             |
| 2   | 68R        | 1206    | R1, R2 | Resistor                        |
| 1   | 1K5        | 1206    | R3     | Resistor                        |
| 2   | 3.3v Zener | SOD-123 | D1, D2 | Diode                           |
| 1   | 4.7uf      | C1206   | C2     | Capacitor                       |
| 2   | WS2812B    | WS2812B | D3, D4 | WS2812B SMD addressable RGB LED |

# Embedded software

## Micronucleus bootloader
The [micronucleus](https://github.com/micronucleus/micronucleus) bootloader allows the Arduino IDE to program the chip over USB.

When the chip boots, it first presents itself as a usb device that can be programmed. After a short timeout the bootloader starts the Arduino code.

To flash the micronucleus bootloader you can use [avrdude](https://www.nongnu.org/avrdude/). I have programmed mine using a arduino uno with the ISP sketch written to it.

## Arduino

To program the device using Arduino you can follow the steps describe in the [tutorial](http://digistump.com/wiki/digispark/tutorials/connecting).

The provided application used the DigisparkCDC library to emulate a serial device.

> Note that the latest micronucleus version currently available (v2.04) is not supported by the arduino plugin. I replaced the micronucleus application with the [latest version](https://github.com/micronucleus/micronucleus/tree/master/commandline), which works fine.

# Host software

The provided arduino code acts as a serial interface. When 3 bytes are sent, they will be read as RGB value and both led's will be set to that color. If 6 bytes are sent, the first 3 are for the top LED, the second 3 for the bottom.

## Shell script
The provided `colors.sh` shows how to set a color. An other option is to send some random values using `head -c3 /dev/urandom > /dev/ttyACM0`

## Native messaging
The goal of this project is to be able to give a notification from within a browser. One of the ways to communicate from a browser to the host is using [native messaging](https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions/Native_messaging). This allows extensions to send messages to an executable on the host.

This project contains a small GO application that can read the messages from the browser, and send the 3 byte values to the serial interface.

## Browser integration

