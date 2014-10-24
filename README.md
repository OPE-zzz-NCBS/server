## Quick Start

Since the server is written in Go you want to make sure the `$GOPATH` environment variable is properly set up. You can learn more about it in the [How to Write Go Code](https://golang.org/doc/code.html) article. But here is a condensed version of it.

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

In *Windows* go to `Start`, right-click `Computer` and choose `Properties`. Open up the `Advanced system settings` window and click the `Environment Variables...` button. Click `New...`, type in `GOPATH` in the variable name field, and enter a valid path in the variable value field. Close all the windows by clicking `OK`.

Open a terminal and type in:

```bash
echo $GOPATH
```

You should see a full path.
