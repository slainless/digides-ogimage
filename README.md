# digides-ogimage

<p align="center">
  <img src="./assets/example.jpeg" width="80%"/>
</p>

A cloudflare worker to generate OpenGraph banner, configured specifically for Digital Desa Profile Website.

Using Hono for edge framework, while Go is used to draw the image.

Live at https://og-image-generator.digital-desa.workers.dev/.

## Performance

To be honest, the result is a bit disappointing, taking ~2s average to render. Well, I'm not expecting too much with this kind of setup.

## Future plan

This repo is production ready, with WASM loaded correctly and service running. However, the performance is too lacking.

Considering that [cloudflare WebGPU API is not ready yet](https://developers.cloudflare.com/durable-objects/api/webgpu/) 
(and even then, it's only supported on Durable Objects), the only choice is to render using CPU. Not that I know how to program using 
WebGPU (I wish I have time to learn it), but it's the best option for graphic drawing.

I'm thinking of incorporating libvips to do filtering and image drawing, if possible. 
I'm also planning to improve the drawing logic, I think I might have wasted too much processing power doing 
filtering four times, perhaps there is better a way.

Doing the processing purely in javascript environment is also an option, using [wasm-vips](https://github.com/kleisauke/wasm-vips) or something.