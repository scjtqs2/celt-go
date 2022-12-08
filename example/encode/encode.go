package main

import (
	"bufio"
	"github.com/scjtqs2/celt-go/celt_codec"
	"io"
	"os"
)

// 相关参数，encode和decode必须一致
const (
	channels      = 1
	frameSize     = 480
	samplerate    = 48000
	bitsPerSample = 16
	frameBytes    = frameSize * channels * (bitsPerSample / 8)
)

const CELT_EOF = "iamiscelteof"

func main() {
	f, _ := os.OpenFile("example/pcm/test.pcm", os.O_RDONLY, os.ModePerm)
	defer f.Close()
	r := bufio.NewReader(f)
	codec, err := celt_codec.NewEncoder(channels, frameSize, samplerate, bitsPerSample)
	if err != nil {
		panic(err)
	}
	w, _ := os.OpenFile("example/pcm/test.celt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer w.Close()
	for {
		buf := make([]byte, frameBytes)
		_, err := io.ReadFull(r, buf)
		if err == io.EOF {
			break
		}
		celt := codec.Encode(buf)
		w.Write(celt)
		w.Write([]byte(CELT_EOF)) // 添加分隔符，pcm每一帧是定长的，celt的不是，不加分隔符就无法识别了。这里分隔符仅供测试decode用。
	}
}
