# ZenStore - Distributed (p2p) file storage 

This project demonstrates basic distributed systems concepts including peer discovery, message serialization, file replication, and encrypted data transfer.

## Features

Peer-to-Peer Architecture
Each node can act as both client and server, storing and serving files.

Custom Binary Messaging
Uses Go’s gob encoding for structured message exchange (MessageStoreFile, MessageGetFile).

Decentralized File Replication
Nodes automatically broadcast file store and fetch requests to all connected peers.

Content Addressable Storage (CAS)
Files are stored using a SHA-1 hash–based directory structure for efficient organization and deduplication.

Encrypted Transfers
File data is encrypted and decrypted during transmission using a symmetric encryption key.

## Architecture

### Core Components
| Component | Responsibility |
|------------|----------------|
| **FileServer** | Manages storage, peers, and message routing |
| **Store** | Handles disk I/O and hashed file path mapping |
| **Transport** | TCP-based abstraction for peer communication (`p2p.Transport`) |


## Key Pkgs

### FileServer
Handles:
- Peer connection management (`OnPeer`)
- Broadcasting messages
- Processing file store and fetch requests
- Encrypting/decrypting data during transfer

### Store
Handles:
- Reading/writing/deleting files from disk
- Deterministic CAS-based file paths


A key "picture_1.png" → sha1 hash → folder path like:

root/peerID/abcde/fghij/klmno/.../abcdef1234...

