# startup-pushover

This tool sends the hostname and all (hopefully) relevant IP addresses via [pushover](https://pushover.net/).
See below for setup instructions for [systemd](https://freedesktop.org/wiki/Software/systemd/)
based Linux distributions with a network configuration backend with proper support for
`network-online.target` (like **Network Manager** or **systemd-networkd**). Instructions for other operating systems are welcome, please send
pull requests .

# Building

Install [Go](https://golang.org).

Get the source and all dependecies:
```
go get gitlab.com/hreese/startup-pushover
```

Build for local plattform:
```
go build -ldflags="-s -w" gitlab.com/hreese/startup-pushover
```

To cross compile, find [the proper set of variables](https://golang.org/doc/install/source#environment)
and prepend them to the `build` command. For example, to build for Raspberry Pi 1, use
```
GOOS=linux GOARCH=arm GOARM=5 go build -ldflags="-s -w" gitlab.com/hreese/startup-pushover
```

# Installation

Run everything as root (or prepend `sudo` to all commands).

Install the executable:
```
cp startup-pushover /usr/local/bin/
```

Copy the example config file:
```
cp example.startup-pushover.json /etc/startup-pushover.json
```

Get an API token and your recipient ID and fill them in:
```
editor /etc/startup-pushover.json
```

Copy the systemd service and enable it:
```
cp startup-pushover.service /etc/systemd/system
systemctl daemon-reload
systemctl enable --now startup-pushover.service
```
