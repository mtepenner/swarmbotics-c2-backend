import argparse
import json
from pathlib import Path


def build_sync_plan(target: str) -> dict:
    snapshot_root = Path("telemetry_pipeline") / "state_storage"
    files = [
        snapshot_root / "redis_live_state.conf",
        snapshot_root / "timescaledb_metrics.sql",
    ]
    return {
        "target": target,
        "mode": "dry-run",
        "artifacts": [str(path) for path in files],
        "steps": [
            "freeze planner writes",
            "export redis hot state",
            "replay latest TimescaleDB rollups",
            "promote backup edge node",
        ],
    }


def main() -> None:
    parser = argparse.ArgumentParser(description="Prepare or execute regional swarm state synchronization.")
    parser.add_argument("--target", required=True, help="Edge node or region that should receive the warm state.")
    parser.add_argument("--once", action="store_true", help="Emit one sync plan and exit.")
    args = parser.parse_args()

    print(json.dumps(build_sync_plan(args.target), indent=2))


if __name__ == "__main__":
    main()