DROP TABLE IF EXISTS todo;

CREATE TABLE todo (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    due_date TIMESTAMP,
    priority INTEGER NOT NULL,
    status CHAR(1) NOT NULL
);
