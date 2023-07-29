# BareBone App Structure

You can use this as base app structure for building another service

#### Note:

- This app developed using:
    - Go v1.20.3
    - MySQL v8.0.33
    - Redis 6.0.16

## How to run

```shell
$ make run
```

### Configuration

This app have some configurable setting that lies in `conf.yaml`. But you must generate it first.

To generate config file from config file template, run

```shell
$ make generate-config
```

#### Database Config

You can choose Database Engine by setting `database.engine` in `conf.yaml`. The value must be one of:
- `inmemory`
- `mysql`

Notes:

- For MySQL, as long as the MySQL config is correct, the schema migration will run automatically if you run the app successfully

#### Cache Config

You can choose Cache Engine by setting `cache.engine` in `conf.yaml`. The value must be one of:
- `inmemory`
- `redis`

### Docs

Generate docs using this command:

```shell
$ make generate-docs
```

And access it here http://localhost:1323/docs
