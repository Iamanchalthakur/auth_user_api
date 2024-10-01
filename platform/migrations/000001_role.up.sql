CREATE TABLE "role" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar UNIQUE NOT NULL,
  "description" varchar
);
