package postgres

import (
	"chat-server/internal/models"
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	db *sql.DB
}

func New(dsn string) (*Storage, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) CreateChat(ctx context.Context, usernames []string) (int64, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	// defer tx.Rollback()

	var chatID int64
	err = tx.QueryRow(`INSERT INTO chats DEFAULT VALUES RETURNING id`).Scan(&chatID)
	if err != nil {
		return 0, err
	}

	stmt, err := tx.Prepare(`INSERT INTO chat_user (chat_id, user_email) VALUES ($1, $2)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	for _, email := range usernames {
		if _, err = stmt.Exec(chatID, email); err != nil {
			return 0, err
		}
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return chatID, nil
}

func (s *Storage) DeleteChat(chatID int64) error {
	_, err := s.db.Exec(`DELETE FROM chats WHERE id = $1`, chatID)
	return err
}

func (s *Storage) SendMessage(msg models.Message) error {
	_, err := s.db.Exec(
		`INSERT INTO message (id, chat_id, sender, text, timestamp) VALUES ($1, $2, $3, $4, $5)`,
		msg.ID, msg.ChatID, msg.Sender, msg.Text, msg.Timestamp,
	)
	return err
}

func (s *Storage) GetMessages(chatID int64) ([]models.Message, error) {
	rows, err := s.db.Query(`SELECT id, chat_id, sender, text, timestamp 
	FROM message WHERE chat_id = $1 ORDER BY timestamp`, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err = rows.Scan(&msg.ID, &msg.ChatID, &msg.Sender, &msg.Text, &msg.Timestamp); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, rows.Err()
}
