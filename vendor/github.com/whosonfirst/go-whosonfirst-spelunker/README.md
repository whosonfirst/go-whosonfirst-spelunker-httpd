# go-whosonfirst-spelunker

## Documentation is incomplete at this time.

## Motivation

This is a refactoring of both the [whosonfirst/whosonfirst-www-spelunker](github.com/whosonfirst/whosonfirst-www-spelunker) and [whosonfirst/go-whosonfirst-browser](github.com/whosonfirst/go-whosonfirst-browser) packages.

Specifically, the former (`whosonfirst-www-spelunker`) is written in Python and has a sufficiently complex set of requirements that spinning up a new instance is difficult. By rewriting the spelunker tool in Go the hope is to eliminate or at least minimize these external requirements and to make it easier to deploy the spelunker to "serverless" environments like AWS Lambda or Function URLs. The latter (`go-whosonfirst-browser`) has developed a sufficiently large and complex code base that starting from scratch and simply copying, and adapting, existing functionality seemed easier than 

There are three "classes" of `go-whosonfirst-spelunker` packages:

### go-whosonfirst-spelunker

That would be the package that you are looking at right now. It defines the [Spelunker](#) interface and exposes a package library for use by other packages to create a "spelunker-like" command-line tool.

This package does not export any working implementations of the `Spelunker` interface. It simply defines the interface and other associated types.

### go-whosonfirst-spelunker-httpd

The [whosonfirst/go-whosonfirst-spelunker-httpd](github.com/whosonfirst/go-whosonfirst-spelunker-httpd) package provides libraries for implementing a web-based spelunker service. While it does define a working `cmd/server` tool demonstrating how those libraries can be used, like the `go-whosonfirst-spelunker` package it does not export any working implementations of the `Spelunker` interface. 

### go-whosonfirst-spelunker-sqlite

The [whosonfirst/go-whosonfirst-spelunker-sqlite](github.com/whosonfirst/go-whosonfirst-spelunker-sql) package implements the `Spelunker` interface using a Go `database/sql` relational database source, for example SQLite databases produced by the [whosonfirst/go-whosonfirst-sqlite-features-index](https://github.com/whosonfirst/go-whosonfirst-sqlite-features-index) package.

It imports both the `go-whosonfirst-spelunker` and `go-whosonfirst-spelunker-httpd` and exports local instances of the "spelunker" command-line tool and web-based server. For example, to create a database for use by the SQLite implementation of the `Spelunker` interface:

```
$> cd /usr/local/whosonfirst/go-whosonfirst-sqlite-features-index
$> ./bin/wof-sqlite-index-features-mattn \
	-timings \
	-database-uri mattn:///usr/local/data/ca.db \
	-spatial-tables \
	-ancestors \
	-search \
	/usr/local/data/whosonfirst-data-admin-ca
```

And then to use that database with a local (`go-whosonfirst-spelunker-sql`) instance of server code exported by the `go-whosonfirst-spelunker-httpd` package:

```
$> cd /usr/local/whosonfirst/go-whosonfirst-spelunker-sql
$> ./bin/server \
	-server-uri http://localhost:8080 \
	-spelunker-uri sql://sqlite3?dsn=file:/usr/local/data/ca.db
```

This is what the code for the server tool looks like (with error handling omitted for the sake of brevity):

```
package main

import (
        "context"
        "log/slog"

        _ "github.com/mattn/go-sqlite3"
        "github.com/whosonfirst/go-whosonfirst-spelunker-httpd/app/server"
        _ "github.com/whosonfirst/go-whosonfirst-spelunker-sql"
)

func main() {
        ctx := context.Background()
        logger := slog.Default()
        server.Run(ctx, logger)
}
```

## See also

* github.com/whosonfirst/go-whosonfirst-spelunker-httpd
* github.com/whosonfirst/go-whosonfirst-spelunker-sql