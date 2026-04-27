CREATE EXTENSION IF NOT EXISTS timescaledb;

CREATE TABLE IF NOT EXISTS mission_metrics (
    observed_at TIMESTAMPTZ NOT NULL,
    mission_id TEXT NOT NULL,
    vehicle_id TEXT NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    battery_percent DOUBLE PRECISION NOT NULL,
    payload_mode TEXT NOT NULL,
    PRIMARY KEY (observed_at, vehicle_id)
);

SELECT create_hypertable('mission_metrics', 'observed_at', if_not_exists => TRUE);

CREATE INDEX IF NOT EXISTS mission_metrics_mission_time_idx
    ON mission_metrics (mission_id, observed_at DESC);

CREATE MATERIALIZED VIEW IF NOT EXISTS mission_health_rollup
WITH (timescaledb.continuous) AS
SELECT
    time_bucket(INTERVAL '1 minute', observed_at) AS bucket,
    mission_id,
    AVG(battery_percent) AS average_battery,
    COUNT(DISTINCT vehicle_id) AS active_vehicles
FROM mission_metrics
GROUP BY bucket, mission_id;