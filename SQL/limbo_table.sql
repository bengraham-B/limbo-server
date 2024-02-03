CREATE TABLE IF NOT EXISTS limbo(
    id SERIAL PRIMARY KEY,
	user_id VARCHAR(255),
	void VARCHAR(255),
    file_name VARCHAR(255),
	file_type VARCHAR(255),
    file_Content TEXT
)