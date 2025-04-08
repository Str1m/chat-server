-- +goose Up
-- +goose StatementBegin
CREATE TABLE chats (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE chat_user (
    chat_id BIGINT NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    user_email VARCHAR(50) NOT NULL,
    PRIMARY KEY (chat_id, user_email)
);

CREATE TABLE message (
    id BIGSERIAL PRIMARY KEY,
    chat_id BIGINT NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    sender VARCHAR(50) NOT NULL,
    text TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL
);

CREATE INDEX idx_chat_user_user_email ON chat_user(user_email);
CREATE INDEX idx_message_chat_id ON message(chat_id);
CREATE INDEX idx_message_timestamp ON message(timestamp);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_message_timestamp;
DROP INDEX IF EXISTS idx_message_chat_id;
DROP INDEX IF EXISTS idx_chat_user_user_email;
DROP TABLE chat_user;
DROP TABLE message;
DROP TABLE chats;
-- +goose StatementEnd

