package facades

import "github.com/mewway/go-laravel/contracts/mail"

func Mail() mail.Mail {
	return App().MakeMail()
}
