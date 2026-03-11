const loginForm = document.querySelector("#login");
const registerForm = document.querySelector("#register");

loginForm.querySelector("button").onclick = async (e) => {
	e.preventDefault();
	const credentials = {
		handle: loginForm.querySelector('input[name="handle"]').value,
		password: loginForm.querySelector('input[name="password"]').value,
	};
	const response = await fetch("/api/user/login", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(credentials),
	});
	const json = await response.json();
	if (response.ok) {
		document.querySelector("#message").textContent =
			"Login success! " + json.session_id;
		await cookieStore.set({
			name: "session_id",
			value: json.session_id,
			path: "/",
		});
		await cookieStore.set({
			name: "handle",
			value: credentials.handle,
			path: "/",
		});
		window.location.href = "/";
	} else {
		document.querySelector("#message").textContent =
			"Login failed: " + json.message;
	}
};

registerForm.querySelector("button").onclick = async (e) => {
	e.preventDefault();
	const credentials = {
		name: registerForm.querySelector('input[name="name"]').value,
		handle: registerForm.querySelector('input[name="handle"]').value,
		email: registerForm.querySelector('input[name="email"]').value,
		password: registerForm.querySelector('input[name="password"]').value,
	};
	const response = await fetch("/api/user/register", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(credentials),
	});
	const json = await response.json();
	if (response.ok) {
		document.querySelector("#message").textContent =
			"Registration success! " + json.session_id;
		await cookieStore.set({
			name: "session_id",
			value: json.session_id,
			path: "/",
		});
		await cookieStore.set({
			name: "handle",
			value: credentials.handle,
			path: "/",
		});
		window.location.href = "/";
	} else {
		document.querySelector("#message").textContent =
			"Registration failed: " + json.message;
	}
};
