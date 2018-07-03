# Call Go binary in Ruby using ffi sample

## Environment

macOS High Sierra
Version 10.13.3

```bash
$ ruby -v                                                                                                                                                                     Tue Jul  3 23:51:45 2018
ruby 2.5.0p0 (2017-12-25 revision 61468) [x86_64-darwin17]
```

```bash
$ go version                                                                                                                                                          161ms  Tue Jul  3 23:52:06 2018
go version go1.10.3 darwin/amd64
```

## Build

```bash
$ go build -o sum.so -buildmode=c-shared sum.go
```

## Execute

```bash
$ ruby sum.rb
```
