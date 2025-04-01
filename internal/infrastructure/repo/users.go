package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	AwaitingName = iota
	AwaitingGroup
	AwaitingQuestion
	Finished
)

type UsersRepo struct {
	pool *pgxpool.Pool
}

func NewUsersRepo(pool *pgxpool.Pool) *UsersRepo {
	return &UsersRepo{pool: pool}
}

func (r *UsersRepo) AddUser(ctx context.Context, tgID int) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO users (tg_id,cur_question, score, state) VALUES ($1, $2,$3,$4)", tgID, 1, 0, AwaitingName)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) GetName(ctx context.Context, tgID int) (string, error) {
	row := r.pool.QueryRow(ctx, "SELECT name FROM users WHERE tg_id = $1", tgID)
	var name string
	if err := row.Scan(&name); err != nil {
		return "", err
	}

	return name, nil
}

func (r *UsersRepo) SetName(ctx context.Context, tgID int, name string) error {
	_, err := r.pool.Exec(ctx, "UPDATE users SET name = $1 WHERE tg_id = $2", name, tgID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UsersRepo) GetGroup(ctx context.Context, tgID int) (string, error) {
	row := r.pool.QueryRow(ctx, "SELECT group_name FROM users WHERE tg_id = $1", tgID)
	var groupName string
	if err := row.Scan(&groupName); err != nil {
		return "", err
	}

	return groupName, nil
}

func (r *UsersRepo) SetGroup(ctx context.Context, tgID int, group string) error {
	_, err := r.pool.Exec(ctx, "UPDATE users SET group_name = $1 WHERE tg_id = $2", group, tgID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UsersRepo) GetState(ctx context.Context, tgID int) (int, error) {
	row := r.pool.QueryRow(ctx, "SELECT state FROM users WHERE tg_id = $1", tgID)

	var state int

	if err := row.Scan(&state); err != nil {
		return -1, err
	}

	return state, nil
}

func (r *UsersRepo) SetState(ctx context.Context, tgID int, state int) error {
	_, err := r.pool.Exec(ctx, "UPDATE users SET state = $1 WHERE tg_id = $2", state, tgID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) GetCurQuestion(ctx context.Context, tgID int) (int, error) {
	row := r.pool.QueryRow(ctx, "SELECT cur_question FROM users WHERE tg_id = $1", tgID)
	var curQuestion int
	if err := row.Scan(&curQuestion); err != nil {
		return -1, err
	}
	return curQuestion, nil
}

func (r *UsersRepo) SetCurQuestion(ctx context.Context, tgID int, numQuestion int) error {
	_, err := r.pool.Exec(ctx, "UPDATE users SET cur_question = $1 WHERE tg_id = $2", numQuestion, tgID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) GetUserScore(ctx context.Context, tgID int) (int, error) {
	row := r.pool.QueryRow(ctx, "SELECT score FROM users WHERE tg_id = $1", tgID)
	var score int
	if err := row.Scan(&score); err != nil {
		return -1, err
	}
	return score, nil
}

func (r *UsersRepo) SetUserScore(ctx context.Context, tgID int, newScore int) error {
	_, err := r.pool.Exec(ctx, "UPDATE users SET score = $1 WHERE tg_id = $2", newScore, tgID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) GetLastMessageID(ctx context.Context, tgID int) (int, error) {
	row := r.pool.QueryRow(ctx, "SELECT last_message_id FROM users WHERE tg_id = $1", tgID)
	var lastMessageId int
	if err := row.Scan(&lastMessageId); err != nil {
		return -1, err
	}
	return lastMessageId, nil
}

func (r *UsersRepo) SetLastMessageID(ctx context.Context, tgID int, newMessageID int) error {
	_, err := r.pool.Exec(ctx, "UPDATE users SET last_message_id = $1 WHERE tg_id = $2", newMessageID, tgID)
	if err != nil {
		return err
	}

	return nil
}
