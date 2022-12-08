# pcm文件转mp3
```shell
ffmpeg -y -f s16be -ac 1 -ar 48000 -acodec pcm_s16le -i test.pcm test.mp3
```
