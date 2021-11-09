package main

import (
	"fmt"
	"strings"
)

// Example - Builder Parameter

type email struct {
	from, to, subject, body string
}

type action func(builder *EmailBuilder)

//EmailBuilder struct
type EmailBuilder struct {
	email email
}

func sendEmail(email *email) {
	// logic to send email
	fmt.Println("email sent ...")
	fmt.Printf("To:%s,\nSubject:%s\n%s\nRegards,\n%s\n", email.to, email.subject, email.body, email.from)
}

//SendEmail function is for client to send email
func SendEmail(action action) {
	builder := EmailBuilder{}
	action(&builder)
	sendEmail(&builder.email)
}

//To sets the email's "To" address
func (eb *EmailBuilder) To(value string) *EmailBuilder {
	//basic validation
	if !strings.Contains(value, "@") {
		panic("Invalid email")
	}
	eb.email.to = value
	return eb

}

//From sets the email's "From" address
func (eb *EmailBuilder) From(value string) *EmailBuilder {
	//basic validation
	if !strings.Contains(value, "@") {
		panic("Invalid email")
	}
	eb.email.from = value
	return eb
}

//Body sets the email's "Body"
func (eb *EmailBuilder) Body(value string) *EmailBuilder {
	eb.email.body = value
	return eb
}

//Subject sets the email's "Subject"
func (eb *EmailBuilder) Subject(value string) *EmailBuilder {
	eb.email.subject = value
	return eb
}

//RunBuilderParamter example
// func RunBuilderParamter() {
func main() {
	SendEmail(func(b *EmailBuilder) {
		b.
			To("World@gmail.com").
			From("me@gmail.com").
			Subject("Hello world").
			Body("Sample body for a dummy email")
	})

	fmt.Println("++++++++++++++++++++++++++")

	SendEmail(func(b *EmailBuilder) {
		b.
			From("me@gmail.com").
			Body("Sample body for a dummy email #1").
			To("World@gmail.com").
			Subject("Hello world").
			Body("Sample body for a dummy email #2")
	})
}
