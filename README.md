# audio-stream-golang
Audio streaming service for educational purposes

# How to start

```bash
docker compose up
```
`First start can take a while!`
## Endpoints
You can reach all endpoints and test it via swagger
`localhost:8080/swagger/`

# Helpful commands
- Update swagger docs
``` bash
swag init -g .\routes\routes.go
```
or
```bash
swag init
```

- Start project (without hot reload)
```bash
go run .\main.go
```
- Start project (with hot reload)
```bash
air
```
- Start project (with all services + hot reload)
```bash
docker compose up
```
