# Testing out the JSON Web Tokens for golang

start postgres docker

```bash
docker run -P -e POSTGRES_PASSWORD=mysecretpassword -d postgres
```

Local port

```bash
PG_PORT=$(docker container inspect $(docker ps --filter ancestor=postgres --format '{{.ID}}') | jq -r '.[0].NetworkSettings.Ports."5432/tcp"[0].HostPort')
```

cli postgres

```bash
docker exec -it $(docker ps --filter ancestor=postgres --format '{{.Names}}') psql -U postgres
```


```bash
go run *.go
```

## View the main page

```bash
curl localhost:5000
```

## Login as a user

```bash
TOKEN=$(curl -u cmcauliffe:c -X POST localhost:5000/login)
```

## Retrieve resource for logged in user

```bash
curl -H "Authorization: Bearer $TOKEN" localhost:5000/my-page
```
