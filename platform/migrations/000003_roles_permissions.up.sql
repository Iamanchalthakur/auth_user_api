CREATE TABLE "roles_permission" (
  "id" SERIAL PRIMARY KEY,
  "role_id" INTEGER NOT NULL,
  "permission_id" INTEGER NOT NULL,  
 CONSTRAINT fk_role_permissions_roles FOREIGN KEY (role_id) REFERENCES role(id)
 ON UPDATE CASCADE ON DELETE RESTRICT,
 CONSTRAINT fk_permission_roles_permissions FOREIGN KEY (permission_id) REFERENCES permission(id)
 ON UPDATE CASCADE ON DELETE RESTRICT
 );