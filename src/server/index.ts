import router from 'bun-modular-router'
import index from "../client/index.html"

Bun.serve({
	port: 8080,
	routes: router({
		"/": index,
		"/api": {
			"/greet": () => new Response("Hello from the API!")
		},
		"/animals": {
			"/feline": {
				"/adult": () => new Response("Mraow!"),
				"/kitten": () => new Response("mew!")
			},
			"/cow/:hasMilk": (req: Bun.BunRequest) =>
			{
				const hasMilk = req.params.hasMilk === "dairy"
				return new Response("moo " + (hasMilk ? "(with milk)" : ""))
			}
		}
	}),
})