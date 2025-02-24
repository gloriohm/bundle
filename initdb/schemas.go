package initdb

const (
	// goals table
	CreateGoalsTable = `
	CREATE TABLE IF NOT EXISTS goals (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		complete INTEGER DEFAULT 0
	)
	`
	// quests table
	CreateQuestsTable = `
	CREATE TABLE IF NOT EXISTS quests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP,
		complete INTEGER DEFAULT 0,
		priority INTEGER DEFAULT 50,
		deadline INTEGER DEFAULT NULL
	`
	// dailies table
	CreateDailiesTable = `
	CREATE TABLE IF NOT EXISTS quests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP,
		priority INTEGER DEFAULT 50,
		next_occurrence INTEGER DEFAULT (strftime('%s', 'now')),
		days INTEGER DEFAULT 127,
		freq_type INTEGER DEFAULT 0,
		before_time TEXT DEFAULT NULL,
		after_time TEXT DEFAULT NULL,
	`
	// junction table
	CreateJunctionTable = `
	CREATE TABLE IF NOT EXISTS goal_quest (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
	)
	`
	CreateCoreJunctionTable = `
	CREATE TABLE IF NOT EXISTS core_junction (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		core_id INTEGER NOT NULL,
		goal_id INTEGER NOT NULL,
		core_score INTEGER NOT NULL
	)
	`
	// cores table
	CreateCoresTable = `
	CREATE TABLE IF NOT EXISTS cores (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		core_health INTEGER DEFAULT 50
	)
	`
	CreateIndexes = `CREATE INDEX idx_complete ON goals(complete);`
)
