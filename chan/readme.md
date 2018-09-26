Use channels to orchestrate and coordinate goroutines
focus on the signaling semantics and not the sharing of data
signaling with data or without data

### Unbuffered Channels

- receive happens before the send
- guarantee the signal has been received
- Unknown latency on when the signal will be received

### Buffered Channels

- Send happens before the receive
- reduce blocking latency between signaling
- No guarantee when the signal has been received
  - buffer of 1 can give you one delayed send of guarantee

### Closing Channels

- close happens before the receive
- signaling without data
- perfect for signaling cancellation and deadlines

### nil channels

- send and receive block
- turn off signaling
- perfect for rate limting or short term stoppages



### Design Philosophy:

Depending on the problem you are solving, you may require different channel semantics. Depending on the semantics you need, different architectural choices must be taken.

- If any given Send on a channel CAN cause the sending goroutine to block:
  - Not allowed to use a Buffered channel larger than 1.
     Buffers larger than 1 must have reason/measurements.
  - Must know what happens when the sending goroutine blocks.

- If any given Send on a channel WON'T cause the sending goroutine to block:
  - You have the exact number of buffers for each send.
    - Fan Out pattern
  - You have the buffer measured for max capacity.
    - Drop pattern
- Less is more with buffers.
  - Don’t think about performance when thinking about buffers.
  - Buffers can help to reduce blocking latency between signaling.
    - Reducing blocking latency towards zero does not necessarily mean better throughput.
    - If a buffer of one is giving you good enough throughput then keep it.
    - Question buffers that are larger than one and measure for size.
    - Find the smallest buffer possible that provides good enough throughput.