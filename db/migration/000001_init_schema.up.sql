CREATE TABLE "fruit" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "color" varchar NOT NULL,
  "price" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "fruit" ("name");
