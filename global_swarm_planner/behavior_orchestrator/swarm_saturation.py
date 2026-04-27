import argparse
import json


def build_saturation_wave() -> dict:
    return {
        "behavior": "swarm-saturation",
        "waves": [
            {"wave": 1, "vehicles": ["ugv-01", "ugv-02"], "timeOffsetSeconds": 0},
            {"wave": 2, "vehicles": ["ugv-03", "ugv-04"], "timeOffsetSeconds": 18},
            {"wave": 3, "vehicles": ["ugv-05"], "timeOffsetSeconds": 30},
        ],
    }


def main() -> None:
    parser = argparse.ArgumentParser(description="Render a coordinated swarm saturation plan.")
    parser.add_argument("--once", action="store_true", help="Print one coordinated attack wave plan.")
    parser.parse_args()
    print(json.dumps(build_saturation_wave(), indent=2))


if __name__ == "__main__":
    main()