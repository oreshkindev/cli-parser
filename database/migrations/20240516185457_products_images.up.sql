CREATE TABLE
    IF NOT EXISTS public.products_images (
        id SERIAL PRIMARY KEY,
        products_id INT NOT NULL,
        href TEXT[] NOT NULL,
        CONSTRAINT fk_products_id FOREIGN KEY (products_id) REFERENCES products (id)
    );
