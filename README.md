# Prerequistes

```bash
wrk - https://github.com/wg/wrk

Node - Fastify
Bun - ElysiaJS
Go
Zig
```

## Benchmark Condition - Local

- MacBook Pro M1 Max 64 GB
- macOS Ventura 13.5.2
- Node 20.6.1
- Pnpm 8.6.12
- Bun 1.0.4

### Test commands

```bash
-- Node - Fastify
wrk -t12 -c400 -d30s http://localhost:3005/

-- Bun - ElysiaJS
wrk -t12 -c400 -d30s http://localhost:3006/
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
```
