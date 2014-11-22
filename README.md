# Processes go package

[![GoDoc](https://godoc.org/github.com/iltempo/processes?status.svg)](https://godoc.org/github.com/iltempo/processes)

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

