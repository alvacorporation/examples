# Alva examples

Worked examples, each one a single file you can read in a sitting.

| Example | Shows |
|---|---|
| [`settlement/`](settlement) | A loop over external calls that must not double-execute |
| [`onboarding/`](onboarding) | Durable sleep across days, surviving deploys |
| [`fanout/`](fanout) | Parallel steps, and how failure of one is handled |

Every example runs against `alva dev` with no other setup:

```bash
alva dev &
go run ./settlement
```

## Reading order

Start with `settlement`. It is the shortest and it demonstrates the single
property everything else is built on: a step that has completed is never
executed twice, even if the process dies mid-loop.

`onboarding` is next — it shows durable sleep, which is the feature people are
usually surprised by. A workflow can wait three weeks without holding a worker
or a database connection.

`fanout` is last because parallelism is where the determinism rule stops being
obvious. Read the comments in that one carefully.
