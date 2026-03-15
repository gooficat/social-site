async function sessionValid() {
	const userIdCookie = await cookieStore.get("userId")
	const sessionIdCookie = await cookieStore.get("sessionId")

	if (!userIdCookie || !sessionIdCookie) return false

	const sessionCookie = {
		userId: userIdCookie.value,
		sessionId: sessionIdCookie.value
	}

	const res = await fetch('/api/user/session-validate', {
		method: 'POST',
		headers: {
			"Content-type": "application/json; charset=UTF-8"
		},
		body: JSON.stringify(sessionCookie)
	})

	return res.ok
}