CREATE DATABASE dosa_ops;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    api_key TEXT NOT NULL,
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


CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

CREATE INDEX idx_metrics_service_metric_time
ON metrics(service_id, metric_type_id, timestamp DESC);
