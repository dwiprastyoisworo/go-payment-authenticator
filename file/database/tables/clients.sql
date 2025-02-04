CREATE TABLE public.clients (
    id SERIAL PRIMARY KEY,
    client_id VARCHAR(128) UNIQUE NOT NULL,       -- ID unik client
    client_secret VARCHAR(128) NOT NULL,          -- Harus di-hash (misal: bcrypt)
    name VARCHAR(255) NOT NULL,                   -- Nama aplikasi/client
    redirect_uris TEXT[] NOT NULL,                -- Daftar URI yang diizinkan
    enabled BOOLEAN NOT NULL DEFAULT TRUE,        -- Status aktif/tidak
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Index untuk pencarian cepat
CREATE INDEX idx_clients_client_id ON clients(client_id);