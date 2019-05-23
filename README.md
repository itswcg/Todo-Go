# Todo-Go
使用beego重构以前的[todo](https://github.com/itswcg/Todo-Python)项目

# 构建
```bash
# 交叉编译,用于docker Scratch镜像

$ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o Todo-Go .
```

# Usage
```bash
$ git clone git@github.com:itswcg/Todo-Go.git
$ cd Todo-Go
$ docker build -t todo-go .
$ docker run -p 8512:8512 -d todo-go
```


# 欢迎使用
<https://gotodo.itswcg.com>