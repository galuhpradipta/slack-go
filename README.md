# Guide to Run a Go Program

This guide provides step-by-step instructions to run a Go program on your local machine.

## Prerequisites

Before starting, ensure that you have the following prerequisites:

- Go version 1.13 or later installed on your machine
- Slack Developer Access

## Steps

Follow the below steps to run a Go program:
### Generate Slack BOT token and also create dummy channel in your workspace. [Tutorial](https://help.capenetworks.com/en/articles/2361824-creating-a-slack-api-token#:~:text=Click%20on%20the%20Add%20Bot,API%20test%20on%20the%20dashboard.)

### Clone Repo
```bash
$ git clone https://github.com/galuhpradipta/slack-go
$ cd slack-go
$ cp .env.example .env
```

### Replace env variable with your own
```bash
SLACK_TOKEN=<your-slack-token>
SLACK_CHANNEL=<your-slack-channel>
```
### build and run go program
```bash
$ go build
$ ./slack-go
```

### Test with curl
program should be running on port 3000 and `/alert` path with POST method

```bash
# Request
curl --location 'http://localhost:3000/alert' \
--header 'Content-Type: application/json' \
--data-raw '{
    "RecordType": "Bounce",
    "Type": "SpamNotification",
    "TypeCode": 512,
    "Name": "Spam notification",
    "Tag": "",
    "MessageStream": "outbound",
    "Description": "The message was delivered, but was either blocked by the user, or classified as spam, bulk mail, or had rejected content.",
    "Email": "zaphod@example.com",
    "From": "notifications@honeybadger.io",
    "BouncedAt": "2023-02-27T21:41:30Z"
}'
```

```bash
# Response
{"message":"success"}
```
