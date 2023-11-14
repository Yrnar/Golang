CREATE INDEX IF NOT EXISTS plantseed_name_idx ON plantseed USING GIN(to_tsvector('simple', name));
CREATE INDEX IF NOT EXISTS plantseed_family_idx ON plantseed USING GIN(to_tsvector('simple', family));
