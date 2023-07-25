# lc-code2md
将力扣美站的代码链接转为力扣中站 md 支持的格式。

## 目录结构
```text
|-------- lc-code2md
  |----- answer
  |----- config
    |----- config.go
    |----- congig.txt
    |----- replace.json
  |----- file
  |----- logic
    |----- replace.go
    |----- request
      |----- graph.go
  |----- main.go
  |----- README.md
```

## 🚀快速开始

1. 将美站 cookie 填充至 config.txt 中
2. 将需要转换的文件复制到 file 中
3. 等待转换完成



## 🎈一些补充
可自定义添加字段至 replace.json 中

| 转换字段                | 备注           |
| ----------------------- | -------------- |
| 具有 frame 标签的代码块 |                |
| 公式符号                | 即 $$ 转为 $   |
| 分隔符                  | --- 转为 \n--- |



1. 如果代码实现部分被替换为空，请检查 cookie 以及解决方案状态
