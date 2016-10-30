# go-xplane

This is an interface to X-Plane for Golang.

## License

This code is licensed under the Apache License 2.0.

## Example
```go
package main

import (
        "github.com/ornen/go-xplane"
        "github.com/ornen/go-xplane/messages"
        "log"
)

func main() {
        x := xplane.New("127.0.0.1:49000", ":49003")
        x.Connect()

        x.Send(messages.ThrottleCommandMessage{
                Throttle: 1,
        })

        x.Send(messages.FlightControlMessage{
                Elevator: 1,
                Aileron:  -1,
                Rudder:   0,
        })

        go x.Receive()

        for {
                message := <-x.Messages

                if message != nil {
                        log.Printf("%+v\n", message)
                }
        }
}
```