# gobocopy

Aplicacion de consola en Go inspirada en Robocopy para copias robustas, concurrentes y automatizables.

## Caracteristicas implementadas

- Copia robusta y tolerante a fallos con:
  - `-restartable` para reanudar archivos parciales (`.part`) cuando aplica.
  - `-backup` para intentar recuperacion en errores comunes de permisos.
- Copia multihilo con goroutines mediante `-threads`.
- Integrable en scripts y tareas programadas.
- Preservacion de metadatos:
  - Windows: copia tiempos, modo y ACL (via PowerShell `Get-Acl | Set-Acl`).
  - Linux/Unix: copia permisos y, opcionalmente, uid/gid.
- Filtros avanzados:
  - extension (`-exclude-ext`, `-include-ext`)
  - tamano (`-min-size`, `-max-size`)
  - antiguedad (`-min-age`, `-max-age`)
  - patron de nombre (`-exclude-pattern`)
- Sincronizacion espejo con `-mirror` eliminando en destino lo que no exista en origen.
- Copia segura:
  - escribe en temporal y renombra de forma atomica
  - verifica integridad con SHA-256 (sin compresion)
- Compresion opcional con `-compress` (guarda archivos destino con extension `.gz`).
- Logs y reportes:
  - log detallado en archivo (`-log-file`)
  - reporte JSON (`-report-json`)
- Recuperacion automatica de errores con `-retries` y `-retry-delay`.

## Uso rapido

```powershell
go run . -source "C:/origen" -dest "D:/destino"
```

## Ejemplos

```powershell
# espejo multihilo con exclusiones
go run . -source "C:/datos" -dest "D:/backup" -threads 24 -mirror -exclude-ext .tmp,.bak -exclude-pattern "*.cache"

# filtrar por tamano y antiguedad
go run . -source "C:/datos" -dest "D:/backup" -min-size 10M -max-size 2G -min-age 24h -max-age 720h

# reporte JSON para pipeline/scripts
go run . -source "C:/datos" -dest "D:/backup" -report-json "./out/report.json" -fail-on-error

# modo simulacion
go run . -source "C:/datos" -dest "D:/backup" -dry-run
```

## Banderas principales

- `-source`, `-dest`: origen y destino (obligatorias).
- `-threads`: workers concurrentes.
- `-mirror`: modo espejo.
- `-restartable`: reanuda archivos parciales.
- `-backup`: tolerancia de permisos/fallos.
- `-compress`: salida comprimida (`.gz`).
- `-preserve-security`: conserva ACL/metadatos de seguridad segun SO.
- `-preserve-owner`: conserva uid/gid en Unix.
- `-exclude-ext`, `-include-ext`, `-exclude-pattern`: filtros.
- `-min-size`, `-max-size`: filtro por tamano (`K`, `M`, `G`).
- `-min-age`, `-max-age`: filtro por antiguedad (`time.Duration`, ejemplo `24h`).
- `-retries`, `-retry-delay`: recuperacion automatica.
- `-log-file`: archivo de log.
- `-report-json`: salida reporte JSON.
- `-fail-on-error`: retorna codigo error si hubo fallos.
- `-dry-run`: no modifica archivos.

## Notas

- En Windows, para copiar ACL de forma completa se recomienda ejecutar en consola con permisos suficientes.
- Si usas `-compress`, los nombres en destino terminan en `.gz` y la comparacion espejo se hace sobre ese nombre final.
