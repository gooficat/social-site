

const db = new Bun.SQL({ url: "sqlite://db.sqlite" })

await db`
CREATE TABLE IF NOT EXISTS users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	handle TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
)
`

export default db