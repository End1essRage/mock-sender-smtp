package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	s "github.com/end1essrage/mock-sender-smtp/smtp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*
{
	"from":"test@test.com",
	"rcpt":"test@test.com",
	"data":"hello \n ."
}
*/

type Request struct {
	From string
	Rcpt string
	Data string
}

type Api struct {
	gin    *gin.Engine
	Client *s.Client
}

func NewApi(Client *s.Client) *Api {
	api := &Api{gin: gin.Default(), Client: Client}
	api.gin.MaxMultipartMemory = 8 << 20 // 8MiB
	api.initRoutes()

	return api
}

func (a *Api) Start(host string) {
	a.gin.Run(host)
}

func (a *Api) initRoutes() {
	a.gin.GET("/status", func(c *gin.Context) {
		logrus.Info("hitted /status")
		c.JSON(http.StatusAccepted, nil)
	})

	a.gin.POST("/send", a.sendMessage)
}

func (a *Api) sendMessage(c *gin.Context) {
	logrus.Info("hitted send Message")

	req, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Error(err)
	}

	var data Request

	if err := json.Unmarshal(req, &data); err != nil {
		logrus.Error(err)
		c.String(http.StatusBadRequest, fmt.Sprintf("bad request"))
	}

	//send smtp request
	if err := a.Client.Send(data.From, data.Rcpt, data.Data); err != nil {
		logrus.Error(err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("error sending request"))
	}

	c.String(http.StatusOK, fmt.Sprintf("sended!"))
}
