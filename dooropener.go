package dooropener

import (
	"log"
	"sync"
	"time"

	"dooropener/config"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	servoConfig  *config.Servo
	httpConfig   *config.HTTP
	telegramBot  *tb.Bot
	telegramChat *tb.Chat
)

var doorMux sync.Mutex
var doorInProgress bool

func setDoorInProgress(b bool) {
	doorMux.Lock()
	defer doorMux.Unlock()
	doorInProgress = b
}

func getDoorInProgress() bool {
	doorMux.Lock()
	defer doorMux.Unlock()
	return doorInProgress
}

func Start(c *config.Config) {
	servoConfig = c.Servo
	httpConfig = c.HTTP
	telegramChat = &tb.Chat{ID: c.Telegram.Chat}

	// Restore servo position
	setAngle(servoConfig.AngleInactive)

	// Connect to Telegram
	poller := &tb.LongPoller{Timeout: 10 * time.Second}
	middlewarePoller := tb.NewMiddlewarePoller(poller, func(upd *tb.Update) bool {
		return upd.Message.Chat.ID == c.Telegram.Chat
	})
	var err error
	telegramBot, err = tb.NewBot(tb.Settings{
		Token:  c.Telegram.Token,
		Poller: middlewarePoller,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Start web server on separate routine
	go startHTTPServer()

	// Start listening Telegram and block main goroutine
	telegramBot.Handle("/open", commandOpen)
	telegramBot.Start()
}

func commandOpen(m *tb.Message) {
	if getDoorInProgress() {
		sendTelegram("Already in progress...")
		return
	}

	setDoorInProgress(true)
	sendTelegram("Opening!")
	openDoor()
	setDoorInProgress(false)
	sendTelegram("Done!")
}

func sendTelegram(msg string) {
	telegramBot.Send(telegramChat, msg, tb.Silent)
}
