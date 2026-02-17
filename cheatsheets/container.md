# Container Fundamentals 

## Containers and Containerization 

### What is a container?
A container is an isolated environment which has its own processes, networks and mounts, but shares the same OS kernel with other containers. It consists of an application along with its dependencies packaged together such that they can run in any environment.

### Why is a container used?
A container solves two issues :
    - **"But it works on my system!"** -> Packages applications and their dependencies together so that the application would work when run anywhere, in any environment.
    - **Provides isolation while being lightweight** -> Containers act like their own entities, while sharing the same OS kernel, thereby providing isolation (Like VMs), but are lighter.

### How to make a container?
A container can be be made using tools known as container engines, such as Docker, Podman, etc., which package an application and its dependencies together and containerize it. It requires a set of instructions known as the Containerfile, which describes the steps and layers required to build a container image.

### Image vs Container
An image is a blueprint of a container, which can be built (using a containerfile) or pulled from an image registry. A container is a running instance of an image.

## Containerflow and Containerfiles

### What is Containerflow?
Building a container follows a specific 3-step lifecycle:
1. **Write**: Create a Containerfile defining your environment.
2.  **Build**: Use the engine to compile the file into an Image.
3.  **Run**: Instantiate the image into a living Container.

### What is a Containerfile?
A Containerfile is a document consisting of the instructions/commands used to build a container image.

### Essential commands for writing a containerfile
| Command | Description |
| ---- | ---- |
| `FROM` | Sets the Base Image (e.g., `FROM golang:1.21`) |
| `WORKDIR` | Sets the working directory inside the container |
| `COPY` | Copies files from your local machine to the container |
| `RUN` | Executes commands during the build phase |
| `ENV` | Sets environment variables |
| `EXPOSE` | Informs the container which port to listen on |
| `CMD` | The default command to run when the container starts |
| `ENTRYPOINT` | Configures a container that will run as an executable |

## Container engines : Podman

### What is Podman?
It is a daemonless container engine which is used tfor developing, managing, and running containers.

### Essential Podman commands
### Podman Common Commands

| Goal | Command |
| ---- | ---- |
| **Build an image** | `podman build -t my-app .` |
| **Run a container** | `podman run -d --name app-instance -p 8080:8080 my-app` |
| **List images** | `podman images` |
| **List running containers** | `podman ps` |
| **Stop a container** | `podman stop <name>` |
| **Remove a container** | `podman rm <name>` |
| **See logs** | `podman logs -f <name>` |






