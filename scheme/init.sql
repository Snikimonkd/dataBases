CREATE EXTENSION IF NOT EXISTS CITEXT;
CREATE UNLOGGED TABLE IF NOT EXISTS Users (
    nickname CITEXT UNIQUE,
    fullname CITEXT,
    about TEXT,
    email CITEXT UNIQUE
);
CREATE UNLOGGED TABLE IF NOT EXISTS Forums (
    slug CITEXT UNIQUE,
    posts INT DEFAULT 0,
    threads INT DEFAULT 0,
    title CITEXT,
    user_nickname CITEXT,
    FOREIGN KEY (user_nickname) REFERENCES users (nickname) ON DELETE CASCADE
);
CREATE UNLOGGED TABLE IF NOT EXISTS Threads (
    id SERIAL PRIMARY KEY,
    title CITEXT,
    author CITEXT,
    FOREIGN KEY (author) REFERENCES users (nickname) ON DELETE CASCADE,
    forum CITEXT,
    FOREIGN KEY (forum) REFERENCES forums (slug) ON DELETE CASCADE,
    message CITEXT,
    votes int DEFAULT 0,
    slug CITEXT,
    created TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);