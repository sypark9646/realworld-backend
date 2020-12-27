create schema realworld;
use realworld;

CREATE TABLE users (
    id SERIAL NOT NULL,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    bio VARCHAR(255),
    image VARCHAR(255),
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE follow (
    id SERIAL NOT NULL,
    user_id BIGINT NOT NULL,
    following_id BIGINT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id)
        REFERENCES users (id),
    FOREIGN KEY (following_id)
        REFERENCES users (id)
);
  
CREATE TABLE article (
    id SERIAL NOT NULL,
    user_id BIGINT NOT NULL,
    slug VARCHAR(65) NOT NULL,
    title VARCHAR(55) NOT NULL,
    description VARCHAR(255) NOT NULL,
    body VARCHAR(255) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id)
        REFERENCES users (id)
);

CREATE TABLE comment (
    id SERIAL NOT NULL,
    article_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    body VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (article_id)
        REFERENCES article (id),
    FOREIGN KEY (user_id)
        REFERENCES users (id)
);
  
CREATE TABLE tag (
    id SERIAL NOT NULL,
    article_id BIGINT NOT NULL,
    tag_name VARCHAR(50) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (article_id)
        REFERENCES article (id)
);
  
CREATE TABLE favorites (
    id SERIAL NOT NULL,
    article_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (article_id)
        REFERENCES article (id),
    FOREIGN KEY (user_id)
        REFERENCES users (id)
);
