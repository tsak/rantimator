# Rantimator

Let's howl at the moon!

Message [@RantimatorBot](http://t.me/RantimatorBot) on Telegram.

Find your howl on [rants.tsak.net](https://rants.tsak.net)

## Config

Copy `.env.sample` to `.env` and set the following environment variables:

```bash
TOKEN=<TELEGRAM BOT TOKEN>
ADDRESS=<LISTEN ADDRESS WITH PORT>
DEBUG=1 # Enables debug mode if DEBUG is set to a non-empty value
```

## Build

```bash
go build
```

## Run

```bash
./rantimator
```

## Run as a systemd service

See [rantimator.service](rantimator.service) systemd service definition.

To install (tested on Ubuntu 16.04):

1. `adduser rantimator`
2. copy `rantimator` binary as well as `.env` file to `/home/rantimator`
3. place systemd service script in `/lib/systemd/system/`
4. `sudo systemctl enable rantimator.service`
5. `sudo systemctl start rantimator`
6. `sudo journalctl -f -u rantimator`

The last command will show if the service was started.
