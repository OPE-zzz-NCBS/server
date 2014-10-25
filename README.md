# Quick Start

The OPENCBS server is written in Go and this section assumes that you have installed it on your computer. If it is not the case go to the [official Go website](http://golang.org), download, and install it.

To get started, you need to:

1. Set up the `$GOPATH` environment variable
2. Build
3. Configure
4. Make sure MS SQL Server allows remote connections
5. Run the server

## Setting up `$GOPATH`

Configuring the Go development environment is pretty much setting up the `$GOPATH` variable. You can learn more about it in the [How to Write Go Code](https://golang.org/doc/code.html) article. But here is a condensed version of it.

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

In *Windows* go to *Start*, right-click *Computer* and choose *Properties*. Open up the *Advanced system settings* window and click the *Environment Variables* button. Click *New*, type in `GOPATH` in the variable name field, and enter a valid path in the variable value field. Close all the windows by clicking *OK*.

## Building

Getting the source code and building it could not be easier with Go. Just run this command:

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

When the server starts it reads in configuration options from the `conf.json` file. This file *does not* exist by default, but you can create it by copying from the template:

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

As you see, you can adjust the port that the server will listen on and the database connection.

There is a couple of things to be aware of when configuring the server. First, if the database host includes a backslash, you should escape it with an extra backslash as in the sample. Second, you have to make sure that the MS SQL Server instance on your computer accepts external connections.

## Allowing remote connections

Just before running the OPENCBS server you want to make sure that the MS SQL Server instance allows remote connections. This involves a couple of steps.

First, if you run MS SQL Server 2008 or 2008 R2 you have to find out if it is upgraded to Service Pack 2. Run this query in the MS SQL Management Studio:

```sql
select @@version
```

If it says "Service Pack 2" then you are good. Otherwise, you have to download and install it.

Second, you want to explicitly allow remote connections to the MS SQL Server instance.

## Running

To run the server, execute this command:

```bash
$ ./server
```

If everything went smoothly, the server will start accepting connections. Otherwise, it will print out an error message and exit.

