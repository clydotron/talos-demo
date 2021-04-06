package ws

import (
	"context"
	"fmt"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

//@todo add a callback or the event bus to handle incoming data
type TestClientX struct {
}

func (tcx *TestClientX) Open(addr string) {

	fmt.Println("TCX.Open:", addr)

	//ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	//defer cancel()
	ctx := context.Background()

	di := websocket.DialOptions{
		Subprotocols: []string{"echo"},
	}
	c, _, err := websocket.Dial(ctx, addr, &di)
	if err != nil {
		// ...
		fmt.Println("Failed to open websocket")
		return
	}

	// send a test message (remove later)
	err = wsjson.Write(ctx, c, "greetings")
	if err != nil {
		fmt.Println(err)
		return
	}
	// @todo design mechanism for client to request data...

	// infinite read loop:
	// will exit when the websocket connection is closed
	go func() {
		defer fmt.Println("TCX.reader >done<")

		//ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
		//defer cancel()

		for {
			//wait?

			var v interface{}
			err := wsjson.Read(ctx, c, &v)
			if err != nil {
				// check for a clean close
				if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
					fmt.Println("websocket is closed")
					return
				}
				fmt.Println("ERROR", err)
				return
			}
			//log.Printf("received: %v", v)
			//@todo do something with the data

			if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
				return
			}
		}
	}()
}

func (tcx *TestClientX) Close() {

	// do something (eventually)

	// tcx.done <- true
	fmt.Println("TCX -- Close")
}
