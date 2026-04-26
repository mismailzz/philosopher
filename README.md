```bash
# -- initial Commit -- 
➜  philosopher git:(main) go run main.go
Philosopher 4 is eating now.
Philosopher 2 is eating now.
Philosopher 3 is eating now.
Philosopher 1 is eating now.
Philosopher 0 is eating now.
fatal error: all goroutines are asleep - deadlock!
```

Fix: add wg.Done on go-routine completion  
We used wg.Add(1), but never make the wg.Done(), due to which the Waitgroup counter becomes 5 but never decreased. Due to which the go-routines ended, but the main goroutine is waiting, and no other goroutine can make progress.

```
Philosopher 3 is eating now.
Philosopher 2 is eating now.
Philosopher 1 is eating now.
Philosopher 0 is eating now.
Philosopher 4 is eating now.
All philosophers end eating
```

Fix: Added visibility of deadlock

```bash
## In Go, the runtime panic tells you which goroutines are blocked, but it usually does not directly tell you which mutex/fork caused it. So its better to have name for the resouces. Currently we are using only sync.Mutex. Secondly its better to add logging before and after lock.

For deadlock debugging, log these two moments:
before Lock  = waiting/requesting
after Lock   = acquired/holding
```
