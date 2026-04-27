import argparse
import json
from dataclasses import dataclass


@dataclass(frozen=True)
class MissionTask:
    task_id: str
    priority: int
    distance_km: float


@dataclass(frozen=True)
class Bidder:
    vehicle_id: str
    battery_percent: float
    payload_ready: bool


def score_bid(task: MissionTask, bidder: Bidder) -> float:
    readiness_bonus = 25.0 if bidder.payload_ready else -10.0
    battery_bonus = bidder.battery_percent * 0.55
    distance_penalty = task.distance_km * 1.4
    priority_bonus = task.priority * 7.5
    return round(priority_bonus + readiness_bonus + battery_bonus - distance_penalty, 2)


def run_cycle() -> dict:
    tasks = [
        MissionTask("route-alpha", priority=5, distance_km=1.8),
        MissionTask("route-bravo", priority=4, distance_km=2.6),
        MissionTask("route-charlie", priority=3, distance_km=3.1),
    ]
    bidders = [
        Bidder("ugv-01", battery_percent=91.0, payload_ready=True),
        Bidder("ugv-02", battery_percent=76.0, payload_ready=True),
        Bidder("ugv-03", battery_percent=62.0, payload_ready=False),
    ]

    assignments = []
    available = bidders[:]
    for task in tasks:
        winner = max(available, key=lambda bidder: score_bid(task, bidder))
        assignments.append(
            {
                "taskId": task.task_id,
                "winner": winner.vehicle_id,
                "score": score_bid(task, winner),
            }
        )
        available.remove(winner)

    return {"planner": "dynamic-auctioneer", "assignments": assignments}


def main() -> None:
    parser = argparse.ArgumentParser(description="Run a single auction cycle for swarm tasking.")
    parser.add_argument("--once", action="store_true", help="Execute one planner cycle and exit.")
    parser.parse_args()
    print(json.dumps(run_cycle(), indent=2))


if __name__ == "__main__":
    main()