# To-Do-List

A REST API to manage a to-do list.


## Prerequisites

[Install Go](https://go.dev/doc/install)

Clone the repository:

```
$ git clone https://github.com/kdonthi/To-Do-List.git
```

## Unit Test

```
$ make test
```

## Start Application

```
$ make run
```
or 

```
$ go run main.go [OPITONAL_PORT]
```

The port defaults to 9000.

## Endpoints

### POST Create

**Request**
```
$ curl -X POST http://localhost:9000/create -H "Content-Type: application/json" -d '{"item":"Do the dishes"}' && echo ""
$ curl -X POST http://localhost:9000/create -H "Content-Type: application/json" -d '{"item":"Mow the lawn"}' && echo ""
$ curl -X POST http://localhost:9000/create -H "Content-Type: application/json" -d '{"item":"Feed the dog"}' && echo ""
```

**Response**
```
$ {"id":1,"item":"Do the dishes"}
$ {"id":2,"item":"Mow the lawn"}
$ {"id":3,"item":"Feed the dog"}
```

### GET Read (requires path parameter)

**Request**
```
$ curl -X GET http://localhost:9000/read/1 -H "Content-Type: application/json" && echo ""
$ curl -X GET http://localhost:9000/read/3 -H "Content-Type: application/json" && echo ""
```

**Response**
```
$ {"id":1,"item":"Do the dishes"}
$ {"id":3,"item":"Feed the dog"}
```

### GET Read All

**Request**
```
$ curl -X GET http://localhost:9000/read -H "Content-Type: application/json" && echo ""
```

**Response**
```
$ [{"id":1,"item":"Do the dishes"},{"id":2,"item":"Mow the lawn"}]
```

### PUT Update (requires path parameter)

**Request**
```
$ curl -X PUT http://localhost:9000/update/1 -H "Content-Type: application/json" -d '{"item":"Wipe the windows"}' && echo ""
```

**Response**
```
$ {"id":1,"item":"Wipe the windows"}
```

### DELETE Delete (requires path parameter)

**Request**
```
$ curl -X DELETE http://localhost:9000/delete/2 -H "Content-Type: application/json" && echo ""
```

**Response**
```
$ {"id":2,"item":"Mow the lawn"}
```

### DELETE Delete All

**Request**
```
$ curl -X DELETE http://localhost:9000/delete -H "Content-Type: application/json" && echo ""
```

**Response**
```
$ [{"id":1,"item":"Wipe the windows"},{"id":2,"item":"Feed the dog"}]
```
  
