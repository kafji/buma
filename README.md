# Buma

Buma is a web server (backend server) for an RSS feed reader.

## Test

To run tests you need Docker and `postgres:14` image, then run `./test.sh`.

## Dev Run

### Create configuration file

```bash
cp ./example.buma.toml ./buma.toml
```

### Run Postgres container

```bash
docker run -e POSTGRES_PASSWORD=password -p5432:5432 -it postgres:14
```

### Run `buma-http`

```bash
go run ./cmd/buma-http
```
