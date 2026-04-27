# The Command & Control (C2) Backend

## Description
`swarmbotics-c2-backend` serves as the "Hive Mind" for massive-scale robotic swarm operations. It provides a highly scalable Command and Control (C2) architecture designed to handle distributed task allocation, collaborative routing, and high-throughput telemetry ingestion. Built for tactical edge networks, it ensures zero-trust security and seamless failover across distributed environments.

## Table of Contents
- [Features](#-features)
- [Technologies Used](#%EF%B8%8F-technologies-used)
- [Installation](#%EF%B8%8F-installation)
- [Usage](#-usage)
- [Contributing](#-contributing)
- [License](#-license)

## 🚀 Features

* **Global Swarm Planner**: Utilizes auction-based algorithms for distributed tasking and assigns roles based on the tactical situation. Includes high-performance C++ waypoint generation to prevent bottlenecks and collisions among hundreds of UGVs.
* **Tactical Behavior Orchestration**: Executes "multiple dilemmas" logic, coordinating multi-axis flanking maneuvers and swarm saturation attacks to overwhelm adversary defenses.
* **High-Throughput Telemetry Pipeline**: Ingests swarm-wide data using Kafka and MQTT. Real-time stream processors written in Go handle health monitoring and payload aggregation.
* **Ultra-Fast State Storage**: Leverages Redis for sub-5ms instant UGV location lookups and TimescaleDB for time-series mission playback and post-action review.
* **Secure API Gateway**: Bridges the tactical network with operator interfaces via ultra-fast gRPC and REST endpoints. Protected by Mutual TLS (mTLS) for zero-trust access and JWT token management.
* **Resilient Infrastructure**: Kubernetes-native deployment with built-in regional synchronization scripts for disaster recovery and distributed redundancy.

## 🛠️ Technologies Used
* **Languages**: Go, Python, C++
* **APIs & Auth**: gRPC, REST, mTLS, JWT
* **Data & Messaging**: Kafka, MQTT, Redis, TimescaleDB
* **DevOps & Infrastructure**: Kubernetes, Docker

## ⚙️ Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/mtepenner/swarmbotics-c2-backend.git
   cd swarmbotics-c2-backend
   ```
2. Provision your cluster infrastructure (ensure Redis, Kafka, and TimescaleDB are running).
3. Apply the Kubernetes manifests to deploy the C2 microservices:
   ```bash
   kubectl apply -f infrastructure/k8s-manifests/
   ```

## 💻 Usage

### Managing the API Gateway
The API gateway serves as the primary entry point for ATAK and Web interfaces. Ensure the mTLS validator is properly configured with your organization's root certificates:
```bash
# Check gateway deployment status
kubectl get deployments | grep api-gateway
```

### Monitoring the Hive Mind State
To verify the telemetry pipeline is processing UGV health pings and payload data:
```bash
# Tail logs for the stream processors
kubectl logs -l app=telemetry-pipeline --tail=100 -f
```

### Regional Failover
In the event of a C2 node failure, trigger the state sync script to push the latest Redis states to a backup edge server:
```bash
python3 infrastructure/failover/regional-sync.py --target edge-node-b
```

## 🤝 Contributing
We welcome contributions aimed at optimizing the C++ conflict resolution algorithms, expanding the Python-based tactical behaviors, or hardening the mTLS auth service. Please ensure all pull requests pass the CI pipeline and include updated protobuf definitions if modifying the gRPC streams.

## 📄 License
This project is licensed under the [MIT License](LICENSE) - see the LICENSE file for details. Copyright (c) 2026 Matthew Penner.
