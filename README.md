Forked from https://github.com/gmcbay/callouts, all credit to them for the idea

# callouts
Vow of the Disciple callout helper for console, with modifications for Trash Panda "unique" callouts.

Runs on Raspberry Pi devices that support OTG-USB on /dev/hidg0.  

Connect Raspberry Pi's USB-OTG port to console USB port to act as a USB keyboard.

clone this repo to /usr/local/bin, and set up a cronjob to "@reboot sudo go run /usr/local/bin/callouts.go&"

Run callouts server on Raspberry Pi, connect to server from any browser 
(PC/mobile/tablet) to get web UI to select icons to transmit callouts over text chat without having to type them in.



Developed and tested using a Raspberry Pi Zero W and Xbox Series X using a standard USB cable.



