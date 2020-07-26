#!/bin/bash
set -euo pipefail
c=""
if [ $# -eq 0 ]; then
	c="ff00ff"
else
	case $1 in
		black)
			;&
		off)
			c="000000"
			;;
		silver)
			c="c0c0c0"
			;;
		gray)
			;&
		grey)
			c="808080"
			;;
		white)
			c="ffffff"
			;;
		maroon)
			c="800000"
			;;
		red)
			c="ff0000"
			;;
		purple)
			c="800080"
			;;
		fuchsia)
			c="ff00ff"
			;;
		green)
			c="008000"
			;;
		lime)
			c="00ff00"
			;;
		olive)
			c="808000"
			;;
		yellow)
			c="ffff00"
			;;
		navy)
			c="000080"
			;;
		blue)
			c="0000ff"
			;;
		teal)
			c="008080"
			;;
		aqua)
			c="00ffff"
			;;
		duo)
			c="00ffffffff00"
			;;
		*)
			if [ ${#1} == 6 ]; then
				c=$1
			else
				echo "Invalid length" >&2
				exit -1
			fi
	esac
fi

for dev in /dev/serial/by-id/usb-digistump.com_Digispark_Serial-*; do
	echo $c | xxd -r -p > $dev
done

