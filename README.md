# eKYC-backend

1. [ Requirements and Assumptions](#requirements)
2. [ Getting Started ](#getting-started)
3. [ Project Structure. ](#project-structure)
4. [ Swagger Document. ](#swagger-document)

<a name="getting-started"></a>
## Getting Started

### Pre-requisite
1. GoLang version 1.18+
2. Docker

### Running as Binary
1. To run this application you need to install go 1.18
2. Execute `go install` in your terminal to download all dependenc
3. To run as standalone you can run `go run .`
4. Or you can build the binary by executing `go build -o app` and this will generate `app` binary (it will be app.exe if you run this in windows)
5. Then you can execute the binary `./app`

### Running as Docker

1. Build the docker image using multistage build  
`docker build -f Dockerfile.multistage -t seek-checkout .`

2. Run docker image. Note that the default port is 4545 if you like to change the port please read the __Application Config__ section
`docker run -p 8080:8080 seek-checkout`

### System Endpoint
1. `/health` use for health check.
2. `/swaggerui` to open up swagger document. To read more about how this project configure the swagger doc please read the __Swagger Document__ section

### Sample Run

After you run all the components with `make up` try the following curl command
```
curl --location --request POST 'http://localhost:4545/v1/session/start' \
--header 'x-b3-traceId: asd314' \
--header 'x-b3-spanid: 123' \
--header 'x-b3-parentspanid: zxc123' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "alex@gmail.com",
    "reference": "000000000"
}'
```
You should be able to get response header as following
```
Expiration-Time : 1658794375
Set-Cookie: eKYC-Session=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmaXJzdE5hbWUiOiJhbGV4IiwiY291bnRyeSI6IkFVIiwiZXhwIjoxNjU4Nzk0Mzc1fQ.D0f6GNU53foJi_8PhC9PmGblWKF8AvMBHc5l22t6xG4; Path=/; Domain=localhost; Max-Age=1658794375; HttpOnly

```

<a name="running-tests"></a>
## Running Tests
### Unit Test

1. `make test-unit` run unit test and will generate coverage report 'coverage.out' in the root folder
2. `make test-html-result` this will generate human readable html report base on 'coverage.out'. Therefore you need to run `make test-unit` before running this command
2. `make test-api` run API test. This will do following things:
   1. stop/remove existing old running docker containers
   2. build/running new containers 
   3. running the API testing
3. `make retest-api` re-run the API tests in existing running docker   
4. `make test-all` run all testings include API testing
5. `make test-html-result` build/open up html coverage report

<a name="project-structure"></a>
## Project Structure

`/api_test` contains APT test files
`/ci` contains Codefresh CI scripts
`/client` contains all http client with downstream services eg. COBRA
`/controller` contains all gin RESTFul HTTP handler
`/middleware` contains all gin middleware handler
`/mocks` contains all mock classes that are used for testing. Some are generated using mockgen
`/repository` contains all objects relate to database access
`/resource` contains request and response objects
`/security` contains certs and key
`/security_local` contains local self-signed cert and key
`/sequence_diagram` contains sequence diagram
`/server` contains service route configration and server startup sequence
`/service` contains services function that help HTTP handler
`/stub` contains stub services use for API testing
`/swagger` contains swagger files
`/utils` contains utilites/helper function

<a name="swagger-document"></a>
## Swagger Document

### Open Swagger UI
1. To open swagger document go to this path `/swaggerui/` with your web browser.
For example `https://localhost:4545/swaggerui/`
2. This path is configured under `router.go`. Take a look at the file and you will find the configuration for `/swaggerui/*any` route