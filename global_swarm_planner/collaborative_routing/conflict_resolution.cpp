#include <algorithm>
#include <iostream>
#include <string>
#include <vector>

struct RouteWindow {
    std::string vehicle_id;
    int slot_seconds;
    int duration_seconds;
};

std::vector<RouteWindow> deconflict(std::vector<RouteWindow> windows) {
    std::sort(windows.begin(), windows.end(), [](const RouteWindow& left, const RouteWindow& right) {
        return left.slot_seconds < right.slot_seconds;
    });

    int next_open_slot = 0;
    for (auto& window : windows) {
        if (window.slot_seconds < next_open_slot) {
            window.slot_seconds = next_open_slot;
        }
        next_open_slot = window.slot_seconds + window.duration_seconds;
    }

    return windows;
}

int main() {
    std::vector<RouteWindow> windows{{"ugv-02", 15, 20}, {"ugv-01", 0, 20}, {"ugv-03", 10, 25}};
    const auto resolved = deconflict(windows);

    for (const auto& window : resolved) {
        std::cout << window.vehicle_id << ":t+" << window.slot_seconds << "s for " << window.duration_seconds
                  << "s\n";
    }

    return 0;
}