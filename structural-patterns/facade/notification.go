package facade

import (
	"fmt"
	"log"
)

type Notification struct {
	provider string
	email    string
	phone    string
}

func (notification *Notification) notifyByEmail(msg string) error {
	fmt.Println(msg)
	log.Println("notification to email was successfully delivered")
	return nil
}
