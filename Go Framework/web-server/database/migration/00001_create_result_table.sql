-- +goose Up
CREATE TABLE IF NOT EXISTS result (
    id            SERIAL PRIMARY KEY,
    user_id       INTEGER NOT NULL,
    total_words   INTEGER NOT NULL,
    total_letters INTEGER NOT NULL,
    total_spaces  INTEGER NOT NULL,
    total_lines   INTEGER NOT NULL,
    total_special_char  INTEGER NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS result;
