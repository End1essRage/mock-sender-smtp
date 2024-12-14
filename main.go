package main

import (
	"flag"
	"os"

	a "github.com/end1essrage/mock-sender-smtp/api"
	s "github.com/end1essrage/mock-sender-smtp/smtp"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	Mode string // api/cli

	SmtpUrl string
	From    string
	Rcpt    string // if many - item@exm.com;item@exm.com
	Data    string
	//add Auth
)

const (
	SMTP_URL = "SMTP_URL" //from env

	FLAG_MODE = "mode"
	FLAG_FROM = "from"
	FLAG_RCPT = "rcpt"
	FLAG_DATA = "data"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})

	if err := godotenv.Load(); err != nil {
		logrus.Warning("error while reading environment %s", err.Error())
	}

	SmtpUrl = os.Getenv(SMTP_URL)

	loadFlags()
}

func main() {
	client := s.NewClient(SmtpUrl)
	if Mode == "api" {
		logrus.Info("staring api server")

		api := a.NewApi(client)

		api.Start(os.Getenv("HOST"))
	} else {
		logrus.Info("doing cli request")

		client.Send(From, Rcpt, Data)
	}
}

func loadFlags() {
	flag.StringVar(&Mode, FLAG_MODE, "", "Mode api/cli")
	flag.StringVar(&From, FLAG_FROM, "example@test.com", "From")
	flag.StringVar(&Rcpt, FLAG_RCPT, "test@example.com", "To")
	flag.StringVar(&Data, FLAG_DATA, "Hello \n .", "Data")
	flag.Parse()

	logrus.Info(Mode)
	logrus.Info(From)
	logrus.Info(Rcpt)
	logrus.Info(Mode)
}
