# RootedFS 单题评测仓库

本仓库提供可组合的 Go 文件系统接口，包含操作系统文件系统、内存文件系统、chroot 隔离、挂载、符号链接和临时文件工具。

本题只评测临时文件工具在非根文件系统中的默认目录行为。项目不依赖外部服务，测试使用内存文件系统和本地 Docker 工具链。

## 工具链

- `go.mod`：Go 1.21
- Docker 基础镜像：`golang:1.23.12`
- `GOTOOLCHAIN=local`
- `CGO_ENABLED=0`

## 专项验证

```bash
go test ./util -run '^TestTempDefaultsRespectChroot$' -count=1 -v
```

基线应失败，修复后应通过。该命令是本题唯一的 Bug 验证命令。

## 完整验收

```bash
go test ./... -count=1
go vet ./...
go build ./...
```

双架构构建入口：

```bash
./build_benzhi_docker.sh linux/amd64
./build_benzhi_docker.sh linux/arm64
```
