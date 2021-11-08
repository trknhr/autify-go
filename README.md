# autify-go

This is a Go wrapper for working with [Autify's Web API](https://autifyhq.github.io/autify-api/).

This project tries to connect the Web API Endpoint easily by using this library.

# Installation

To install the library, you can run this command.

```
go install github.com/trknhr/auify
```

# How to use

## Authentication 

You have to prepare API key and your projectID to use this library. To generate or manage API keys、please visit your account page.

Then please set an environment variable like `AUTIFY_API_KEY`. You can use the key as the below example.

```go
autifyApiKey := os.Getenv("AUTIFY_API_KEY")

ctx := context.Background()
client := autify.New(autify.Config{Token: autifyApiKey})

// Do your task !!
```
