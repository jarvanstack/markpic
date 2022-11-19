# du

将 markdown 中的所有本地图片通过 picgo 上传到图床

## 安装

```bash
go install github.com/dengjiawen8955/du@latest
```

或者下载编译好的二进制文件并放到 PATH 中

下载地址: <https://github.com/dengjiawen8955/du/releases>

## 使用

```bash
$ du -h
markdown 图片下载上传工具. 例如:

# 将图片下载到本地文件夹中      
du d --from test.md --dir images/
# 将图片上传到图床并自动替换到新文件夹
du u --from test.local.md
# 一键上传和下载
du du --from test.md

Usage:
  du [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  d           将 markdown 中的图片下载到本地
  u           将 markdown 中的所有本地图片通过 picgo 上传到图床
  du          du 一键下载 markdown 中图片, 并上传图片到图床并替换链接
  help        Help about any command

Flags:
  -d, --dir string    图片存放的目录 (default "images/")
  -f, --from string   需要处理的文件 (default "README.md")
  -h, --help          help for du

Use "du [command] --help" for more information about a command.
```

## 使用实例

我们在 windows 有一个 markdown 文件 test.md

```markdown
六边形架构

![](https://cdn.nlark.com/yuque/0/2022/png/2397010/1668752212281-41b93ae1-c4c0-4af4-befa-ceaf1f4efb05.png#averageHue=%23f9f9f8&clientId=u39f9e662-ddb8-4&crop=0&crop=0&crop=1&crop=1&from=paste&id=u7149c748&margin=[object Object]&originHeight=279&originWidth=448&originalType=url&ratio=1&rotation=0&showTitle=false&status=done&style=none&taskId=u7e969d8b-9118-433a-86db-806375faecf&title=)

```

* 里面的网络图片因为防盗链机制, 无法再其他博客平台上同步
* 即使在 typroa 中设置插入图片时候对网络位置的图片上传也无法上传

![无法上传](https://markdown-1304103443.cos.ap-guangzhou.myqcloud.com/2022-02-0420221119171918.png)

* 所以我们需要将图片下载到本地, 然后再上传到图床
* 但是这个操作如果图片比较多的话比较繁琐, 而且容易遗漏
* 现在 **du 可以帮助我们一键下载和上传**

```bash
PS C:\c_code\du> du -f test.md
[下载-上传]
[下载-上传完成] .\test.md.download.md .\test.md.upload.md
PS C:\c_code\du> ls
-a---          2022/11/19    17:36            473 test.md
-a---          2022/11/19    17:36            108 test.md.download.md
-a---          2022/11/19    17:36            156 test.md.upload.md
```

test.md.download.md 内容如下

```markdown
六边形架构

![image.png](C:\c_code\du\images\)

```

test.md.upload.md 内容如下

```markdown
六边形架构

![image.png](https://myweb/1_1668752375501-1bb18730-73da-421a-8c35-de8d29029919.png.png)

```

因为自己搭建的图床没有防盗链机制, 所以可以直接在其他平台同步
