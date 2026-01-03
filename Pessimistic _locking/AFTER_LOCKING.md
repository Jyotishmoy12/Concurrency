#### so to make the count++ is atomic:
- When one thread is executing count++, others should wait for that thread to finish.
- Pessimistic locking: You are locking a resource, one who has the lock processes the request, others wait
this implementation is called mutex basically mutually exclusive lock
- But it has performance issues like think we have only one thread which is doing all the tasks and others are just waiting

PS C:\Users\ASUS\OneDrive\Desktop\concurrency_in_depth\pesimistic_locking> go run main.go
1000000
PS C:\Users\ASUS\OneDrive\Desktop\concurrency_in_depth\pesimistic_locking> go run main.go
1000000
PS C:\Users\ASUS\OneDrive\Desktop\concurrency_in_depth\pesimistic_locking> go run main.go
1000000
PS C:\Users\ASUS\OneDrive\Desktop\concurrency_in_depth\pesimistic_locking> 