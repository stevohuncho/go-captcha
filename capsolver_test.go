package captcha_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stevohuncho/go-captcha"
	"github.com/stevohuncho/go-captcha/capsolver"
)

func loadCapsolverHarvester() (*captcha.Harvester, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	APIKEY := os.Getenv("CAPSOLVER_APIKEY")
	h := captcha.CreateHarvester(
		captcha.CapsolverHarvesterOpt(
			capsolver.NewClient(
				APIKEY,
			),
		),
	)
	return h, nil
}

func TestCapsolverBalance(t *testing.T) {
	c, err := loadCapsolverHarvester()
	if err != nil {
		t.Error(err)
	}
	bal, err := c.CapsolverBalance()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("balance:", bal)
}

func TestCapsolverReCaptchaV2(t *testing.T) {
	c, err := loadCapsolverHarvester()
	if err != nil {
		t.Error(err)
	}
	res, err := c.Capsolver(capsolver.ReCaptchaV2(
		"https://2captcha.com/demo/recaptcha-v2",
		"6LfD3PIbAAAAAJs_eEHvoOl75_83eXSqpPSRFJ_u",
	))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("res:", res)
}

func TestCapsolverHCaptcha(t *testing.T) {
	c, err := loadCapsolverHarvester()
	if err != nil {
		t.Error(err)
	}
	res, err := c.Capsolver(capsolver.HCaptcha(
		"https://democaptcha.com/demo-form-eng/hcaptcha.html",
		"338af34c-7bcb-4c7c-900b-acbec73d7d43",
	))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("res:", res)
}

func TestCapsolverCFTurnstile(t *testing.T) {
	c, err := loadCapsolverHarvester()
	if err != nil {
		t.Error(err)
	}
	res, err := c.Capsolver(capsolver.CFTurnstile(
		"https://peet.ws/turnstile-test/non-interactive.html",
		"0x4AAAAAAABS7vwvV6VFfMcD",
		capsolver.Proxy(os.Getenv("PROXY")),
	))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("res:", res)
}

func TestCapsolverCFChallenge(t *testing.T) {
	c, err := loadCapsolverHarvester()
	if err != nil {
		t.Error(err)
	}
	res, err := c.Capsolver(capsolver.CFChallenge(
		"https://cfschl.peet.ws/",
		capsolver.Proxy(os.Getenv("PROXY")),
	))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("res:", res)
}

func TestCapsolverDatadome(t *testing.T) {
	c, err := loadCapsolverHarvester()
	if err != nil {
		t.Error(err)
	}
	res, err := c.Capsolver(capsolver.Datadome(
		"https://bck.websiteurl.com/registry",
		"https://geo.captcha-delivery.com/captcha/?initialCid=AHrlqAAAAAMA1QGvUmJwyYoAwpyjNg%3D%3D&hash=789361B674144528D0B7EE76B35826&cid=6QAEcL8coBYTi9tYLmjCdyKmNNyHz1xwM2tMHHGVd_Rxr6FsWrb7H~a04csMptCPYfQ25CBDmaOZpdDa4qwAigFnsrzbCkVkoaBIXVAwHsjXJaKYXsTpkBPtqJfLMGN&t=fe&referer=https%3A%2F%2bck.websiteurl.com%2Fclient%2Fregister%2FYM4HJV%3Flang%3Den&s=40070&e=3e531bd3b30650f2e810ac72cd80adb5eaa68d2720e804314d122fa9e84ac25d",
		capsolver.Proxy(os.Getenv("PROXY")),
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
	))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("res:", res)
}

func TestCapsolverImperva(t *testing.T) {
	c, err := loadCapsolverHarvester()
	if err != nil {
		t.Error(err)
	}
	res, err := c.Capsolver(capsolver.Imperva(
		"https://example.com/",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
		capsolver.Proxy(os.Getenv("PROXY")),
	))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("res:", res)
}

func TestCapsolverAkamaiWeb(t *testing.T) {
	c, err := loadCapsolverHarvester()
	if err != nil {
		t.Error(err)
	}
	res, err := c.Capsolver(capsolver.AkamaiWeb(
		"https://www.microsoft.com/en-us/",
		capsolver.BmszAkamaiWebOpt("E7DFAA75B9AD0D3F4D78C6DE5810D9D0~000000000000000000000000000000~YAAQRz83Fx1zA4mKAQAALflxxhU7MrGCQB8mSFISvQhk20KWs9oZ0qSylvE7W7b2w93vV7Sm1IzdpSfFaOvpdhy1sUYdHlAxYhFoXcCZjuhy8Us3PQZBh1ymDNqnjDCNQqRtJo7xzIdtrYhzIUN0SE2X680f+XOC2qIpcn0Tjj3PXpgb7RAW/gf9bR65dW/NhvNuolhqGbRmgphwpC8zMzSeiLlDM7sxxMMYWRy8KEK/c4XYeGHAN0j/4FUJfDVl3Yq/8F4LmBE3WiOEAEFjU35yo1rXGYlacimjM+nYaELVE23rZhDfikKy1dgRmXfGSBy/4TDfv7N6Q/2Dgd3PMSkn0q4HQl03Pb2WdtEEb9K8t9UHU9xI80AJ7hHw1jkiqAPvdLi24pkSGyP3dg=="),
	))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("res:", res)
}
