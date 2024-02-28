# Movies Microservices with Go

Application for movie lovers.

## Services

```
                       +----------------------+
                       |                      |
          +-------------    Movie Service     |-------------+
          |            |                      |             |
          |            +----------------------+             |
+-------------------+                             +-------------------+
|                   |                             |                   |
| Metadata Service  |                             |  Rating Service   |
|                   |                             |                   |
+-------------------+                             +-------------------+

```

- **Movie metadata service**: Store and retrieve the movie metadata records by movie IDs.
- **Rating service**: Store ratings for different types of records and retrieve aggregated ratings for records.
- **Movie service**: Provide complete information to the callers about a movie or a set of movies, including the movie metadata and its rating.

## Layers

```
    +--------------+      +----------------+      +-------------+
--->| API Handler  | ---> | Business Logic | ---> | Repository  | ---> DB(postgres)
    +--------------+      +----------------+      +-------------+
                              (controller)
```

### Consul-based service discovery

Run locally hashicorp consul:

```
docker run -d -p 8500:8500 -p 8600:8600/udp --name=dev-consul hashicorp/consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
```

Run each microservice by executing this command inside each cmd directory:

```
go run main.go
```

Visit Consul UI via http://localhost:8500.

### Proto VS Json

```
$ make proto
$ go mod tidy
$ make bench
```
