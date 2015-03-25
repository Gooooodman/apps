Go语言包管理工具(Go Package Manager)
====================================

设计目标
========

主要解决以下问题:

- 在断网时依然可用
- 支持按照版本或标签进行获取
- 源码打包发布是, 自动提取第三方依赖
- 将第三方依赖的包路径转为相对路径导入
- 在没有git/hg等版本工具时依然可用(zip导入)

断网时的处理策略
================

在  `$GOPKG_REPO` 目录缓存已经下载的库 (默认在 `$GOROOT/src.gopkg` 目录).

目录结构如下:

```
$GOROOT/src.gopkg
 |
 +- git
 +- hg
 +- bzr
 +- zip
```

当需要获取依赖的库时, 先从本地的缓存获取, 然后再从远程获取.

对于 github/googlecode 等常见的托管网站进行定制, 如果没有git工具
则根据版本号下载对应的zip文件.

版本依赖管理
============

对于 app 应用一般需要设置依赖的版本. 工作模式类似 gopm 工具.

间接依赖的版本将被忽略, 因为会出现不同版本同时导入的问题.

自动提取第三方依赖
==================

对于 app 应用, 在对应的 `./Gopkgs` 目录保存第三方依赖的信息.
格式和 godep 类似.

依赖信息在 `Gopkgs/Gopkgs.json` 文件中.

当前app可以是普通目录(不受git等管理).

相对路径导入第三方依赖
======================

提供路径修复命令, 在 app 路径发生变化是, 可以将第三方库和子库
转为当前新的导入路径, 完美支持 `go build` 零依赖编译.

提供zip文件导入
===============

当没有 git 等版本工具时, 尝试直接获取对应版本的 zip 文件.

如果是自己假设的包服务, 提供以下的扩展:

```
<meta name="go-import" content="import-prefix vcs repo-root">
```

其中 vcs 部分除了标准的 "git", "hg", "svn" 工具外, 还支持 **zip** 格式.

安装
====

	$ go get github.com/gopkg/gopkg

起步
====

如何向 gopkg 添加项目.

假设你的应用代码完成到一定阶段, 想要使用 `go install` 和 `go test` 完成构建,
只需要一条 gopkg 命令:

	$ gopkg save

此命令会保存版本依赖关系表保存到文件 Gopkgs/Gopkgs.json, 并复制源代码到 Gopkgs/_workspace.
确保总有一份依赖源码可用.

恢复
====

`gopkg restore` 命令是 `gopkg save` 的反向操作. 它将 Gopkgs/Gopkgs.json 中的依赖包安装到你的 GOPATH 路径中.

测试
====

	$ gopkg go test

增加依赖
========

当你的项目源码中 import 改变后再次使用 `gopkg save` 即可.

更新依赖
========

当你需要更新依赖包, 依次执行命令:

	1. `go get -u pathto/Dependency`
	2. `gopkg update pathto/Dependency`

多包支持
========

你可以将多个项目的依赖指向同一个 Gopkgs/Gopkgs.json 和 Gopkgs/_workspace.

	`gopkg save pathto/app1 pathto/app2`
	`gopkg restore pathto/app1 pathto/app2`

文件格式
========

gopkg 使用 JSON 格式.

类型定义:
```go
type Gopkgs struct {
	ImportPath string
	GoVersion  string
	Packages   []string `json:",omitempty"` // Arguments to save, if any.
	Deps       []Dependency
}

type Dependency struct {
	ImportPath string
	Comment    string `json:",omitempty"` // Description of commit, if present.
	Rev        string // VCS-specific commit ID.
}
```

Gopkgs.json 样例:

```json
{
    "ImportPath": "github.com/kr/hk",
    "GoVersion": "go1.1.2",
    "Deps": [
        {
            "ImportPath": "code.google.com/p/go-netrc/netrc",
            "Rev": "28676070ab99"
        },
        {
            "ImportPath": "github.com/kr/binarydist",
            "Rev": "3380ade90f8b0dfa3e363fd7d7e941fa857d0d13"
        }
    ]
}
```

补充
====

欢迎讨论.
