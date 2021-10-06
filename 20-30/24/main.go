package main

import (
	"log"
	"net/smtp"
)

func main() {
	var (
		from       = "demo@lan.ikunal.in"
		msg        = []byte("Message to be sent.")
		recipients = []string{"mail@example.com"}
	)
	// func PlainAuth(identity, username, password, host string) Auth
	// Usually identity should be the empty string, to act as username.
	// Setup Plain Auth details

	// If done over ssl/tls
	//smtpAuth := smtp.PlainAuth("", "demo@lan.ikunal.in", "p4ssw0rd", "lan.ikunal.in")
	// If done over non-secure connection
	smtpAuth := smtp.CRAMMD5Auth("demo@lan.ikunal.in", "p4ssw0rd")

	err := smtp.SendMail("lan.ikunal.in:25", smtpAuth, from, recipients, msg)

	if err != nil {
		log.Panicln(err)
	}
}
