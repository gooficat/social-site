import db from "../../db"

async function login(req: Bun.BunRequest)
{
	const credentials = await req.body?.json()
	console.log(credentials)

	const storedUser = await db`
			SELECT * FROM users WHERE handle = ${credentials.handle}
			`

	if (storedUser.length === 0)
	{
		return new Response("User with that handle does not exist", { status: 400 })
	}

	if (credentials.password !== storedUser[0].password)
	{
		return new Response("Incorrect password", { status: 400 })
	}

	console.log(storedUser)

	return Response.json({
		session_id: "placeholder session id",
		user_id: storedUser[0].id
	})
}

async function register(req: Bun.BunRequest)
{
	const credentials = await req.body?.json()
	console.log(credentials)

	const storedUser = await db`
			SELECT * FROM users WHERE handle = ${credentials.handle}
			`

	if (storedUser.length !== 0)
	{
		return new Response("User with that handle already exists", { status: 400 })
	}

	await db`
		INSERT INTO users (handle, password) VALUES (${credentials.handle}, ${credentials.password})
	`

	const newUser = await db`
			SELECT * FROM users WHERE handle = ${credentials.handle}
			`

	return Response.json({
		session_id: "placeholder session id",
		user_id: newUser[0].id
	})
}

export default {
	"/login": {
		POST: login
	},
	"/register": {
		POST: register
	}
}


