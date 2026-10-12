# gestion_personas_crud

El API provee la gestión de proveedores para el sistema de gestión de personas Ágora con su información financiera y de contacto.

## Especificaciones Técnicas

### Tecnologías y Versiones

- [Golang](https://github.com/udistrital/lineamientos_oas/blob/master/instalacion_de_herramientas/golang.md)
- [Beego](https://github.com/udistrital/lineamientos_oas/blob/master/instalacion_de_herramientas/beego.md)
- [PostgreSQL](https://github.com/udistrital/lineamientos_oas/blob/master/instalacion_de_herramientas/postgres.md)
- [UtilsOAS](https://github.com/udistrital/lineamientos_oas/blob/master/generacion_de_apis/control_error_json_crud.md)

### Variables de Entorno

Configurar antes de ejecutar (ver también `.env.example` y `conf/app.conf`).

```shell
GESTION_PERSONAS_CRUD_API_NAME=[nombre de la API]
GESTION_PERSONAS_CRUD_SERVICE_NAME=[nombre del servicio]
GESTION_PERSONAS_CRUD_API_BASE_DIR=[directorio base de la API con respecto a $GOPATH/src]
GESTION_PERSONAS_CRUD_HTTP_PORT=[puerto de la API]
GESTION_PERSONAS_CRUD_RUN_MODE=[dev|prod]
GESTION_PERSONAS_CRUD_PGUSER=[usuario BD]
GESTION_PERSONAS_CRUD_PGPASS=[password BD]
GESTION_PERSONAS_CRUD_PGHOST=[host BD]
GESTION_PERSONAS_CRUD_PGPORT=[puerto BD]
GESTION_PERSONAS_CRUD_PGDB=[nombre BD]
GESTION_PERSONAS_CRUD_PGSCHEMA=[esquema BD]
```

### Ejecución del Proyecto

#### Clonar el repositorio

```shell
mkdir -p $GOPATH/src/github.com/udistrital
cd $GOPATH/src/github.com/udistrital
git clone https://github.com/udistrital/gestion_personas_crud.git
cd gestion_personas_crud
```

#### Ejecución para Desarrollo

Requiere una base de datos [PostgreSQL](https://github.com/udistrital/lineamientos_oas/blob/master/instalacion_de_herramientas/postgres.md) con el esquema y las tablas requeridas. Ejecute `database/agora_proveedores.sql` para prepararla y luego:

```shell
cp .env.example .env # Configurar las variables de entorno
set -a && source .env && set +a

go mod tidy
bee run
```

#### Ejecución con Docker

Compile el proyecto y la imagen:

```shell
go mod tidy
go build -o main
docker build -t gestion_personas_crud .
```

Para ejecutar el servicio en un contenedor. Requiere una base de datos [PostgreSQL](https://github.com/udistrital/lineamientos_oas/blob/master/instalacion_de_herramientas/postgres.md) con el esquema y las tablas requeridas. Ejecute `database/agora_proveedores.sql` para prepararla y luego:

```shell
cp .env.example .env # Configurar las variables de entorno
set -a && source .env && set +a
docker run --name gestion_personas_crud -p "$GESTION_PERSONAS_CRUD_HTTP_PORT:$GESTION_PERSONAS_CRUD_HTTP_PORT" --env-file .env gestion_personas_crud
```

### Ejecución con Docker Compose

Para consumo local durante el desarrollo de otros servicios, ejecute Docker Compose después de [compilar el proyecto y la imagen](#ejecución-con-docker). Incluye la base de datos y ejecuta `database/agora_proveedores.sql` al inicializarla:

```shell
# Creando red bridge para consumo local
docker network create \
  --driver bridge \
  --subnet 172.28.0.0/16 \
  --gateway 172.28.0.1 \
  back_end

cp .env.example .env # Configurar las variables de entorno
docker compose --env-file .env up
```

El script se ejecuta únicamente cuando se crea el volumen de PostgreSQL. Para repetir la inicialización, elimine el volumen `postgres_data` y levante Compose nuevamente.

### Ejecución Pruebas

Pruebas unitarias

```shell
# En Proceso
```

## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/gestion_personas_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/gestion_personas_crud/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/gestion_personas_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/gestion_personas_crud/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/gestion_personas_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/gestion_personas_crud/) |

## Modelo de Datos

[Modelo de Datos API CRUD Gestion Personas](https://github.com/udistrital/gestion_personas_crud/blob/develop/database/agora_proveedores.svg)

## Licencia

This file is part of gestion_personas_crud.

gestion_personas_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

gestion_personas_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with novedades_crud. If not, see [https://www.gnu.org/licenses/](https://www.gnu.org/licenses/).
