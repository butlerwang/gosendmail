package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"gopkg.in/gomail.v2"
)

func main() {

	port := flag.Int("p", 465, "smtp server port")
	host := flag.String("h", "smtp.example.com", "smtp server host")
	user := flag.String("u", "admin@example.com", "username")
	pass := flag.String("P", "password", "email password")
	rec := flag.String("t", "tosomebody@whatever.com", "reciever email address")
	sub := flag.String("s", "hello", "subject")
	body := flag.String("b", "test", "body")
	flag.Parse()

	m := gomail.NewMessage()
	m.SetHeader("From", *user)
	m.SetHeader("To", *rec)
	m.SetHeader("Subject", *sub)
	m.SetBody("text/html", *body)
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(*host, *port, *user, *pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	//Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Printf("Successfully send email to %v", *rec)
}
