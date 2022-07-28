# REST API & gRPC For Creating TODO Lists on Go

## The following concepts that embodied in the project:
- Developing Go Web Applications Following REST API Design (port 8000).
- Developing gRPC Server (port 9090).
- Registration and authentication. Work with JWT. Middleware.
- The Clean Architecture approach.
- Working with framework <a href="https://github.com/gin-gonic/gin">gin</a>.
- Work with Postgres DB. Launch on Docker. Migration files generation.
- Configure app with library <a href="https://github.com/spf13/viper">viper</a>. Working with environment variables.
- Work on DB with library <a href="https://github.com/jmoiron/sqlx">sqlx</a>.
- Writing SQL queries.
- Writing unit tests with <a href="https://github.com/stretchr/testify">testify</a> and <a href="https://github.com/golang/mock">gomock</a>.
- Swagger with <a href="https://github.com/swaggo/swag">swag</a>.
- Graceful Shutdown

### To build in Docker Compose:

```
make build && make run
```

For first run of app use:

```
make migrate
```

Swagger at (Btw smth wrong with header {only in swagger} - if you know how to fix it, let me know):

```
localhost:8000/swagger/index.html
```

Also don`t forget to create .env file in root directory with some variables:

```
DB_PASSWORD={your DB password}
PASSWORD_SALT={your password salt for sha1 encryption}
JWT_SIGNING_KEY={your JWT signing key}
```