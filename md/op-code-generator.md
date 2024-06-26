### 生成 op 代码

下载官方生成库 `git clone git@github.com:kubernetes/code-generator.git`

此时拥有的文件各文件路径为：

- 开发库目录与名称为`op-crd`
- 开发库`go.mod`内项目名称为`op-crd`
- 官方生成库项目目录为`/home/gocode/code-generator`
- 开发库需生成代码的`gvk`路径为`op-crd/pkg/apis`
- `op-crd/pkg/apis/task/v1`内已经定义好`doc.go`、`register.go`、`types.go`

在开发项目目录执行:

```
/home/gocode/code-generator/generate-groups.sh all op-crd/pkg/generated op-crd/pkg/apis crd.example.com:v1 -v 10 --go-header-file=/home/gocode/code-generator/hack/boilerplate.go.txt --output-base ../
```

如果报错，可以是因为安装命令没有安装，需要在`code-generator`目录下安装 `go install cmd/*`，一般情况应该会自动安装。

如果是成功的应该会在控制台输出很多很多日志，至少几百行的数量级，如果日志有打印，却没有看到在`op-crd/pkg/generated`生成文件，可以去上几级目录找下，有时候是目录层级输出错了位置