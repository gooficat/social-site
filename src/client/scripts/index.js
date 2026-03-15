(async () => {
	const valid = await sessionValid()
	console.log(valid)
	if (!await sessionValid()) {
		window.location.href = "/login.html"
	}
})()