# Bloom Go

Simple bloom filter in go.

## Sample
```go
    package main

    import (
        "fmt"
        "github.com/oldcodemonk/bloomgo"
    )

    func main() {
        bloomFilt := New(10000, 0.01)
	    bloomFilt.Add([]byte("bob"))
        bloomFilt.Add([]byte("sam"))
        bloomFilt.Add([]byte("mary"))

        fmt.Println(bloomFilt.Exists([]byte("bob")))
        fmt.Println(bloomFilt.Exists([]byte("abc")))
    }
```
