# Distributed Trie Service with gRPC & Kubernetes

An implementation of a **Trie (Prefix Tree)**–based string management system implemented in **Go**, exposed via **gRPC**, containerized using **Podman**, and deployed on a local **Kubernetes (Minikube)** cluster.

The system consists of:
- **Server:** Hosts the in‑memory Trie and exposes gRPC endpoints.
- **Client:** Interactive CLI that communicates with the server.

---

## 1. Features

- **Add** a string  
- **Remove** a string  
- **Check** if a string exists  
- **List** all stored strings  
- gRPC communication over HTTP/2  
- Fully containerized & deployable on Kubernetes  

---

## 2. Technology Stack

| Component | Version / Details |
|----------|-------------------|
| Language | Go 1.25+ |
| gRPC | Protocol Buffers v3 |
| Containerization | Podman |
| Orchestration | Minikube (Kubernetes 1.31+) |

---

## 3. Setup

1. Start minikube with Podman driver

```bash
minikube start --driver=podman
```
2. Generate grpc code

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/trie.proto
```
3. Build and load containers into minikube

```bash
# Build images
podman build -t trie-server:latest -f server.Dockerfile .
podman build -t trie-client:v4 -f client.Dockerfile .

# Save to tar and load into Minikube
podman save -o server.tar localhost/trie-server:latest
podman save -o client.tar localhost/trie-client:v3

minikube image load server.tar
minikube image load client.tar
```

4. Deploy to kubernetes

```bash
kubectl apply -f deploy.yaml
```

5. Attach to client terminal to access the running program

```bash
kubectl attach -it trie-client-pod
```

---