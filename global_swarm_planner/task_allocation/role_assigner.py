import json
from dataclasses import dataclass


@dataclass(frozen=True)
class VehicleState:
    vehicle_id: str
    battery_percent: float
    payload_type: str
    comms_margin: float


def assign_role(state: VehicleState) -> str:
    if state.payload_type == "isr" and state.comms_margin >= 0.8:
        return "forward-observer"
    if state.payload_type == "kinetic" and state.battery_percent >= 60.0:
        return "breach-lead"
    if state.battery_percent < 40.0:
        return "reserve-relay"
    return "support-screen"


def main() -> None:
    vehicles = [
        VehicleState("ugv-01", 88.0, "isr", 0.91),
        VehicleState("ugv-02", 67.0, "kinetic", 0.74),
        VehicleState("ugv-03", 38.0, "cargo", 0.83),
    ]
    assignment = {vehicle.vehicle_id: assign_role(vehicle) for vehicle in vehicles}
    print(json.dumps({"roleAssignments": assignment}, indent=2))


if __name__ == "__main__":
    main()