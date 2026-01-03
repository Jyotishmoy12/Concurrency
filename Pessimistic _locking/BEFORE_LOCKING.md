#### Point to be noted: It should be 1 million, but it's not true.

--- This shows that count++ is not atomic in nature

```
PS C:\Users\ASUS\OneDrive\Desktop\concurrency_in_depth\pesimistic_locking> go run main.go
948611
PS C:\Users\ASUS\OneDrive\Desktop\concurrency_in_depth\pesimistic_locking> go run main.go
947150
PS C:\Users\ASUS\OneDrive\Desktop\concurrency_in_depth\pesimistic_locking> go run main.go
947862
PS C:\Users\ASUS\OneDrive\Desktop\concurrency_in_depth\pesimistic_locking>

```
