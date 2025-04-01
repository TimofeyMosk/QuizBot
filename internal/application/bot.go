package application

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"PGKQuizBot/internal/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Константы состояний пользователя
const (
	StateAwaitingName = iota
	StateAwaitingGroup
	StateAwaitingQuestion
	StateAwaitingAnswer
	StateFinished
)

// Интерфейс для работы с данными пользователей
type UserRepo interface {
	AddUser(ctx context.Context, tgID int) error
	GetState(ctx context.Context, tgID int) (int, error)
	SetState(ctx context.Context, tgID int, state int) error
	GetName(ctx context.Context, tgID int) (string, error)
	SetName(ctx context.Context, tgID int, name string) error
	GetGroup(ctx context.Context, tgID int) (string, error)
	SetGroup(ctx context.Context, tgID int, groupName string) error
	GetCurQuestion(ctx context.Context, tgID int) (int, error)
	SetCurQuestion(ctx context.Context, tgID int, questionID int) error
	GetUserScore(ctx context.Context, tgID int) (int, error)
	SetUserScore(ctx context.Context, tgID int, score int) error
	GetLastMessageID(ctx context.Context, tgID int) (int, error)
	SetLastMessageID(ctx context.Context, tgID int, messageID int) error
}

// Интерфейс для работы с вопросами викторины
type QuestionRepo interface {
	GetQuestion(ctx context.Context, questionID int) (domain.Question, error)
	GetNextQuestion(ctx context.Context, currentID int) (domain.Question, error)
}

type QuizBot struct {
	tgAPI        *tgbotapi.BotAPI
	userRepo     UserRepo
	questionRepo QuestionRepo
}

func NewQuizBot(tgAPI *tgbotapi.BotAPI, userRepo UserRepo, questionRepo QuestionRepo) *QuizBot {
	return &QuizBot{
		tgAPI:        tgAPI,
		userRepo:     userRepo,
		questionRepo: questionRepo,
	}
}

// Обработка входящих текстовых сообщений
func (b *QuizBot) HandleMessage(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}
	chatID := update.Message.Chat.ID
	tgID := int(chatID)
	text := update.Message.Text
	ctx := context.Background()

	if text == "/start" {
		if err := b.userRepo.AddUser(ctx, tgID); err != nil {
			log.Printf("Ошибка регистрации пользователя tgID=%d: %v", tgID, err)
			b.sendMessage(chatID, "Ошибка. Попробуйте позже.")
			return
		}
		if err := b.userRepo.SetState(ctx, tgID, StateAwaitingName); err != nil {
			log.Printf("Ошибка установки состояния пользователя tgID=%d: %v", tgID, err)
			b.sendMessage(chatID, "Ошибка. Попробуйте позже.")
			return
		}
		b.sendMessage(chatID, "Добро пожаловать в викторину!\nКак вас зовут?")
		return
	}

	state, err := b.userRepo.GetState(ctx, tgID)
	if err != nil {
		log.Printf("Ошибка получения состояния пользователя tgID=%d: %v", tgID, err)
		b.sendMessage(chatID, "Ошибка. Попробуйте позже.")
		return
	}

	switch state {
	case StateAwaitingName:
		if err := b.userRepo.SetName(ctx, tgID, text); err != nil {
			log.Printf("Ошибка сохранения имени пользователя tgID=%d: %v", tgID, err)
			b.sendMessage(chatID, "Ошибка. Попробуйте позже.")
			return
		}
		if err := b.userRepo.SetState(ctx, tgID, StateAwaitingGroup); err != nil {
			log.Printf("Ошибка смены состояния для пользователя tgID=%d: %v", tgID, err)
			b.sendMessage(chatID, "Ошибка. Попробуйте позже.")
			return
		}
		b.sendMessage(chatID, "Введите название вашей группы:")
	case StateAwaitingGroup:
		if err := b.userRepo.SetGroup(ctx, tgID, text); err != nil {
			log.Printf("Ошибка сохранения группы для пользователя tgID=%d: %v", tgID, err)
			b.sendMessage(chatID, "Ошибка. Попробуйте позже.")
			return
		}
		if err := b.userRepo.SetState(ctx, tgID, StateAwaitingQuestion); err != nil {
			log.Printf("Ошибка смены состояния для пользователя tgID=%d: %v", tgID, err)
			b.sendMessage(chatID, "Ошибка. Попробуйте позже.")
			return
		}
		b.sendQuestion(ctx, tgID, chatID)
	default:
		b.sendMessage(chatID, "Неизвестная команда. Для начала введите /start")
	}
}

// Обработка callback-запросов
func (b *QuizBot) HandleCallback(callback *tgbotapi.CallbackQuery) {
	chatID := callback.Message.Chat.ID
	tgID := int(chatID)
	data := callback.Data
	ctx := context.Background()

	parts := strings.Split(data, ":")
	if len(parts) != 4 {
		log.Printf("Неверный формат callback данных: %s", data)
		b.answerCallback(callback.ID, "Ошибка")
		return
	}

	questionID, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Printf("Ошибка преобразования questionID: %v", err)
		b.answerCallback(callback.ID, "Ошибка")
		return
	}

	answerIndex, err := strconv.Atoi(parts[3])
	if err != nil {
		log.Printf("Ошибка преобразования answerIndex: %v", err)
		b.answerCallback(callback.ID, "Ошибка")
		return
	}

	question, err := b.questionRepo.GetQuestion(ctx, questionID)
	if err != nil {
		log.Printf("Вопрос не найден (questionID=%d): %v", questionID, err)
		b.answerCallback(callback.ID, "Ошибка")
		return
	}

	correct := (answerIndex + 1) == question.TrueAns
	var response string
	if correct {
		response = "Верно!"
		score, err := b.userRepo.GetUserScore(ctx, tgID)
		if err != nil {
			log.Printf("Ошибка получения счета пользователя tgID=%d: %v", tgID, err)
		} else {
			if err := b.userRepo.SetUserScore(ctx, tgID, score+1); err != nil {
				log.Printf("Ошибка обновления счета пользователя tgID=%d: %v", tgID, err)
			}
		}
	} else {
		response = "Неверно!"
		log.Printf("Неверный ответ пользователя tgID=%d: выбран %d, верный %d", tgID, answerIndex+1, question.TrueAns)
	}

	b.answerCallback(callback.ID, response)

	curQ, err := b.userRepo.GetCurQuestion(ctx, tgID)
	if err != nil {
		log.Printf("Ошибка получения текущего вопроса для пользователя tgID=%d: %v", tgID, err)
		b.sendMessage(chatID, "Ошибка. Попробуйте позже.")
		return
	}

	nextQuestion, err := b.questionRepo.GetNextQuestion(ctx, curQ)
	if err != nil {
		log.Printf("Викторина окончена для пользователя tgID=%d: %v", tgID, err)
		lastMsgID, err := b.userRepo.GetLastMessageID(ctx, tgID)
		if err == nil && lastMsgID != 0 {
			delCfg := tgbotapi.DeleteMessageConfig{
				ChatID:    chatID,
				MessageID: lastMsgID,
			}
			if _, err := b.tgAPI.Request(delCfg); err != nil {
				log.Printf("Ошибка удаления предыдущего сообщения для tgID=%d: %v", tgID, err)
			}
		}
		b.sendMessage(chatID, "Викторина окончена!")
		b.userRepo.SetState(ctx, tgID, StateFinished)
		return
	}

	if err := b.userRepo.SetCurQuestion(ctx, tgID, nextQuestion.ID); err != nil {
		log.Printf("Ошибка обновления текущего вопроса для пользователя tgID=%d: %v", tgID, err)
		b.sendMessage(chatID, "Ошибка. Попробуйте позже.")
		return
	}

	b.sendQuestion(ctx, tgID, chatID)
}

// Отправка вопроса с inline-кнопками и удалением предыдущего сообщения
func (b *QuizBot) sendQuestion(ctx context.Context, tgID int, chatID int64) {
	lastMsgID, err := b.userRepo.GetLastMessageID(ctx, tgID)
	if err == nil && lastMsgID != 0 {
		delCfg := tgbotapi.DeleteMessageConfig{
			ChatID:    chatID,
			MessageID: lastMsgID,
		}
		if _, err := b.tgAPI.Request(delCfg); err != nil {
			log.Printf("Ошибка удаления предыдущего сообщения для tgID=%d: %v", tgID, err)
		}
	}

	curQ, err := b.userRepo.GetCurQuestion(ctx, tgID)
	if err != nil {
		log.Printf("Ошибка получения номера текущего вопроса для пользователя tgID=%d: %v", tgID, err)
		b.sendMessage(chatID, "Ошибка. Попробуйте позже.")
		return
	}

	question, err := b.questionRepo.GetQuestion(ctx, curQ)
	if err != nil {
		log.Printf("Вопрос не найден (questionID=%d): %v", curQ, err)
		b.sendMessage(chatID, "Ошибка. Попробуйте позже.")
		return
	}

	var keyboardButtons [][]tgbotapi.InlineKeyboardButton
	for i, answer := range question.Answers {
		callbackData := fmt.Sprintf("q:%d:a:%d", question.ID, i)
		btn := tgbotapi.NewInlineKeyboardButtonData(answer, callbackData)
		keyboardButtons = append(keyboardButtons, tgbotapi.NewInlineKeyboardRow(btn))
	}

	msg := tgbotapi.NewMessage(chatID, question.Question)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(keyboardButtons...)
	sentMsg, err := b.tgAPI.Send(msg)
	if err != nil {
		log.Printf("Ошибка отправки вопроса для пользователя tgID=%d: %v", tgID, err)
		return
	}

	if err := b.userRepo.SetLastMessageID(ctx, tgID, sentMsg.MessageID); err != nil {
		log.Printf("Ошибка сохранения ID сообщения для пользователя tgID=%d: %v", tgID, err)
	}
	if err := b.userRepo.SetState(ctx, tgID, StateAwaitingAnswer); err != nil {
		log.Printf("Ошибка смены состояния на ожидание ответа для пользователя tgID=%d: %v", tgID, err)
	}
}

// Вспомогательная функция для отправки текстовых сообщений
func (b *QuizBot) sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := b.tgAPI.Send(msg); err != nil {
		log.Printf("Ошибка отправки сообщения в чатID=%d: %v", chatID, err)
	}
}

// Вспомогательная функция для ответа на callback-запрос
func (b *QuizBot) answerCallback(callbackID string, text string) {
	cfg := tgbotapi.NewCallback(callbackID, text)
	if _, err := b.tgAPI.Request(cfg); err != nil {
		log.Printf("Ошибка отправки callback-ответа (callbackID=%s): %v", callbackID, err)
	}
}
