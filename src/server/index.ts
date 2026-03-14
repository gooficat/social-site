import router from 'bun-modular-router'

Bun.serve({
	port: 8080,
	routes: router({
		"/": () => new Response("Index"),
		"/greet": () => new Response("Hello"),
		"/meow": () => new Response("Meow"),
	}),
})