# REST API #
[![Go Report Card](https://goreportcard.com/badge/github.com/ingvaar/grapi)](https://goreportcard.com/report/github.com/ingvaar/grapi)
[![Build Status](https://travis-ci.com/ingvaar/grapi.svg?branch=master)](https://travis-ci.com/ingvaar/grapi)

A simple REST API in Go that works with SQL database and a basic cache handling (Redis).
**Work in progress**.

## Usage ##

Configure it via the `config.json` file.

To build the routes, you need to edit `routes.json`.

There are 3 keywords to use in the url for the routes:
* `{type}`: used by the cache handlers.
* `{table}`: tell the API to search for the table passed here, used by the database handlers.
* `{id}`: tell the API to search for the id passed here, used by the database and cache handlers. 

### Handlers ###

Currently, there are *12* handlers:

  ***Database handlers***:

   * **status**: returns the status of the api as basic html text.
   * **index**: says welcome !
   * **select**: returns a SQL line, defined by the id and the table passed in the url, as a json array and html code `200`       if successful.
   * **insert**: creates a new SQL line in the table passed in the url and returns html code `201` if successful.
   * **delete**: deletes the SQL line located at the id in the table passed in the url and returns html code `201` if             successful.
   * **update**: updates the SQL line located at the id in the table passed in the url and returns html code `200` if             successful.
   * **show**: describe the SQL table passed in the url.
   * **login**: returns a json token uppon connection.
    
  ***Cache handlers***:
  
   * **set**: add or update an entry in redis located by `{type}/{id}` passed in the url with an `HMSET` command and returns     html code `200` if successful.
   * **exists**: check if the specified entry exists and return `200` if successful.
   * **read**: returns the entry located by `{type}/{id}` passed in the url as a json array.
   * **remove**: deletes the entry located by `{type}/{id}` passed in the url as a json array.


If you want to know more about the possible flags, simply use:

```sh
$ ./grapi --help
```
