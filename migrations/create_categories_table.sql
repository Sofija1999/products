-- Drop table

-- DROP TABLE public.categories;

CREATE TABLE public.categories (
	category_id bigserial NOT NULL,
	category_name varchar NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	CONSTRAINT categories_pk PRIMARY KEY (category_id)
);

ALTER TABLE products
ADD COLUMN category_id int4 NULL;

ALTER TABLE products
ADD CONSTRAINT products_fk FOREIGN KEY (category_id) REFERENCES categories(category_id);