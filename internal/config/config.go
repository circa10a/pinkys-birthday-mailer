package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

const (
	envVar = "MAILFORM_API_TOKEN"
)

// Config is used to config the application at runtime.
type Config struct {
	Config struct {
		// APIToken is the mailform api token
		APIToken string `yaml:"apiToken"`
		// Count is the number of postcards to send
		Count int `yaml:"count" validate:"required,gt=0"`
		// DryRun is set to true when image -> pdf generation is only desired. Prevents sending orders.
		DryRun bool `yaml:"dryrun"`
		// Image directory is the directory to find images in to put on the postcards.
		ImageDirectory string `yaml:"imageDirectory" validate:"required,dir"`
		// Directory to write generated pdfs to.
		OutputDirectory string `yaml:"outputDirectory" validate:"required,dir"`
		// Mail holds the details needed to mail the postcards
		Mail struct {
			To      MailingAddress `yaml:"to" validate:"required"`
			From    MailingAddress `yaml:"from" validate:"required"`
			Service string         `yaml:"service" validate:"required"`
		} `yaml:"mail" validate:"required"`
	} `yaml:"config"`
}

// MailingAddress holds information needed for sender/receiver of mail.
type MailingAddress struct {
	Name       string `yaml:"name" validate:"required"`
	Address    string `yaml:"address" validate:"required"`
	City       string `yaml:"city" validate:"required"`
	State      string `yaml:"state" validate:"required"`
	PostalCode int    `yaml:"postalCode" validate:"required"`
	Country    string `yaml:"country" validate:"required,iso3166_1_alpha2"`
}

// Validate validates the config.
func (c *Config) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) Read(file string) error {
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}

	apiTokenEnvVar := os.Getenv(envVar)
	if apiTokenEnvVar != "" {
		c.Config.APIToken = apiTokenEnvVar
	}

	return nil
}
