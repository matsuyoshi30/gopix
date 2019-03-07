# Happiness Filter in Go

### Installation

- Go
- [Face API in Azure](https://azure.microsoft.com/ja-jp/services/cognitive-services/face/)

### Usage

1. `go get` this repository.

```
$ go get github.com/matsuyoshi30/hfg
```

2. Write env file.

```.env
URL=https://[location].api.cognitive.microsoft.com/face/v1.0
KEY1=[Your SubscriptionKey]
KEY2=[Your SubscriptionKey]
```

3. Build.

```
$ go build hfg.go
```

4. Run command.

```
$ ./hfg [path/to/image]
```

Output image is named `output_[yyyymmddhhmmss].jpeg`.  

This is `./hfg test/Lena.png` result.  

![](./test/Lena_output.jpeg)

