#!/usr/bin/env bash

if [ ! -d "pkg/webp/libwebp" ]; then
  git clone https://chromium.googlesource.com/webm/libwebp pkg/webp/libwebp
fi

echo "~ Building libwebp..."
emcc pkg/webp/webp.c -Os -o pkg/webp/webp.generated.wasm --no-entry \
  -sTOTAL_MEMORY=1024MB \
  -I \
    pkg/webp/libwebp \
    pkg/webp/libwebp/src/{dec,dsp,demux,enc,mux,utils}/*.c \
    pkg/webp/libwebp/sharpyuv/*.c
if [ $? -eq 0 ]; then
    echo "~ Done building libwebp"
else
    echo "~ Failed building libwebp"
    exit $?
fi
