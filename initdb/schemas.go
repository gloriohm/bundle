package initdb

var CreateTables = []string{
	// goals table
	`
	CREATE TABLE IF NOT EXISTS goals (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		complete INTEGER DEFAULT 0
	);`,

	// quests table
	`
	CREATE TABLE IF NOT EXISTS quests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP,
		complete INTEGER DEFAULT 0,
		priority INTEGER DEFAULT 50,
		deadline INTEGER DEFAULT NULL
	);`,

	// dailies table
	`
	CREATE TABLE IF NOT EXISTS dailies (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP,
		complete INTEGER DEFAULT 0,
		priority INTEGER DEFAULT 50,
		next_occurrence INTEGER DEFAULT (strftime('%s', 'now')),
		days INTEGER DEFAULT 127,
		freq_type INTEGER DEFAULT 0,
		before_time TEXT DEFAULT NULL,
		after_time TEXT DEFAULT NULL
	);`,

	// dailies log
	`CREATE TABLE IF NOT EXISTS dailies_log (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		daily_name TEXT NOT NULL,
		daily_id INTEGER NOT NULL,
		completed_at TEXT DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (daily_name) REFERENCES dailies(name) ON UPDATE CASCADE,
		FOREIGN KEY (daily_id) REFERENCES dailies(id) ON DELETE CASCADE,
	)`,

	// goal - quest/daily junction tables
	`
	CREATE TABLE IF NOT EXISTS goal_quest (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		goal_id INTEGER NOT NULL,
		quest_id INTEGER NOT NULL,
		FOREIGN KEY (goal_id) REFERENCES goals(id) ON DELETE CASCADE,
		FOREIGN KEY (quest_id) REFERENCES quests(id) ON DELETE CASCADE
	);`,
	`
	CREATE TABLE IF NOT EXISTS goal_daily (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		goal_id INTEGER NOT NULL,
		daily_id INTEGER NOT NULL,
		FOREIGN KEY (goal_id) REFERENCES goals(id) ON DELETE CASCADE,
		FOREIGN KEY (daily_id) REFERENCES dailies(id) ON DELETE CASCADE
	);`,

	// core - goal/quest/daily junction table - takes "goal", "quest", or "daily" as variable
	`
	CREATE TABLE IF NOT EXISTS core_goal (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		core_id INTEGER NOT NULL,
		goal_id INTEGER NOT NULL,
		weight INTEGER NOT NULL,
		FOREIGN KEY (core_id) REFERENCES cores(id) ON DELETE CASCADE,
		FOREIGN KEY (goal_id) REFERENCES goals(id) ON DELETE CASCADE
	);`,

	`
	CREATE TABLE IF NOT EXISTS core_quest (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		core_id INTEGER NOT NULL,
		quest_id INTEGER NOT NULL,
		weight INTEGER NOT NULL,
		FOREIGN KEY (core_id) REFERENCES cores(id) ON DELETE CASCADE,
		FOREIGN KEY (quest_id) REFERENCES quests(id) ON DELETE CASCADE
	);`,

	`
	CREATE TABLE IF NOT EXISTS core_daily (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		core_id INTEGER NOT NULL,
		daily_id INTEGER NOT NULL,
		weight INTEGER NOT NULL,
		FOREIGN KEY (core_id) REFERENCES cores(id) ON DELETE CASCADE,
		FOREIGN KEY (daily_id) REFERENCES dailies(id) ON DELETE CASCADE
	);`,

	// cores table
	`
	CREATE TABLE IF NOT EXISTS cores (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		core_health INTEGER DEFAULT 50,
		last_updated INTEGER DEFAULT (strftime('%s', 'now'))
	);`,

	//create indexes
	`CREATE INDEX idx_goal_complete ON goals(complete);`,
	`CREATE INDEX idx_quest_complete ON quests(complete);`,
}
