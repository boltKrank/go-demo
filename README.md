# go-demo

Demonstrating golang 

Tutorial page: "https://golang.org/doc/tutorial/handle-errors"

## Module path structrure

Not sure how the modules are arranged - need to look at the way dependancy trees are calculated.

NOTES:

Module structure

```text

    /home/user/go/
        src/
            crash/
                bang/              (go code in package bang)
                    b.go
            foo/                   (go code in package foo)
                f.go
                bar/               (go code in package bar)
                    x.go
                vendor/
                    crash/
                        bang/      (go code in package bang)
                            b.go
                    baz/           (go code in package baz)
                        z.go
                quux/              (go code in package main)
                    y.go

```
