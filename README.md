# BareBone App Structure

This is my example barebone app structure for building backend service.

This example app has the following features:

- HTTP Server (using [Echo](https://echo.labstack.com/))
- Database (using [GORM](https://gorm.io/))
- Cache (using [go-redis](https://github.com/go-redis/redis))
- Scheduler (using [Cron](https://github.com/robfig/cron))
- Queue (TODO)

## How to run

```shell
$ go run cmd/main.go
```

### Configuration

This app have some configurable setting that lies in `.env`.
You can take a look at `.env.example` first or [env config](internal/config/env.go).
