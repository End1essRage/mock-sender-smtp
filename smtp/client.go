package smtp

import (
	"fmt"
	"net/smtp"
)

type Client struct {
	Url string
}

func NewClient(Url string) *Client {
	return &Client{Url: Url}
}

// add multiple rcpt
// объединить входные параметры в один тип
func (cl *Client) Send(From, Rcpt, Data string) error {
	c, err := smtp.Dial(cl.Url)
	if err != nil {
		return fmt.Errorf("EHLO request Error :", err)
	}

	// Set the sender and recipient first
	if err := c.Mail(From); err != nil {
		return fmt.Errorf("FROM request Error :", err)
	}
	if err := c.Rcpt(Rcpt); err != nil {
		return fmt.Errorf("RCPT request Error :", err)
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		return fmt.Errorf("Body creating Error :", err)
	}
	_, err = fmt.Fprintf(wc, Data)
	if err != nil {
		return fmt.Errorf("Body writing Error :", err)
	}
	err = wc.Close()
	if err != nil {
		return fmt.Errorf("Closing io Error :", err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		return fmt.Errorf("END Error :", err)
	}

	return nil
}
