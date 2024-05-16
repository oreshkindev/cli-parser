CREATE TABLE
    IF NOT EXISTS public.characteristics_extended (
        id INT PRIMARY KEY,
        name TEXT UNIQUE NOT NULL,
        name_translated TEXT DEFAULT NULL
    );
