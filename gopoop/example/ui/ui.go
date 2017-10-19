package main

import (
	"github.com/pdt256/canipoop/gopoop"
	"github.com/gizak/termui"
)

var rooms map[string]*termui.Par

func main() {
	configuration := gopoop.GetFlagConfiguration()
	canIPoop := gopoop.NewCanIPoop(configuration)

	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	p := termui.NewPar("Are the rooms open?")
	p.Height = 3
	p.Width = 50
	p.TextFgColor = termui.ColorWhite
	p.BorderFg = termui.ColorCyan
	termui.Render(p)

	rooms = make(map[string]*termui.Par)

	y := 4

	for _, roomId := range configuration.GetRooms() {
		rooms[roomId] = termui.NewPar(``)
		rooms[roomId].Height = 3
		rooms[roomId].Width = 20
		rooms[roomId].TextFgColor = termui.ColorWhite
		rooms[roomId].BorderLabel = roomId
		rooms[roomId].BorderFg = termui.ColorCyan
		rooms[roomId].Y = y
		termui.Render(rooms[roomId])
		y += rooms[roomId].Height
	}

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		canIPoop.Stop()
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/C-c", func(termui.Event) {
		canIPoop.Stop()
		termui.StopLoop()
	})

	go canIPoop.Process(roomInfoCallback)
	termui.Loop()
}

func roomInfoCallback(roomId string, roomName string, roomInfo *gopoop.RoomInfo) {
	var text string
	if roomInfo.IsOpen {
		text = `[                  ](fg-white,bg-green)`
	} else {
		text = `[                  ](fg-white,bg-red)`
	}
	rooms[roomId].BorderLabel = roomName
	rooms[roomId].Text = text
	termui.Render(rooms[roomId])
}
