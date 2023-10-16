<p align="center">
<img src="logo.png" width="200px"/>
<br>
<p align="center">
 <img src="https://img.shields.io/github/stars/jarvanstack/markpic" />
 <img src="https://img.shields.io/github/issues/jarvanstack/markpic" />
 <img src="https://img.shields.io/github/forks/jarvanstack/markpic" />
</p>
</p>

# markpic

一键将 markdown 中的所有图片下载到本地, 并通过 picgo 上传到图床


## 安装

(1) 方法1 通过 go install 直接安装

```bash
go install github.com/jarvanstack/markpic@latest
```

(2) 方法2 下载编译好的二进制文件并放到 PATH 中

下载地址: <https://github.com/jarvanstack/markpic/releases>

## 使用

```bash
# 一键将 markdown 中的所有图片下载到本地, 并通过 picgo 上传到图床
markpic test.md
```

```bash
$ markpic --help
一键将 markdown 中的所有图片下载到本地, 并通过 picgo 上传到图床

Usage:
  markpic [command] [flags]
  markpic [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  d           将 markdown 中的图片下载到本地
  du          先 d(下载) 再 u(上传)
  help        Help about any command
  u           将 markdown 中的所有本地图片通过 picgo 上传到图床

Flags:
  -d, --dir string    图片存放的目录 (default "images/")
  -f, --from string   需要处理的源文件 (default "source.md")
  -h, --help          help for markpic

Use "markpic [command] --help" for more information about a command.
```

## 使用实例

我们在 windows 有一个 markdown 文件 test.md

```markdown
六边形架构

![](https://cdn.nlark.com/yuque/0/2022/png/2397010/1668752212281-41b93ae1-c4c0-4af4-befa-ceaf1f4efb05.png#averageHue=%23f9f9f8&clientId=u39f9e662-ddb8-4&crop=0&crop=0&crop=1&crop=1&from=paste&id=u7149c748&margin=[object Object]&originHeight=279&originWidth=448&originalType=url&ratio=1&rotation=0&showTitle=false&status=done&style=none&taskId=u7e969d8b-9118-433a-86db-806375faecf&title=)

```

* 里面的网络图片因为防盗链机制, 无法再其他博客平台上同步
* 即使在 typroa 中设置插入图片时候对网络位置的图片上传也无法上传

![无法上传](https://cdn.jarvans.com/blog/2023/2022-02-0420221119171918.png)

* 所以我们需要将图片下载到本地, 然后再上传到图床
* 但是这个操作如果图片比较多的话比较繁琐, 而且容易遗漏
* 现在 **markpic 可以帮助我们一键将 markdown 中的所有图片下载到本地, 并通过 picgo 上传到图床**

```bash
PS C:\c_code\markpic> markpic  test.md
[下载-上传]
[下载-上传完成] .\test.md.download.md .\test.md.upload.md
```

test.md.download.md 内容如下

```markdown
六边形架构

![image.png](C:\c_code\markpic\images\)

```

test.md.upload.md 内容如下

```markdown
六边形架构

![image.png](https://myweb/1_1668752375501-1bb18730-73da-421a-8c35-de8d29029919.png.png)

```

因为自己搭建的图床没有防盗链机制, 所以可以直接在其他平台同步

## 注意

* 该程序调用 PicGo 客户端的接口上传图片, 使用前需要下载 PicGo 客户端并配置好图床, 保持 PicGo 客户端正常运行


