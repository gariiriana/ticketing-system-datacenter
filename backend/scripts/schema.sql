-- Schema for Ticketing System Datacenter

-- Users Table
CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    role TEXT NOT NULL DEFAULT 'engineer'
);

-- Sites Table
CREATE TABLE IF NOT EXISTS sites (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    address TEXT
);

-- Tickets Table
CREATE TABLE IF NOT EXISTS tickets (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id),
    site_id TEXT NOT NULL REFERENCES sites(id),
    description TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending',
    photo_url TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    approved_by TEXT REFERENCES users(id),
    approved_at TIMESTAMP WITH TIME ZONE
);

-- Initial Mock Data
INSERT INTO users (id, name, email, role) VALUES 
('admin-001', 'Admin Utama', 'admin@example.com', 'admin'),
('engineer-001', 'Engineer Satu', 'engineer1@example.com', 'engineer')
ON CONFLICT (id) DO NOTHING;

INSERT INTO sites (id, name, address) VALUES 
('site-001', 'Datacenter Jakarta', 'Jl. Kuningan No. 1'),
('site-002', 'Datacenter BSD', 'Jl. BSD Raya No. 10')
ON CONFLICT (id) DO NOTHING;
