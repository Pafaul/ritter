CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    nickname VARCHAR(100) UNIQUE NOT NULL,
    passwd VARCHAR(256) NOT NULL ,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_data (
    id INTEGER PRIMARY KEY,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    content VARCHAR(140) NOT NULL,
    user_id INTEGER NOT NULL,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);