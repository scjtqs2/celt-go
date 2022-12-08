# pcm文件转mp3
```shell
ffmpeg -y -f s16be -ac 1 -ar 48000 -acodec pcm_s16le -i test.pcm test.mp3
```

# 对比 test.pcm 和test.celt 压缩效果还是很明显的。用来推流降低带宽效果还行。