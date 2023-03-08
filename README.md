# Guide to Run a Go Program

This guide provides step-by-step instructions to run a Go program on your local machine.

## Prerequisites

Before starting, ensure that you have the following prerequisites:

- Go version 1.13 or later installed on your machine
- Slack Developer Access

## Steps

Follow the below steps to run a Go program:
1. Generate Slack BOT token and also create dummy channel in your workspace. [Tutorial](https://help.capenetworks.com/en/articles/2361824-creating-a-slack-api-token#:~:text=Click%20on%20the%20Add%20Bot,API%20test%20on%20the%20dashboard.)
2. Clone Repo
```bash
$ git clone https://github.com/galuhpradipta/slack-go
$ cd slack-go
$ cp .env.example .env
```
3. Replace env variable with your own
```bash
SLACK_TOKEN=<your-slack-token>
SLACK_CHANNEL=<your-slack-channel>
```
4. build and run go program
```bash
$ go build
$ ./slack-go
```

