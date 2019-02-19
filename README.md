# Rate Converter

e.g, Service A is generating 1.2 errors per second, how many errors per day is that?
```bash
-> rate-converter 1.2/s d
-> 103,680
```
How about per 5 minutes?
```bash
-> rate-converter 1.2/s 5m
-> 360
```

e.g, Running a t2.large EC2 instance will cost me how much per 30 days?
```bash
-> rate-converter 0.0928/h 30d
-> 66.816
```

#### _neat._

## Usage

```
-> rate-converter SOURCE_RATE TARGET_RATE
-> OUTPUT
```

### Help
```bash
-> rate-converter --help
```

## Development
1. Fetch the single dependency used with `go get golang.org/x/text/message`
2. Use `goreleaser` to package distribution binaries