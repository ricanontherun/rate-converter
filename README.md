# Rate Converter

e.g, Service A is generating 1.2 errors per second, how many errors per day is that?
```bash
-> rate-converter -source=1.2/s -target=d
-> 103,680
```
How about per 5 minutes?
```bash
-> rate-converter -source=1.2/s -target=5m
-> 360
```

e.g, Running a t2.large EC2 instance will cost me how much per 30 days?
```bash
-> rate-converter -source=0.0928/h -target=30d
-> 66.816
```
I want to print the result with 4 decimal places of precision.
```bash
-> rate-converter -source=0.0144/h -target=30d -precision=4
-> 10.3680
```

#### _neat._

## Usage

First, grab the binaries from the releases page.

```
-> rate-converter -source=SOURCE_RATE -target=TARGET_RATE
-> OUTPUT
```

## Development
2. Use `goreleaser` to package distribution binaries
