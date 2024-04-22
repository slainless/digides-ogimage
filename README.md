# digides-ogimage

<p align="center">
  <img src="./assets/example.jpeg" width="80%"/>
</p>

A cloudflare worker to generate OpenGraph banner, configured specifically for Digital Desa Profile Website.

Using Hono for edge framework, while Go is used to draw the image.

## Benchmark

| Benchmark                                               | Result           |
|---------------------------------------------------------|------------------|
| [Drawing directly](./pkg/ogimage/draw_test.go)          | 0.262145631 s/op |
| [Base64 decoding + drawing](./pkg/bridge/start_test.go) | 0.314484976 s/op |
| [E2E (direct request)](./src/app.benchmark.test.ts)  | 1.271148737 s/op |

## To-do

There is a huge inefficiency in data conversion and bridge usage on E2E process, hence the need to improve the performance of the implementation, specifically on how javascript and Go passing data to each other.