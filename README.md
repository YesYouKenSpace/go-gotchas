# go-snippets

This repository contains snippets of Go code, serving as experiments to demonstrate specific behaviors. Each `**/main.go` file is an independent experiment, accompanied by a `sample.out` file that includes both stdout and stderr outputs.

## Table of Contents

<details>
<summary><b>1. RWMutex</b></summary>

- 1.1. [No read preference](#no-read-preference)

</details>

<details>
<summary><b>2. Maps</b></summary>

- 1.1. [Potential memory leak](#potential-memory-leak)

</details>

## 1. RWMutex

1. [No read preference](rwmutex/no-read-preference/main.go) - Demonstrates the lack of read preference to prevent recursive read locking.
   > A blocked Lock call excludes new readers from acquiring the lock.
   > -- [Source](https://pkg.go.dev/sync#RWMutex.RLock:~:text=a%20blocked%20Lock%20call%20excludes%20new%20readers%20from%20acquiring%20the%20lock)
   - Learn more about the [readers-writers problem](https://en.wikipedia.org/wiki/Readers%E2%80%93writers_problem).

## 2. Maps

1. [Potential memory leak](maps/potential-memory-leak/main.go) 
   - Further discussion at [GitHub Issue #20135](https://github.com/golang/go/issues/20135).
   - This snippet shows that even after deleting all keys and attempting to reuse them, there is still a growing heap allocation. See [sample.out](maps/potential-memory-leak/sample.out) for details on memory loss per loop.
