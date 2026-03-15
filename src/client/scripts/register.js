
let form = document.querySelector('#sign')
form.addEventListener('submit', async (e) => {
	e.preventDefault()
	let formData = new FormData(form)
	const credentials = { handle: formData.get('handle'), password: formData.get('password') }

	if (formData.get('password-confirm') !== credentials.password) {
		document.querySelector('#error').innerText = "Passwords do not match"
		return;
	}

	let res = await fetch('/api/user/register', {
		method: 'POST',
		headers: {
			"Content-type": "application/json; charset=UTF-8"
		},
		body: JSON.stringify(credentials)
	})
	let session_id = await res.text()

	console.log(session_id)
})