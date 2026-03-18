# Mini-CI Orchestrator
### A Cloud-Native Security Scanning Pipeline on Kubernetes

A compact and powerful orchestration system built in **Go**, designed to manage the full lifecycle of security scans within a Kubernetes cluster. This project demonstrates dynamic resource management, CI/CD automation, and distributed microservices architecture.

---

## The Concept

Instead of running heavy, resource-intensive scans on the backend server, this system leverages Kubernetes as an **Execution Engine**. 

The **Go Orchestrator** receives API requests and triggers real-time, isolated **Kubernetes Jobs**. These jobs perform the scan, persist artifacts to a shared volume, and ensure resource cleanup upon completion.

---

## Tech Stack

* **Backend:** [Go (Golang)](https://go.dev/) utilizing `client-go` for Kubernetes API interaction.
* **Infrastructure:** [Kubernetes](https://kubernetes.io/) (Jobs, Deployments, Services, RBAC).
* **Security:** [TruffleHog](https://github.com/trufflesecurity/trufflehog) (Secret & Key scanning).
* **Storage:** Persistent Volume Claims (PVC) for inter-pod data sharing.
* **DevOps:** Docker (Multi-stage builds), Containerization, YAML manifests.

---

## System Architecture

2.  **Dynamic Job Controller:** A custom mechanism that constructs K8s Job objects on-the-fly, including specific Resource limits and Timeouts.
3.  **Shared Storage Pipeline:** Utilizes a shared PVC allowing **Init Containers** to clone code and the **Scanner Container** to analyze it within the same volume.
4.  **Artifact Manager:** Logic for fetching the latest results from the volume and serving them as JSON via the API.
