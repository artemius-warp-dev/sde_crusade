# sde_crusade

My long-running practice log for algorithms, data structures and software
design, in three languages.

| folder | contents |
|--------|----------|
| [`golang/educative`](golang/educative) | Solutions to coding-interview pattern exercises (two pointers, sliding window, heaps, graphs, DP, …), grouped by pattern, with my study notes. |
| [`golang/patterns`](golang/patterns) | GoF design patterns (creational, structural, behavioral), GRASP principles and concurrency patterns (barrier, fan-out/fan-in, future, …), each as a small runnable program. |
| [`golang/common_sorts`](golang/common_sorts) | Bubble, selection, insertion, merge, quick and radix sort. |
| [`golang/homework`](golang/homework) | Course homework: CLI utilities (`copy_util`, `envdir`, a telnet-like client/server) and a calendar service with a layered architecture, Docker Compose and integration tests. |
| [`elixir`](elixir) | Exercism-style exercises plus graph algorithms (BFS, DFS, Dijkstra, Prim, Floyd–Warshall, topological sort, union–find). |
| [`geeksforgeek`](geeksforgeek) | GeeksforGeeks problems in JavaScript and Python. |

Most Go programs are self-contained `main` packages:

```sh
go run golang/common_sorts/merge.go
```
