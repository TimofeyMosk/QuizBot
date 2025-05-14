package main

import (
	"PGKQuizBot/internal/application"
	"PGKQuizBot/internal/infrastructure/repo"
	"context"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/time/rate"
)

func NewPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("create pool: %w", err)
	}

	return pool, nil
}

func main() {

	botToken := "YOUR_TOKEN"
	tgAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return
	}

	ctx := context.Background()
	connStr := "postgres://postgres:postgres@localhost:5432/pgkquizbot?pool_max_conns=100"

	pool, err := NewPool(ctx, connStr)
	if err != nil {
		fmt.Printf("create pool: %v", err)
		return
	}

	userRepo := repo.NewUsersRepo(pool)
	questionRepo := repo.NewQuestionsRepo(pool)

	quizBot := application.NewQuizBot(tgAPI, userRepo, questionRepo)

	// Настраиваем получение обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := tgAPI.GetUpdatesChan(u)

	limiter := rate.NewLimiter(rate.Every(time.Second/30), 30)

	for update := range updates {
		limiter.Wait(ctx)
		if update.Message != nil {
			go quizBot.HandleMessage(update)
		} else if update.CallbackQuery != nil {
			go quizBot.HandleCallback(update.CallbackQuery)
		}
	}
}

