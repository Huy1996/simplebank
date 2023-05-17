# GRPC

## What is GRPC

### Remote procedure Call Framework

- The client can execute a remote procedure on the server
- The remote interaction code is handled by gRPC
- The API & data structure code is automatically generated
- Support multiple programming languages

## How it works?

1. Define API & data structure
The RPC and its request/response structure are difined using protobuf

2. Generate gRPC stubs
Generate codes for the server and client in the language of your choice

3. Implement the server
Implement the RPC handler on the server side

4. Use the client
Use the generated client stubs to call the RPC on the server

## Why GRPC?

**High performance**
HTTP/2: binary framing, multiplexing, header compression, bidirectional communication.

**Strong API contract**
Server & client share the same protobuf RPC definition with strongly typed data

**Automatic code generation**
Codes that serialize/deserialize data, or transfer data between client & server are automatically generated

## 4 Types of GRPC

- Unary gRPC
- Client streaming gRPC
- Server streaming gRPC
- Bidirectional streaming gRPC

## GRPC Gateway

**Serve both GRPC and HTTP request at once**

- A plugin of protobuf compiler
- Generate proxy codes from protobuf
- Translate HTTP JSON calls to gRPC
  - In-process translation: only for unary
  - Separate proxy server: both unary and streaming
- Write code once, serve both gRPC and HTTP requests

---

## gRPC Gateway

- A plugin of protobuf compiler
- Generate proxy codes from protobuf
- Translate HTTP JSON calls to gRPC
  - In-process translation: only for unary
  - Separate proxy server: both unary and streaming

  