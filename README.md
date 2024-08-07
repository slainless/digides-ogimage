# digides-ogimage

> [!WARNING]
> Deprecated: This implementation has memory leak.
> WASI version available which is more stable and truly production ready.
> https://github.com/slainless/digides-ogimaker.

<p align="center">
  <img src="./assets/example.jpeg" width="80%"/>
</p>

A cloudflare worker to generate OpenGraph banner, configured specifically for Digital Desa Profile Website.

Using Hono for edge framework, while Go is used to draw the image.

Live at https://og-image-generator.digital-desa.workers.dev/.

[Click here to view render example.](https://og-image-generator.digital-desa.workers.dev/?d=JfDxtkTsSNXx8lDF4BqMVcHUla4WE59lF6SeUCmRr61FvhaaPrmDfkz41Zt0C1lUKGX0X-ywirXMyPLTW78arco6dyhnnIl4OWgi6g1Evf8rFSh6iMrn1OPB7U7OKAb7-2lLFa5RfmmmAQ2U9wrtIn5EHVFGi8jnrUewocOO5vK2pdSxXaFJnSlM47ULN3fHxuqnYGdi0KT7RjyZ7et1ZAOzA-GD9PZTLhZBlc-H48FOIx9zmLo-E8UBacdm3hHDKXUuociO8e8VJzEaZ1bLhhb8ttipzIzn7wgrQ3PUUjHCar8eCd06jg)

## Performance

Managed to reduce drawing time to around 0.8~0.9s/op. I think i can get more performance gain by optimizing the string drawing but I'm pretty
satisfied with the current drawer. Now, I just need to pass the image to a WEBP or AVIF decoder.

## Future plan

Considering that [cloudflare WebGPU API is not ready yet](https://developers.cloudflare.com/durable-objects/api/webgpu/) 
(and even then, it's only supported on Durable Objects), the only choice is to render using CPU. Not that I know how to program using 
WebGPU (I wish I have time to learn it), but it's the best option for graphic drawing in Cloudflare Worker.

I'm thinking of incorporating libvips to do filtering and image drawing, if possible. 
I'm also planning to improve the drawing logic, I think I might have wasted too much processing power doing 
filtering four times, perhaps there is a better way.

Doing the processing purely in javascript environment is also an option, using [wasm-vips](https://github.com/kleisauke/wasm-vips) or something.