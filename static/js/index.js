(async () => {
	const session_id = await cookieStore.get({ name: "session_id", path: "/" });
	if (!session_id) {
		window.location.href = "/sign.html";
	}
	const handle = await cookieStore.get({ name: "handle", path: "/" });

	document.querySelectorAll(".handle").forEach((el) => {
		el.textContent = handle.value;
	});
})();

document.getElementById("logout").addEventListener("click", async (e) => {
	e.preventDefault();
	await cookieStore.delete({ name: "session_id", path: "/" });
	await cookieStore.delete({ name: "handle", path: "/" });
	await fetch("/api/user/logout", { method: "POST" });
	window.location.href = "/sign.html";
});
