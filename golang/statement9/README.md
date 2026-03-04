# Statement 9: Distributed Trie Service

This project implements a Trie (prefix tree) data structure as a containerized gRPC service. It is designed to be deployed on Kubernetes using Minikube and utilizes Rook-Ceph for persistent data storage.

## Overview

The service provides a centralized system to store and manage strings using a Trie. This allows for efficient prefix-based searching and data organization. The service supports the following core operations:

*   **Add**: Insert a new string into the Trie.
*   **Check**: Verify if a specific string exists in the system.
*   **List**: Retrieve all strings currently stored in the Trie.
*   **Remove**: Delete a specific string from the Trie.

## Technical Architecture

The system is built using the following components:

1.  **gRPC**: Used for communication between the client and the server. This ensures fast and typed interactions based on the `trie.proto` definition.
2.  **Podman**: Used to build the container images for both the client and the server.
3.  **Kubernetes (Minikube)**: The platform where the service is deployed and managed.
4.  **Rook-Ceph Persistence**: Storage is handled through a Persistent Volume Claim (PVC). Rook-Ceph manages the underlying storage, ensuring that the Trie data is preserved even if the server pod is restarted or moved.

---

## Deployment Instructions

### 1. Build the Images
Use Podman to build the server and client images from the provided Containerfiles.

```bash
# Build the server image
podman build -t trie-server:latest -f Containerfile.server .

# Build the client image
podman build -t trie-client:latest -f Containerfile.client .
```

### 2. Export Images to Tarballs
Since Minikube operates in an isolated environment, the images must be saved as files to be transferred.

```bash
podman save trie-server:latest > trie-server.tar
podman save trie-client:latest > trie-client.tar
```

### 3. Load Images into Minikube
Load the generated tar files into the Minikube internal image registry.

```bash
minikube image load trie-server.tar
minikube image load trie-client.tar
```

### 4. Apply Kubernetes Configurations
Deploy the storage components followed by the application pods.

```bash
# Step A: Initialize Storage
kubectl apply -f k8s/storageclass.yaml
kubectl apply -f k8s/pvc.yaml

# Step B: Deploy the Server
kubectl apply -f k8s/server_pod.yaml

# Step C: Deploy the Client
kubectl apply -f k8s/client_pod.yaml
```

---

## File Structure

*   **client/**: Contains the client-side logic and `main.go`.
*   **server/**: Contains the server-side logic and `main.go`.
*   **pkg/trie.go**: The core Trie data structure implementation.
*   **proto/**: Contains the generated Go code from the Protobuf definitions.
*   **k8s/**: Kubernetes manifest files for pods, storage classes, and volume claims.
*   **Containerfile.client / Containerfile.server**: Build instructions for the container images.
*   **trie.proto**: The service contract defining the gRPC methods.