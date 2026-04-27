#include <array>
#include <iomanip>
#include <iostream>
#include <string>
#include <vector>

struct Waypoint {
    std::string vehicle_id;
    double latitude;
    double longitude;
    double altitude_m;
};

std::vector<Waypoint> generate_waypoints(const std::vector<std::string>& vehicle_ids) {
    std::vector<Waypoint> waypoints;
    waypoints.reserve(vehicle_ids.size());

    constexpr std::array<double, 2> anchor{33.6500, -117.8400};
    for (std::size_t index = 0; index < vehicle_ids.size(); ++index) {
        waypoints.push_back(Waypoint{
            vehicle_ids[index],
            anchor[0] + static_cast<double>(index) * 0.0014,
            anchor[1] - static_cast<double>(index) * 0.0011,
            18.0 + static_cast<double>(index),
        });
    }

    return waypoints;
}

int main() {
    const std::vector<std::string> vehicle_ids{"ugv-01", "ugv-02", "ugv-03"};
    const auto waypoints = generate_waypoints(vehicle_ids);

    std::cout << std::fixed << std::setprecision(4);
    for (const auto& waypoint : waypoints) {
        std::cout << waypoint.vehicle_id << ',' << waypoint.latitude << ',' << waypoint.longitude << ','
                  << waypoint.altitude_m << '\n';
    }

    return 0;
}