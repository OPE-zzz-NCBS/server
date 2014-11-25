# Quick Start

The OPENCBS server is written in Go and therefore you must have Go installed on your system. If you don't, head to the [official Go website](http://golang.org) to download and install it.

To get started, you need to:

1. Set up the `$GOPATH` environment variable
2. Build
3. Configure
4. Make sure MS SQL Server allows remote connections
5. Run the server

## Setting up `$GOPATH`

Configuring the Go development environment is pretty much the same as setting up the `$GOPATH` variable. For an in depth discussion of the topic, check out this article: [How to Write Go Code](https://golang.org/doc/code.html). In the meantime, what follows are the condensed instructions.

If you are on *Linux* or *OS X* run these commands:

```bash
$ mkdir $HOME/go
$ export GOPATH=$HOME/go
$ export PATH=$PATH:$GOPATH/bin
```

To save the variable accross terminal sessions add these lines into the `.bash_profile` file:

```bash
export GOPATH=$HOME/go
export PATH="$PATH:$GOPATH/bin"
```

In *Windows* go to *Start*, right-click *Computer* and choose *Properties*. Open up the *Advanced system settings* window and click the *Environment Variables* button. Click *New*, type `GOPATH` in the variable name field, and enter a valid path in the variable value field. Close all windows by clicking *OK*.

## Building

Getting and rebuilding source code could not be any easier with Go. Just run this command:

```bash
$ go get github.com/OPENCBS/server
```

This will clone the repository, fetch all the dependencies, and build the project.

The repository resides in `$GOPATH/src/github.com/OPENCBS/server`. If you want to rebuild it, run these:

```bash
$ cd $GOPATH/src/github.com/OPENCBS/server
$ go build .
```

## Configuring

When the server starts it reads the configuration options int the `conf.json` file. This file *does not* exist by default, but you can create it by copying from the template:

```bash
$ cp conf.sample.json conf.json
```

The file looks like this:

```
{
  "Server": {
    "Port": 8080
  },
  "Database": {
    "Host": "localhost\\sqlexpress",
    "Username": "sa",
    "Password": "opencbs",
    "Name": "opencbs"
  }
}
```

As you see, you can adjust the port that the server will listen on as well as the database connection. Please note that if the database hostname contains a backslash, it has to be escaped by an extra backslash.

## Allowing remote connections

Just before running the OPENCBS server, you'll want to verify that the MS SQL Server instance allows remote connections. This involves a couple of steps.

First, if you run MS SQL Server 2008 or 2008 R2 you have to find out if it is upgraded to Service&nbsp;Pack&nbsp;2. Run this query in the MS SQL Management Studio:

```sql
select @@version
```

If it says "Service&nbsp;Pack&nbsp;2" then you are good. Otherwise, you have to download and install it.

Second, you need to explicitly allow remote connections to the MS SQL Server instance. A quick search on Google reveales a number of [detailed articles](http://blogs.msdn.com/b/walzenbach/archive/2010/04/14/how-to-enable-remote-connections-in-sql-server-2008.aspx) on [how to do this](http://stackoverflow.com/questions/11278114/enable-remote-connections-for-sql-server-express-2012).

## Running

To run the server, execute this command:

```bash
$ ./server
```

If there weren't any issues, the server will start accepting connections. If there were issues, it will display an error message and exit.

