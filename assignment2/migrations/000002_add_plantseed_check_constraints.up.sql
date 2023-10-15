ALTER TABLE plantseed ADD CONSTRAINT plantseed_amount_check CHECK (amount > 0);
ALTER TABLE plantseed ADD CONSTRAINT plantseed_price_check CHECK (price > 0);
