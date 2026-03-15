

const db = new Bun.SQL({ url: "sqlite://db.sqlite" })

await db`
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	handle TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	createdAt DATETIME DEFAULT CURRENT_TIMESTAMP
)
`

await db`
CREATE TABLE IF NOT EXISTS sessions (
	id TEXT NOT NULL PRIMARY KEY,
	userId INT NOT NULL,
	createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
	lastUsage DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (userId) REFERENCES users(id)
)
`

export default db