CREATE TABLE "permission" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar UNIQUE NOT NULL,
  "description" varchar
);
