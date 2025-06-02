ALTER TABLE accounts ALTER COLUMN id SET DEFAULT gen_random_uuid();
ALTER TABLE transfers ALTER COLUMN id SET DEFAULT gen_random_uuid();
ALTER TABLE entries ALTER COLUMN id SET DEFAULT gen_random_uuid();