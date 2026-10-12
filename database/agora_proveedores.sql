CREATE SCHEMA "proveedores";

CREATE TABLE "proveedores"."perfil" (
  "id" serial PRIMARY KEY,
  "nombre" varchar(50) NOT NULL,
  "descripcion" varchar(250),
  "codigo_abreviacion" varchar(20),
  "activo" boolean NOT NULL DEFAULT true,
  "numero_orden" numeric(5,2)
);

CREATE TABLE "proveedores"."prefijo_facturacion" (
  "id" serial PRIMARY KEY,
  "nombre" varchar(50) NOT NULL,
  "descripcion" varchar(250),
  "codigo_abreviacion" varchar(20),
  "activo" boolean NOT NULL DEFAULT true,
  "numero_orden" numeric(5,2)
);

CREATE TABLE "proveedores"."rango_facturacion" (
  "id" serial PRIMARY KEY,
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
  "descripcion_portafolio" varchar(250),
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer
);

CREATE TABLE "proveedores"."proveedor_natural" (
  "id" serial PRIMARY KEY,
  "perfil_declarado" integer NOT NULL,
  "experiencia_laboral_meses" integer NOT NULL DEFAULT 0,
  "experiencia_profesional_meses" integer NOT NULL DEFAULT 0,
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

CREATE TABLE "proveedores"."contacto" (
  "id" serial PRIMARY KEY,
  "proveedor_id" integer NOT NULL,
  "nombre_contacto" varchar(100),
  "finalidad" varchar(50) NOT NULL,
  "tipo_canal" varchar(50) NOT NULL,
  "valor_especifico" varchar(100) NOT NULL,
  "extension" varchar(10),
  "info_complementaria_tercero_id" integer,
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

CREATE TABLE "proveedores"."perfil_financiero" (
  "id" serial PRIMARY KEY,
  "proveedor_id" integer NOT NULL,
  "gran_contribuyente" boolean NOT NULL DEFAULT false,
  "autorretenedor" boolean NOT NULL DEFAULT false,
  "exencion_ica" boolean NOT NULL DEFAULT false,
  "obligado_facturar_electronicamente" boolean NOT NULL DEFAULT false,
  "prefijo_facturacion_id" integer NOT NULL,
  "rango_facturacion_id" integer NOT NULL,
  "resolucion_facturacion" varchar(50),
  "opera_moneda_extranjera" boolean NOT NULL DEFAULT false,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer
);

CREATE TABLE "proveedores"."responsabilidad_fiscal_perfil" (
  "id" serial PRIMARY KEY,
  "perfil_financiero_id" integer NOT NULL,
  "responsabilidad_fiscal_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer
);

CREATE TABLE "proveedores"."informacion_financiera" (
  "id" serial PRIMARY KEY,
  "proveedor_id" integer NOT NULL,
  "fecha_corte" date NOT NULL DEFAULT (date_trunc('year', now()) + interval '1 year - 1 day'),
  "moneda_id" integer,
  "capital_autorizado" numeric(20,7) NOT NULL,
  "activos_totales" numeric(20,7),
  "pasivos_totales" numeric(20,7),
  "indice_liquidez" numeric(5,4) NOT NULL,
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

CREATE TABLE "proveedores"."declaracion" (
  "id" serial PRIMARY KEY,
  "tipo_declaracion_id" integer NOT NULL,
  "texto" varchar(250) NOT NULL,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer
);

CREATE TABLE "proveedores"."tipo_declaracion_proveedor" (
  "id" serial PRIMARY KEY,
  "proveedor_id" integer NOT NULL,
  "declaracion_id" integer NOT NULL,
  "activo" boolean NOT NULL DEFAULT true,
  "fecha_creacion" timestamp NOT NULL DEFAULT (now()),
  "fecha_modificacion" timestamp,
  "autor_modificacion_id" integer
);

CREATE INDEX "idx_perfil_activo" ON "proveedores"."perfil" ("activo");

CREATE INDEX "idx_prefijo_facturacion_activo" ON "proveedores"."prefijo_facturacion" ("activo");

CREATE INDEX "idx_rango_facturacion_activo" ON "proveedores"."rango_facturacion" ("activo");

CREATE UNIQUE INDEX "uq_proveedor_tercero_id" ON "proveedores"."proveedor" ("tercero_id");

CREATE INDEX "idx_proveedor_tipo_registro_id" ON "proveedores"."proveedor" ("tipo_registro_id");

CREATE INDEX "idx_proveedor_activo" ON "proveedores"."proveedor" ("activo");

CREATE INDEX "idx_proveedor_fecha_creacion" ON "proveedores"."proveedor" ("fecha_creacion");

CREATE INDEX "idx_proveedor_fecha_modificacion" ON "proveedores"."proveedor" ("fecha_modificacion");

CREATE INDEX "idx_proveedor_autor_modificacion_id" ON "proveedores"."proveedor" ("autor_modificacion_id");

CREATE INDEX "idx_proveedor_natural_perfil_declarado" ON "proveedores"."proveedor_natural" ("perfil_declarado");

CREATE INDEX "idx_proveedor_natural_activo" ON "proveedores"."proveedor_natural" ("activo");

CREATE INDEX "idx_proveedor_natural_fecha_creacion" ON "proveedores"."proveedor_natural" ("fecha_creacion");

CREATE INDEX "idx_proveedor_natural_fecha_modificacion" ON "proveedores"."proveedor_natural" ("fecha_modificacion");

CREATE INDEX "idx_proveedor_natural_autor_modificacion_id" ON "proveedores"."proveedor_natural" ("autor_modificacion_id");

CREATE UNIQUE INDEX "uq_proveedor_juridico_matricula_mercantil" ON "proveedores"."proveedor_juridico" ("matricula_mercantil");

CREATE INDEX "idx_proveedor_juridico_activo" ON "proveedores"."proveedor_juridico" ("activo");

CREATE INDEX "idx_proveedor_juridico_fecha_creacion" ON "proveedores"."proveedor_juridico" ("fecha_creacion");

CREATE INDEX "idx_proveedor_juridico_fecha_modificacion" ON "proveedores"."proveedor_juridico" ("fecha_modificacion");

CREATE INDEX "idx_proveedor_juridico_autor_modificacion_id" ON "proveedores"."proveedor_juridico" ("autor_modificacion_id");

CREATE INDEX "idx_contacto_info_complementaria_tercero_id" ON "proveedores"."contacto" ("info_complementaria_tercero_id");

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

CREATE INDEX "idx_perfil_financiero_activo" ON "proveedores"."perfil_financiero" ("activo");

CREATE INDEX "idx_perfil_financiero_fecha_creacion" ON "proveedores"."perfil_financiero" ("fecha_creacion");

CREATE INDEX "idx_perfil_financiero_fecha_modificacion" ON "proveedores"."perfil_financiero" ("fecha_modificacion");

CREATE INDEX "idx_perfil_financiero_autor_modificacion_id" ON "proveedores"."perfil_financiero" ("autor_modificacion_id");

CREATE INDEX "idx_responsabilidad_fiscal_perfil_activo" ON "proveedores"."responsabilidad_fiscal_perfil" ("activo");

CREATE INDEX "idx_responsabilidad_fiscal_perfil_fecha_creacion" ON "proveedores"."responsabilidad_fiscal_perfil" ("fecha_creacion");

CREATE INDEX "idx_responsabilidad_fiscal_perfil_fecha_modificacion" ON "proveedores"."responsabilidad_fiscal_perfil" ("fecha_modificacion");

CREATE INDEX "idx_responsabilidad_fiscal_perfil_autor_modificacion_id" ON "proveedores"."responsabilidad_fiscal_perfil" ("autor_modificacion_id");

CREATE INDEX "idx_informacion_financiera_activo" ON "proveedores"."informacion_financiera" ("activo");

CREATE INDEX "idx_informacion_financiera_fecha_creacion" ON "proveedores"."informacion_financiera" ("fecha_creacion");

CREATE INDEX "idx_informacion_financiera_fecha_modificacion" ON "proveedores"."informacion_financiera" ("fecha_modificacion");

CREATE INDEX "idx_informacion_financiera_autor_modificacion_id" ON "proveedores"."informacion_financiera" ("autor_modificacion_id");

CREATE INDEX "idx_representacion_activo" ON "proveedores"."representacion" ("activo");

CREATE INDEX "idx_representacion_fecha_creacion" ON "proveedores"."representacion" ("fecha_creacion");

CREATE INDEX "idx_representacion_fecha_modificacion" ON "proveedores"."representacion" ("fecha_modificacion");

CREATE INDEX "idx_representacion_autor_modificacion_id" ON "proveedores"."representacion" ("autor_modificacion_id");

CREATE INDEX "idx_documento_activo" ON "proveedores"."documento" ("activo");

CREATE INDEX "idx_documento_fecha_creacion" ON "proveedores"."documento" ("fecha_creacion");

CREATE INDEX "idx_documento_fecha_modificacion" ON "proveedores"."documento" ("fecha_modificacion");

CREATE INDEX "idx_documento_autor_modificacion_id" ON "proveedores"."documento" ("autor_modificacion_id");

CREATE INDEX "idx_declaracion_activo" ON "proveedores"."declaracion" ("activo");

CREATE INDEX "idx_declaracion_fecha_creacion" ON "proveedores"."declaracion" ("fecha_creacion");

CREATE INDEX "idx_declaracion_fecha_modificacion" ON "proveedores"."declaracion" ("fecha_modificacion");

CREATE INDEX "idx_declaracion_autor_modificacion_id" ON "proveedores"."declaracion" ("autor_modificacion_id");

CREATE INDEX "idx_tipo_declaracion_proveedor_activo" ON "proveedores"."tipo_declaracion_proveedor" ("activo");

CREATE INDEX "idx_tipo_declaracion_proveedor_fecha_creacion" ON "proveedores"."tipo_declaracion_proveedor" ("fecha_creacion");

CREATE INDEX "idx_tipo_declaracion_proveedor_fecha_modificacion" ON "proveedores"."tipo_declaracion_proveedor" ("fecha_modificacion");

CREATE INDEX "idx_tipo_declaracion_proveedor_autor_modificacion_id" ON "proveedores"."tipo_declaracion_proveedor" ("autor_modificacion_id");

COMMENT ON TABLE "proveedores"."perfil" IS 'Contratista, docente, investigador, etc. Posiblemente Terceros GrupoInfoComplementariaId 5: Tipo Perfil';

COMMENT ON TABLE "proveedores"."prefijo_facturacion" IS 'SATT, SATC, etc.';

COMMENT ON TABLE "proveedores"."rango_facturacion" IS '0 a 100, 101 a 500, etc.';

COMMENT ON COLUMN "proveedores"."proveedor"."tercero_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."proveedor"."tipo_registro_id" IS 'Referencia externa a Terceros CRUD: info complementaria, tipo contribuyente';

COMMENT ON COLUMN "proveedores"."proveedor"."autor_modificacion_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."proveedor_natural"."autor_modificacion_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."proveedor_juridico"."autor_modificacion_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."contacto"."nombre_contacto" IS 'Juan Pérez';

COMMENT ON COLUMN "proveedores"."contacto"."finalidad" IS 'Representante legal, contacto comercial, notificaciones, contacto de emergencia, etc.';

COMMENT ON COLUMN "proveedores"."contacto"."info_complementaria_tercero_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."contacto"."autor_modificacion_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."cuenta_bancaria"."entidad_bancaria_id" IS 'Referencia externa a Parametros CRUD';

COMMENT ON COLUMN "proveedores"."cuenta_bancaria"."tipo_cuenta_id" IS 'Referencia externa a Parametros CRUD';

COMMENT ON COLUMN "proveedores"."cuenta_bancaria"."ciudad_apertura_id" IS 'Referencia externa a Ubicaciones CRUD';

COMMENT ON COLUMN "proveedores"."cuenta_bancaria"."autor_modificacion_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."actividad_economica_proveedor"."actividad_economica_id" IS 'Referencia externa a Parametros CRUD';

COMMENT ON COLUMN "proveedores"."actividad_economica_proveedor"."autor_modificacion_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."perfil_financiero"."autor_modificacion_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."responsabilidad_fiscal_perfil"."responsabilidad_fiscal_id" IS 'Referencia externa a Parametros CRUD';

COMMENT ON COLUMN "proveedores"."responsabilidad_fiscal_perfil"."autor_modificacion_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."informacion_financiera"."fecha_corte" IS 'Por defecto se toma el último día del año en curso';

COMMENT ON COLUMN "proveedores"."informacion_financiera"."moneda_id" IS 'Referencia externa a Parametros CRUD';

COMMENT ON COLUMN "proveedores"."informacion_financiera"."autor_modificacion_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."representacion"."representante_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON COLUMN "proveedores"."representacion"."tipo_representacion_id" IS 'Referencia externa a Parametros CRUD';

COMMENT ON COLUMN "proveedores"."representacion"."cargo_id" IS 'Referencia externa a Parametros CRUD';

COMMENT ON COLUMN "proveedores"."representacion"."autor_modificacion_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON TABLE "proveedores"."documento" IS 'Se incluye certificado de registro.';

COMMENT ON COLUMN "proveedores"."documento"."tipo_documento_id" IS 'Referencia externa a Parametros CRUD';

COMMENT ON COLUMN "proveedores"."documento"."enlace" IS 'Archivo de registro único tributario (RUT)';

COMMENT ON COLUMN "proveedores"."documento"."autor_modificacion_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON TABLE "proveedores"."declaracion" IS 'Versionado de las declaraciones que puede firmar el proveedor.';

COMMENT ON COLUMN "proveedores"."declaracion"."tipo_declaracion_id" IS 'Referencia externa a Parametros CRUD';

COMMENT ON COLUMN "proveedores"."declaracion"."texto" IS 'Declaración de renta, declaración de IVA, declaración de retención en la fuente, etc.';

COMMENT ON COLUMN "proveedores"."declaracion"."autor_modificacion_id" IS 'Referencia externa a Terceros CRUD';

COMMENT ON TABLE "proveedores"."tipo_declaracion_proveedor" IS 'Tabla de rompimiento entre un texto de declaración firmada y el proveedor que la firmó.';

COMMENT ON COLUMN "proveedores"."tipo_declaracion_proveedor"."autor_modificacion_id" IS 'Referencia externa a Terceros CRUD';

ALTER TABLE "proveedores"."proveedor_natural" ADD CONSTRAINT "fk_proovedor_natural_proveedor" FOREIGN KEY ("id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."proveedor_natural" ADD CONSTRAINT "fk_proovedor_natural_perfil" FOREIGN KEY ("perfil_declarado") REFERENCES "proveedores"."perfil" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."proveedor_juridico" ADD CONSTRAINT "fk_proovedor_juridico_proveedor" FOREIGN KEY ("id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."contacto" ADD CONSTRAINT "fk_contacto_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."cuenta_bancaria" ADD CONSTRAINT "fk_cuenta_bancaria_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."actividad_economica_proveedor" ADD CONSTRAINT "fk_actividad_economica_proveedor_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."perfil_financiero" ADD CONSTRAINT "fk_perfil_financiero_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor_juridico" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."perfil_financiero" ADD CONSTRAINT "fk_perfil_financiero_juridico_prefijo_facturacion" FOREIGN KEY ("prefijo_facturacion_id") REFERENCES "proveedores"."prefijo_facturacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."perfil_financiero" ADD CONSTRAINT "fk_perfil_financiero_juridico_rango_facturacion" FOREIGN KEY ("rango_facturacion_id") REFERENCES "proveedores"."rango_facturacion" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."responsabilidad_fiscal_perfil" ADD CONSTRAINT "fk_responsabilidad_fiscal_perfil_perfil_financiero" FOREIGN KEY ("perfil_financiero_id") REFERENCES "proveedores"."perfil_financiero" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."informacion_financiera" ADD CONSTRAINT "fk_informacion_financiera_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."representacion" ADD CONSTRAINT "fk_representacion_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor_juridico" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."documento" ADD CONSTRAINT "fk_documento_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."tipo_declaracion_proveedor" ADD CONSTRAINT "fk_tipo_declaracion_proveedor_proveedor" FOREIGN KEY ("proveedor_id") REFERENCES "proveedores"."proveedor" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "proveedores"."tipo_declaracion_proveedor" ADD CONSTRAINT "fk_tipo_declaracion_proveedor_declaracion" FOREIGN KEY ("declaracion_id") REFERENCES "proveedores"."declaracion" ("id") DEFERRABLE INITIALLY IMMEDIATE;
