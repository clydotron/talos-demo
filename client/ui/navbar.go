package ui

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// NavBar ...
type NavBar struct {
	app.Compo

	upx UpdateButton
}

// Render ...
func (c *NavBar) Render() app.UI {
	return app.Nav().Class("bg-indigo-200 fixed w-screen").
		Body(
			app.Div().Class("px-6 mx-auto flex justify-between").
				Body(
					app.Div().Class("flex space-x-2").
						Body(
							app.A().Class("py-5 px-3 text-gray-700").Href("/clusters").Text("Clusters"),
							//app.A().Class("py-5 px-3 text-gray-700").Href("/nodes").Text("Nodes"),
							app.A().Class("py-5 px-3 text-gray-700").Href("/machines").Text("Machines"),
							app.A().Class("py-5 px-3 text-gray-700").Href("/events").Text("Events"),
							app.A().Class("py-5 px-3 text-gray-700").Href("/charts").Text("Charts"),
						),
					app.Div().Class("flex items-center space-x-1").
						Body(
							&c.upx, //right justify this
						),
				),
		)
}

/*
<!-- Navbar goes here -->
<nav class="bg-purple-100">
  <div class="px-8 mx-auto">
    <div class="flex justify-between">

      <!-- logo -->
      <div class="flex items-centered">
        <a href="#" class="flex items-center py-5 px-3 text-gray-700">
                <svg class="w-6 h-6 mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15.546c-.523 0-1.046.151-1.5.454a2.704 2.704 0 01-3 0 2.704 2.704 0 00-3 0 2.704 2.704 0 01-3 0 2.704 2.704 0 00-3 0 2.704 2.704 0 01-3 0 2.701 2.701 0 00-1.5-.454M9 6v2m3-2v2m3-2v2M9 3h.01M12 3h.01M15 3h.01M21 21v-7a2 2 0 00-2-2H5a2 2 0 00-2 2v7h18zm-3-9v-2a2 2 0 00-2-2H8a2 2 0 00-2 2v2h12z" />
</svg>
          <span class="font-bold">Fun!</span>
        </a>


              <!-- primary nav -->
        <div class="flex items-center space-x-1">
          <a href="#" class="py-5 px-3 text-gray-700">Features</a>
          <a href="#" class="py-5 px-3 text-gray-700">Pricing</a>
       </div>
              </div>
      <!-- secondary nav -->
      <div class="flex items-center space-x-1">
        <a href="#">Login</a>
        <a href="#" class="py-2 px-3 bg-yellow-400 text-yellow-900 rounded hover:bg-yellow-300 hover:text-yellow-800 transition duration-300">Signup</a>
      </div>

    </div>
  </div>
</nav>
*/
