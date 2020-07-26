#!/bin/bash
set -exuo pipefail

exec avrdude \
	-c arduino \
	-b 19200 \
	-p t85 \
	-U flash:w:t85_default.hex \
	-U lfuse:w:0xe1:m \
	-U hfuse:w:0xdd:m \
	-U efuse:w:0xfe:m \
	-P /dev/ttyUSB0
