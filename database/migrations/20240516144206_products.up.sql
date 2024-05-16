CREATE TABLE
    IF NOT EXISTS public.products (
        id INT PRIMARY KEY,
        brands_id INT NOT NULL,
        characteristics INT[] DEFAULT NULL,
        created_at TIMESTAMP DEFAULT now (),
        depth DOUBLE PRECISION DEFAULT 0.0,
        description TEXT DEFAULT NULL,
        description_translated TEXT DEFAULT NULL,
        height DOUBLE PRECISION DEFAULT 0.0,
        href TEXT NOT NULL,
        marketplaces_id INT NOT NULL,
        markup DOUBLE PRECISION DEFAULT 0.0,
        name TEXT NOT NULL,
        name_translated TEXT DEFAULT NULL,
        price DOUBLE PRECISION DEFAULT 0.0,
        quantity INT DEFAULT 0,
        updated_at TIMESTAMP DEFAULT now (),
        weight DOUBLE PRECISION DEFAULT 0.0,
        width DOUBLE PRECISION DEFAULT 0.0,
        CONSTRAINT fk_brands_id FOREIGN KEY (brands_id) REFERENCES brands (id),
        CONSTRAINT fk_marketplaces_id FOREIGN KEY (marketplaces_id) REFERENCES marketplaces (id)
    );
