#include "emscripten.h"
#include "./libwebp/src/webp/encode.h"

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
uint8_t* encode(uint8_t* buffer, int width, int height, int quality) {
  uint8_t* out;
  size_t size = WebPEncodeRGBA(buffer, width, height, width * 4, quality, &out);

  uint8_t* result = malloc(2 * sizeof(int));
  result[0] = (int) out;
  result[1] = size;
  
  return result;
}

EMSCRIPTEN_KEEPALIVE
void free_result(uint8_t* resultPtr) {
  WebPFree((uint8_t*) resultPtr[0]);
  free(resultPtr);
}