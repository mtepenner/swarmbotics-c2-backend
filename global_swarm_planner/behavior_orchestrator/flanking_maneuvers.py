import json
import math
from dataclasses import dataclass


@dataclass(frozen=True)
class Axis:
    label: str
    bearing_degrees: float
    offset_m: float


def build_flank_plan() -> dict:
    axes = [
        Axis("northwest", 315.0, 250.0),
        Axis("center-fix", 0.0, 100.0),
        Axis("southeast", 135.0, 275.0),
    ]
    assignments = []
    for index, axis in enumerate(axes, start=1):
        assignments.append(
            {
                "axis": axis.label,
                "vehicleId": f"ugv-0{index}",
                "offsetVector": {
                    "x": round(math.cos(math.radians(axis.bearing_degrees)) * axis.offset_m, 2),
                    "y": round(math.sin(math.radians(axis.bearing_degrees)) * axis.offset_m, 2),
                },
            }
        )
    return {"behavior": "flanking-maneuvers", "assignments": assignments}


if __name__ == "__main__":
    print(json.dumps(build_flank_plan(), indent=2))