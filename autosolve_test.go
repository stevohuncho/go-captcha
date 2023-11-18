package captcha_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stevohuncho/go-captcha"
	"github.com/stevohuncho/go-captcha/autosolve"
)

func loadAutosolveHarvester() (*captcha.Harvester, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	APIKEY := os.Getenv("AUTOSOLVE_APIKEY")
	c := autosolve.NewClient(APIKEY)
	err = c.Auth()
	if err != nil {
		return nil, err
	}
	h := captcha.CreateHarvester(captcha.AutosolveHarvesterOpt(c))
	return h, nil
}

func TestAutosolveClient(t *testing.T) {
	h, err := loadAutosolveHarvester()
	if err != nil {
		t.Error(err)
	}
	_ = h
}
