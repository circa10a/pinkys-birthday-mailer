# pinkys-birthday-mailer

A tool to send bulk postcards from local images + random dog facts using https://mailform.io

I used this to mail 125 postcards to a friend of mine that included a bunch of embarrassing pictures of me and some random dog facts for the messages (cause she :heart: 's :dog2:'s). The order ended up costing $217. Hilarious.

[![Go Report Card](https://goreportcard.com/badge/github.com/circa10a/pinkys-birthday-mailer)](https://goreportcard.com/report/github.com/circa10a/pinkys-birthday-mailer)
![Build Status](https://github.com/circa10a/pinkys-birthday-mailer/workflows/Test/badge.svg)

![alt text](https://media.tenor.com/jw_lB19FDokAAAAC/happy-birthday-dog.gif)

## Usage

Simply drop a bunch of images in the `./images` directory or update `imageDirectory` key in the config.

```sh
# requires local ./config.yaml
make run
```

## Config

Here's a sample config:

```yaml
config:
  apiToken: '' # MAILFORM_API_TOKEN environment variable takes precedence if set
  count: 10
  dryrun: false
  imageDirectory: ./images
  outputDirectory: ./output
  mail:
    to:
      name: redacted
      address: redacted
      city: Seattle
      state: WA
      postalCode: 00000
      country: US
    from:
      name: redacted
      address: redacted
      city: Seattle
      state: WA
      postalCode: 00000
      country: US
    service: USPS_POSTCARD
```
