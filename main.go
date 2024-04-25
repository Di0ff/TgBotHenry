package main

import (
	"Henry/keyboard"
	"Henry/structures"
	tokenBot "Henry/token"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(tokenBot.HenryToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		handle(bot, update)
	}
}

func handle(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch {
	case strings.HasPrefix(update.Message.Text, "/start"):
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, выберите институт:")
		msg.ReplyMarkup = keyboard.KB()
		_, err := bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	case strings.HasPrefix(update.Message.Text, "И"):
		handleInstitute(bot, update)
	default:
		handleCourse(bot, update)
	}
}

func handleInstitute(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	institute := update.Message.Text
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, выберите курс:")
	switch institute {
	case "ИПТИП", "ИТУ", "ИИТ", "ИИИ", "ИКБ", "ИРИ", "ИТХТ":
		switch institute {
		case "ИПТИП":
			msg.ReplyMarkup = keyboard.IATIP()
		case "ИТУ":
			msg.ReplyMarkup = keyboard.IMT()
		case "ИИТ":
			msg.ReplyMarkup = keyboard.IIT()
		case "ИИИ":
			msg.ReplyMarkup = keyboard.IAI()
		case "ИКБ":
			msg.ReplyMarkup = keyboard.CI()
		case "ИРИ":
			msg.ReplyMarkup = keyboard.IRECS()
		case "ИТХТ":
			msg.ReplyMarkup = keyboard.IFCT()
		}
		_, err := bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	default:
		msg.Text = "Выберите институт из предложенных вариантов."
		_, err := bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func handleCourse(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	course := update.Message.Text
	found := false
	for _, courses := range structures.CourseMap {
		for _, info := range courses {
			if info.Name == course {
				found = true
				file(bot, update.Message.Chat.ID, info.URL)
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите курс из предложенных вариантов.")
		_, err := bot.Send(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func file(bot *tgbotapi.BotAPI, chatID int64, fileURL string) {
	resp, err := http.Get(fileURL)
	if err != nil {
		log.Println("Ошибка при загрузке файла:", err)
		return
	}
	defer resp.Body.Close()

	tempFile, err := os.CreateTemp("", "tempfile-*.xlsx")
	if err != nil {
		log.Println("Ошибка при создании временного файла:", err)
		return
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		log.Println("Ошибка при копировании данных в файл:", err)
		return
	}

	doc := tgbotapi.NewDocumentUpload(chatID, tempFile.Name())
	doc.Caption = "Ваш файл"

	_, err = bot.Send(doc)
	if err != nil {
		log.Println("Ошибка при отправке файла:", err)
	}

	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Println("Ошибка при удалении временного файла:", err)
	}
}
