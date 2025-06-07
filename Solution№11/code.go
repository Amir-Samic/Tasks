package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time" // Добавлен недостающий импорт

	"github.com/Knetic/govaluate"
	"github.com/fogleman/gg"
	"gopkg.in/telebot.v3"
)

// Config содержит настройки бота
type Config struct {
	TelegramToken string
}

func main() {
	// Загружаем конфигурацию
	cfg := Config{
		TelegramToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
	}
	if cfg.TelegramToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is required")
	}

	// Настройки бота
	pref := telebot.Settings{
		Token:  cfg.TelegramToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	/start
	bot.Handle("/start", func(c telebot.Context) error {
		helpText := `Привет! Я бот для вычисления площади под кривой функции.

Отправь мне математическую функцию в формате:
"функция a b"
где:
- функция - выражение с x (например, x^2, sin(x), 2*x+3)
- a - начало интервала
- b - конец интервала

Примеры:
x^2 0 5
sin(x) 0 3.14
2*x+3 -1 1

Поддерживаемые операции: +, -, *, /, ^
Поддерживаемые функции: sin, cos, tan, sqrt, log, exp`
		return c.Send(helpText)
	})

	// Обработчик текстовых сообщений
	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		text := c.Text()
		parts := strings.Fields(text)
		if len(parts) != 3 {
			return c.Send("Неверный формат. Отправь функцию и интервал в формате: x^2 0 5")
		}

		expr := parts[0]
		a, err1 := strconv.ParseFloat(parts[1], 64)
		b, err2 := strconv.ParseFloat(parts[2], 64)
		if err1 != nil || err2 != nil {
			return c.Send("Неверный формат чисел a и b. Используй формат: x^2 0 5")
		}

		if a >= b {
			return c.Send("Начало интервала (a) должно быть меньше конца интервала (b)")
		}

		area, err := calculateArea(expr, a, b)
		if err != nil {
			return c.Send(fmt.Sprintf("Ошибка вычисления: %v", err))
		}

		img, err := plotFunction(expr, a, b, area)
		if err != nil {
			return c.Send(fmt.Sprintf("Ошибка генерации графика: %v", err))
		}

		caption := fmt.Sprintf("Функция: %s\nИнтервал: [%.2f, %.2f]\nПлощадь: %.4f", expr, a, b, area)
		return c.Send(&telebot.Photo{File: telebot.FromReader(img), Caption: caption})
	})

	log.Println("Бот запущен...")
	bot.Start()
}

func calculateArea(expr string, a, b float64) (float64, error) {
	const n = 1000
	h := (b - a) / n
	sum := 0.0

	for i := 0; i <= n; i++ {
		x := a + float64(i)*h
		y, err := evaluateFunction(expr, x)
		if err != nil {
			return 0, err
		}

		if i == 0 || i == n {
			sum += y
		} else {
			sum += 2 * y
		}
	}

	return sum * h / 2, nil
}

func evaluateFunction(expr string, x float64) (float64, error) {
	expr = strings.ReplaceAll(expr, "^", "**")

	evalExpr, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return 0, fmt.Errorf("ошибка парсинга функции: %v", err)
	}

	parameters := make(map[string]interface{}, 1)
	parameters["x"] = x

	result, err := evalExpr.Evaluate(parameters)
	if err != nil {
		return 0, fmt.Errorf("ошибка вычисления функции: %v", err)
	}

	return result.(float64), nil
}

func plotFunction(expr string, a, b, area float64) (*bytes.Buffer, error) {
	const width, height = 800, 600
	const padding = 50.0

	dc := gg.NewContext(width, height)

	// Фон
	dc.SetColor(color.White)
	dc.Clear()

	// Оси координат
	dc.SetColor(color.Black)
	dc.SetLineWidth(2)
	dc.DrawLine(padding, height-padding, width-padding, height-padding)
	dc.DrawLine(padding, padding, padding, height-padding)
	dc.Stroke()

	// Вычисляем значения функции
	points := make([]float64, width-int(2*padding))
	for i := range points {
		x := a + (b-a)*float64(i)/float64(len(points)-1)
		y, err := evaluateFunction(expr, x)
		if err != nil {
			return nil, err
		}
		points[i] = y
	}

	// Находим min и max для масштабирования
	minY, maxY := math.Inf(1), math.Inf(-1)
	for _, y := range points {
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	// Масштабируем
	yScale := (height - 2*padding) / (maxY - minY)
	_ = (width - 2*padding) / (b - a) // xScale не используется, но оставлен для ясности

	// Закрашиваем область под кривой
	dc.SetColor(color.RGBA{100, 200, 255, 100})
	dc.MoveTo(padding, height-padding)
	for i, y := range points {
		x := padding + float64(i)
		yy := height - padding - (y-minY)*yScale
		dc.LineTo(x, yy)
	}
	dc.LineTo(width-padding, height-padding)
	dc.ClosePath()
	dc.Fill()

	// Рисуем саму функцию
	dc.SetColor(color.RGBA{0, 0, 200, 255})
	dc.SetLineWidth(2)
	for i, y := range points {
		x := padding + float64(i)
		yy := height - padding - (y-minY)*yScale
		if i == 0 {
			dc.MoveTo(x, yy)
		} else {
			dc.LineTo(x, yy)
		}
	}
	dc.Stroke()

	// Подписи (без использования DefaultFontFace)
	dc.SetColor(color.Black)
	dc.DrawStringAnchored(fmt.Sprintf("y = %s", expr), width/2, padding+20, 0.5, 0.5)
	dc.DrawStringAnchored(fmt.Sprintf("Площадь: %.4f", area), width/2, height-padding-20, 0.5, 0.5)

	// Сохраняем в буфер
	buf := new(bytes.Buffer)
	if err := dc.EncodePNG(buf); err != nil {
		return nil, err
	}

	return buf, nil
}
