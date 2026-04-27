PYTHON ?= python
GO ?= go

.PHONY: validate python-check go-check smoke cpp-check

validate: go-check python-check smoke

go-check:
	$(GO) test ./...

python-check:
	$(PYTHON) -m compileall global_swarm_planner infrastructure

smoke:
	$(PYTHON) global_swarm_planner/task_allocation/dynamic_auctioneer.py --once
	$(PYTHON) global_swarm_planner/behavior_orchestrator/swarm_saturation.py --once
	$(PYTHON) infrastructure/failover/regional_sync.py --target edge-node-b --once

cpp-check:
	g++ -std=c++17 -fsyntax-only global_swarm_planner/collaborative_routing/waypoint_generator.cpp
	g++ -std=c++17 -fsyntax-only global_swarm_planner/collaborative_routing/conflict_resolution.cpp