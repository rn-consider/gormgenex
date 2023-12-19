CREATE TABLE users (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    age INT,
    gender BOOLEAN
);

CREATE TABLE user_infos (
    user_id INT,
    id INT PRIMARY KEY,
    addr VARCHAR(255),
    like VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE articles (
    id INT PRIMARY KEY,
    title VARCHAR(255),
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE tags (
    id INT PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE article_tags (
    article_id INT,
    tag_id INT,
    PRIMARY KEY (article_id, tag_id),
    FOREIGN KEY (article_id) REFERENCES articles(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);