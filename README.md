# svchk

svchk is a tool written to monitor the status of a list of websites and send notifications when they're offline.


### Usage
Download the release binary and add it to your path.

Change the config in the `~/.svchk.yml` file.

Whenever the program start without a config it will create populate the `~/.svchk.yml` file with the example config.

### Development
Clone the project and run `go get` inside the repo to install the dependencies.

Run it with `go run main.go`

Build with `go build main.go`
