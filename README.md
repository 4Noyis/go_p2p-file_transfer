# P2P File Transfer

## Peer A (Sender)

- Waits connection to given port with using TCP listener ( `net.Listen` ).

- When connection is accepted, opens the file and write file context to connection (sends).

## Peer B (Receiver)

- Makes connection to Peer A with using TCP via given address (localhost:8080).

- Creates new file and reads(receive) files data with `io.Copy` and writes this data to file.
