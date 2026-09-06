ALTER TABLE skills ADD COLUMN IF NOT EXISTS effect_image TEXT NOT NULL DEFAULT '';
UPDATE skills SET effect_image = effects WHERE effect_image = '' AND effects LIKE 'data:image/%';
