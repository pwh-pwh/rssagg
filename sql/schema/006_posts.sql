-- +goose Up
CREATE TABLE posts (
                              id UUID PRIMARY KEY,
                              created_at TIMESTAMP NOT NULL,
                              updated_at TIMESTAMP NOT NULL,
                              title text,
                              url text not null unique ,
                              description text,
                              published_at timestamp not null ,
                              feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;