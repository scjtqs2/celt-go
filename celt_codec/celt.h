#ifndef CELT_H
#define CELT_H

#include "celt_types.h"

#ifdef __cplusplus
extern "C" {
#endif

#if defined(__GNUC__) && defined(CELT_BUILD)
#define EXPORT __attribute__ ((visibility ("default")))
#elif defined(WIN32)
#define EXPORT __declspec(dllexport)
#else
#define EXPORT
#endif

#define _celt_check_int(x) (((void)((x) == (celt_int32)0)), (celt_int32)(x))
#define _celt_check_mode_ptr_ptr(ptr) ((ptr) + ((ptr) - (CELTMode**)(ptr)))

#define CELT_OK                0
#define CELT_BAD_ARG          -1
#define CELT_INVALID_MODE     -2
#define CELT_INTERNAL_ERROR   -3
#define CELT_CORRUPTED_DATA   -4
#define CELT_UNIMPLEMENTED    -5
#define CELT_INVALID_STATE    -6
#define CELT_ALLOC_FAIL       -7

#define CELT_GET_MODE_REQUEST    1
#define CELT_GET_MODE(x) CELT_GET_MODE_REQUEST, _celt_check_mode_ptr_ptr(x)
#define CELT_SET_COMPLEXITY_REQUEST    2
#define CELT_SET_COMPLEXITY(x) CELT_SET_COMPLEXITY_REQUEST, _celt_check_int(x)
#define CELT_SET_PREDICTION_REQUEST    4

#define CELT_SET_PREDICTION(x) CELT_SET_PREDICTION_REQUEST, _celt_check_int(x)
#define CELT_SET_VBR_RATE_REQUEST    6
#define CELT_SET_VBR_RATE(x) CELT_SET_VBR_RATE_REQUEST, _celt_check_int(x)
#define CELT_RESET_STATE_REQUEST        8
#define CELT_RESET_STATE       CELT_RESET_STATE_REQUEST

#define CELT_GET_FRAME_SIZE   1000
#define CELT_GET_LOOKAHEAD    1001
#define CELT_GET_SAMPLE_RATE  1003

#define CELT_GET_BITSTREAM_VERSION 2000

typedef struct CELTEncoder CELTEncoder;

typedef struct CELTDecoder CELTDecoder;

typedef struct CELTMode CELTMode;

EXPORT CELTMode *celt_mode_create(celt_int32 Fs, int frame_size, int *error);

EXPORT void celt_mode_destroy(CELTMode *mode);

EXPORT int celt_mode_info(const CELTMode *mode, int request, celt_int32 *value);

EXPORT CELTEncoder *celt_encoder_create(const CELTMode *mode, int channels, int *error);

EXPORT void celt_encoder_destroy(CELTEncoder *st);

EXPORT int celt_encode_float(CELTEncoder *st, const float *pcm, float *optional_synthesis, unsigned char *compressed, int nbCompressedBytes);

EXPORT int celt_encode(CELTEncoder *st, const celt_int16 *pcm, celt_int16 *optional_synthesis, unsigned char *compressed, int nbCompressedBytes);

EXPORT int celt_encoder_ctl(CELTEncoder * st, int request, ...);

EXPORT CELTDecoder *celt_decoder_create(const CELTMode *mode, int channels, int *error);

EXPORT void celt_decoder_destroy(CELTDecoder *st);

EXPORT int celt_decode_float(CELTDecoder *st, const unsigned char *data, int len, float *pcm);

EXPORT int celt_decode(CELTDecoder *st, const unsigned char *data, int len, celt_int16 *pcm);

EXPORT int celt_decoder_ctl(CELTDecoder * st, int request, ...);


EXPORT const char *celt_strerror(int error);


#ifdef __cplusplus
}
#endif

#endif /*CELT_H */
