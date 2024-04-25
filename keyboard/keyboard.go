package keyboard

import "github.com/go-telegram-bot-api/telegram-bot-api"

func KB() tgbotapi.ReplyKeyboardMarkup {
	btn1 := tgbotapi.NewKeyboardButton("ИПТИП")
	btn2 := tgbotapi.NewKeyboardButton("ИТУ")
	btn3 := tgbotapi.NewKeyboardButton("ИИТ")
	btn4 := tgbotapi.NewKeyboardButton("ИИИ")
	btn5 := tgbotapi.NewKeyboardButton("ИКБ")
	btn6 := tgbotapi.NewKeyboardButton("ИРИ")
	btn7 := tgbotapi.NewKeyboardButton("ИТХТ")

	row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2, btn3)
	row2 := tgbotapi.NewKeyboardButtonRow(btn4, btn5, btn6, btn7)

	return tgbotapi.NewReplyKeyboard(row1, row2)
}

func IATIP() tgbotapi.ReplyKeyboardMarkup {
	btn1 := tgbotapi.NewKeyboardButton("1 Курс")
	btn2 := tgbotapi.NewKeyboardButton("2 Курс")
	btn3 := tgbotapi.NewKeyboardButton("3 Курс")
	btn4 := tgbotapi.NewKeyboardButton("4 Курс")
	btn5 := tgbotapi.NewKeyboardButton("5 Курс")

	row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2)
	row2 := tgbotapi.NewKeyboardButtonRow(btn3, btn4, btn5)

	return tgbotapi.NewReplyKeyboard(row1, row2)
}

func IMT() tgbotapi.ReplyKeyboardMarkup {
	btn1 := tgbotapi.NewKeyboardButton("1 Курс")
	btn2 := tgbotapi.NewKeyboardButton("2 Курс")
	btn3 := tgbotapi.NewKeyboardButton("3 Курс")
	btn4 := tgbotapi.NewKeyboardButton("4 Курс")

	row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2)
	row2 := tgbotapi.NewKeyboardButtonRow(btn3, btn4)

	return tgbotapi.NewReplyKeyboard(row1, row2)
}

func IIT() tgbotapi.ReplyKeyboardMarkup {
	btn1 := tgbotapi.NewKeyboardButton("1 Курс")
	btn2 := tgbotapi.NewKeyboardButton("2 Курс")
	btn3 := tgbotapi.NewKeyboardButton("3 Курс")

	row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2, btn3)

	return tgbotapi.NewReplyKeyboard(row1)
}

func IAI() tgbotapi.ReplyKeyboardMarkup {
	btn1 := tgbotapi.NewKeyboardButton("1 Курс")
	btn2 := tgbotapi.NewKeyboardButton("2 Курс")
	btn3 := tgbotapi.NewKeyboardButton("3 Курс")
	btn4 := tgbotapi.NewKeyboardButton("4 Курс")
	btn5 := tgbotapi.NewKeyboardButton("5 Курс")

	row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2)
	row2 := tgbotapi.NewKeyboardButtonRow(btn3, btn4, btn5)

	return tgbotapi.NewReplyKeyboard(row1, row2)
}

func CI() tgbotapi.ReplyKeyboardMarkup {
	btn1 := tgbotapi.NewKeyboardButton("1 Курс")
	btn2 := tgbotapi.NewKeyboardButton("2 Курс")
	btn3 := tgbotapi.NewKeyboardButton("3 Курс")
	btn4 := tgbotapi.NewKeyboardButton("4 Курс")
	btn5 := tgbotapi.NewKeyboardButton("5 Курс")

	row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2)
	row2 := tgbotapi.NewKeyboardButtonRow(btn3, btn4, btn5)

	return tgbotapi.NewReplyKeyboard(row1, row2)
}

func IRECS() tgbotapi.ReplyKeyboardMarkup {
	btn1 := tgbotapi.NewKeyboardButton("1 Курс")
	btn2 := tgbotapi.NewKeyboardButton("2 Курс")
	btn3 := tgbotapi.NewKeyboardButton("3 Курс")
	btn4 := tgbotapi.NewKeyboardButton("4 Курс")
	btn5 := tgbotapi.NewKeyboardButton("5 Курс")

	row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2)
	row2 := tgbotapi.NewKeyboardButtonRow(btn3, btn4, btn5)

	return tgbotapi.NewReplyKeyboard(row1, row2)
}

func IFCT() tgbotapi.ReplyKeyboardMarkup {
	btn1 := tgbotapi.NewKeyboardButton("1 Курс")
	btn2 := tgbotapi.NewKeyboardButton("2 Курс")
	btn3 := tgbotapi.NewKeyboardButton("3 Курс")
	btn4 := tgbotapi.NewKeyboardButton("4 Курс")

	row1 := tgbotapi.NewKeyboardButtonRow(btn1, btn2)
	row2 := tgbotapi.NewKeyboardButtonRow(btn3, btn4)

	return tgbotapi.NewReplyKeyboard(row1, row2)
}
