CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" VARCHAR(255) NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "role_id" INTEGER NOT NULL,
  CONSTRAINT fk_role_users FOREIGN KEY (role_id) REFERENCES role(id)
 ON UPDATE CASCADE ON DELETE RESTRICT
 );
  

