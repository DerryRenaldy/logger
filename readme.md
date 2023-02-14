# 🗳 Logger Package

## Installation
```go
go get github.com/DerryRenaldy/logger
```

## Leveled Logging
With this Package you have 6 logging level that you can implement (from highest to lowest) :
* panic / panicf (constant.PanicLevel)
* fatal / fatalf (constant.FatalLevel)
* error / errorf (constant.ErrorLevel)
* warn / warnf (constant.WarnLevel)
* info / infof (constant.InfoLeve)
* debug / debugf (constant.DebugLevel)

## How to use

```go
// Create New Logger Object
log := logger.New("serviceName", "environment", "logLevel")
```

**Log Debug Level**
```go
log.Debug("Your Message")
```

**Output Log**
```json
{"level":"debug","service":"test","env":"development","file":"/home/derryrenaldy/Desktop/zerolog_test/test_log_package/test_logger/main.go","function":"main.main","line":10,"time":"2023-02-14T10:42:58+07:00","message":"Your Message"}
```

**Log Debug Level with formatting**
```go
message := "Your Message"
log.Debugf("Your Message is : [%s]",message)
```

**Output Log**
```json
{"level":"debug","service":"test","env":"development","file":"/home/derryrenaldy/Desktop/zerolog_test/test_log_package/test_logger/main.go","function":"main.main","line":11,"time":"2023-02-14T10:57:17+07:00","message":"Your Message is : [Your Message]"}
```