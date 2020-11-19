package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

var responses = []string{"Glad you got that off your chest.", "Feel better now?", "I feel with you."}

type pile []string

func (p *pile) String() string {
	var s string
	for i := len(*p) - 1; i >= 0; i-- {
		s += (*p)[i] + "\n\n"
	}
	return s
}

var rants pile

func main() {
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")
	address := os.Getenv("ADDRESS")
	debug := os.Getenv("DEBUG")
	if debug != "" {
		log.SetLevel(log.DebugLevel)
	}

	log.WithFields(log.Fields{
		"Token":   token,
		"Address": address,
	}).Debug("Config")

	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.WithError(err).Fatal("Unable to create bot")
		return
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		log.WithFields(log.Fields{
			"ID":      m.ID,
			"Sender":  m.Sender.Username,
			"Message": m.Text,
		}).Info("Rant received")

		msg := strings.TrimSpace(m.Text)
		if msg != "" {
			rants = append(rants, m.Sender.FirstName+" howled at the moon: \""+msg+"\"")
		}

		if _, err := b.Send(m.Sender, responses[rand.Intn(len(responses))]); err != nil {
			log.WithError(err).Error("Unable to respond")
		}
	})

	http.HandleFunc("/", handler)
	go func() {
		log.Fatal(http.ListenAndServe(address, nil))
	}()

	b.Start()
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/plain")
	if _, err := fmt.Fprintf(w, "RANTIMATOR\n\n"+rants.String()); err != nil {
		log.WithError(err).Error("Unable to handle request")
	}
}
