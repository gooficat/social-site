import { clearUser } from "./session.mjs"

(async () => {
	const sessionId = (await cookieStore.get('sessionId')).value
	if (sessionId) {
		await fetch('/api/user/logout', {
			method: 'POST',
			headers: {
				"Content-type": "application/json; charset=UTF-8"
			},
			body: JSON.stringify({
				userId: (await cookieStore.get('userId')).value,
				sessionId: sessionId
			})
		})

		await clearUser()
	}
	window.location.href = "/login.html"
})()