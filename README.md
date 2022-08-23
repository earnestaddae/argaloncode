## A small Golang application that provides a JSON API.

#### Description
The API should let the user POST two numbers (in the following they're referred to as 'a' and 'b'), and return a response that contains two results: (1) 'a' divided by 'b', and (2) 'b' divided by 'a'.

Examples:

1. POST 100 and 2 -> return 50 (because 100 / 2) and 0.02 (because 2 / 100)
2. POST 1 and 1 -> return 1 (because 1 / 1) and 1 (because 1 / 1)
3. POST 3 and 9 -> return 0.3333... (because 3 / 9) and 3 (because 9 / 3)

The layout of the project and code, structure and design of the endpoint, and format of the request and response are up to you.

– You can use any project template, framework, or library, ... as you like.
– Consider things like edge cases, error handling, logging, and tests if you can.
– Please include instructions for building the project, running the tests (if there are any), and how to call the endpoint.