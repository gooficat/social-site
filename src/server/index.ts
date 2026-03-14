import router from 'bun-modular-router'
import routes from './routes/routes'
import path from 'path'

const clientPath = path.join(import.meta.dirname, "../client")

Bun.serve({
	port: 8080,
	routes: router(routes),
	fetch(req)
	{
		let fpath = new URL(req.url).pathname
		if (fpath[fpath.length - 1] == '/') fpath += "index.html"
		return new Response(Bun.file(path.join(clientPath, fpath!)))
	}
})