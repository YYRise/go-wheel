##

### 1. 生成动态库
```shell
gcc -fPIC -shared -o lib/libadd.so include/add.c
```

```shell
xcrun cc -arch arm64 -fPIC -shared -o lib/libadd-arm64.so include/add.c
xcrun cc -arch x86_64 -fPIC -shared -o lib/libadd-x86_64.so include/add.c
```

### 2. 链接动态库生成可执行文件
```shell
gcc include/test.c -L lib/ -ladd -o test
```
- `-L .` :表示搜索要链接的库文件时包含当前目录
- `-ladd` :表示要链接动态库 libadd.so (备注：默认lib + xxx + .so ,中间的xxx就是库名)
- `-o test` :生成可执行文件 test

### 3. 执行 test

```shell
./test
```

#### 3.1 出现找不到动态库错误时，先设置环境变量 
```shell
export LD_LIBRARY_PATH=lib/ 

./test
```

### 4. 执行 main.go
```shell
export LD_LIBRARY_PATH=lib/ 

CGO_ENABLED=1 go run main.go
```