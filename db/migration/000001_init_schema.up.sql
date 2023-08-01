CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "photo" varchar DEFAULT 'profile.jpg',
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "customers" (
  "id" bigserial PRIMARY KEY,
  "customer_name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "shops" (
  "id" bigserial PRIMARY KEY,
  "shop_name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "phone_number" varchar NOT NULL
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "product_name" varchar NOT NULL,
  "price" integer NOT NULL,
  "product_image" varchar DEFAULT 'product_image.jpg'
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "customer_id" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()',
  "total_price" integer NOT NULL,
  "cash" integer NOT NULL,
  "return" integer NOT NULL
);

CREATE TABLE "order_details" (
  "id" bigserial PRIMARY KEY,
  "order_id" integer NOT NULL,
  "product_id" integer NOT NULL,
  "price" integer NOT NULL,
  "quantity" integer NOT NULL,
  "total" integer NOT NULL
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "customers" ("customer_name");

CREATE INDEX ON "products" ("product_name");

CREATE INDEX ON "orders" ("customer_id");

CREATE INDEX ON "orders" ("username");

CREATE INDEX ON "orders" ("created_at");

ALTER TABLE "orders" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "orders" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");