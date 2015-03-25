# 实用小工具

## lsdir: 目录列表

```
// List files, support file/header regexp.
//
// Example:
//	lsdir dir
//	lsdir dir "\.go$"
//	lsdir dir "\.go$" "chaishushan"
//	lsdir dir "\.tiff?|jpg|jpeg$"
//	lsdir dir "\.(par|eip)$"
//
// Help:
//	lsdir -h
```

## cpdir: 复制目录

```
// Copy dir, support regexp.
//
// Example:
//	cpdir src dst
//	cpdir src dst "\.go$"
//	cpdir src dst "\.tiff?$"
//	cpdir src dst "\.tiff?|jpg|jpeg$"
//	cpdir src dst "\.(par|eip)$"
//
// Help:
//	cpdir -h
```

## md5: 计算 MD5

```
// Cacl dir or file MD5Sum, support regexp.
//
// Example:
//	md5 file
//	md5 dir "\.go$"
//	md5 dir "\.tiff?$"
//	md5 dir "\.tiff?|jpg|jpeg$"
//	md5 dir "\.(par|eip)$"
//
// Help:
//	cpdir -h
```

## fixpath: 复制Go语言fork包路径


```
// Fix import path.
//
// Example:
//	go run fixpath.go -old=old-path-prefix -new=new-path-prefix
```


## lookpath: 简化版 witch 命令

```
// Lookpath is a simple which.
```

