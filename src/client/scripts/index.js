(async () => {
	if (!await sessionValid()) {
		window.location.href = "/login.html"
	}
})()