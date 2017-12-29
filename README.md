# startup-pushover

This tool sends the hostname and all (hopefully) relevant IP addresses via [pushover](https://pushover.net/).
See below for setup instructions for [systemd](https://freedesktop.org/wiki/Software/systemd/)
based Linux distributions. Instructions for other operating systems are welcome, please send
pull requests .

# Building

Install [Go](https://golang.org).

Get the source and all dependecies: `go get gitlab.com/hreese/startup-pushover`

Build for local plattform: `go build -ldflags="-s -w" gitlab.com/hreese/startup-pushover`

To cross compile, find the proper set of variables and prepend them to the `build` command. For example,
to build for Raspberry Pi 1, use
`GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-s -w" gitlab.com/hreese/startup-pushover`
