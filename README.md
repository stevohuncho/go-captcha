# go-captcha
A wrapper library for 3rd party captcha solvers in Go

## Usage
To add this package to your project:
```go get github.com/stevohuncho/go-captcha```

## Getting Started
This basic example shows you how to instanciate your harvester with authentication:
```go
err := godotenv.Load()
if err != nil {
    panic(err)
}
APIKEY := os.Getenv("CAPSOLVER_APIKEY")
harvester := captcha.CreateHarvester(
    captcha.CapsolverHarvesterOpt(
        capsolver.NewClient(
            APIKEY,
        ),
    ),
)
bal, err := harvester.CapsolverBalance()
if err != nil {
    panic(err)
}
fmt.Println("balance:", bal)
```
If no errors are encountered when grabbing the harvesters balance for a specific module this means the harvester is properly authenticated.

## Solving Captchas
Solving captchas is straightforward as using the same parameters that available to the 3rd party service's documentation:
```go
res, err := harvester.Capsolver(capsolver.ReCaptchaV2(
    "https://2captcha.com/demo/recaptcha-v2",
    "6LfD3PIbAAAAAJs_eEHvoOl75_83eXSqpPSRFJ_u",
))
if err != nil {
    panic(err)
}
fmt.Println("res:", res)
```
To view more examples for solving captchas view the `*_test.go` files in the main directory of the repository. Additionally, responses are not always just a token string and some captcha responses have multiple parameters and will be returned in JSON format as a string.

## Supported Modules
| Module    | Status      | Last Updated        |
|-----------|-------------|---------------------|
| Capsolver | Completed   | November 18th, 2023 |
| AYCD      | In Progress | November 18th, 2023 |
| 2Captcha  | N/A         | N/A                 |

## Feedback
Add your issue here on GitHub. Feel free to get in touch if you have any questions.