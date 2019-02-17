# Event Rate Converter

e.g, Service A is generating 1.2 errors per second, how many errors per day is that?
```bash
-> event-rate-converter 1.2/s d
-> 103,680
```
How about per 5 minutes?
```bash
-> event-rate-converter 1.2/s 5m
-> 360
```

#### _neat._

## Usage

```
-> event-rate-converter SOURCE_RATE TARGET_RATE
-> OUTPUT
```