# Story

I'm exploring different alterantives to what i'm used to Node / Express / Fastify.
Currently i'm starting with MVP which is `hello world` simple response, this will give me
a base / starting point to compare to more advance and real life use cases.

# Prerequisites

```bash
wrk - https://github.com/wg/wrk

Node - Fastify
Bun - ElysiaJS
Go - Fiber
Zig
```

## Benchmark Condition - Local

- MacBook Pro M1 Max 64 GB
- macOS Ventura 13.5.2
- Node 20.6.1
- Pnpm 8.6.12
- Bun 1.0.4

### Commands to build and run

```bash
-- Node - Fastify
pnpm i
pnpm build
pnpm start

-- Bun - ElysiaJS
bun i
bun run build
bun run start

-- Go - Fiber
go build app.go && ./app
```

### Test commands

```bash
-- Node - Fastify
wrk -t12 -c400 -d30s http://localhost:3005/

-- Bun - ElysiaJS
wrk -t12 -c400 -d30s http://localhost:3006/

-- Go - Fiber
wrk -t12 -c400 -d30s http://localhost:3007/
```

### Results

```bash
-- Node - Fastify
Running 30s test @ http://localhost:3005
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.81ms    6.22ms 300.69ms   99.72%
    Req/Sec     5.62k     1.25k   10.48k    68.72%
  2019851 requests in 30.10s, 362.14MB read
  Socket errors: connect 157, read 300, write 0, timeout 0
Requests/sec:  67097.86
Transfer/sec:     12.03MB

-- Bun - ElysiaJS
Running 30s test @ http://localhost:3006
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.76ms  624.12us  12.82ms   64.11%
    Req/Sec     7.25k     1.84k   13.90k    63.96%
  2602515 requests in 30.10s, 310.24MB read
  Socket errors: connect 157, read 107, write 0, timeout 0
Requests/sec:  86448.98
Transfer/sec:     10.31MB

-- Go - Fiber
Running 30s test @ http://localhost:3007/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.72ms    3.22ms 123.03ms   95.31%
    Req/Sec    12.25k     4.37k   46.82k    71.04%
  4387485 requests in 30.10s, 523.03MB read
  Socket errors: connect 157, read 169, write 0, timeout 0
Requests/sec: 145786.66
Transfer/sec:     17.38MB
```
