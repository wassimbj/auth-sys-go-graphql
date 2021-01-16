> this is a simple signup/login and logout system, **using GO, GraphQL, Postgres and Docker**

#### Development
first of all, clone the repo.  

cd into the `/server` and then run

```bash
go mod download
```

now cd into `/server/docker` and run

```bash
docker-compose up
```

if everything goes fine you can go to `http://localhost:8080` and you shall see the GraphQL interface

here is an example of mutations available
