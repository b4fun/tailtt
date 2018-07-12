# tailtt

Tail, then that.

## Usage

```
$ tailtt \
  -f /var/log/nginx.log \
  -w "holy keyword" \
  -notify-bearychat-channel "<BEARYCHAT-CHANNEL>" \
  -notify-bearychat-rtm "<BEARYCHAT-RTM-TOKEN>"
```

## Build

```
$ make build
```

### Build with env vars

```
$ export BEARYCHAT_RTM_TOKEN=xxx
$ export BEARYCHAT_RTM_CHANNEL=yyy
$ make build
```

## LICENSE

MIT
