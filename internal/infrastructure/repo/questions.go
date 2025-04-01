package repo

import (
	"PGKQuizBot/internal/domain"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type QuestionsRepository struct {
	pool *pgxpool.Pool
}

func NewQuestionsRepo(pool *pgxpool.Pool) *QuestionsRepository {
	return &QuestionsRepository{pool: pool}
}

func (r *QuestionsRepository) AddQuestion(ctx context.Context, question string, answers []string, trueAns int) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO questions(question, answers, true_ans) VALUES ($1, $2, $3)", question, answers, trueAns)
	if err != nil {
		return err
	}

	return nil
}

func (r *QuestionsRepository) GetQuestion(ctx context.Context, questionID int) (domain.Question, error) {
	row := r.pool.QueryRow(ctx, "SELECT * FROM questions WHERE id = $1", questionID)

	var question domain.Question

	if err := row.Scan(&question.ID, &question.Question, &question.Answers, &question.TrueAns); err != nil {
		return domain.Question{}, err
	}

	return question, nil
}

func (r *QuestionsRepository) GetNextQuestion(ctx context.Context, currentID int) (domain.Question, error) {
	nextID := currentID + 1

	row := r.pool.QueryRow(ctx, "SELECT * FROM questions WHERE id= $1", nextID)

	var question domain.Question
	if err := row.Scan(&question.ID, &question.Question, &question.Answers, &question.TrueAns); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Question{}, errors.New("больше вопросов нет")
		}

		return domain.Question{}, err
	}

	return question, nil
}
