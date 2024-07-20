package main

import (
	"time"
)

type User struct {
	Email string
}

type UserRepository interface {
	CreateUserAccount(u User) error
}

type NotificationsClient interface {
	SendNotification(u User) error
}

type NewsletterClient interface {
	AddToNewsletter(u User) error
}

type Handler struct {
	repository          UserRepository
	newsletterClient    NewsletterClient
	notificationsClient NotificationsClient
}

func NewHandler(
	repository UserRepository,
	newsletterClient NewsletterClient,
	notificationsClient NotificationsClient,
) Handler {
	return Handler{
		repository:          repository,
		newsletterClient:    newsletterClient,
		notificationsClient: notificationsClient,
	}
}

func (h Handler) SignUp(u User) error {
	go func() {
		for {
			err := h.repository.CreateUserAccount(u)
			if err != nil {
				time.Sleep(1 * time.Millisecond)
				continue
			}

			return
		}
	}()

	go func() {
		for {
			err := h.newsletterClient.AddToNewsletter(u)
			if err != nil {
				time.Sleep(1 * time.Millisecond)
				continue
			}

			return
		}
	}()

	go func() {
		for {
			err := h.notificationsClient.SendNotification(u)
			if err != nil {
				time.Sleep(1 * time.Millisecond)
				continue
			}

			return
		}
	}()

	time.Sleep(1 * time.Second) // No way this passes and wait group does not.
	return nil
}
