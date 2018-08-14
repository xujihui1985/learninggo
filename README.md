learninggo
==========

### Setup go developing environment

#### install golang 

[download from here](https://golang.org/dl/)

#### setup GOPATH and GOROOT env

```
export GOPATH=$HOME/code/go
export PATH=$GOPATH/bin:$PATH
```

`$GOPATH` is where you put all of your go code


#### initialize gopath

```
mkdir -p $GOPATH/{bin,pkg,src}
```

congratulation, you are good to go

#### run your hello world

```
mkdir -p $GOPATH/src/hellogo

cat > $GOPATH/src/hellogo/main.go <<EOF
package main

import "fmt"

func main() {
  fmt.Println("hello world")
}

EOF
``

go run $GOPATH/src/hellogo/main.go