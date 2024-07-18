#include "emscripten.h"
#include "./libwebp/src/webp/encode.h"

int result[2];

EMSCRIPTEN_KEEPALIVE
uint8_t* create_buffer(int width, int height) {
  return malloc(width * height * 4 * sizeof(uint8_t));
}

EMSCRIPTEN_KEEPALIVE
void free_buffer(uint8_t* buffer) {
  free(buffer);
}

EMSCRIPTEN_KEEPALIVE
void encode(uint8_t* buffer, int width, int height, int quality) {
  uint8_t* out;
  size_t size = WebPEncodeRGBA(buffer, width, height, width * 4, quality, &out);

  result[0] = (int) out;
  result[1] = size;
}

EMSCRIPTEN_KEEPALIVE
int get_result_ptr() {
  return result[0];
}

EMSCRIPTEN_KEEPALIVE
int get_result_size() {
  return result[1];
}

EMSCRIPTEN_KEEPALIVE
void free_result() {
  WebPFree((uint8_t*) result[0]);
}