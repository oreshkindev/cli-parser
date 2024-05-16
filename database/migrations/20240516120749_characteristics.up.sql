CREATE TABLE
    IF NOT EXISTS public.characteristics (
        id INT PRIMARY KEY,
        name TEXT UNIQUE NOT NULL,
        name_translated TEXT DEFAULT NULL
    );
