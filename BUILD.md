## 书籍

### 构建
1. 新建WORKSPACE
2. 新建BUILD.bazel
3. 更新仓库
```
bazel run //:gazelle -- update-repos -from_file=go.mod
```
4. 更新各目录下的BUILD.bazel文件
```
bazel run //:gazelle
```

5. 运行项目
```
bazel run /cmd/booksvc:booksvc
```

6. 构建
编译后生成的二进制文件在 bazel-bin/darwin_amd64_stripped 路径下
```
bazel build /cmd/booksvc:booksvc
```

参考： https://mp.weixin.qq.com/s/fxxNj61pO80XY8EwDiRM9A