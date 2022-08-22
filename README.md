# wordCounter

This service will be used to return a word count of a webpage given a URL.
Currently, this service only has the ability to fetch the word count of a URL returning a text file or equivalent.
But there is scope to start fetching the relevant textual content from all webpages by parsing html.

## Documentation
TODO: Add useful links here e.g. to confluence docs

## Owned By

Eunice B

## Dashboards & Metrics

TODO: Add links to dashboards for dev, qa and prod. Add links for metrics to track for this service.

## Production Alerts

TODO: Link and explain how the prod alerts work.

## Project Structure

TODO: explain the project structure.

## Dependencies

TODO: Add dependencies

## How it works

This service scans the web response word by word and preprocesses them before adding to a counter.

## Spec

GET `/count` returns the word counts of a webpage from a URL passed through the request body


## How to run locally

Execute `make run` on your local then run GET request via Postman or curl request like below.

#### cURL request Example
`curl --location --request GET 'http://localhost:8080/count' \
--header 'Authorization: some-token' \
--header 'Content-Type: text/plain' \
--data-raw '{"url": "https://norvig.com/big.txt"}'`

## Make commands

`make lint` analyses  code for potential errors.

`make run` runs the server.

## Contribution guidelines

Any steps required before submitting a pull request for this service. This includes the following:

*  running all tests (unit and integration where applicable)
*  running `go mod tidy`
*  running `make lint`

## Limitations

- add more tests
- add mock for sample web response
- containerise with Docker

## Roadmap

- Optimise with goroutines and file chunking, perhaps store in DB first?
- Test other approaches efficiency e.g. reading a file in chunks concurrently,reading the entire file into memory,splitting long strings into words
- HTML parsing for relevant text only
- pagination of response
- sorting of response
