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
	Env  string
	Mode string // api/cli

	SmtpUrl string
	From    string
	Rcpt    string // if many - item@exm.com;item@exm.com
	Data    string
	//add Auth
)

const (
	ENV_DEBUG = "ENV_DEBUG" //Для локального запуска в дебаг режиме
	ENV_LOCAL = "ENV_LOCAL" //Для локального запуска
	ENV_POD   = "ENV_POD"   //Для запуска в контейнере

	SMTP_URL = "SMTP_URL" //from env

	FLAG_MODE = "mode"
	FLAG_FROM = "from"
	FLAG_RCPT = "rcpt"
	FLAG_DATA = "data"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})

	Env = os.Getenv("ENV")
	if Env == "" {
		if err := godotenv.Load(); err != nil {
			logrus.Warning("error while reading environment %s", err.Error())
		}
	}

	Env = os.Getenv("ENV")
	if Env == "" {
		logrus.Warn("cant set environment, setting to local by default")
		Env = ENV_LOCAL
	}

	logrus.Info("ENVIRONMENT IS " + Env)
	SmtpUrl = os.Getenv(SMTP_URL)

	loadFlags()
}

func main() {
	client := s.NewClient(SmtpUrl)
	if Env == ENV_POD || Mode == "api" {
		logrus.Info("staring api server")

		api := a.NewApi(client)

		api.Start(os.Getenv("HOST"))
	} else {
		logrus.Info("doing cli request")

		if err := client.Send(From, Rcpt, Data); err != nil {
			logrus.Error(err)
		}
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
	logrus.Info(Data)

	logrus.Info(SmtpUrl)
}
