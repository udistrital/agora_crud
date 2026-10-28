CREATE SCHEMA "proveedores";

CREATE TABLE "proveedores"."area_conocimiento_snies" (
  "id" serial PRIMARY KEY,
  "nombre" varchar(50) NOT NULL,
  "descripcion" varchar(250),
  "codigo_abreviacion" varchar(20),
  "activo" boolean NOT NULL DEFAULT true,
  "numero_orden" numeric(5,2)
);

CREATE TABLE "proveedores"."nucleo_basico_conocimiento_snies" (
  "id" serial PRIMARY KEY,
  "area_conocimiento_snies_id" integer NOT NULL,
  "nombre" varchar(50) NOT NULL,
  "descripcion" varchar(250),
  "codigo_abreviacion" varchar(20),
  "activo" boolean NOT NULL DEFAULT true,
  "numero_orden" numeric(5,2)
);

CREATE TABLE "proveedores"."proveedor" (
  "id" serial PRIMARY KEY,
  "tercero_id" integer NOT NULL,
  "tipo_registro_id" integer NOT NULL,
  "ciudad_origen_id" integer,
  "declara_renta" boolean NOT NULL DEFAULT false,
  "consentimiento_datos" boolean NOT NULL DEFAULT false,
  "descripcion_portafolio" varchar(250),
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer
);

CREATE TABLE "proveedores"."proveedor_natural" (
  "id" serial PRIMARY KEY,
  "experiencia_laboral_meses" integer NOT NULL DEFAULT 0,
  "experiencia_profesional_meses" integer NOT NULL DEFAULT 0,
  "interes_vivienda_afc" numeric(5,4) NOT NULL DEFAULT 0,
  "cargo" varchar(100) NOT NULL,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer
);

CREATE TABLE "proveedores"."titulo_proveedor_natural" (
  "id" serial PRIMARY KEY,
  "proveedor_natural_id" integer NOT NULL,
  "perfil_id" integer NOT NULL,
  "area_conocimiento_snies_id" integer NOT NULL,
  "nucleo_basico_conocimiento_snies_id" integer NOT NULL,
  "titulo_recibido" varchar(100) NOT NULL,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer
);

CREATE TABLE "proveedores"."proveedor_juridico" (
  "id" serial PRIMARY KEY,
  "nombre_comercial" varchar(50) NOT NULL,
  "matricula_mercantil" integer NOT NULL,
  "fecha_constitucion" date NOT NULL,
  "fecha_renovacion" date NOT NULL,
  "reporta_beneficios_finales" boolean NOT NULL DEFAULT false,
  "cotiza_bolsa" boolean NOT NULL DEFAULT false,
  "requiere_revisor_fiscal" boolean NOT NULL DEFAULT false,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer
);

CREATE TABLE "proveedores"."representacion" (
  "id" serial PRIMARY KEY,
  "proveedor_id" integer NOT NULL,
  "representante_id" integer NOT NULL,
  "tipo_representacion_id" integer NOT NULL,
  "cargo_id" integer NOT NULL,
  "tiene_limitacion_cuantia" boolean NOT NULL DEFAULT false,
  "descripcion_facultades" varchar(250),
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer
);

CREATE TABLE "proveedores"."contacto" (
  "id" serial PRIMARY KEY,
  "proveedor_id" integer NOT NULL,
  "nombre_contacto" varchar(100),
  "finalidad" varchar(50) NOT NULL,
  "tipo_canal" varchar(50) NOT NULL,
  "valor_especifico" varchar(100) NOT NULL,
  "extension" varchar(10),
  "es_principal" boolean NOT NULL DEFAULT false,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer
);

CREATE TABLE "proveedores"."cuenta_bancaria" (
  "id" serial PRIMARY KEY,
  "proveedor_id" integer NOT NULL,
  "entidad_bancaria_id" integer NOT NULL,
  "tipo_cuenta_id" integer NOT NULL,
  "numero_cuenta" varchar(50) NOT NULL,
  "nombre_titular" varchar(100) NOT NULL,
  "ciudad_apertura_id" integer,
  "capital_autorizado" numeric(20,7) NOT NULL,
  "es_principal" boolean DEFAULT false,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer
);

CREATE TABLE "proveedores"."actividad_economica_proveedor" (
  "id" serial PRIMARY KEY,
  "proveedor_id" integer NOT NULL,
  "actividad_economica_id" integer NOT NULL,
  "es_principal" boolean NOT NULL DEFAULT false,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer
);

CREATE TABLE "proveedores"."documento" (
  "id" serial PRIMARY KEY,
  "proveedor_id" integer NOT NULL,
  "tipo_documento_id" integer NOT NULL,
  "enlace" uuid,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer
);

CREATE INDEX "idx_area_conocimiento_snies_activo" ON "proveedores"."area_conocimiento_snies" ("activo");

CREATE INDEX "idx_nucleo_basico_conocimiento_snies_area_conocimiento_snies_id" ON "proveedores"."nucleo_basico_conocimiento_snies" ("area_conocimiento_snies_id");

CREATE INDEX "idx_nucleo_basico_conocimiento_snies_activo" ON "proveedores"."nucleo_basico_conocimiento_snies" ("activo");

CREATE UNIQUE INDEX "uq_proveedor_tercero_id" ON "proveedores"."proveedor" ("tercero_id");

CREATE INDEX "idx_proveedor_tipo_registro_id" ON "proveedores"."proveedor" ("tipo_registro_id");

CREATE INDEX "idx_proveedor_activo" ON "proveedores"."proveedor" ("activo");

CREATE INDEX "idx_proveedor_fecha_creacion" ON "proveedores"."proveedor" ("fecha_creacion");

CREATE INDEX "idx_proveedor_fecha_modificacion" ON "proveedores"."proveedor" ("fecha_modificacion");

CREATE INDEX "idx_proveedor_autor_modificacion_id" ON "proveedores"."proveedor" ("autor_modificacion_id");

CREATE INDEX "idx_proveedor_natural_activo" ON "proveedores"."proveedor_natural" ("activo");

CREATE INDEX "idx_proveedor_natural_fecha_creacion" ON "proveedores"."proveedor_natural" ("fecha_creacion");

CREATE INDEX "idx_proveedor_natural_fecha_modificacion" ON "proveedores"."proveedor_natural" ("fecha_modificacion");

CREATE INDEX "idx_proveedor_natural_autor_modificacion_id" ON "proveedores"."proveedor_natural" ("autor_modificacion_id");

CREATE INDEX "idx_titulo_proveedor_natural_proveedor_natural_id" ON "proveedores"."titulo_proveedor_natural" ("proveedor_natural_id");

CREATE INDEX "idx_titulo_proveedor_natural_perfil_id" ON "proveedores"."titulo_proveedor_natural" ("perfil_id");

CREATE INDEX "idx_titulo_proveedor_natural_area_conocimiento_snies_id" ON "proveedores"."titulo_proveedor_natural" ("area_conocimiento_snies_id");

CREATE INDEX "idx_titulo_proveedor_natural_nucleo_basico_conocimiento_snies_id" ON "proveedores"."titulo_proveedor_natural" ("nucleo_basico_conocimiento_snies_id");

CREATE INDEX "idx_titulo_proveedor_natural_activo" ON "proveedores"."titulo_proveedor_natural" ("activo");

CREATE INDEX "idx_titulo_proveedor_natural_fecha_creacion" ON "proveedores"."titulo_proveedor_natural" ("fecha_creacion");

CREATE INDEX "idx_titulo_proveedor_natural_fecha_modificacion" ON "proveedores"."titulo_proveedor_natural" ("fecha_modificacion");

CREATE INDEX "idx_titulo_proveedor_natural_autor_modificacion_id" ON "proveedores"."titulo_proveedor_natural" ("autor_modificacion_id");

CREATE UNIQUE INDEX "uq_proveedor_juridico_matricula_mercantil" ON "proveedores"."proveedor_juridico" ("matricula_mercantil");

CREATE INDEX "idx_proveedor_juridico_activo" ON "proveedores"."proveedor_juridico" ("activo");

CREATE INDEX "idx_proveedor_juridico_fecha_creacion" ON "proveedores"."proveedor_juridico" ("fecha_creacion");

CREATE INDEX "idx_proveedor_juridico_fecha_modificacion" ON "proveedores"."proveedor_juridico" ("fecha_modificacion");

CREATE INDEX "idx_proveedor_juridico_autor_modificacion_id" ON "proveedores"."proveedor_juridico" ("autor_modificacion_id");

CREATE INDEX "idx_representacion_activo" ON "proveedores"."representacion" ("activo");

CREATE INDEX "idx_representacion_fecha_creacion" ON "proveedores"."representacion" ("fecha_creacion");

CREATE INDEX "idx_representacion_fecha_modificacion" ON "proveedores"."representacion" ("fecha_modificacion");

CREATE INDEX "idx_representacion_autor_modificacion_id" ON "proveedores"."representacion" ("autor_modificacion_id");

CREATE INDEX "idx_contacto_activo" ON "proveedores"."contacto" ("activo");

CREATE INDEX "idx_contacto_fecha_creacion" ON "proveedores"."contacto" ("fecha_creacion");

CREATE INDEX "idx_contacto_fecha_modificacion" ON "proveedores"."contacto" ("fecha_modificacion");

CREATE INDEX "idx_contacto_autor_modificacion_id" ON "proveedores"."contacto" ("autor_modificacion_id");

CREATE INDEX "idx_cuenta_bancaria_activo" ON "proveedores"."cuenta_bancaria" ("activo");

CREATE INDEX "idx_cuenta_bancaria_fecha_creacion" ON "proveedores"."cuenta_bancaria" ("fecha_creacion");

CREATE INDEX "idx_cuenta_bancaria_fecha_modificacion" ON "proveedores"."cuenta_bancaria" ("fecha_modificacion");

CREATE INDEX "idx_cuenta_bancaria_autor_modificacion_id" ON "proveedores"."cuenta_bancaria" ("autor_modificacion_id");

CREATE INDEX "idx_actividad_economica_proveedor_activo" ON "proveedores"."actividad_economica_proveedor" ("activo");

CREATE INDEX "idx_actividad_economica_proveedor_fecha_creacion" ON "proveedores"."actividad_economica_proveedor" ("fecha_creacion");

CREATE INDEX "idx_actividad_economica_proveedor_fecha_modificacion" ON "proveedores"."actividad_economica_proveedor" ("fecha_modificacion");

CREATE INDEX "idx_actividad_economica_proveedor_autor_modificacion_id" ON "proveedores"."actividad_economica_proveedor" ("autor_modificacion_id");

CREATE INDEX "idx_documento_activo" ON "proveedores"."documento" ("activo");

CREATE INDEX "idx_documento_fecha_creacion" ON "proveedores"."documento" ("fecha_creacion");

CREATE INDEX "idx_documento_fecha_modificacion" ON "proveedores"."documento" ("fecha_modificacion");

CREATE INDEX "idx_documento_autor_modificacion_id" ON "proveedores"."documento" ("autor_modificacion_id");

COMMENT ON TABLE "proveedores"."area_conocimiento_snies" IS 'AGRONOMÍA, VETERINARIA Y AFINES; BELLAS ARTES; CIENCIAS DE LA EDUCACIÓN; etc.';

COMMENT ON TABLE "proveedores"."nucleo_basico_conocimiento_snies" IS 'Agronomía, Artes Plásticas, Educación, etc.';

COMMENT ON COLUMN "proveedores"."proveedor"."tercero_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."proveedor"."tipo_registro_id" IS 'Referencia externa a Terceros CRUD: info complementaria, tipo contribuyente';

COMMENT ON COLUMN "proveedores"."proveedor"."ciudad_origen_id" IS 'Referencia externa a Ubicaciones CRUD';

COMMENT ON COLUMN "proveedores"."proveedor"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."proveedor_natural"."cargo" IS 'Campo abierto para migración';

COMMENT ON COLUMN "proveedores"."proveedor_natural"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON TABLE "proveedores"."titulo_proveedor_natural" IS 'En sistema sólo se permitía un título y estaba en el proveedor natural. Trasladar registro único aquí en migración.';

COMMENT ON COLUMN "proveedores"."titulo_proveedor_natural"."perfil_id" IS 'Referencia externa a Terceros CRUD: info complementaria, grupo tipo de perfil';

COMMENT ON COLUMN "proveedores"."titulo_proveedor_natural"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."proveedor_juridico"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."representacion"."representante_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."representacion"."tipo_representacion_id" IS 'Referencia externa a Parametros CRUD';

COMMENT ON COLUMN "proveedores"."representacion"."cargo_id" IS 'Referencia externa a Parametros CRUD';

COMMENT ON COLUMN "proveedores"."representacion"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."contacto"."nombre_contacto" IS 'Juan Pérez';

COMMENT ON COLUMN "proveedores"."contacto"."finalidad" IS 'Representante legal, contacto comercial, notificaciones, contacto de emergencia, etc.';

COMMENT ON COLUMN "proveedores"."contacto"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."cuenta_bancaria"."entidad_bancaria_id" IS 'Referencia externa a Parametros CRUD';

COMMENT ON COLUMN "proveedores"."cuenta_bancaria"."tipo_cuenta_id" IS 'Referencia externa a Parametros CRUD';

COMMENT ON COLUMN "proveedores"."cuenta_bancaria"."ciudad_apertura_id" IS 'Referencia externa a Ubicaciones CRUD';

COMMENT ON COLUMN "proveedores"."cuenta_bancaria"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."actividad_economica_proveedor"."actividad_economica_id" IS 'Referencia externa a Parametros CRUD';

COMMENT ON COLUMN "proveedores"."actividad_economica_proveedor"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

COMMENT ON TABLE "proveedores"."documento" IS 'Se incluye certificado de registro.';

COMMENT ON COLUMN "proveedores"."documento"."tipo_documento_id" IS 'Referencia externa a Parametros CRUD';

COMMENT ON COLUMN "proveedores"."documento"."enlace" IS 'Archivo de registro único tributario (RUT)';

COMMENT ON COLUMN "proveedores"."documento"."autor_modificacion_id" IS 'Posible eferencia externa a Terceros CRUD';

ALTER TABLE "proveedores"."nucleo_basico_conocimiento_snies" ADD CONSTRAINT "fk_nucleo_basico_conocimiento_snies_area_conocimiento_snies" FOREIGN KEY ("area_conocimiento_snies_id") REFERENCES "proveedores"."area_conocimiento_snies" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."proveedor_natural" ADD CONSTRAINT "fk_proovedor_natural_proveedor" FOREIGN KEY ("id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."titulo_proveedor_natural" ADD CONSTRAINT "fk_titulo_proveedor_natural_proveedor_natural" FOREIGN KEY ("proveedor_natural_id") REFERENCES "proveedores"."proveedor_natural" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."titulo_proveedor_natural" ADD CONSTRAINT "fk_titulo_proveedor_natural_area_conocimiento_snies" FOREIGN KEY ("area_conocimiento_snies_id") REFERENCES "proveedores"."area_conocimiento_snies" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."titulo_proveedor_natural" ADD CONSTRAINT "fk_titulo_proveedor_natural_nucleo_basico_conocimiento_snies" FOREIGN KEY ("nucleo_basico_conocimiento_snies_id") REFERENCES "proveedores"."nucleo_basico_conocimiento_snies" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."proveedor_juridico" ADD CONSTRAINT "fk_proovedor_juridico_proveedor" FOREIGN KEY ("id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."representacion" ADD CONSTRAINT "fk_representacion_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor_juridico" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."contacto" ADD CONSTRAINT "fk_contacto_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."cuenta_bancaria" ADD CONSTRAINT "fk_cuenta_bancaria_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."actividad_economica_proveedor" ADD CONSTRAINT "fk_actividad_economica_proveedor_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."documento" ADD CONSTRAINT "fk_documento_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;
