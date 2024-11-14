package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-park-mail-ru/2024_2_EaglesDesigner/internal/messages/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const pageSize = 25

type MessageRepositoryImpl struct {
	pool *pgxpool.Pool
}

func NewMessageRepositoryImpl(pool *pgxpool.Pool) MessageRepository {
	return &MessageRepositoryImpl{
		pool: pool,
	}
}

func (r *MessageRepositoryImpl) GetFirstMessages(ctx context.Context, chatId uuid.UUID) ([]models.Message, error) {
	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Repository: не удалось установить соединение: %v", err)
		return nil, err
	}
	defer conn.Release()

	log.Printf("Repository: соединение успешно установлено")

	rows, err := conn.Query(context.Background(),
		`SELECT
	m.id,
	m.author_id,
	m.message,
	m.sent_at, 
	m.is_redacted
	FROM public.message AS m
	WHERE m.chat_id = $1
	ORDER BY sent_at DESC
	LIMIT $2;`,
		chatId,
		pageSize,
	)
	if err != nil {
		log.Printf("Repository: Unable to SELECT chats: %v\n", err)
		return nil, err
	}
	log.Println("Repository: сообщения получены")

	messages := []models.Message{}
	for rows.Next() {
		var messageId uuid.UUID
		var authorID uuid.UUID
		var message string
		var sentAt time.Time
		var isRedacted bool

		err = rows.Scan(&messageId, &authorID, &message, &sentAt, &isRedacted)
		if err != nil {
			log.Printf("Repository: unable to scan: %v", err)
			return nil, err
		}

		messages = append(messages, models.Message{
			MessageId:  messageId,
			AuthorID:   authorID,
			Message:    message,
			SentAt:     sentAt,
			IsRedacted: isRedacted,
		})
	}

	log.Printf("Repository: сообщения успешно найдеты. Количество сообшений: %d", len(messages))
	return messages, nil
}

func (r *MessageRepositoryImpl) AddMessage(message models.Message, chatId uuid.UUID) error {
	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Repository: не удалось установить соединение: %v", err)
		return err
	}
	defer conn.Release()
	log.Printf("Repository: соединение успешно установлено")

	// нужно чё-то придумать со стикерами
	row := conn.QueryRow(context.Background(),
		`INSERT INTO public.message (id, chat_id, author_id, message, sent_at, is_redacted)
	VALUES ($1, $2, $3, $4, $5, false) RETURNING id;`,
		uuid.New(),
		chatId,
		message.AuthorID,
		message.Message,
		message.SentAt,
	)

	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		log.Printf("Repository: не удалось добавить сообщение: %v", err)
		return err
	}

	return nil
}

func (r *MessageRepositoryImpl) GetLastMessage(chatId uuid.UUID) (models.Message, error) {
	conn, err := r.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("Repository: не удалось установить соединение: %v", err)
		return models.Message{}, err
	}
	defer conn.Release()

	// нужно чё-то придумать со стикерами
	row := conn.QueryRow(context.Background(),
		`SELECT
	m.id,
	m.author_id,
	m.message,
	m.sent_at, 
	m.is_redacted
	FROM public.message AS m
	WHERE m.chat_id = $1
	ORDER BY sent_at DESC
	LIMIT 1;`,
		chatId,
	)

	var messageId uuid.UUID
	var authorID uuid.UUID
	var message string
	var sentAt time.Time
	var isRedacted bool

	err = row.Scan(&messageId, &authorID, &message, &sentAt, &isRedacted)

	if errors.Is(err, pgx.ErrNoRows) {
		return models.Message{}, nil
	}
	if err != nil {
		log.Printf("Repository: unable to scan: %v", err)

		return models.Message{}, err
	}

	messageModel := models.Message{
		MessageId:  messageId,
		AuthorID:   authorID,
		Message:    message,
		SentAt:     sentAt,
		IsRedacted: isRedacted,
	}

	return messageModel, nil
}

func (r *MessageRepositoryImpl) GetAllMessagesAfter(ctx context.Context, chatId uuid.UUID, lastMessageId uuid.UUID) ([]models.Message, error) {
	conn, err := r.pool.Acquire(ctx)
	if err != nil {
		log.Printf("Repository: не удалось установить соединение: %v", err)
		return nil, err
	}
	defer conn.Release()
	log.Printf("Repository: соединение успешно установлено")

	rows, err := conn.Query(ctx,
		`SELECT
	m.id,
	m.author_id,
	m.message,
	m.sent_at, 
	m.is_redacted
	FROM public.message AS m
	WHERE m.chat_id = $1 AND m.sent_at <= (SELECT sent_at FROM message WHERE id = $2) AND m.id != $2
	ORDER BY sent_at DESC
	LIMIT $3;`,
		chatId,
		lastMessageId,
		pageSize,
	)

	if err != nil {
		log.Printf("Repository: Unable to SELECT chats: %v\n", err)
		return nil, err
	}
	log.Println("Repository: сообщения получены")

	messages := []models.Message{}
	for rows.Next() {
		var messageId uuid.UUID
		var authorID uuid.UUID
		var message string
		var sentAt time.Time
		var isRedacted bool

		err = rows.Scan(&messageId, &authorID, &message, &sentAt, &isRedacted)
		if err != nil {
			log.Printf("Repository: unable to scan: %v", err)
			return nil, err
		}

		messages = append(messages, models.Message{
			MessageId:  messageId,
			AuthorID:   authorID,
			Message:    message,
			SentAt:     sentAt,
			IsRedacted: isRedacted,
		})
	}

	log.Printf("Repository: сообщения успешно найдеты. Количество сообшений: %d", len(messages))
	return messages, nil
}
