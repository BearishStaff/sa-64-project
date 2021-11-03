BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "employees" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"tel"	text,
	"email"	text,
	"password"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "balances" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"type"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "statuses" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"detail"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "room_types" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"detail"	text,
	"price"	integer,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "rooms" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"location"	text,
	"roomnumber"	text,
	"type_id"	integer,
	"recorder_id"	integer,
	"status_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_room_types_rooms" FOREIGN KEY("type_id") REFERENCES "room_types"("id"),
	CONSTRAINT "fk_statuses_rooms" FOREIGN KEY("status_id") REFERENCES "statuses"("id"),
	CONSTRAINT "fk_employees_rooms" FOREIGN KEY("recorder_id") REFERENCES "employees"("id")
);
CREATE TABLE IF NOT EXISTS "payments" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"method"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "customers" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	"email"	text,
	"password"	text,
	"tel"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "reservations" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"people"	integer,
	"date_and_time"	datetime,
	"customer_id"	integer,
	"room_id"	integer,
	"payment_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_customers_reservations" FOREIGN KEY("customer_id") REFERENCES "customers"("id"),
	CONSTRAINT "fk_rooms_reservations" FOREIGN KEY("room_id") REFERENCES "rooms"("id"),
	CONSTRAINT "fk_payments_reservations" FOREIGN KEY("payment_id") REFERENCES "payments"("id")
);
CREATE TABLE IF NOT EXISTS "room_payments" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"payment_date"	datetime,
	"amount"	integer,
	"recorder_id"	integer,
	"reservation_id"	integer,
	"balance_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_reservations_room_payments" FOREIGN KEY("reservation_id") REFERENCES "reservations"("id"),
	CONSTRAINT "fk_balances_room_payments" FOREIGN KEY("balance_id") REFERENCES "balances"("id"),
	CONSTRAINT "fk_employees_room_payments" FOREIGN KEY("recorder_id") REFERENCES "employees"("id")
);
CREATE TABLE IF NOT EXISTS "check_ins" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"date_time"	datetime,
	"customer_id"	integer,
	"room_id"	integer,
	"room_payment_id"	integer,
	"employee_id"	integer,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_employees_records" FOREIGN KEY("employee_id") REFERENCES "employees"("id"),
	CONSTRAINT "fk_check_ins_room_payment" FOREIGN KEY("room_payment_id") REFERENCES "room_payments"("id"),
	CONSTRAINT "fk_rooms_records" FOREIGN KEY("room_id") REFERENCES "rooms"("id"),
	CONSTRAINT "fk_employees_check_ins" FOREIGN KEY("employee_id") REFERENCES "employees"("id"),
	CONSTRAINT "fk_customers_check_ins" FOREIGN KEY("customer_id") REFERENCES "customers"("id"),
	CONSTRAINT "fk_rooms_check_in" FOREIGN KEY("room_id") REFERENCES "rooms"("id"),
	CONSTRAINT "fk_customers_record" FOREIGN KEY("customer_id") REFERENCES "customers"("id")
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
	"condition"	text,
	PRIMARY KEY("id"),
	CONSTRAINT "fk_check_ins_check_out" FOREIGN KEY("check_in_id") REFERENCES "check_ins"("id"),
	CONSTRAINT "fk_employees_check_outs" FOREIGN KEY("employee_id") REFERENCES "employees"("id"),
	CONSTRAINT "fk_customers_check_outs" FOREIGN KEY("customer_id") REFERENCES "customers"("id")
);
CREATE TABLE IF NOT EXISTS "equipment" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"name"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "problems" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"value"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "urgencies" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"value"	text,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "repair_informations" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"datetime"	datetime,
	"check_in_id"	integer,
	"equipment_id"	integer,
	"problem_id"	integer,
	"urgency_id"	integer,
	CONSTRAINT "fk_problems_repair_informations" FOREIGN KEY("problem_id") REFERENCES "problems"("id"),
	CONSTRAINT "fk_equipment_repair_informations" FOREIGN KEY("equipment_id") REFERENCES "equipment"("id"),
	CONSTRAINT "fk_urgencies_repair_informations" FOREIGN KEY("urgency_id") REFERENCES "urgencies"("id"),
	CONSTRAINT "fk_check_ins_repair_informations" FOREIGN KEY("check_in_id") REFERENCES "check_ins"("id"),
	PRIMARY KEY("id")
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_employees_email" ON "employees" (
	"email"
);
CREATE INDEX IF NOT EXISTS "idx_employees_deleted_at" ON "employees" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_balances_deleted_at" ON "balances" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_statuses_deleted_at" ON "statuses" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_room_types_deleted_at" ON "room_types" (
	"deleted_at"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_rooms_roomnumber" ON "rooms" (
	"roomnumber"
);
CREATE INDEX IF NOT EXISTS "idx_rooms_deleted_at" ON "rooms" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_payments_deleted_at" ON "payments" (
	"deleted_at"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_customers_email" ON "customers" (
	"email"
);
CREATE INDEX IF NOT EXISTS "idx_customers_deleted_at" ON "customers" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_reservations_deleted_at" ON "reservations" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_room_payments_deleted_at" ON "room_payments" (
	"deleted_at"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_check_ins_room_payment_id" ON "check_ins" (
	"room_payment_id"
);
CREATE INDEX IF NOT EXISTS "idx_check_ins_deleted_at" ON "check_ins" (
	"deleted_at"
);
CREATE UNIQUE INDEX IF NOT EXISTS "idx_check_outs_check_in_id" ON "check_outs" (
	"check_in_id"
);
CREATE INDEX IF NOT EXISTS "idx_check_outs_deleted_at" ON "check_outs" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_equipment_deleted_at" ON "equipment" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_problems_deleted_at" ON "problems" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_urgencies_deleted_at" ON "urgencies" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_repair_informations_deleted_at" ON "repair_informations" (
	"deleted_at"
);
COMMIT;
