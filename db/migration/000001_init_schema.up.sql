CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION public.uuid_if_empty(id uuid) RETURNS uuid 
    LANGUAGE plpgsql
    AS $$BEGIN 
        IF id = uuid_nil() THEN 
            RETURN uuid_generate_v4();
        ELSE 
            RETURN id;
        END IF;
    END$$;

CREATE TABLE "fruit" (
  "id" varchar PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "name" varchar NOT NULL,
  "color" varchar NOT NULL,
  "price" bigint NOT NULL,
  "quantity" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "fruit" ("name");
