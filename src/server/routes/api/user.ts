import type { BunRequest } from "bun"
import db from "../../db"
import { createSession, getSessions, isValidSession } from "./session"

async function login(req: Bun.BunRequest)
{
	const credentials = await req.body?.json()

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


	return Response.json({
		sessionId: await createSession(storedUser[0].id),
		userId: storedUser[0].id
	})
}

async function register(req: Bun.BunRequest)
{
	const credentials = await req.body?.json()

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
		sessionId: await createSession(newUser[0].id),
		userId: newUser[0].id
	})
}

async function validateSession(req: BunRequest)
{
	const body = await req.body?.json()

	const isValid = await isValidSession(body)

	if (isValid)
	{
		console.log("valid ", body)
		return new Response(null, { status: 200 })
	}
	console.log("not valid ", body)
	return new Response(null, { status: 400 })
}

export default {
	"/login": {
		POST: login
	},
	"/register": {
		POST: register
	},
	"/session-validate": {
		POST: validateSession
	}
}


