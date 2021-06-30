package controller
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/websocket"
	"fmt"
	"math/rand"
	"time"
	"encoding/json"
)

func Websocket(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	websocket.Handler(ReceiveSend).ServeHTTP(res,req)
  }
func ReceiveSend(ws *websocket.Conn) {
	//fmt.Printf("recvSendServer %#v\n", ws)
	for {
		var buf string
		err := websocket.Message.Receive(ws, &buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv:%q\n", buf)
		for {
			mapD := map[string]int{
				"Roll":rand.Intn(65535), 
				"Pitch":rand.Intn(65535),
				"Yaw":rand.Intn(65535),
				"Heave":rand.Intn(65535),
				"Surge":rand.Intn(65535),
				"Sway":rand.Intn(65535),
			}
			mapB, _ := json.Marshal(mapD)

			err = websocket.Message.Send(ws, string(mapB))
			if err != nil {
				fmt.Println(err)
				break
			}
			//fmt.Printf("send:%q\n", string(mapB))
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("recvSendServer finished")
}