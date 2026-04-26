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
