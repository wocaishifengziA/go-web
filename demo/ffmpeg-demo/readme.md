# 视频抽帧
```
ffmpeg -i aaa.mp4 -r 1 -q:v 2 -f image2 /tmp/%08d.jpg
ffmpeg -i aaa.mp4 -vf fps=fps=1/10 -q:v 2 -f image2 /tmp/%08d.jpg
```
* -i 输入视频
* -f 类型 image2 表示图片
* -r 一秒截取多少张 （如果这里指定了 -r 那么后面 -vf 指定隔几秒取帧就无效了）
* -q:v2 表示截图画面的质量
* -vf fps=1/20 每隔20秒截取一张

# 获取视频时长
```
ffmpeg -i aaa.mp4 2>&1 | grep 'Duration' | cut -d ' ' -f 4 | sed s/,//
```

```
ffmpeg -i aaa.mp4 -r 1 -q:v 2  -s 640x360  -f image2 /tmp/%08d.jpg
```

# 获取视频完整信息
```
ffprobe bbb.mp4 -show_streams -select_streams v -print_format json
```

## 获取视频首帧
```
ffmpeg -i aaa.mp4 -r 1 -ss 00:00:00 -vframes 1 preview.jpg
```