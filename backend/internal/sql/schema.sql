CREATE DATABASE metrics;

CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE metric_types (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE
);

CREATE TABLE metrics (
    id BIGSERIAL PRIMARY KEY,
    service_id INT REFERENCES services(id),
    metric_type_id INT REFERENCES metric_types(id),
    value DOUBLE PRECISION,
    timestamp TIMESTAMP NOT NULL
);

CREATE INDEX idx_metrics_service_metric_time
ON metrics(service_id, metric_type_id, timestamp DESC);
