CREATE TABLE IF NOT EXISTS plantseed(
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    family text NOT NULL,
    amount integer NOT NULL,
    price integer NOT NULL
);
