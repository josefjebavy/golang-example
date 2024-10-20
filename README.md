# golang-example
example for programing language Go

## build && run
``
go run .
``
or
`` 
go run hello.go
``

Hello, World!

## build
``
go build .
``
or
``
go build hello.go
``

## run
``./hello``



## 04-rest-api
How to manage items in the app:


Create new item:

``
curl -X POST http://localhost:8080/items -H "Content-Type: application/json" -d '{"name":"Item1", "value":"Value1"}'
``

List all items:

``
curl http://localhost:8080/items
``

Show item with ID 1:

``
curl http://localhost:8080/items/1
``

Update item ID 1:

``
curl -X PUT http://localhost:8080/items/1 -H "Content-Type: application/json" -d '{"name":"UpdatedItem1", "value":"UpdatedValue1"}'
``

Deltete item ID 1:

``
curl -X DELETE http://localhost:8080/items/1
``

