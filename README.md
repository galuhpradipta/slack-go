# Guide to Run a Go Program

This guide provides step-by-step instructions to run a Go program on your local machine.

## Prerequisites

Before starting, ensure that you have the following prerequisites:

- Go version 1.13 or later installed on your machine
- Slack Developer Access

## Steps

Follow the below steps to run a Go program:
1. Generate Slack BOT token and also create dummy channel in your workspace. [Tutorial](https://help.capenetworks.com/en/articles/2361824-creating-a-slack-api-token#:~:text=Click%20on%20the%20Add%20Bot,API%20test%20on%20the%20dashboard.)
2. Open your terminal and navigate to the directory where the Go program is located using the `cd` command.
3. Build the Go program using the following command:
    
    ```
    go build
    
    ```
    
    This command will compile the Go program and generate an executable file with the same name as the directory.
    
3. Run the Go program using the following command:
    
    ```
    ./<executable-file-name>
    
    ```
    
    Replace `<executable-file-name>` with the name of the executable file generated in step 2.
    
4. If the program runs successfully, you should see the output printed in your terminal.

Congratulations! You have successfully run a Go program on your local machine.