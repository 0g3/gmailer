package gmailer

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"
)

type GmailSender struct {
	GmailAddr     string
	GmailPassword string
	To            []string
	CC            []string
	BCC           []string
	Subject       string
	Body          []byte
}

func (g *GmailSender) build() ([]byte, bool) {
	if g.GmailAddr == "" || len(g.To) == 0 {
		return nil, false
	}

	// RFC 822-style
	ret := ""
	ret += fmt.Sprintf("From: %s\r\n", g.GmailAddr)
	ret += fmt.Sprintf("To: %s\r\n", strings.Join(g.To, ";"))
	if len(g.CC) > 0 {
		ret += fmt.Sprintf("Cc: %s\r\n", strings.Join(g.CC, ";"))
	}
	ret += fmt.Sprintf("Subject: %s\r\n", g.Subject)
	ret += "\r\n" + string(g.Body)

	return []byte(ret), true
}

func (g *GmailSender) Send() (error) {
	auth := smtp.PlainAuth("", g.GmailAddr, g.GmailPassword, "smtp.gmail.com")
	msg, ok := g.build()
	if !ok {
		return errors.New("failed to build a mail")
	}
	return smtp.SendMail( "smtp.gmail.com:587", auth, g.GmailAddr, append(append(g.To, g.CC...), g.BCC...), msg)
}

func (g *GmailSender) Print() {
	fmt.Println("From:", g.GmailAddr)
	fmt.Println("Cc:", g.CC)
	fmt.Println("Bcc:", g.BCC)
	fmt.Println("Subject:", g.Subject)
	fmt.Println("Body:\n", string(g.Body))
}
