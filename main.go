package logger

import (
	"endcoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"my_bot/betypes"
	"my_bot/logger"

	tgbotapi "gitHub.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	NewBot, BotErr = tgbotapi.NewBotAPI(betypes.BOT_TOKEN)
)

func setWebhook(bot *tgbotapi.BotAPI) {
	webHookInfo := tgbotapi.NewWerbhookWithCert(fmt.Sprintf("https://%s:%s/%s", betypes.BOT_ADDRES, betypes.BOT_PORT, betypes.BOT_TOKEN), betypes.CERT_PATH)
	_, err := bot.SetWebhook(webHookInfo)
	logger.ForError(err)
}

func main() {
	logger.ForError(BotErr)
	setWebhook(NewBot)
	message := func(w http.ResponseWriter, r *http.Request) {
		text, err := ioutil.ReadAll(r.Body)
		logger.ForError(err)
		var botText betypes.BotMessage
		err := json.Unmarshal(text, &botText)
		logger.ForError(err)
		fmt.Println(fmt.Sprintf("%s, text"))
		logger.LogFile.Println(fmt.Sprintf("%s, text"))

		username := botText.Message.From.Username
		chatUser := botText.Message.From.Id
		chatGroup := botText.Message.Chat.Id
		messageID := botText.Message.Message_id
		botCommand := strings.Split(botText.Message.Text, "@")[0]
		commandText := strings.Split(botText.Message.Text, " ")

		fmt.Println(username, chatUser, chatGroup, messageID, botCommand, commandText)

	}

	http.HandleFunc("/", message)
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf("%s:%s", betypes.BOT_ADDRES, betypes.BOT_PORT), betypes.CERT_PATH, betypes.KEY_PATH, nil))
}
