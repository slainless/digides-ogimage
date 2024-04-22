#!/usr/bin/env node

import app from '.'
import { add, complete, configure, cycle, suite } from 'benny'
import { getPlatformProxy } from 'wrangler'

const platform = await getPlatformProxy()
const parameter = 'lg58eGIKWrp8NtUIqmkzTF7RXA9C2SuMg0-RWo0uS5c9pO3igVnHOxKPoEzcQFsO3BAxnu4hpMAdr2EtinVoGXF5r-CBRQji9nEnDCdi31VFFjPE_UczrJSnN1XT7mvP5FZTS1ffddOLYNnUx72sUnSg9rrmGBv29Fn4FcWKX1mkMUHY36L8_PiIyAC_0FZ79mJNjBciJg5_P1ermsusCG0y6ITRqhwl41JGFM-FspO2gWY15ErQetEX-lrhphb2vAJYcozle81ywWRPDfNsIgnVolFWGFh2LMqaVVO5K-1hjhazjOHDiA'

suite(
  "app.main benchmarking",
  configure({
    minDisplayPrecision: 5
  }),

  add("Default", () => app.request(`/?d=${parameter}`, {}, platform.env)),
  cycle(),
  complete(),
)