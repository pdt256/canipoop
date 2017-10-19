package main

import (
	"fmt"

	"github.com/pdt256/canipoop/gopoop"
)

func main() {
	configuration := gopoop.GetFlagConfiguration()
	canIPoop := gopoop.NewCanIPoop(configuration)

	fmt.Printf("Are the rooms open?\n")

	canIPoop.Process(roomInfoCallback)
}

func roomInfoCallback(room string, roomName string, roomInfo *gopoop.RoomInfo) {
	fmt.Printf("%s: %t\n", roomName, roomInfo.IsOpen)
}
