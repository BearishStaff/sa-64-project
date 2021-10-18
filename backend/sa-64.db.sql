BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "customers" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"tel"	text,
	"email"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "employees" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"tel"	text,
	"email"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "check_ins" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"date_time"	datetime,
	"customer_id"	integer,
	"employee_id"	integer,
	"room"	text,
	CONSTRAINT "fk_employees_check_ins" FOREIGN KEY("employee_id") REFERENCES "employees"("id"),
	CONSTRAINT "fk_check_ins_customer" FOREIGN KEY("customer_id") REFERENCES "customers"("id"),
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "check_outs" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"check_in_id"	integer,
	"customer_id"	integer,
	"employee_id"	integer,
	"check_out_time"	datetime,
	CONSTRAINT "fk_customers_check_outs" FOREIGN KEY("customer_id") REFERENCES "customers"("id"),
	CONSTRAINT "fk_employees_check_outs" FOREIGN KEY("employee_id") REFERENCES "employees"("id"),
	CONSTRAINT "fk_check_outs_check_in" FOREIGN KEY("check_in_id") REFERENCES "check_ins"("id"),
	PRIMARY KEY("id")
);
CREATE INDEX IF NOT EXISTS "idx_customers_deleted_at" ON "customers" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_employees_deleted_at" ON "employees" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_check_ins_deleted_at" ON "check_ins" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_check_outs_deleted_at" ON "check_outs" (
	"deleted_at"
);
COMMIT;
