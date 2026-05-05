# Go Learning Roadmap for Node.js Engineers
> From Node.js Senior Engineer to Go Backend Developer in 20 weeks  
> Time commitment: ~9 hrs/week (1hr/day weekdays + 3+3 hrs weekends)

---

## Why Go (Context)

Go occupies a specific and lucrative niche: **high-throughput, low-latency backend infrastructure**. The supply of Go engineers is still meaningfully lower than demand, especially in distributed systems and platform engineering. For a Node.js backend engineer, the mental model transfer is smooth — you already think in async I/O, event-driven services, HTTP servers, and microservices.

---

## Weekly Time Budget

| Session | Time | Focus |
|---|---|---|
| Mon–Fri | 1 hr/day | Concept study + small exercises |
| Saturday | 3 hrs | Structured project building |
| Sunday | 3 hrs | Review, debug, write tests, read source code |

---

## Phase 1 — Go Fundamentals (Weeks 1–3)

> **Goal**: Get comfortable with Go's syntax, toolchain, and mental model differences from JS.

### Week 1 — Syntax & Tooling ✅
- ✅ Install Go, understand `GOPATH`, `go mod`, `go run`, `go build`
- ✅ Types, variables, short declaration (`:=`), zero values
- ✅ Functions — multiple return values, named returns
- ✅ Pointers — the one thing JS engineers underestimate; spend real time here
  - ✅ **(Extended)** Pointer exercises: swap function, struct mutation, nil guard

### Week 2 — Data Structures & Control Flow ✅
- ✅ Arrays vs slices (richer than JS arrays — understand `append`, capacity, underlying array)
- ✅ Maps, structs, methods on structs
- ✅ `for` as the only loop, `range`, `switch`
- ✅ No classes — understand the struct + method pattern

### Week 3 — Interfaces & Error Handling ✅
- ✅ Interfaces — implicit satisfaction, `io.Reader`/`io.Writer` as the canonical example
- ✅ Error handling — `error` as a value, `if err != nil`, custom error types
- ✅ `defer`, `panic`, `recover` — when and why
- ✅ Compare to JS: no try/catch as a crutch, errors are explicit

### Resources (Phase 1)
- [Tour of Go](https://go.dev/tour) — do the whole thing
- [Go by Example](https://gobyexample.com) — use as a reference alongside
- *The Go Programming Language* — Donovan & Kernighan, Chapters 1–5

---

## Phase 2 — Concurrency (Weeks 4–7)

> **Goal**: Rebuild your concurrency mental model. This is where Go separates itself — your Node.js event loop intuition needs to be extended, not just ported.

### Week 4 — Goroutines & the Scheduler ✅
- ✅ `go` keyword, goroutine lifecycle
- ✅ Understand Go's M:N scheduler vs Node's single-threaded event loop
- ✅ Why goroutines are cheap (2KB stack vs OS threads)
- ✅ `sync.WaitGroup` for coordination

### Week 5 — Channels ✅
- ✅ Unbuffered vs buffered channels
- ✅ Directional channels (`chan<-`, `<-chan`)
- ✅ `select` statement — Go's equivalent of `Promise.race` but more powerful
- ✅ Channel closing, ranging over channels

### Week 6 — Sync Primitives & Patterns
- ✅ `sync.Mutex`, `sync.RWMutex` — when channels aren't the right tool
- ✅ `sync.Once`, `sync.Pool`
- ✅  Worker pool pattern — you'll use this constantly
- ✅ Fan-out / fan-in pattern

### Week 7 — Context Package
- ✅ `context.Context` — cancellation, deadlines, timeouts
- ✅ Propagating context through call chains (equivalent to AsyncLocalStorage but idiomatic)
- Why every HTTP handler and DB call should accept a context

### Resources (Phase 2)
- *Concurrency in Go* — Katherine Cox-Buday (read the whole book this phase)
- [Go Concurrency Patterns](https://go.dev/blog/pipelines) — official blog post
- Rob Pike's [Concurrency is not Parallelism](https://go.dev/blog/waza-talk) talk

---

## Phase 3 — Backend Services (Weeks 8–13)

> **Goal**: Build real things. This phase is where your Node.js experience pays dividends.

### Week 8 — HTTP & Routing
- ✅ `net/http` standard library — understand `Handler`, `HandlerFunc`, `ServeMux`
- ✅ Pick [Chi](https://github.com/go-chi/chi) or [Gin](https://gin-gonic.com) — Chi is closer to Express in philosophy
- ✅ Middleware chains — replicate what you'd do in Express
- ✅ **Build**: a REST API with CRUD endpoints

### Week 9 — Database (PostgreSQL)
- `database/sql` — understand the interface
- Use [pgx](https://github.com/jackc/pgx) as the driver (better than lib/pq for Postgres)
- Connection pooling — `pgxpool`
- Use [sqlc](https://sqlc.dev) — write SQL, get type-safe Go code generated (idiomatic, not an ORM)
- **Build**: add persistence to your Week 8 API

### Week 10 — Configuration, Logging & Structure
- Project layout — understand the [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- Config management with [Viper](https://github.com/spf13/viper) or env-based with `os.Getenv`
- Structured logging with `slog` (stdlib since Go 1.21) — compare to Winston/Pino
- **Build**: make your API production-structured

### Week 11 — Testing in Go
- `testing` package, table-driven tests (idiomatic Go — learn to love it)
- `httptest` for handler testing
- Mocking with interfaces — no need for heavy mocking libraries
- Benchmarks with `testing.B`
- **Build**: get your API to 80%+ test coverage

### Week 12 — gRPC & Protobuf
- Why gRPC matters for internal microservice communication
- Define a `.proto` file, generate Go code with `protoc`
- Implement a gRPC server and client
- Streaming RPCs — unary vs server-side vs bidirectional
- **Build**: convert one of your REST endpoints to gRPC

### Week 13 — Kafka Integration
- Use [franz-go](https://github.com/twmb/franz-go) (better than confluent's client)
- Producer with retries and idempotency
- Consumer groups, offset management
- **Build**: add an async event pipeline to your service (publish on write, consume for side effects)

---

## Phase 4 — Distributed Systems in Go (Weeks 14–17)

> **Goal**: This is where your compensation ceiling rises. Tie your DS theory to Go implementations.

### Week 14 — Redis with Go
- Use [go-redis](https://github.com/redis/go-redis)
- Implement: rate limiter (token bucket, Lua script)
- Implement: distributed lock with `SET NX PX`
- Implement: leaderboard with `ZSET`

### Week 15 — Distributed Patterns
- Implement circuit breaker with [gobreaker](https://github.com/sony/gobreaker)
- Retry with exponential backoff — build your own, then see [go-retry](https://github.com/sethvargo/go-retry)
- Implement a simple consistent hashing ring from scratch (great interview prep)
- Health checks and graceful shutdown — signals, `os.Signal`, draining connections

### Week 16 — Observability
- OpenTelemetry Go SDK — traces, metrics, logs
- Instrument your service end-to-end: HTTP → DB → Kafka
- Export to Jaeger (traces) and Prometheus (metrics)
- Build a simple Grafana dashboard for your service

### Week 17 — Capstone Project
Pick one and build it properly in Go:

| Option | Stack |
|---|---|
| URL Shortener | Redis + Postgres + rate limiting + metrics |
| Job Queue | Postgres-backed persistence (mini Temporal) |
| Notification Fanout | Kafka consumer + WebSocket push |

---

## Phase 5 — Production Readiness (Weeks 18–20)

> **Goal**: Differentiate yourself. Most Go engineers stop at Phase 3.

### Week 18 — Docker & Kubernetes Basics
- Multi-stage Docker builds for Go (tiny final images ~10MB vs Node's 300MB+)
- Write a `Dockerfile`, understand scratch vs alpine base
- Basic K8s: `Deployment`, `Service`, `ConfigMap`, `HorizontalPodAutoscaler`
- Deploy your capstone project to a local Kind cluster

### Week 19 — Performance & Profiling
- `pprof` — CPU and memory profiling (this is a superpower)
- `go tool trace`
- Escape analysis — understanding heap vs stack allocation
- Benchmarking with `benchstat`
- Profile your capstone service and fix one real bottleneck

### Week 20 — Go Internals & Interview Prep
- Go memory model — happens-before, why it matters for concurrent code
- GC internals at a high level — tri-color mark-and-sweep
- Common Go interview patterns: implement a concurrent map, a semaphore, a pipeline
- Goroutine leaks and how to detect them with `goleak`

---

## Guiding Principles

**Don't fight Go's idioms.**  
The temptation is to write JavaScript in Go syntax — async patterns, callbacks, chaining. Resist it. Lean into explicit errors, table-driven tests, and simple structs.

**Read the standard library.**  
Go's stdlib is exceptionally well-written. Reading `net/http` or `encoding/json` source teaches you idiomatic Go faster than most books.

**Build, don't just read.**  
Every phase has a build component. A working, tested, observed Go service at the end of 20 weeks is worth more on a resume than 40 weeks of passive learning.

**Phases 4 and 5 are the differentiator.**  
Most Go tutorial graduates stop at Phase 3. Distributed patterns + profiling + K8s is what puts you in the top tier of Go candidates.

---

## Full Resource List

| Resource | Type | Phase |
|---|---|---|
| [Tour of Go](https://go.dev/tour) | Interactive | 1 |
| [Go by Example](https://gobyexample.com) | Reference | 1 |
| *The Go Programming Language* — Donovan & Kernighan | Book | 1 |
| *Concurrency in Go* — Katherine Cox-Buday | Book | 2 |
| [Go Concurrency Patterns](https://go.dev/blog/pipelines) | Article | 2 |
| [Concurrency is not Parallelism](https://go.dev/blog/waza-talk) — Rob Pike | Talk | 2 |
| [Standard Go Project Layout](https://github.com/golang-standards/project-layout) | Reference | 3 |
| [sqlc docs](https://sqlc.dev) | Docs | 3 |
| [pgx](https://github.com/jackc/pgx) | Library | 3 |
| [Chi router](https://github.com/go-chi/chi) | Library | 3 |
| [franz-go](https://github.com/twmb/franz-go) | Library | 3 |
| [go-redis](https://github.com/redis/go-redis) | Library | 4 |
| [gobreaker](https://github.com/sony/gobreaker) | Library | 4 |
| [OpenTelemetry Go](https://opentelemetry.io/docs/languages/go/) | Docs | 4 |