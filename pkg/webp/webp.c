#include "emscripten.h"
#include "./libwebp/src/webp/encode.h"
#include <stdio.h>

typedef struct {
    uint8_t *buffer;
    size_t size;
} EncodeResult;

EMSCRIPTEN_KEEPALIVE
int version() {
  return WebPGetEncoderVersion();
}

EMSCRIPTEN_KEEPALIVE
uint8_t* create_buffer(int width, int height) {
  return malloc(width * height * 4 * sizeof(uint8_t));
}

EMSCRIPTEN_KEEPALIVE
void free_buffer(uint8_t* buffer) {
  free(buffer);
}

EMSCRIPTEN_KEEPALIVE
EncodeResult* encode(uint8_t* buffer, int width, int height, float quality) {
  uint8_t* out = malloc(1024 * 1024 * sizeof(uint8_t));
  size_t size;

  size = WebPEncodeRGBA(buffer, width, height, width * 4, quality, &out);

  EncodeResult *result = malloc(sizeof(EncodeResult));
  result->buffer = out;
  result->size = size;
  
  return result;
}

EMSCRIPTEN_KEEPALIVE
void free_result(EncodeResult* result) {
  WebPFree(result->buffer);
  free(result);
}