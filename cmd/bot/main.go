package main

import (
	"PGKQuizBot/internal/application"
	"PGKQuizBot/internal/infrastructure/repo"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/time/rate"
	"time"
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

	botToken := "7160142882:AAEetML9kg7Ep3jwIBHoOE4IQt8-9kevsK8"
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

	// Ограничитель по скорости для предотвращения спама
	limiter := rate.NewLimiter(rate.Every(time.Second/30), 30)

	// Основной цикл обработки обновлений
	for update := range updates {
		limiter.Wait(context.Background())
		if update.Message != nil {
			go quizBot.HandleMessage(update)
		} else if update.CallbackQuery != nil {
			go quizBot.HandleCallback(update.CallbackQuery)
		}
	}
	//userRepo.AddUser(ctx, 123)
	//err = userRepo.UpdateName(ctx, 123, "Nbvjatq Vjcrfkd")
	//if err != nil {
	//	fmt.Printf("update group: %v", err)
	//}

	//qRepo := repo.NewQuestionsRepo(pool)
	//question := "Высшей военной наградой Советского союза, которой были удостоены менее 20 человек, являлся:"
	//answers := []string{"Орден Ленина",
	//	"Орден «Победа»",
	//	"Орден красного знамени"}
	//trueAns := 2
	//err = qRepo.AddQuestion(ctx, question, answers, trueAns)
	//if err != nil {
	//	fmt.Printf("add question: %v", err)
	//}
}
