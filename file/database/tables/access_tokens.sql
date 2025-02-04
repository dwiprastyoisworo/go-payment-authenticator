CREATE TABLE public.access_tokens (
      token VARCHAR(512) PRIMARY KEY,               -- Access token (harus di-hash)
      client_id INT NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
      user_id UUID NOT NULL,                        -- ID pengguna (resource owner)
      expires_at TIMESTAMPTZ NOT NULL,              -- Waktu kedaluwarsa (misal: 1 jam)
      refresh_token VARCHAR(512) UNIQUE,            -- Refresh token (harus di-hash)
      revoked BOOLEAN NOT NULL DEFAULT FALSE,       -- Apakah token dicabut?
      created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
      updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Index untuk query performa
CREATE INDEX idx_access_tokens_client_user ON access_tokens(client_id, user_id);
CREATE INDEX idx_access_tokens_expires ON access_tokens(expires_at);