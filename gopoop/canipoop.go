package gopoop

import (
	"sync"
	"fmt"
	"log"
	"strconv"
	"encoding/json"

	"github.com/ereyes01/firebase"
)

type RoomInfo struct {
	Location   string `json:"location,omitempty"`
	IsOpen     ConvertibleBoolean      `json:"isOpen"`
	LastChange uint   `json:"lastChange,omitempty"`
	LastUpdate uint   `json:"lastUpdate,omitempty"`
}

type RoomInfoCallback func(room string, roomName string, roomInfo *RoomInfo)

type CanIPoop struct {
	configuration Configuration
}

func NewCanIPoop(configuration Configuration) *CanIPoop {
	return &CanIPoop{configuration: configuration}
}

func (p *CanIPoop) Process(roomInfoCallback RoomInfoCallback) {
	var wg sync.WaitGroup

	for _, room := range p.configuration.GetRooms() {
		p.watchRoom(room, &wg, roomInfoCallback)
	}

	wg.Wait()
	fmt.Println(`Done`)

}

func (p *CanIPoop) watchRoom(room string, wg *sync.WaitGroup, roomInfoCallback RoomInfoCallback) {
	roomName := room

	wg.Add(1)
	go func() {
		stop := make(chan bool)
		client := p.getFirebaseClient()
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

			roomInfoCallback(room, roomName, roomInfo)
		}

		fmt.Printf("Notifications have stopped")
		wg.Done()
	}()
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

func logEventError(event firebase.StreamEvent) {
	if event.Error != nil {
		log.Println(`Stream error:`, event.Error)
	} else if event.UnmarshallerError != nil {
		log.Println(`Malformed event:`, event.UnmarshallerError)
	}
}

func (p *CanIPoop) getFirebaseClient() firebase.Client {
	url := fmt.Sprintf(`https://%s.firebaseio.com`, p.configuration.storageBucket)
	return firebase.NewClient(url, "", nil)
}
