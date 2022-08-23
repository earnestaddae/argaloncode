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

#### Installation 
- Ensure you have the latest version of [golang installed](https://go.dev/dl/)
- Ensure you have make install for 
    - Ubuntu - `sudo apt install make`
    - MacOS - `brew install make`
    - Windows - `choco install make`
- Ensure you have [curl](https://curl.se/download.html) or [httpie](https://httpie.io/docs/cli/installation)
- Clone the repo into a directory of your choice (GOPATH recommended)
- Install [Staticcheck](https://github.com/dominikh/go-tools/releases). Alternative run `go install honnef.co/go/tools/cmd/staticcheck@latest` to install staticcheck.
- Run `make help` to view the make commands
- To run tests `make audit`
- To run vendor `make vendor`
- Application runs on `localhost:8081`

#### Buidling & Running application
- To build binary `make build`
- To start application `make start`
- To stop application `make stop`

#### Curl / httpie commands - CLI test cases
- CURL
```shell
BODY='{"a":100,"b":2 }'
curl -i -d "$BODY" http://localhost:8081/v1/divisions

BODY='{"a":1,"b":1 }'
curl -i -d "$BODY" http://localhost:8081/v1/divisions

BODY='{"a":3,"b":9 }'
curl -i -d "$BODY" http://localhost:8081/v1/divisions

BODY='{"a":3,"b":0 }'
curl -i -d "$BODY" http://localhost:8081/v1/divisions
```

- HTTPIE 
```shell 
echo '{"a": 100, "b": 2 }' | http -v post http://localhost:8081/v1/divisions

echo '{"a": 1, "b": 1 }' | http -v post http://localhost:8081/v1/divisions

echo '{"a": 3, "b": 9 }' | http -v post http://localhost:8081/v1/divisions

echo '{"a": 3, "b": 0 }' | http -v post http://localhost:8081/v1/divisions
```

#### When all fails
- The binaries for Windows and Linux are located
- `.bin/api/linux_amd64/api/argalon`
- `.bin/api/windowns_amd64/api/argalon`