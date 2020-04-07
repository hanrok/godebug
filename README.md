
# Go Debugging for Intellij GoLand

## Example
Clone the repo
```shell script
git clone https://github.com/hanrok/godebug.git
```

Build the micro-service example
```shell script
docker build -t godebug .
```

Run the service
```shell script
docker run --rm -v $PWD:/build -p8080:8080 -p40000:40000 godebug
```

You should see output like this:
```text
Running server
API server listening at: [::]:40000
Delve PID: 8, Server PID: 17
Setting up watches.  Beware: since -r was given, this may take a while!
Watches established.
```

After you see this output, you need to create a new remote debug configuration
in Goland (described in details in [here](https://medium.com/@hananrok/debugging-hot-reloading-go-app-within-docker-container-b44d2929e8bd)).

Then run the debug (by click green beetle or Ctrl-D). Container output should now 
be with this line:
```text
2020/03/05 20:13:32 Starting up repeater
```

Whenever you change your code (`main.go` file) it will rebuild the image and rerun the server.

