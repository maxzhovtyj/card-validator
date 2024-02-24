### Card-Validator by Maksym Zhovtaniuk

---

#### In case of any questions tg: @maxzhovtyj

Run grpc server in docker

```shell
make start 
```

### Test API:

---

grpc_cli:
```shell
grpc_cli call 127.0.0.1:7799 Validate 'card: {number: "1111" expirationYear: 2023 expirationMonth: "12"}'
```

#### OR

grpcurl
```shell
grpcurl -plaintext -d '{"card": {"number": "1111", "expirationYear": 2023, "expirationMonth": "12"}}' 127.0.0.1:7799 CardService/Validate
```

#### OR
```shell
make test-api
```
or use `cmd/client/main.go` file
