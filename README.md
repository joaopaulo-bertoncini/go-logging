# Logging in Go

Logging is an important aspect of software development, and there are many libraries available in Go for logging, such as Zap, Logrus, Zerolog, and go-kit/log. Each of these libraries has its own unique features and trade-offs, so choosing the right library for your needs depends on the specific requirements of your application.

Zap is a high-performance logging library that provides a simple and flexible API, making it a popular choice for logging in Go.

Logrus is a logging library that provides a more feature-rich API than Zap, and also supports structured logging, making it easier to search and analyze log data.

Zerolog is a logging library that provides a simple and flexible API, and is designed for high-performance logging with JSON output.

go-kit/log is a logging library that is part of the go-kit toolkit, and provides a flexible and scalable logging framework for Go applications.

In general, choosing the right logging library for your needs depends on the specific requirements of your application, such as performance, logging format, and the level of customization you need.

## running 

````
go run logrus/logrus.go 
````

````
go run zerolog/zerolog.go
````

````
go run zap/zap.go
````

````
go run standard/standard.go
````

````
go run go-kit-log/go-kit-log.go
````