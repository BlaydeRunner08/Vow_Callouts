# callouts
Vow of the Disciple callout helper for console

Runs on Raspberry Pi devices that support OTG-USB on /dev/hidg0.  

Connect Raspberry Pi's USB-OTG port to console USB port to act as a USB keyboard.

Run callouts server on Raspberry Pi, connect to server from any browser 
(PC/mobile/tablet) to get web UI to select icons to transmit callouts over text chat without having to type them in.

Text chat in console Destiny 2 is currently a bit buggy and will occasionally just drop messages without sending them 
(even though you see the message being input into the text chat box). This happens with real keyboards and human 
typists as well, I don't think there's anything that can be done to avoid the situation with this code.  

Various attempts to throttle input speed did not stop the problem from occuring.  Hopefully Bungie fixes this at some point.

Developed and tested using a Raspberry Pi Zero WH and Xbox Series X using this USB dongle for the RPI (power and USB connection to Xbox):

https://www.amazon.com/gp/product/B07NKNBZYG

I'm unlikely to maintain this until such a time as Raspberry Pi devices become more widely available.  For now I'm focusing on the solution using Adafruit Feather 32u4 Bluefruit LE.  See:

https://github.com/gmcbay/Transmat_32u4

https://github.com/gmcbay/Transmat_Android

