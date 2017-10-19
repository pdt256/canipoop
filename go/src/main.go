package main

import (
	"fmt"
	"log"
	"sync"
	"encoding/json"

	"github.com/ereyes01/firebase"
	"strconv"
)

var configuration Configuration

type RoomInfo struct {
	Location   string `json:"location,omitempty"`
	IsOpen     ConvertibleBoolean      `json:"isOpen"`
	LastChange uint   `json:"lastChange,omitempty"`
	LastUpdate uint   `json:"lastUpdate,omitempty"`
}

func main() {
	configuration = getConfiguration()

	fmt.Printf("Are the rooms open?\n")

	var wg sync.WaitGroup

	for _, room := range configuration.GetRooms() {
		watchRoom(room, &wg)
	}

	wg.Wait()
	fmt.Println(`Done`)
}

func watchRoom(room string, wg *sync.WaitGroup) {
	roomName := room

	wg.Add(1)
	go func() {
		stop := make(chan bool)
		client := getFirebaseClient()
		events, err := client.Child(room).Watch(roomInfoParser, stop)
		if err != nil {
			log.Fatal(err)
		}

		for event := range events {
			if event.Error != nil || event.UnmarshallerError != nil {
				logEventError(event)
				continue
			}

			roomInfo := event.Resource.(*RoomInfo)
			if roomInfo.Location != `` {
				roomName = roomInfo.Location
			}
			fmt.Printf("%s: %t\n", roomName, roomInfo.IsOpen)
		}

		fmt.Printf("Notifications have stopped")
		wg.Done()
	}()
}

func logEventError(event firebase.StreamEvent) {
	if event.Error != nil {
		log.Println(`Stream error:`, event.Error)
	} else if event.UnmarshallerError != nil {
		log.Println(`Malformed event:`, event.UnmarshallerError)
	}
}

func roomInfoParser(path string, data []byte) (interface{}, error) {
	var roomInfo *RoomInfo
	//log.Println(string(data))

	// Handle firebase returning single digit instead of full object
	if len(string(data)) == 1 {
		isOpen, err := strconv.ParseBool(string(data))
		if err == nil {
			roomInfo = &RoomInfo{IsOpen: ConvertibleBoolean(isOpen)}
			return roomInfo, nil
		}
	}

	err := json.Unmarshal(data, &roomInfo)
	return roomInfo, err
}

func getFirebaseClient() firebase.Client {
	url := fmt.Sprintf(`https://%s.firebaseio.com`, configuration.storageBucket)
	return firebase.NewClient(url, "", nil)
}
