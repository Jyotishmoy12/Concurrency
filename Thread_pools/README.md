## Problem with Creating a Thread per Request

Say we have a web server, and whenever a client connects to it, we fork a new thread to handle the request.  
So, if there are **n concurrent requests**, we would have **n threads** handling them.

### So, what’s the problem with this approach?

What happens when **n shoots up**?

- A large number of threads start running
- Threads consume a lot of system resources
- Hardware gets overwhelmed
- The system starts hanging and may eventually crash

Because of this, we **cannot keep creating threads endlessly**.  
We need to **limit the maximum number of threads** that can be created.

This is exactly what **thread pools** solve.

---

## Real-World Use Cases

- Web servers handling multiple clients simultaneously
- Asynchronous processing of messages from a message broker

---

## What Is a Thread Pool?

A **thread pool** is a collection of worker threads used to execute tasks concurrently.

### How it works:

- When a task needs to be executed, a thread is picked from the pool
- The task is delegated to that thread
- Once the task is complete, the thread is returned back to the pool for reuse
