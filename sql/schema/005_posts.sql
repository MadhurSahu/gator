-- +goose Up
CREATE TABLE posts
(
    id           uuid PRIMARY KEY,
    feed_id      uuid      NOT NULL REFERENCES feeds (id) ON DELETE CASCADE,
    title        TEXT      NOT NULL,
    url          TEXT      NOT NULL UNIQUE,
    description  TEXT      NOT NULL,
    published_at TIMESTAMP,
    created_at   TIMESTAMP NOT NULL,
    updated_at   TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE posts;