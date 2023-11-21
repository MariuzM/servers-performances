# Story

I'm exploring different alterantives to what i'm used to Node / Express / Fastify. Currently i'm
starting with MVP which is `hello world` simple response, this will give me a base / starting point
to compare to more advance and real life use cases.

# Prerequisites

```bash
wrk - https://github.com/wg/wrk
```

## Benchmark Condition - Local

- MacBook Pro M1 Max 64 GB
- macOS Ventura 13.6.2
- Node 20.9.1
- Pnpm 8.9.2
- Bun 1.0.13

### Commands to build and run

```bash
-- bun-elysia
bun i
bun st

-- bun-fastify
bun i
bun st

-- go-fiber
go build app.go && ./app

-- node-fastify
pnpm i
pnpm b
pnpm st
```

### Test commands

```bash
-- bun-elysia
wrk -t12 -c400 -d30s http://localhost:3005

-- bun-fastify
wrk -t12 -c400 -d30s http://localhost:3006

-- go-fiber
wrk -t12 -c400 -d30s http://localhost:3007

-- node-fastify
wrk -t12 -c400 -d30s http://localhost:3008
```

### Results

```bash
-- bun-elysia
Running 30s test @ http://localhost:3005
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.43ms  570.08us  16.02ms   57.70%
    Req/Sec     8.21k     2.45k   16.24k    67.50%
  2950042 requests in 30.10s, 351.67MB read
  Socket errors: connect 157, read 155, write 0, timeout 0
Requests/sec:  95998.41
Transfer/sec:     11.68MB

-- bun-fastify
Running 30s test @ http://localhost:3006
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.49ms  377.33us  15.01ms   83.06%
    Req/Sec     5.73k     1.80k   15.76k    58.29%
  2054935 requests in 30.10s, 274.36MB read
  Socket errors: connect 157, read 111, write 0, timeout 0
Requests/sec:  68763.99
Transfer/sec:      9.11MB

-- go-fiber
Running 30s test @ http://localhost:3007
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.72ms    3.22ms 123.03ms   95.31%
    Req/Sec    12.25k     4.37k   46.82k    71.04%
  4387485 requests in 30.10s, 523.03MB read
  Socket errors: connect 157, read 169, write 0, timeout 0
Requests/sec: 147895.34
Transfer/sec:     17.38MB

-- node-fastify
Running 30s test @ http://localhost:3008
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.65ms    5.56ms 299.27ms   99.79%
    Req/Sec     5.79k     1.59k    9.91k    74.70%
  2082358 requests in 30.10s, 373.35MB read
  Socket errors: connect 157, read 202, write 0, timeout 0
Requests/sec:  68172.32
Transfer/sec:     12.40MB
```
