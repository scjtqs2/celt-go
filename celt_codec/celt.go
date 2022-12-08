package celt_codec

/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib -lcelt0  -lm -ldl -lstdc++

#ifndef ___SCJTQSCLASS_CELT_CODEC_2022_12_8__H_
#define ___SCJTQSCLASS_CELT_CODEC_2022_12_8__H_
#include <stdlib.h>
#include "celt.h"
// EXPORT const unsigned char *make_const_unsigned_char(unsigned char *data);
const unsigned char *make_const_unsigned_char(unsigned char *data) {
	return (const unsigned char*)data;
}
int decode(CELTDecoder* st, char* src, long srclen, char* dst) {
    return celt_decode(st, (const unsigned char*)src, srclen, (celt_int16 *)dst);
}

int encode(CELTEncoder* st,char* src, size_t srclen, char* dst)
{
    int len;
    len = celt_encode(st, (celt_int16 *)src, NULL, (unsigned char*)dst, (long)srclen);
	return len;
}

#endif
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

// CeltDecoder 解码工具
type CeltDecoder struct {
	Mode          *C.CELTMode
	Decoder       *C.CELTDecoder
	Channels      C.int
	FrameSize     C.int
	SampleRate    C.int
	BitsPerSample C.int
	FrameBytes    C.int
}

// NewDecoder 初始化 CeltDecoder类
func NewDecoder(channels, frameSize, sampleRate, bitsPerSample int) (*CeltDecoder, error) {
	codec := &CeltDecoder{
		Channels:      C.int(channels),
		FrameSize:     C.int(frameSize),
		SampleRate:    C.int(sampleRate),
		BitsPerSample: C.int(bitsPerSample),
		FrameBytes:    C.int(frameSize * channels * (bitsPerSample / 8)),
	}
	err := (C.int)(0)
	codec.Mode = C.celt_mode_create(codec.SampleRate, codec.FrameSize, &err)
	if err != C.CELT_OK {
		return nil, errors.New(fmt.Sprintf("faild to create celt mode errorcode=%d", int(err)))
	}
	codec.Decoder = C.celt_decoder_create(codec.Mode, codec.Channels, &err)
	if err != C.CELT_OK {
		return nil, errors.New(fmt.Sprintf("faild to create celt decoder errorcode=%d", int(err)))
	}
	return codec, nil
}

// Decode 解码 解码一帧celt的二进制数据，解码成pcm数据
func (codec *CeltDecoder) Decode(celtFrameBuf []byte) []byte {
	buflen := len(celtFrameBuf)
	if buflen == 0 {
		return nil
	}
	p := make([]byte, buflen*10+1)
	if C.CELT_OK != C.celt_decode(codec.Decoder, C.make_const_unsigned_char((*C.uchar)(unsafe.Pointer(&celtFrameBuf[0]))), C.int(buflen), (*C.celt_int16)(unsafe.Pointer(&p[0]))) {
		return nil
	}
	// if C.CELT_OK != C.decode(codec.Decoder, (*C.char)(unsafe.Pointer(&celtFrameBuf[0])), C.long(buflen), (*C.char)(unsafe.Pointer(&p[0]))) {
	// 	return nil
	// }
	sdtLenDecoder := int(codec.FrameBytes)
	pcm := p[:sdtLenDecoder]
	return pcm
}

// Init 重新初始化参数
func (codec *CeltDecoder) Init(channels, frameSize, sampleRate, bitsPerSample int) error {
	codec.Destroy()
	codec.Channels = C.int(channels)
	codec.FrameSize = C.int(frameSize)
	codec.SampleRate = C.int(sampleRate)
	codec.BitsPerSample = C.int(bitsPerSample)
	codec.FrameBytes = C.int(frameSize * channels * (bitsPerSample / 8))
	err := C.int(0)
	codec.Mode = C.celt_mode_create(codec.SampleRate, codec.FrameSize, &err)
	if err > 0 {
		return errors.New(fmt.Sprintf("faild to create celt mode errorcode=%d", int(err)))
	}
	codec.Decoder = C.celt_decoder_create(codec.Mode, codec.Channels, &err)
	if err > 0 {
		return errors.New(fmt.Sprintf("faild to create celt decoder errorcode=%d", int(err)))
	}
	return nil
}

// Destroy 销毁celt的mode和decoder
func (codec *CeltDecoder) Destroy() {
	if codec.Mode != nil {
		C.celt_mode_destroy(codec.Mode)
	}
	if codec.Decoder != nil {
		C.celt_decoder_destroy(codec.Decoder)
	}
}

// CeltEncoder 压缩工具
type CeltEncoder struct {
	Mode          *C.CELTMode
	Encoder       *C.CELTEncoder
	Channels      C.int
	FrameSize     C.int
	SampleRate    C.int
	BitsPerSample C.int
	FrameBytes    C.int
}

// NewEncoder 初始化 CeltEncoder
func NewEncoder(channels, frameSize, sampleRate, bitsPerSample int) (*CeltEncoder, error) {
	codec := &CeltEncoder{
		Channels:      C.int(channels),
		FrameSize:     C.int(frameSize),
		SampleRate:    C.int(sampleRate),
		BitsPerSample: C.int(bitsPerSample),
		FrameBytes:    C.int(frameSize * channels * (bitsPerSample / 8)),
	}
	err := (C.int)(0)
	codec.Mode = C.celt_mode_create(codec.SampleRate, codec.FrameSize, &err)
	if err != 0 {
		return nil, errors.New(fmt.Sprintf("faild to create celt mode errorcode=%d", int(err)))
	}
	codec.Encoder = C.celt_encoder_create(codec.Mode, codec.Channels, &err)
	if err != 0 {
		return nil, errors.New(fmt.Sprintf("faild to create celt decoder errorcode=%d", int(err)))
	}
	return codec, nil
}

// Encode 编码 用于压缩一帧的pcm数据
func (codec *CeltEncoder) Encode(pcmFrameBuf []byte) []byte {
	buflen := len(pcmFrameBuf)
	if buflen == 0 {
		return nil
	}
	p := make([]byte, buflen)
	dstlen := (C.int)(0)
	dstlen = C.celt_encode(codec.Encoder, (*C.celt_int16)(unsafe.Pointer(&pcmFrameBuf[0])), nil, (*C.uchar)(unsafe.Pointer(&p[0])), C.int(buflen))
	//dstlen = C.encode(codec.Encoder, (*C.char)(unsafe.Pointer(&pcmFrameBuf[0])), C.ulong(buflen), (*C.char)(unsafe.Pointer(&p[0])))
	if dstlen <= 0 {
		return nil
	}
	celt := p[:int(dstlen)]
	println("celtlen=", dstlen)
	return celt
}

// Init 重新初始化参数
func (codec *CeltEncoder) Init(channels, frameSize, sampleRate, bitsPerSample int) error {
	codec.Destroy()
	codec.Channels = C.int(channels)
	codec.FrameSize = C.int(frameSize)
	codec.SampleRate = C.int(sampleRate)
	codec.BitsPerSample = C.int(bitsPerSample)
	codec.FrameBytes = C.int(frameSize * channels * (bitsPerSample / 8))
	err := C.int(0)
	codec.Mode = C.celt_mode_create(codec.SampleRate, codec.FrameSize, &err)
	if err > 0 {
		return errors.New(fmt.Sprintf("faild to create celt mode errorcode=%d", int(err)))
	}
	codec.Encoder = C.celt_encoder_create(codec.Mode, codec.Channels, &err)
	if err > 0 {
		return errors.New(fmt.Sprintf("faild to create celt decoder errorcode=%d", int(err)))
	}
	return nil
}

// Destroy 销毁celt的mode和encoder
func (codec *CeltEncoder) Destroy() {
	if codec.Mode != nil {
		C.celt_mode_destroy(codec.Mode)
	}
	if codec.Encoder != nil {
		C.celt_encoder_destroy(codec.Encoder)
	}
}
