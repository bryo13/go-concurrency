## Golang Concurrency model
1. There is a difference btn concurrency and parallelism.  
	- Concurrency is a value of code, parallelism is a property of a running program.  
2. Golang's goroutines are coroutines and are non preemtive(cant be interrupted). These means they are concurrent routines(methods) and are different to OS threads and green threads(threads managed by a languages runtime.  
  

3. Golang uses a fork and join model, meaning a flow of execution can spawn a child which can be run concurrently with the parent - the fork point - and a join point where the child and the parent will 'join' together in  future - this is archieved using the waitgroup package. wg.Wait() for joining a branch.    

4. Goroutines work in the same address space they were created in, hence they modify variables - note, on closures, its prudent to pass the variables i.e as 'params' in functions. Example will be in this codebase in a branch called 'one'.  
  

