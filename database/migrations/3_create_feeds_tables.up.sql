CREATE TABLE IF NOT EXISTS feed_sources (
	id SERIAL PRIMARY KEY,
	user_id INTEGER,
	url TEXT,
	name TEXT,
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS feed_items (
	id SERIAL PRIMARY KEY,
	source_id INTEGER,
	title TEXT,
	url TEXT,
	FOREIGN KEY (source_id) REFERENCES feed_sources (id) ON DELETE CASCADE ON UPDATE CASCADE
);
