package ui

import (
	"github.com/clydotron/talos-demo/client/models"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// WorkerNode ...
type WorkerNode struct {
	app.Compo

	WNI models.WorkerNodeInfo
}

// Render ...
func (c *WorkerNode) Render() app.UI {
	return app.Div().
		Class("bg-white shadow p-3 rounded w-64").
		Body(
			app.Div().Class("text-left").Body(
				app.H3().Class("mb-2 text-gray-700").Text(c.WNI.Name).Body(
					app.P().Class("text-grey-600 text-sm").Text(c.WNI.Status),
				),
			),
			app.Div().Class("mt-4").Body(
				app.A().Class("no-underline mr-4 text-blue-500 hover:text-blue-400").Href("#").Text("Words?"),
			),
		)
}

/*
<div class="shadow p-4 bg-white">
    <div class="text-left">
        <h3 class="mb-2 text-gray-700">Card Title</h3>
        <p class="text-grey-600 text-sm">Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor.
            Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus
            mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. </p>
    </div>
​
    <div class="mt-4">
        <a href="#" class="no-underline mr-4 text-blue-500 hover:text-blue-400">Link 1</a>
        <a href="#" class="no-underline mr-4 text-blue-500 hover:text-blue-400">Link 2</a>
    </div>
</div>
*/
/*
<div class="bg-white shadow p-3 rounded lg:w-64">
  <div>
    <div style="background-image: url('')"
         class="bg-cover bg-center bg-gray-300 h-32 rounded">
    </div>
  </div>
  <div class="mt-6">
    <p class="text-lg text-bold tracking-wide text-gray-600 mb-2">
      Title
    </p>
    <p class="text-sm text-gray-600 font-hairline">
      Lorem ipsum dolor sit amet, consectetur adipiscing elit.
    </p>
  </div>
  <div class="mt-6">
    <button class="rounded shadow-md flex items-center shadow bg-blue-500 px-4 py-2 text-white hover:bg-blue-600">
      Default
    </button>
  </div>
</div>

*/

/*
<i class="fas fa-server"></i>
<!-- Column -->
<div class="my-1 px-1 w-full md:w-1/2 lg:my-4 lg:px-4 lg:w-1/3">

		<!-- Article -->
		<article class="overflow-hidden rounded-lg shadow-lg">

				<a href="#">
						<img alt="Placeholder" class="block h-auto w-full" src="https://picsum.photos/600/400/?random">
				</a>

				<header class="flex items-center justify-between leading-tight p-2 md:p-4">
						<h1 class="text-lg">
								<a class="no-underline hover:underline text-black" href="#">
										Article Title 4
								</a>
						</h1>
						<p class="text-grey-darker text-sm">
								11/1/19
						</p>
				</header>

				<footer class="flex items-center justify-between leading-none p-2 md:p-4">
						<a class="flex items-center no-underline hover:underline text-black" href="#">
								<img alt="Placeholder" class="block rounded-full" src="https://picsum.photos/32/32/?random">
								<p class="ml-2 text-sm">
										Author Name
								</p>
						</a>
						<a class="no-underline text-grey-darker hover:text-red-dark" href="#">
								<span class="hidden">Like</span>
								<i class="fa fa-heart"></i>
						</a>
				</footer>

		</article>
		<!-- END Article -->

</div>
<!-- END Column -->
*/
