import { randomUUIDv7 } from "bun"
import db from "../../db"

async function createSession(userId: number)
{
	const sessionId = randomUUIDv7()

	await db`INSERT INTO sessions (id, userId) VALUES (${sessionId}, ${userId})`

	return sessionId
}

async function deleteSession(session: { userId: number, sessionId: string })
{
	await db`DELETE FROM sessions WHERE userId = ${session.userId} AND id = ${session.sessionId}`
}

async function getSessions(userId: number)
{
	const session = await db`SELECT * FROM sessions WHERE userId = ${userId}`

	return session
}

async function isValidSession(session: { userId: string, sessionId: string })
{
	const sessions = await getSessions(parseInt(session.userId))
	let match = false
	sessions.forEach((el: { userId: number, id: string }) =>
	{
		if (el.id == session.sessionId)
		{
			match = true
		}
	})
	return match
}

export { createSession, getSessions, isValidSession, deleteSession }
export default { createSession, getSessions, isValidSession, deleteSession }