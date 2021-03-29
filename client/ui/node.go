package ui

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type Node struct {
	app.Compo
}

func (c *Node) OnNav(ctx app.Context, url *url.URL) {

	fmt.Println("Node: ", url)

	ep := url.EscapedPath()
	s := strings.Split(ep, "/")

	fmt.Println("S:", s)
	fmt.Println("ID:", s[len(s)-1])
	fmt.Println("ep:", ep)

	// get the info for node x
}

func (c *Node) Render() app.UI {
	return app.Div().
		Class("bg-purple-400").
		Body(
			app.Text("Node"),
		)
}
