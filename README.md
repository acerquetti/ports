# Ports API

## How to run

```shell
$ make docker        # Build Docker container
$ docker compose up  # Run service
```

Perform a request to it and check return code

```shell
$ curl -v -d '{"id":"AEAJM","name":"Ajman","city":"Ajman","country":"United Arab Emirates","alias": [],"regions": [],"coordinates": [55.5136433,25.4052165],"province":"Ajman","timezone":"Asia/Dubai","unlocs":["AEAJM"],"code": "52000"}' localhost:8080/ports
```

## Tests to perform on assignment

### Graceful shutdown

```shell
$ go run ./...   # Start service locally
```

Then press `Ctrl + C` to send interrupt signal and check log.

### Check errors when building Port aggregate

Perform following request

```shell
$ curl -v -d '{"id":"AEAJM"}' localhost:8080/ports
```

Expect HTTP code 400 and message `name cannot be empty` (first condition checked when building aggregate).

### Check Docker container memory usage

Run service with docker compose and check memory usage and limit

```shell
$ docker compose up -d
$ docker stats ports --no-stream --format "{{ json . }}"|jq '.MemUsage'
```

## Further improvements

* Add integration tests from controller.
* Extend unit tests for domain validations and in-memory db.
* Expose operation to get port by filtering on any of its fields.
