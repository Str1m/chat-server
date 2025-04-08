-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats(
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE chat_user(
    chat_id BIGINT NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    user_email VARCHAR(50) NOT NULL,
    PRIMARY KEY (chat_id, user_email)
);

CREATE TABLE message(
    id BIGINT PRIMARY KEY,
    chat_id BIGINT NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    sender TEXT NOT NULL,
    text TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE chat_user;
DROP TABLE message;
DROP TABLE chats;
-- +goose StatementEnd

