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

	question6 := "Укажите офицерские звания в Российской армии в порядке старшинства Капитан (1) Майор (2) Старший лейтенант (3)"
	answers6 := []string{"1-2-3",
		"2-1-3",
		"2-3-1"}
	trueAnswer6 := 2

	err = qRepo.AddQuestion(context.TODO(), question6, answers6, trueAnswer6)
	if err != nil {
		fmt.Println(err)
	}

	question7 := "Укажите генеральские звания в Российской армии в порядке старшинства Генерал – полковник (1) Генерал – майор (2) Генерал – лейтенант (3)"
	answers7 := []string{"2-3-1",
		"3-1-2",
		"1-3-2"}
	trueAnswer7 := 3

	err = qRepo.AddQuestion(context.TODO(), question7, answers7, trueAnswer7)
	if err != nil {
		fmt.Println(err)
	}

	question8 := "По окончании русско – турецкой войны 1787 – 1791 годов за Российской империей были окончательно закреплены территории северного Причерноморья, включая Крым, согласно:"
	answers8 := []string{"Кючук-Кайнарджийскому мирному договору",
		"Георгиевскому трактату",
		"Ясскому мирному договору"}
	trueAnswer8 := 3

	err = qRepo.AddQuestion(context.TODO(), question8, answers8, trueAnswer8)
	if err != nil {
		fmt.Println(err)
	}

	question9 := "Высшей военной наградой Советского союза, которой были удостоены менее 20 человек, являлся"
	answers9 := []string{"Орден Ленина",
		"Орден красного знамени",
		"Орден «Победа»"}
	trueAnswer9 := 3

	err = qRepo.AddQuestion(context.TODO(), question9, answers9, trueAnswer9)
	if err != nil {
		fmt.Println(err)
	}

	question10 := "Первым полководцем Древней Руси, одержавшим крупную победу над Византией, был:"
	answers10 := []string{"Олег Вещий",
		"Владимир Мономах",
		"Ярослав Мудрый"}
	trueAnswer10 := 1

	err = qRepo.AddQuestion(context.TODO(), question10, answers10, trueAnswer10)
	if err != nil {
		fmt.Println(err)
	}

	question11 := "Знаменитая крепость, которая выдержала длительную осаду немецких войск в годы Великой Отечественной войны:"
	answers11 := []string{"Брестская крепость",
		"Севастопольская крепость",
		"Кронштадтская крепость"}
	trueAnswer11 := 1

	err = qRepo.AddQuestion(context.TODO(), question11, answers11, trueAnswer11)
	if err != nil {
		fmt.Println(err)
	}

	question12 := "В ночь на 9 мая 1945 года в пригороде Берлина от советской стороны принял капитуляцию нацистской Германии, а 24 июня 1945 года принимал парад Победы на Красной площади:"
	answers12 := []string{"Георгий Жуков",
		"Константин Рокоссовский",
		"Иван Конев "}
	trueAnswer12 := 1

	err = qRepo.AddQuestion(context.TODO(), question12, answers12, trueAnswer12)
	if err != nil {
		fmt.Println(err)
	}

	question13 := "Какие войска сыграли ключевую роль в обороне Ленинграда?"
	answers13 := []string{"Воздушно-десантные войска",
		"Морская пехота",
		"Войска ПВО"}
	trueAnswer13 := 2

	err = qRepo.AddQuestion(context.TODO(), question13, answers13, trueAnswer13)
	if err != nil {
		fmt.Println(err)
	}

	question14 := "Кто стал символом героизма советских солдат в битве за Москву"
	answers14 := []string{"Панфиловцы",
		"Летчик Алексей Маресьев",
		"Матросов Александр"}
	trueAnswer14 := 1

	err = qRepo.AddQuestion(context.TODO(), question14, answers14, trueAnswer14)
	if err != nil {
		fmt.Println(err)
	}

	question15 := "Какое сражение стало крупнейшим танковым сражением в истории?"
	answers15 := []string{"Сталинградская битва  ",
		"Битва за Берлин",
		"Курская битва"}
	trueAnswer15 := 3

	err = qRepo.AddQuestion(context.TODO(), question15, answers15, trueAnswer15)
	if err != nil {
		fmt.Println(err)
	}

	return err
}
