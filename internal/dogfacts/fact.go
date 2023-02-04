package dogfacts

import "github.com/go-resty/resty/v2"

const (
	factEndpoint = "https://dog-api.kinduff.com/api/facts"
)

type factResponse struct {
	Facts   []string `json:"facts"`
	Success bool     `json:"success"`
}

func GetRandom() (string, error) {
	factResp := &factResponse{}
	client := resty.New()
	_, err := client.R().
		SetResult(factResp).
		Get(factEndpoint)
	if err != nil {
		return "", err
	}

	return factResp.Facts[0], nil
}
