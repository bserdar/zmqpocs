# ZeroMQ: Multiple publishers for the same socket

Typical pub-sub scenario has one publisher and multiple
subscribers. Does zmq handle multiple publishers to the same address?

Setup:

 * Server, publishes to a socket. We'll run a few instances of this
 * Client, receives from the socket. We'll test whether all clients
   can receive all server messages
   
So zmq can deal with multiple publishers using the same address.

