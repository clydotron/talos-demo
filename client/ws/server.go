package ws

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/clydotron/talos-demo/client/models"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type TestServerX struct {
	logf func(f string, v ...interface{})
	conn *websocket.Conn
}

//@todo add the sub or
func NewTestServerX() *TestServerX {
	tsx := &TestServerX{
		logf: log.Printf,
	}
	return tsx
}

func (tsx *TestServerX) ServeHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serveHttp - start")
	defer fmt.Println("serveHttp - end")

	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols: []string{"echo"},
	})
	if err != nil {
		tsx.logf("error: websocket.Accept failed: %s", err.Error())
		return
	}

	if c.Subprotocol() != "echo" {
		c.Close(websocket.StatusPolicyViolation, "client must speak the same subprotocol")
		return
	}

	tsx.conn = c
	go tsx.reader(r)
	tsx.writer(r)

	c.Close(websocket.StatusNormalClosure, "")
}

func (tsx *TestServerX) Close() {
	//fmt.Println("TSX_close")
	if tsx.conn != nil {
		tsx.conn.Close(websocket.StatusNormalClosure, "")
	}
}

func (tsx *TestServerX) reader(r *http.Request) {
	//fmt.Println("tsxest.reader >> start")
	//defer fmt.Println("tsxest.reader >> end")

	//ctx, cancel := context.WithTimeout(r.Context(), time.Second*60)
	//defer cancel()
	ctx := context.Background()

	for {
		//wait?

		var v interface{}
		err := wsjson.Read(ctx, tsx.conn, &v)
		if err != nil {
			if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
				tsx.logf("socket closed")
				//signal that we are done?
				return
			}
			tsx.logf("error: %s", err.Error())
			return
		}

		log.Printf("received: %v", v)
	}
}

func (tsx *TestServerX) writer(r *http.Request) {

	//fmt.Println("tsxest.writer - START")
	//defer fmt.Println("tsxest.writer - END")

	// need a context?
	ctx := context.Background()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// write a message every 2 seconds
	for {
		select {
		//case <-done:
		//	return
		case t := <-ticker.C:

			dd := models.TestWebsocketEvent{
				TS:      t,
				Payload: "test",
			}

			err := wsjson.Write(ctx, tsx.conn, dd)
			if err != nil {
				if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
					return
				}
				//log something? (do i actually care)
				return
			}
		}
	}
}
