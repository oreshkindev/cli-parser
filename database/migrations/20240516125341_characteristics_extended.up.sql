CREATE TABLE
    IF NOT EXISTS public.characteristics_extended (
        id INT PRIMARY KEY,
        characteristics_id INT NOT NULL,
        name TEXT NOT NULL,
        name_translated TEXT DEFAULT NULL,
        CONSTRAINT fk_characteristics_id FOREIGN KEY (characteristics_id) REFERENCES characteristics (id)
    );
