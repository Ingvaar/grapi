# REST API
[![Go Report Card](https://goreportcard.com/badge/github.com/Ingvaar/rest)](https://goreportcard.com/report/github.com/Ingvaar/rest)
[![Build Status](https://travis-ci.org/Ingvaar/rest.svg?branch=master)](https://travis-ci.org/Ingvaar/rest)

A simple REST API in Go that works with Redis and/or SQL database.
**Work in progress**.

## Usage

First, build the API with:

```sh
$ make deps && make
```

And configure it via `config.json` file.

To build the routes, you need to edit `routes.json`.

There are 3 keywords to use in the url for the routes:
* `{type}`: used by the redis handlers.
* `{table}`: tell the API to search for the table passed here.
* `{id}`: tell the API to search for the id passed here. 

Currently, there are *9* handlers:
* **status**: returns the status of the api as basic html text.
* **index**: says welcome !
* **getTableSQL**: returns a whole SQL table as a json array and html code `200` if successful.
* **getLineSQL**: returns a SQL line, defined by the id and the table passed in the url, as a json array and html code `200` if successful.
* **createLineSQL**: creates a new SQL line in the table passed in the url from a json array passed in the body of the request, and returns html code `201` if successful.
* **deleteLineSQL**: deletes the SQL line located at the id in the table passed in the url and returns html code `201` if successful.
* **updateLineSQL**: updates the SQL line located at the id in the table passed in the url, with the json passed in the body, and returns html code `200` if successful. It can update only one value (exemple: the table contains an id, a name, and a commentary, you can update only the name with `{"name":"NewValue"}`).
* **addEntryRedis**: add or update an entry in redis located by `{type}/{id}` passed in the url with an `HMSET` command and returns html code `200` if successful.
* **getAllRedis**: get all content from an entry in redis located by `{type}/{id}` passed in the url with an `HGETALL` command and returns a json array if successful.

Finally, to start the API, use:

```sh
$ ./rest_api
```

If you want to know more about the possible flags, simply use:

```sh
$ ./rest_api --help
```
