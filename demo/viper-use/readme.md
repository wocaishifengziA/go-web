# viper

## 配置优先级
显示调用Set设置值
命令行参数（flag）
环境变量
配置文件
key/value存储
默认值

## 注意
1. 当前版本对key不区分大小写（统一读取为小写）
2. Viper将ENV变量视为区分大小写
3. 使用ENV变量时，每次访问该值时都将读取它

## 相关博客
https://www.liwenzhou.com/posts/Go/viper_tutorial/