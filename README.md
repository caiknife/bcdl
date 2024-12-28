# bcdl

**本项目仅供学习，请勿用于其他用途。**

BandCamp 专辑下载、单曲下载，由于BandCamp是被墙了的，所以魔法工具请自备。

### 安装

本地有Go环境，执行

```shell
go install github.com/caiknife/bcdl/bcdl@latest
```
本地没有Go环境，请在[release页面](https://github.com/caiknife/bcdl/releases)下载最新的二进制文件，请注意区分操作系统和CPU架构。并将二进制文件放在曲库文件夹下。

### 下载

**请自行替换链接**

下载单曲

```shell
bcdl "https://shirttailstompers.bandcamp.com/track/sweets"
```

下载专辑

```shell
ncmdl "https://shirttailstompers.bandcamp.com/album/thats-my-kick"
```

仅能下载128码率的MP3文件
