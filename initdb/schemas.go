package initdb

const (
	// goals table
	CreateGoalsTable = `
	CREATE TABLE IF NOT EXISTS goals (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		complete BOOLEAN DEFAULT 0
	)
	`
	// quests table
	CreateQuestsTable = `
	CREATE TABLE IF NOT EXISTS quests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		complete BOOLEAN DEFAULT 0),
		priority INTEGER DEFAULT 50,
		deadline DATETIME
	`
	CreateIndexes = `CREATE INDEX idx_complete ON goals(complete);`
)
