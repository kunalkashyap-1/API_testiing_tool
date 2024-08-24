-- Users Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Authentication Tokens Table
CREATE TABLE auth_tokens (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    refresh_token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Scenarios Table
CREATE TABLE scenarios (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    config JSONB NOT NULL,  -- Stores the configuration like requests_per_second, duration, etc.
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tests Table
CREATE TABLE tests (
    id SERIAL PRIMARY KEY,
    scenario_id INT REFERENCES scenarios(id),
    user_id INT REFERENCES users(id),
    status VARCHAR(50) NOT NULL,  -- e.g., 'pending', 'running', 'completed', 'aborted'
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Test Results Table
CREATE TABLE test_results (
    id SERIAL PRIMARY KEY,
    test_id INT REFERENCES tests(id),
    metric_type VARCHAR(255) NOT NULL,  -- e.g., 'response_time', 'error_rate', 'throughput'
    metric_value DOUBLE PRECISION NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Metrics Table
CREATE TABLE metrics (
    id SERIAL PRIMARY KEY,
    test_id INT REFERENCES tests(id),
    request_count INT NOT NULL,
    avg_response_time DOUBLE PRECISION NOT NULL,
    error_rate DOUBLE PRECISION NOT NULL,
    throughput DOUBLE PRECISION NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Reports Table
CREATE TABLE reports (
    id SERIAL PRIMARY KEY,
    test_id INT REFERENCES tests(id),
    report_data JSONB NOT NULL,  -- Stores report details like charts, summaries, etc.
    generated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Logs Table
CREATE TABLE logs (
    id SERIAL PRIMARY KEY,
    test_id INT REFERENCES tests(id),
    log_message TEXT NOT NULL,
    log_level VARCHAR(50) NOT NULL,  -- e.g., 'INFO', 'WARNING', 'ERROR'
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
