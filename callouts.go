package main

import (
	"fmt"
	"log"
	"net/http"
	"syscall"
	"time"
)

const port = 8080

var keyboardFd int
var err error
var msgChan chan string

func sendKey(key byte) {
	modifier := byte(0x00)
	hidCode := byte(0x00)

	// 'a'-'z'
	if key >= 0x61 && key <= 0x7a {
		hidCode = key - 0x5d
	}

	// 'A'-'Z'
	if key >= 0x41 && key <= 0x5a {
		hidCode = key - 0x3d
		modifier = 0x20
	}

	// space
	if key == 0x20 {
		hidCode = 0x2c
	}

	// minus (used as a divider)
	if key == 0x2d {
		hidCode = 0x2d
	}

	if hidCode != 0 {
		syscall.Write(keyboardFd, []byte{modifier, 0, hidCode, 0, 0, 0, 0, 0})
		releaseKeys()
	}
}

func releaseKeys() {
	time.Sleep(10 * time.Millisecond)
	syscall.Write(keyboardFd, []byte{0, 0, 0, 0, 0, 0, 0, 0})
	time.Sleep(10 * time.Millisecond)
}

func pressKeyCode(keyCode byte) {
	syscall.Write(keyboardFd, []byte{0, 0, 0, keyCode, 0, 0, 0, 0})
	releaseKeys()
}

func pressEnter() {
	pressKeyCode(0x28)
}

func pressEsc() {
	pressKeyCode(0x29)
}

func getTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	msg, ok := r.URL.Query()["msg"]

	if !ok || len(msg[0]) < 1 {
		log.Println("request param 'msg' is missing")
		return
	}

	log.Println("Sending keyboard msg: " + string(msg[0]))

	msgChan <- string(msg[0])
}

func processMessages() {
	for {
		msg := <-msgChan
		byteData := []byte(msg)

		pressEnter()

		for _, key := range byteData {
			sendKey(key)
		}
		
		pressEnter()
		pressEsc()
	}
}

func main() {
	msgChan = make(chan string, 32)

	keyboardFd, err = syscall.Open("/dev/hidg0", syscall.O_RDWR, 06666)

	if err != nil {
		fmt.Print(err.Error(), "\n")
		return
	}

	go processMessages()

	fileServer := http.FileServer(http.Dir("./vow"))
	http.Handle("/", fileServer)
	http.HandleFunc("/send", sendHandler)

	fmt.Printf("callouts server starting on port %d\n", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
