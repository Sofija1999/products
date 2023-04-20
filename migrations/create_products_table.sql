-- Drop table

-- DROP TABLE public.products;

CREATE TABLE public.products (
	id bigserial NOT NULL,
	"name" varchar NULL,
	shortdescription text NULL,
	description text NULL,
	price numeric NULL,
	created timestamp NULL,
	updated timestamp NULL,
	CONSTRAINT products_pk PRIMARY KEY (id)
);
