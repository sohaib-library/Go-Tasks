-- +goose Up
CREATE TABLE IF NOT EXISTS result (
    id            SERIAL PRIMARY KEY,
    total_words   INTEGER NOT NULL,
    total_letters INTEGER NOT NULL,
    total_spaces  INTEGER NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS result;
