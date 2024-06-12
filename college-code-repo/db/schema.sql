CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    student_id VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL
);

CREATE TABLE repositories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    owner_id INTEGER REFERENCES users(id),
    stars INTEGER DEFAULT 0,
    issues INTEGER DEFAULT 0,
    pull_requests INTEGER DEFAULT 0
);

CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    repository_id INTEGER REFERENCES repositories(id)
);

CREATE TABLE submissions (
    id SERIAL PRIMARY KEY,
    repository_id INTEGER REFERENCES repositories(id),
    user_id INTEGER REFERENCES users(id),
    submission_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE evaluations (
    id SERIAL PRIMARY KEY,
    teacher_id INTEGER REFERENCES users(id),
    submission_id INTEGER REFERENCES submissions(id),
    grade INTEGER,
    feedback TEXT
);

CREATE TABLE repository_permissions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    repository_id INTEGER REFERENCES repositories(id),
    permission_level VARCHAR(50) NOT NULL 
);
