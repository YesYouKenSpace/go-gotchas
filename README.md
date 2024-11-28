# go-snippets
This repository stores snippets of golang code. They are experiments to demonstrate certain behaviours.

Every `**/main.go` is an experiment on its own, accompanied with `sample.out` which contains both the stdout and stderr

## 1. RWMutex

1. [No read preference](rwmutex/no-read-preference/main.go) - demonstrates that there is no read preference to prevent recursive read lokcing
   >  a blocked Lock call excludes new readers from acquiring the lock
   > -- [src](https://pkg.go.dev/sync#RWMutex.RLock:~:text=a%20blocked%20Lock%20call%20excludes%20new%20readers%20from%20acquiring%20the%20lock)
   - read more about [readers-writers-problem](https://en.wikipedia.org/wiki/Readers%E2%80%93writers_problem) 

### 2. Maps

1. [Potential memory leak](maps/potential-memory-leak/main.go) 
   - better discussed in https://github.com/golang/go/issues/20135 
   - this snippet demonstrates that despite deleting all the keys and then attempting to reuse them, we still end up with a growing heap allocation - see [sample.out](maps/potential-memory-leak/sample.out) for a glimpse of how much we lose per loop
