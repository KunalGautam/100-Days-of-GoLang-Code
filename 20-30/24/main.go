package main

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
)

func main() {
	// func PlainAuth(identity, username, password, host string) Auth
	// Usually identity should be the empty string, to act as username.
	// Setup Plain Auth details

	// If done over ssl/tls
	//smtpAuth := smtp.PlainAuth("", "demo@lan.ikunal.in", "p4ssw0rd", "lan.ikunal.in")
	// If done over non-secure connection
	smtpAuth := smtp.CRAMMD5Auth("demo@lan.ikunal.in", "p4ssw0rd")

	t, err := template.ParseFiles("email.tpl")
	if err != nil {
		log.Panicln(err)
	}
	var data string
	buf := &bytes.Buffer{}
	if err = t.Execute(buf, data); err != nil {
		log.Panicln(err)
	}

	templateRendered := buf.String()
	const MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	var (
		from       = "demo@lan.ikunal.in"
		subject    = "Welcome Email"
		msg        = []byte(subject + MIME + "\n" + templateRendered)
		recipients = []string{"mail@example.com"}
	)

	err = smtp.SendMail("lan.ikunal.in:25", smtpAuth, from, recipients, msg)

	if err != nil {
		log.Panicln(err)
	}
}
