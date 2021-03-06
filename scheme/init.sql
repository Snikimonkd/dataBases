CREATE EXTENSION IF NOT EXISTS CITEXT;

CREATE UNLOGGED TABLE IF NOT EXISTS users (
    nickname CITEXT UNIQUE,
    fullname CITEXT,
    about TEXT,
    email CITEXT UNIQUE
);

CREATE INDEX index_users_nickname ON users USING HASH (nickname);
CREATE INDEX index_users_email ON users USING HASH (email);
CREATE INDEX index_users_email_nickname ON users (email, nickname);

CREATE UNLOGGED TABLE IF NOT EXISTS forums (
    slug CITEXT UNIQUE,
    posts INT DEFAULT 0,
    threads INT DEFAULT 0,
    title CITEXT,
    user_nickname CITEXT,
    FOREIGN KEY (user_nickname) REFERENCES users (nickname) ON DELETE CASCADE
);

CREATE INDEX index_forums_slug ON forums USING HASH (slug);
CREATE INDEX index_forums_users ON forums USING HASH (user_nickname);

CREATE UNLOGGED TABLE IF NOT EXISTS threads (
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

CREATE INDEX index_threads_id ON threads (id);
CREATE INDEX index_threads_created ON threads (created);
CREATE INDEX index_threads_slug ON threads USING HASH (slug);
CREATE INDEX index_forum ON threads USING HASH (forum);


CREATE UNLOGGED TABLE IF NOT EXISTS posts (
    id BIGSERIAL PRIMARY KEY,
    author CITEXT,
    FOREIGN KEY (author) REFERENCES users (nickname) ON DELETE CASCADE,
    created TIMESTAMP WITH TIME ZONE,
    forum CITEXT,
    isEdited BOOLEAN DEFAULT FALSE,
    message TEXT,
    parent INT,
    thread INT,
    tree BIGINT[]
);

CREATE INDEX index_posts_id ON posts (id);
CREATE INDEX index_posts_thread_id ON posts (thread, id);
CREATE INDEX index_posts_thread_tree ON posts (thread, tree);
CREATE INDEX index_posts_thread_parent_tree ON posts (thread, parent, (tree[1]));
CREATE INDEX index_posts_tree1 ON posts ((tree[1]));

CREATE INDEX index_posts_thread ON posts (thread);
CREATE INDEX index_posts_tree ON posts (tree);

CREATE UNLOGGED TABLE IF NOT EXISTS votes (
    nickname CITEXT,
    FOREIGN KEY (nickname) REFERENCES users (nickname) ON DELETE CASCADE,
    thread_id INT,
    vote INT
);

CREATE INDEX vote on votes (nickname, vote);

CREATE FUNCTION insert_votes()
    RETURNS TRIGGER AS
$insert_votes$
BEGIN
    IF new.vote > 0 THEN
        UPDATE threads SET votes = (votes + 1)
        WHERE id = new.thread_id;
    ELSE
        UPDATE threads SET votes = (votes - 1)
        WHERE id = new.thread_id;
    END IF;
    RETURN new;
END;
$insert_votes$ language plpgsql;

CREATE TRIGGER insert_votes
    BEFORE INSERT
    ON votes
    FOR EACH ROW
EXECUTE PROCEDURE insert_votes();

CREATE FUNCTION update_votes()
    RETURNS TRIGGER AS
$update_votes$
BEGIN
    IF new.vote = 1 THEN
        UPDATE threads
        SET votes = (votes + 2)
        WHERE threads.id = new.thread_id;
    else
        UPDATE threads
        SET votes = (votes - 2)
        WHERE threads.id = new.thread_id;
    END IF;
    RETURN new;
END;
$update_votes$ LANGUAGE plpgsql;

CREATE TRIGGER update_votes
    BEFORE UPDATE
    ON votes
    FOR EACH ROW
EXECUTE PROCEDURE update_votes();

CREATE FUNCTION make_post_tree()
    RETURNS TRIGGER AS
$make_post_tree$
BEGIN
    new.tree = (SELECT tree FROM posts WHERE id = new.parent) || new.id;
    UPDATE forums SET posts = posts + 1 WHERE slug = new.forum;
    RETURN new;
END;
$make_post_tree$ LANGUAGE plpgsql;

CREATE TRIGGER make_post_tree
    BEFORE INSERT
    ON posts
    FOR EACH ROW
EXECUTE PROCEDURE make_post_tree();

CREATE FUNCTION update_threads()
    RETURNS TRIGGER AS
$update_threads$
BEGIN
    UPDATE forums SET threads = threads + 1 WHERE slug = new.forum;
    RETURN new;
END;
$update_threads$ LANGUAGE plpgsql;

CREATE TRIGGER update_threads
    BEFORE INSERT
    ON threads
    FOR EACH ROW
EXECUTE PROCEDURE update_threads();

CREATE UNLOGGED TABLE IF NOT EXISTS forum_participants (
    forum CITEXT,
    user_nickname CITEXT,
    user_fullname CITEXT,
    user_about TEXT,
    user_email CITEXT,
    UNIQUE (forum, user_nickname)
);

CREATE INDEX index_participants_nickname ON forum_participants USING HASH (user_nickname);
CREATE INDEX index_participants_forum ON forum_participants USING HASH (forum);
CREATE INDEX forum_participants_all ON forum_participants (user_nickname, user_fullname, user_about, user_email, forum);

CREATE FUNCTION forum_participant()
    RETURNS TRIGGER AS
$forum_participant$
DECLARE
    buf_nickname CITEXT;
    buf_fullname CITEXT;
    buf_about    TEXT;
    buf_email    CITEXT;
BEGIN
    SELECT nickname, fullname, about, email
    FROM users
    WHERE nickname = new.author
    INTO buf_nickname, buf_fullname, buf_about, buf_email;

    INSERT INTO forum_participants (forum, user_nickname, user_fullname, user_about, user_email)
    VALUES (new.forum, buf_nickname, buf_fullname, buf_about, buf_email)
    ON CONFLICT DO NOTHING;
    RETURN new;
END;
$forum_participant$ LANGUAGE plpgsql;

CREATE TRIGGER forum_participant_thread
    AFTER INSERT
    ON threads
    FOR EACH ROW
EXECUTE PROCEDURE forum_participant();

CREATE TRIGGER forum_participant_post
    AFTER INSERT
    ON posts
    FOR EACH ROW
EXECUTE PROCEDURE forum_participant();