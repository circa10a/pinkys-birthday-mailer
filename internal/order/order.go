package order

import (
	"strconv"

	"github.com/circa10a/go-mailform"
	"github.com/circa10a/postcard-mailer/internal/config"
	"github.com/circa10a/postcard-mailer/internal/dogfacts"
)

func Create(filename string, cfg config.Config) (*mailform.Order, error) {
	mailCfg := cfg.Config.Mail
	client, err := mailform.New(&mailform.Config{
		Token: cfg.Config.APIToken,
	})
	if err != nil {
		return nil, err
	}

	dogFact, err := dogfacts.GetRandom()
	if err != nil {
		return nil, err
	}

	orderInput := &mailform.OrderInput{
		FilePath:     filename,
		Service:      mailCfg.Service,
		ToName:       mailCfg.To.Name,
		ToAddress1:   mailCfg.To.Address,
		ToCity:       mailCfg.To.City,
		ToState:      mailCfg.To.State,
		ToPostcode:   strconv.Itoa(mailCfg.To.PostalCode),
		ToCountry:    mailCfg.To.Country,
		FromName:     mailCfg.From.Name,
		FromAddress1: mailCfg.From.Address,
		FromCity:     mailCfg.From.City,
		FromState:    mailCfg.From.State,
		FromPostcode: strconv.Itoa(mailCfg.From.PostalCode),
		FromCountry:  mailCfg.From.Country,
		Color:        true,
		Message:      dogFact,
	}

	order, err := client.CreateOrder(*orderInput)
	if err != nil {
		return nil, err
	}

	return order, nil
}
