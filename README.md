# Testing out the JSON Web Tokens for golang

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
