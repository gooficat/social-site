import { sessionValid } from "./session.mjs"

(async () => {
	if (!await sessionValid()) {
		window.location.href = "/login.html"
	}
})()