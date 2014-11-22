# Processes go package

Package that provides a list of running processes in the system. It's quite basic and 
illustrates my first steps with go and it's packaging system.

## Installation

    go get github.com/iltempo/processes

## Usage

To retrieve a list of processes of the current user

```go
import "github.com/iltempo/processes"

func main() {
    procs := processes.OfCurrentUser()
    ...
}
```

