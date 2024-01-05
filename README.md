# Global Ip Checker

## What is it?

Simply, this CLI tool is to `watch` your global ip and notify you if there has been any changes.

It uses the [ipify](https://www.ipify.org/) api to check the current expose global ip.

## Dependencies

This is written in `GO 1.20` so you might have to have that installed.

## Build

```
go build .
```

Should compile and give you the correct executable.

But if you dont really care, you can just `go run main.go ~` and that should work fine as well.

## Usage

### Current ip

This will get you your current ip address.

```
$ global-ip-checker current
```

### Watch

This will tell the program to watch for changes in your current ip

```
$ global-ip-checker watch "your.ip.address.here`
```

#### Rate

By default it will poll the `ipify` endpoint every minute (60 seconds)

But you can specify the rate of polling with the `-r` flag

```
$ global-ip-checker watch "your.ip.address.here" -r 60
```

Rate is per seconds. So if you need to limit the polling to 5 minutes you need to set it for `300`

### Notifications

Global Ip Checker uses the [beeep](https://github.com/gen2brain/beeep) library for notifications.  
If it detects any changes you will get a notification message.

## Disclaimer

**This works on my computer (macOS) and I haven't tested it on windows or any other OS so you might find that it doesn't work for you so.... sorry** :bow: :sweat_smile:

**There are probably a bunch of bug with this code so if you would like fix them, please feel free to open a Pull Request :smile:**

## TODO

WIP
