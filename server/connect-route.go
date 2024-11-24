package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
)

// the game session:
type player struct {
	X  int
	Y  int
	Id int
}

var last_id = 0

var entities []player

func checkOrigin(r *http.Request) bool {
	return true
}

var upgrader = websocket.Upgrader{
	CheckOrigin: checkOrigin,
}

var eternal chan (bool)

func Connect(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Something tried connecting!")

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		panic(err)
	}
	// defer conn.Close()

	player_id := last_id
	last_id += 1
	entities = append(entities, player{Id: player_id})
	type control struct {
		X int
		Y int
	}
	var controlls control

	//set up listener for the connection
	go func() {
		for {
			_, content, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Something went wrong: ", err)
				break
			}

			json.Unmarshal(content, &controlls)
		}
	}()

	// updates the player every second
	go func() {
		for {
			time.Sleep(50 * time.Millisecond)

			if controlls.X == 1 {
				entities[player_id].X += 6
			}
			if controlls.X == -1 {
				entities[player_id].X -= 6
			}
			if controlls.Y == 1 {
				entities[player_id].Y += 6
			}
			if controlls.Y == -1 {
				entities[player_id].Y -= 6
			}

			// Works very well =)
			err := conn.WriteJSON(entities)
			if err != nil {
				fmt.Println("Something went wrong: ", err)
				break
			}
		}
	}()

	<-eternal
}
