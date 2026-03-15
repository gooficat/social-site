async function sessionValid() {
	const userIdCookie = await cookieStore.get("userId")
	const sessionIdCookie = await cookieStore.get("sessionId")
	const sessionCookie = {
		userId: userIdCookie.value,
		sessionId: sessionIdCookie.value
	}

	console.log(sessionCookie)

	const res = await fetch('/api/user/session-validate', {
		method: 'POST',
		headers: {
			"Content-type": "application/json; charset=UTF-8"
		},
		body: JSON.stringify(sessionCookie)
	})

	console.log("Session ? ", res.ok)

	return res.ok
}