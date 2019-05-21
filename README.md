# Pixelate Face written in Go

### Requirement

- [Go](https://golang.org/dl/)
- [GoCV and OpenCV](https://gocv.io/getting-started/)

### Usage

1. `go get` this repository.

```
$ go get -u github.com/matsuyoshi30/gopix
```

2. Run command.

```
$ gopix [path/to/image]
```

Output image is named `output_[yyyymmddhhmmss].jpeg`.

This is `gopix test/Lena.png` result.

![](./test/Lena_output.jpeg)


### Thanks

Thanks to developers of [GoCV](https://gocv.io/).

### License

[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0.html).
