CREATE TABLE
    IF NOT EXISTS public.marketplaces (id INT PRIMARY KEY, name TEXT UNIQUE DEFAULT NULL);

INSERT INTO
    marketplaces (id, name)
VALUES
    (1, 'trendyoll'),
    (2, 'ozon'),
    (3, 'toyzshop') ON CONFLICT DO NOTHING;
