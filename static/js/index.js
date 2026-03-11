(async () => {
	const session_id = await cookieStore.get({ name: "session_id", path: "/" });
	if (!session_id.value) {
		window.location.href = "/login.html";
	}
	const handle = await cookieStore.get({ name: "handle", path: "/" });

	document.querySelectorAll(".handle").forEach((el) => {
		el.textContent = handle.value;
	});
})();
