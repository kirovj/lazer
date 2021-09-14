# Lazer - Lazy & Async Logger In Go

> It is my personal project, if you really want to use a log library in go, please use others.

## Features

* Dependency Free
* Use [chanx](https://github.com/smallnest/chanx)
* Only Lazy & Async
* Simple Interface

## Getting Started
```go
package main

import "github.com/kirovj/lazer"

func main() {
    log := lazer.Default()
    for i := 0; i < 100; i++ {
        log.Info("one  " + strconv.Itoa(i))
    }

    time.Sleep(10 * time.Second)
    for i := 0; i < 100; i++ {
        log.Error("two  " + strconv.Itoa(i))
    }
}
```

## Todo
* Time Struct
* Xid: a id to trace msg
* Msg Format
* Pretty Console Writer  
...