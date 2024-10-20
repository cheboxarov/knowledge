CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    date_create TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    date_update TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    author INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    views INT DEFAULT 0,
    is_hide BOOLEAN DEFAULT false,
    likes INT DEFAULT 0,
    dislikes INT DEFAULT 0,
    parent INT REFERENCES articles(id) ON DELETE CASCADE NULL
);

CREATE TABLE articles_groups (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE articles_groups_articles (
    id SERIAL PRIMARY KEY,
    articles_group_id INT REFERENCES articles_groups(id) ON DELETE CASCADE NOT NULL,
    article_id INT REFERENCES articles(id) ON DELETE CASCADE NOT NULL
);

