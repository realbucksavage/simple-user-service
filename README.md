# Simple User Service

A simple program that provides a dummy user data on wire.

```
Usage of simple-user-service:
  -grpc-port int
        Specify the port number on which the gRPC server will listen. 0 or no value means any random free port.

```

The program can be executed with `go run . -grpc-port=<a port number>`. The protocol buffer file used to generate the service is at `./proto`.
However, the generated code can be imported in your programs like this:

```go
import (
    "fmt"
    "github.com/realbucksavage/simple-user-service/generated/users"
)

func GetUserByID(id string) {

    // .. create client connection

    client := users.NewUserServiceClient(conn)
    u, err := client.GetUser(...)
    if err != nil {
        panic(err)
    }

    fmt.Printf("resolved user %q\n", user.Name)
}
```
