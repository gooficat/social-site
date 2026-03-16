import { storeUser } from './session.mjs'

let form = document.querySelector('#sign')
form.addEventListener('submit', async (e) => {
	e.preventDefault()
	let formData = new FormData(form)
	const credentials = { handle: formData.get('handle'), password: formData.get('password') }
	let res = await fetch('/api/user/login', {
		method: 'POST',
		headers: {
			"Content-type": "application/json; charset=UTF-8"
		},
		body: JSON.stringify(credentials)
	})

	if (!res.ok) {
		document.querySelector("#error").text = await res.text()
		return
	}
	let json = await res.json()

	storeUser(credentials.handle, json.userId, json.sessionId)

	window.location.href = "/"
})