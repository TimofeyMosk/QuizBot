package repo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitQuestions(pool *pgxpool.Pool) (err error) {
	qRepo := NewQuestionsRepo(pool)

	question1 := "Главнокомандующим русской армией в русско – турецкой войне 1787 – 1791 годов, по итогам которой за Российской империей были окончательно закреплены территории северного Причерноморья, включая Крым, был:"
	answers1 := []string{"Князь Г.А. Потёмкин - Таврический",
		"Граф П.А. Румянцев - Задунайский",
		"Граф А.В. Суворов – Рымникский"}
	trueAnswer1 := 1

	err = qRepo.AddQuestion(context.TODO(), question1, answers1, trueAnswer1)
	if err != nil {
		fmt.Println(err)
	}

	question2 := "Командующим русским флотом в морском сражении у мыса Калиакра, победа в котором ускорила окончание войны и подписание Ясского мирного договора, был:"
	answers2 := []string{"Адмирал П.С. Нахимов",
		"Адмирал Ф.Ф. Ушаков",
		"Адмирал В.А. Корнилов"}
	trueAnswer2 := 2

	err = qRepo.AddQuestion(context.TODO(), question2, answers2, trueAnswer2)
	if err != nil {
		fmt.Println(err)
	}

	question3 := "Высшим орденом Российской империи до 1917 года являлся:"
	answers3 := []string{"Орден святого Владимира",
		"Орден Святого Апостола Андрея Первозванного",
		"Орден Святого Георгия"}
	trueAnswer3 := 2

	err = qRepo.AddQuestion(context.TODO(), question3, answers3, trueAnswer3)
	if err != nil {
		fmt.Println(err)
	}

	question4 := "Высшей по статусу наградой российской Федерации из указанных является:"
	answers4 := []string{"Орден «За заслуги перед Отечеством»",
		"Орден Мужества",
		"Орден Святого Георгия"}
	trueAnswer4 := 1

	err = qRepo.AddQuestion(context.TODO(), question4, answers4, trueAnswer4)
	if err != nil {
		fmt.Println(err)
	}

	question5 := "Новороссия – территория от Дона до Днестра – вошла в состав Российской империи при:"
	answers5 := []string{"Екатерине I",
		"Екатерине II",
		"Елизавете I"}
	trueAnswer5 := 2

	err = qRepo.AddQuestion(context.TODO(), question5, answers5, trueAnswer5)
	if err != nil {
		fmt.Println(err)
	}

	return err
}
