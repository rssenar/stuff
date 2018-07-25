package main

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {
	e := email.NewEmail()
	e.From = "Richard Senar <rssenar@gmail.com>"
	e.To = []string{"gl.senar@gmail.com"}
	e.Bcc = []string{"rssenar@gmail.com"}
	// e.Cc = []string{"test_cc@example.com"}
	e.AttachFile("/Users/richardsenar/Desktop/IMG_3703-COLLAGE.jpg")
	e.Subject = "Emailing you from command line"
	e.Text = []byte("I am not using Gmail, I am emailing you from my app.  cool right!")
	e.HTML = []byte("<h1>I am not using Gmail, I am emailing you from my app.  cool right!</h1>")
	e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "rssenar@gmail.com", "", "smtp.gmail.com"))
}
