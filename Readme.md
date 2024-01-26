# How To run
run service or code with command ``` go run main.go```

# code coverage
```
?       github.com/aditya37/test-cc-validator   [no test files]
?       github.com/aditya37/test-cc-validator/interface [no test files]
?       github.com/aditya37/test-cc-validator/mock      [no test files]
?       github.com/aditya37/test-cc-validator/model     [no test files]
ok      github.com/aditya37/test-cc-validator/repository        (cached)        coverage: 90.9% of statements
?       github.com/aditya37/test-cc-validator/transport [no test files]
ok      github.com/aditya37/test-cc-validator/service   (cached)        coverage: 100.0% of statements
ok      github.com/aditya37/test-cc-validator/utils     (cached)        coverage: 64.3% of statements
```

# API Contract
Base URL: http://127.0.0.1:1234

## endpoint insert card number
URL: {BASE_URL}/api/card

Payload

```
{
    "number":"41234567890123456789",
    "card_network":"Visa"
}
```

## endpoint get card number
### note: before hit this endpoint please register number with endpoint insert card number

URL: {BASE_URL}/api/check-card

Payload

```
{
    "card_number": "41234567890123456789"
}
```