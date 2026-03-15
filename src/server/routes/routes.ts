import userRoute from "./api/user.ts"

export default {
	"/api": {
		"/greet": new Response("Hello from the API!"),
		"/user": userRoute,
	},
}