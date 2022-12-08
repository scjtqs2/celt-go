package main

import (
	"bufio"
	"bytes"
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
	f, err := os.OpenFile("example/pcm/test.celt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	dcodec, err := celt_codec.NewDecoder(channels, frameSize, samplerate, bitsPerSample)
	if err != nil {
		panic(err)
	}
	w, err := os.OpenFile("example/pcm/test2.pcm", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer w.Close()
	celtbuf, _ := io.ReadAll(r)
	celtBufs := bytes.Split(celtbuf, []byte(CELT_EOF))
	for _, buf := range celtBufs {
		pcm := dcodec.Decode(buf)
		w.Write(pcm)
	}
}
