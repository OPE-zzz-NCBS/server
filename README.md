## Quick Start

Since the server is written in Go you want to make sure the `$GOPATH` environment variable is properly set up. You can learn more about it in the (How to Write Go Code)[https://golang.org/doc/code.html] article. But here is a condensed version of it.

If you are on *Linux* or *OS X* run these commands:

```
$ mkdir $HOME/go
$ export GOPATH=$HOME/go
$ export PATH=$PATH:$GOPATH/bin
```

To save the variable accross terminal sessions add these lines into the .bash_profile file:

```
export GOPATH=$HOME/go
export PATH="$PATH:$GOPATH/bin"
```

