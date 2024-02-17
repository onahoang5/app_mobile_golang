package skiuser

import (
	"Food-delivery/common"
	"log"

	socketio "github.com/googollee/go-socket.io"
	"gorm.io/gorm"
)

type SmallAppcontext interface {
	GetMainDBConnection() *gorm.DB
}

type LocationData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func OnUserUpdateLocation(appCtx SmallAppcontext, requester common.Requester) func(s socketio.Conn, location LocationData) {
	return func(s socketio.Conn, location LocationData) {
		log.Println("User", requester.GetUserId(), "update location", location)
	}
}
