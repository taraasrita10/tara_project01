# Rook Cheatsheet

### 1. Minikube Setup
Run this command to start a virtual machine with 8GB RAM, 4 CPUs, and an extra disk for storage.
```bash
minikube start --driver=qemu --extra-disks=1 --memory=8192 --cpus=4
```
This command gives the VM 8GB RAM and 4 CPUs.

### 2. Install Rook Operator
Apply the common resources and the operator.
```bash
kubectl create -f common.yaml
kubectl create -f crds.yaml
kubectl create -f operator.yaml
```
This sets up the basic framework and the controller that manages the storage.

### 3. Create the Cluster
Create the storage cluster.
```bash
kubectl create -f cluster-test.yaml
```
This starts the Ceph storage nodes and services.

### 4. Storage Class
Create a storage class to define how storage is provisioned.
```bash
kubectl create -f storageclass.yaml
```
This allows you to request storage using a PersistentVolumeClaim (PVC).

### 5. Provisioning a PVC
Request a specific amount of storage.
```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: ceph-pvc
spec:
  storageClassName: rook-ceph-block
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
```
This creates a 1GB volume for your application.

### 6. Snapshots
Create a VolumeSnapshotClass and then a Snapshot.
```bash
kubectl create -f snapshotclass.yaml
kubectl create -f snapshot.yaml
```
This allows you to save the state of your storage at a specific point in time.