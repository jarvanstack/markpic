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


