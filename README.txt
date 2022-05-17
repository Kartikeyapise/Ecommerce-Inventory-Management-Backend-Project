#run postgres docker image
docker run -d -p 5432:5432 --name kartikeya -e POSTGRES_PASSWORD=kartikeya postgres

#commands to test the code coverage
go test ./test/service/product_service_test.go  -v -coverpkg=./main/service    -> 87.0%
go test ./test/repository/product_repository_test.go -v -coverpkg=./main/repository    ->62.5%
go test ./test/controller/product_controller_test.go -v -coverpkg=./main/controller    ->93.0%
go test ./test/view/response_message_test.go -v -coverpkg=./main/view    -> 100.0%
go test ./test/entity/product_test.go -v -coverpkg=./main/entity    -> 100.0%
go test ./test/custum_errors/product_error_test.go -v -coverpkg=./main/custum_errors    -> 100.0%

Libraries Used : Mux for Routing, Gorm for database Interactions

Key Features About Application ->
1. 3-layered Architecture is followed. Presentation layer -> Controller, Business logic Layer -> Service, Data access layer -> Repository
2. Application is Testable without changing the state of the System. All the dependencies are Mocked.
3. Test Coverage is greater than 80%.
4. Application in Independent of frameworks.
5. Application is loosely Coupled.Dependencies talk with each other through Interfaces.
6. Code is Modular, Readable, Maintainable and extensible.
7. Application is stateless and thread-safe to allow concurrent requests.
8. SOLID Design Principles are Followed.
9. Errors are handled Smoothly without letting the application crash