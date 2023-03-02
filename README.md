# wxworkbot
调用企业微信机器人发送消息至群里的命令行工具

# Usage
首先设置环境变量，指定企业微信机器人的key(即创建机器人后，生成的链接中send?key= 后的内容)

```Bash
export BOT_KEY="YOUR_KEY"
```

```Bash
$./wxworkbot -h

wxworkbot is a command-line tool for enterprise WeChat robots

Usage:
  wxworkbot [command]

Available Commands:
  help        Help about any command
  image       wxworkbot send image message
  markdown    wxworkbot send markdown message
  news        wxworkbot send news message
  text        wxworkbot send text message

Flags:
      --debug     Enable debug mode
  -h, --help      help for wxworkbot
  -v, --version   version for wxworkbot

Use "wxworkbot [command] --help" for more information about a command.
```

[企业微信官方文档](https://developer.work.weixin.qq.com/document/path/91770)

## 发送文本(text)类型消息
```Bash
$./wxworkbot text -h

wxworkbot send text message

Usage:
  wxworkbot text [flags]

Flags:
  -c, --content string                 指定要发送的文本内容,最长不超过2048个字节,[必须指定]
  -h, --help                           help for text
  -u, --mentioned_list string          userid的列表,提醒群中的指定成员(@某个成员),@all表示提醒所有人,如果开发者获取不到userid,可以使用mentioned_mobile_list
  -m, --mentioned_mobile_list string   手机号列表，提醒手机号对应的群成员(@某个成员),@all表示提醒所有人

Global Flags:
      --debug   Enable debug mode
```

## 发送markdown(markdown)类型消息
```Bash
$./wxworkbot markdown -h

wxworkbot send markdown message

Usage:
  wxworkbot markdown [flags]

Flags:
  -c, --content string   指定要发送的文本内容,最长不超过2048个字节[必须指定]
  -h, --help             help for markdown

Global Flags:
      --debug   Enable debug mode
```

## 发送图片(image)类型消息
```Bash
$./wxworkbot image -h

wxworkbot send image message

Usage:
  wxworkbot image [flags]

Flags:
  -b, --base64 string   图片内容的base64编码, [必须指定]
  -h, --help            help for image
  -m, --md5 string      图片内容(base64编码前)的md5值,[必须指定]

Global Flags:
      --debug   Enable debug mode
```

## 发送图文(news)类型消息
```Bash
$./wxworkbot news -h

wxworkbot send news message

Usage:
  wxworkbot news [flags]

Flags:
  -d, --description string   描述,不超过512个字节,超过会自动截断,[必须指定]
  -h, --help                 help for news
  -p, --picurl string        图文消息的图片链接,支持JPG、PNG格式,较好的效果为大图 1068*455,小图150*150。
  -t, --title string         标题不超过128个字节,超过会自动截断,[必须指定]
  -u, --url string           点击后跳转的链接。[必须指定]

Global Flags:
      --debug   Enable debug mode
```
