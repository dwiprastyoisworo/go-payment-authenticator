CREATE TABLE public.authorization_codes (
    code VARCHAR(128) PRIMARY KEY,                -- Authorization code (harus di-hash)
    client_id INT NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
    expires_at TIMESTAMPTZ NOT NULL,              -- Waktu kedaluwarsa (misal: 10 menit)
    used BOOLEAN NOT NULL DEFAULT FALSE,          -- Apakah code sudah digunakan?
    redirect_uri TEXT NOT NULL,                   -- Redirect URI saat request
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Index untuk query berdasarkan client dan waktu
CREATE INDEX idx_auth_codes_client_user ON authorization_codes(client_id);
CREATE INDEX idx_auth_codes_expires ON authorization_codes(expires_at);