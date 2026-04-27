# The Command & Control (C2) Backend

`swarmbotics-c2-backend` is the backend scaffold for swarm-scale command and control. This repository now implements the blueprint structure with compileable Go gateway and telemetry packages, Python planning and behavior scripts, standalone C++ routing tools, storage configuration artifacts, and Kubernetes deployment manifests.

## Implemented Layout

- `api_gateway/`
   - `proto/`: gRPC contract definitions for swarm control and telemetry streaming
   - `rest_endpoints/`: route and node-status shaping helpers for operator-facing APIs
   - `auth_service/`: mTLS certificate validation and JWT issuance/verification helpers
- `global_swarm_planner/`
   - `task_allocation/`: auction-based mission assignment and role selection
   - `collaborative_routing/`: standalone C++ waypoint generation and deconfliction tools
   - `behavior_orchestrator/`: flanking and saturation behavior planners
- `telemetry_pipeline/`
   - `ingestion_gateways/`: MQTT broker configuration and Kafka envelope construction
   - `stream_processors/`: health alerting and ISR payload aggregation
   - `state_storage/`: Redis tuning and TimescaleDB schema/bootstrap SQL
- `infrastructure/`
   - `k8s_manifests/`: baseline deployment and stateful workload manifests
   - `failover/`: regional state sync planning script

## Technologies

- Languages: Go, Python, C++17
- Interfaces: gRPC, REST, mTLS, JWT
- Data plane: Kafka, MQTT, Redis, TimescaleDB
- Deployment: Kubernetes manifests and failover utilities

## Validation

From the repository root:

```bash
go test ./...
python -m compileall global_swarm_planner infrastructure
python global_swarm_planner/task_allocation/dynamic_auctioneer.py --once
python infrastructure/failover/regional_sync.py --target edge-node-b --once
```

If a local C++ compiler is available:

```bash
g++ -std=c++17 -fsyntax-only global_swarm_planner/collaborative_routing/waypoint_generator.cpp
g++ -std=c++17 -fsyntax-only global_swarm_planner/collaborative_routing/conflict_resolution.cpp
```

## Deployment

Apply the baseline manifests after wiring real images, secrets, broker addresses, and storage endpoints:

```bash
kubectl apply -f infrastructure/k8s_manifests/
```

The failover utility can be dry-run tested locally:

```bash
python infrastructure/failover/regional_sync.py --target edge-node-b --once
```

## License

This project is licensed under the [MIT License](LICENSE).
