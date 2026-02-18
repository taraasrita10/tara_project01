# Kubernetes Fundamentals

## Kubernetes Primitives

### What is Kubernetes?
Kubernetes is an open-source container orchestration engine used for automatic deployment, scaling, and management of containerized workloads. 

### What is a Node?
A Node is a worker machine in Kubernetes, which can be either a physical or virtual machine. Each node contains the services necessary to run Pods, including the container runtime (like Podman) and a process called the Kubelet that manages the node's health.

### What is a Pod?
A Pod is the smallest deployable unit in Kubernetes and represents a single instance of a running process. It can contain one or more containers that share the same network IP, storage, and port space, allowing them to communicate via localhost.

### What is a Deployment?
A Deployment is a Kubernetes resource that manages the lifecycle of containerized applications by ensuring a specific number of Pod replicas are running at all times. It automatically handles the rollout of new versions and the replacement of failed Pods to maintain system availability.

### What is a Service?
A Service is an abstraction to help you expose groups of Pods over a network.

### What is a ReplicaSet vs. a ReplicationController?
A ReplicaSet is a controller that ensures a specific number of pod replicas are running at all times. ReplicaSets are the modern standard for the same because they allow for more flexible label selectors to identify which pods they should manage.

### What is a DaemonSet?
A DaemonSet ensures that a specific Pod runs on every single Node in the cluster. This is commonly used for system-level tools like log collectors or performance monitors that need to be present on every machine.

### What is a StatefulSet?
A StatefulSet is used for applications that require a unique identity or memory, such as databases. Unlike a regular Deployment where pods are interchangeable, StatefulSet pods are numbered (0, 1, 2) and retain their identity and storage even after a restart.

### What is a ConfigMap?
A ConfigMap stores non-sensitive configuration data, like environment variables or app settings, in key-value pairs. 

### What is a Secret?
A Secret is used for sensitive data like passwords or API keys, keeping them encoded so they aren't stored in plain text.

## kubectl and commands

### What is kubectl?
kubectl is the official command-line interface (CLI) used to communicate with the Kubernetes cluster's control plane. It converts commands into API calls that tell Kubernetes to create, modify, or delete resources like Pods and Deployments.

## Essential Kubectl Commands

| Action | Command | Description |
| ---- | ---- | ---- |
| **Apply Config** | `kubectl apply -f <filename.yaml>` | Creates or updates resources from a yaml file. |
| **List Resources**| `kubectl get pods/service/deployment` | Lists all active resources of a specific type. |
| **Check Details** | `kubectl describe pod <name>` | Shows technical details and events for a pod. |
| **View Logs** | `kubectl logs <name>` | Shows the output (errors/logs) from the container. |
| **Remote Shell** | `kubectl exec -it <name> -- sh` | Opens a terminal directly inside a running pod. |
| **Delete** | `kubectl delete -f <filename.yaml>` | Removes the resources defined in that file. |
| **Node Status** | `kubectl get nodes` | Checks if your worker machines are "Ready." |