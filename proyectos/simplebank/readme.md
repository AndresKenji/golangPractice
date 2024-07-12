# database migrations

Cuando se trabaja con bases de datos es comun tener que realizar migraciones entre diferentes **schemas** para adaptarse a nuevos requerimientos o logicas del negocio.

## Migrar bases de datos con Golang

Haciendo uso de la libreria [Migrate][golang-migrate], podemo usar las opciones de la linea de comandos siguiendo las instrucciones de [instalación](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

Para usar migrate podemos validar la documentación oficial del proyecto en la cual encontramos lo siguiente

```txt
$ migrate -help
Usage: migrate OPTIONS COMMAND [arg...]
       migrate [ -version | -help ]

Options:
  -source          Location of the migrations (driver://url)
  -path            Shorthand for -source=file://path
  -database        Run migrations against this database (driver://url)
  -prefetch N      Number of migrations to load in advance before executing (default 10)
  -lock-timeout N  Allow N seconds to acquire database lock (default 15)
  -verbose         Print verbose logging
  -version         Print version
  -help            Print usage

Commands:
  create [-ext E] [-dir D] [-seq] [-digits N] [-format] NAME
               Create a set of timestamped up/down migrations titled NAME, in directory D with extension E.
               Use -seq option to generate sequential up/down migrations with N digits.
               Use -format option to specify a Go time format string.
  goto V       Migrate to version V
  up [N]       Apply all or N up migrations
  down [N]     Apply all or N down migrations
  drop         Drop everything inside database
  force V      Set version V but don't run migration (ignores dirty state)
  version      Print current migration version
```

Para empezar creamos la carpeta db/migration y vamos a ejecutar el comando `migrate create -ext sql -dir db/migration -seq init_schema` con el cual indicamos lo siguiente:
- create: Crear un nuevos archivos de migración
- -ext: la extensión de los archivos en este caso **sql**
- -dir: el directorio en el cual se van a crear los archivos para este proyecto será **db/migration**
- -seq: indica que se lleve una secuencia para los archivos de migración
- finalmente indicamos el nombre de la migración el cual llamamos **init_schema**

Al correr el comando vemos que crea dos archivos con los siguientes nombres:
- 000001_init_schema.down.sql
- 000001_init_schema.up.sql

Al trabajar con migraciones de bases de datos se usa como buena practica, los scripts de **up** se usan para aplicar o subir cambios al nuevo schema y el **down** se usa para revertir dichos cambios.

Al usar el comando `migrate up` correra los archivos de migración en orden segun su prefijo de secuencia y lo contrario ocurrira con su contraparte `migrate down` el cual corre secuencialmente los archivos para revertir cambios en la nueva base de datos o el nuevo schema

Ahora procedemos con nuestra primera migración para lo cual necesitamos que nuestra db de prueba se encuentre en ejecución revisar el [Makefile][./Makefile]
Para ejecutar la migración lo realizamos con el comando `migrate -path db/migration -database "postgresql://root:admin1234@localhost:5432/simple_bank?sslmode=disable" -verbose up` ahora al revisar la db podemos observar los cambios junto con una nueva tabla schema migrations la cual muestra el resultado de las migraciones que realicemos 


# Generar CRUD desde codigo SQL

Para trabajar con sql desde go existen varias alternativas y así realizar operaciones tipo CRUD.

- Libreria estandar: Se puede recurrir al uso del paquete [database][pkg-database] de la librería estandar el cual brinda herramientas para conexión y para realizar las operaciones dentro de la base de datos, entre sus ventajas tenemos su gran velocidad al realizar las consultas, sin embargo, es necesario escribir en código cada estructura o entidad de la base de datos para así realizar un mapeo manual de estas a variables por lo cual es bastante propenso a errores.
- [GORM][gorm]: Es una libreria de go la cual brinda un enfoque basado en ORM (object relational mapping) brinda bastantes herramientas para realizar las consultas y las operaciones dentro de la base de datos y a nivel de variables, sus principales desventajas son que debes aprender a escribir las consultas de la manera en la que la determina la propia libreria por lo que puede resultar confuso, su velocidad la cual esta muy por debajo de la libreria estandar.
- [SQLX][sqlx]: Es una libreria que brinda un set de herramientas para la libreria estandar agregando la capacidad de crear las estructuras necesaria a través del codigo sql, puede escanear los resultados directo en variables por lo que es mas facil y rápido trabajar con este tipo de herramientas, su principal desventaja es que similar que la librería estandar se debe escribir código muy extenso por lo que puede resultal en posibles errores.
- [SQLC][sqlc]: Es un compililador el cual toma código slql y lo convierte en codigo de go para realizar operaciones con las tablas o las entidades definidas en el código.








[golang-migrate]:https://github.com/golang-migrate/migrate
[pkg-database]:https://pkg.go.dev/database/sql
[gorm]:https://gorm.io/index.html
[sqlx]:https://pkg.go.dev/github.com/jmoiron/sqlx#section-readme
[sqlc]:https://pkg.go.dev/github.com/kyleconroy/sqlc