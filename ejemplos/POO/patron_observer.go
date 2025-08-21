package main

import (
	"fmt"
	"time"
)

// ============= PATRÓN OBSERVER =============

// Interfaz Observer
type Observer interface {
	Actualizar(mensaje string, datos interface{})
	ID() string
}

// Interfaz Subject (Observable)
type Subject interface {
	Suscribir(observer Observer)
	Desuscribir(observer Observer)
	Notificar(mensaje string, datos interface{})
}

// ============= SISTEMA DE NOTICIAS =============

// Implementación de Subject - Centro de Noticias
type CentroNoticias struct {
	observers []Observer
	noticias  []Noticia
}

type Noticia struct {
	ID        int
	Titulo    string
	Contenido string
	Categoria string
	Fecha     time.Time
}

func NuevoCentroNoticias() *CentroNoticias {
	return &CentroNoticias{
		observers: make([]Observer, 0),
		noticias:  make([]Noticia, 0),
	}
}

func (cn *CentroNoticias) Suscribir(observer Observer) {
	cn.observers = append(cn.observers, observer)
	fmt.Printf("    %s se suscribió al centro de noticias\n", observer.ID())
}

func (cn *CentroNoticias) Desuscribir(observer Observer) {
	for i, obs := range cn.observers {
		if obs.ID() == observer.ID() {
			cn.observers = append(cn.observers[:i], cn.observers[i+1:]...)
			fmt.Printf("    %s se desuscribió del centro de noticias\n", observer.ID())
			return
		}
	}
}

func (cn *CentroNoticias) Notificar(mensaje string, datos interface{}) {
	fmt.Printf("    📢 Notificando a %d observers: %s\n", len(cn.observers), mensaje)
	for _, observer := range cn.observers {
		observer.Actualizar(mensaje, datos)
	}
}

func (cn *CentroNoticias) PublicarNoticia(titulo, contenido, categoria string) {
	noticia := Noticia{
		ID:        len(cn.noticias) + 1,
		Titulo:    titulo,
		Contenido: contenido,
		Categoria: categoria,
		Fecha:     time.Now(),
	}

	cn.noticias = append(cn.noticias, noticia)
	cn.Notificar(fmt.Sprintf("Nueva noticia de %s", categoria), noticia)
}

func (cn *CentroNoticias) NumeroObservers() int {
	return len(cn.observers)
}

// Implementación de Observer - Suscriptor de Email
type SuscriptorEmail struct {
	id             string
	email          string
	intereses      []string
	notificaciones []string
}

func NuevoSuscriptorEmail(id, email string, intereses []string) *SuscriptorEmail {
	return &SuscriptorEmail{
		id:             id,
		email:          email,
		intereses:      intereses,
		notificaciones: make([]string, 0),
	}
}

func (se *SuscriptorEmail) ID() string {
	return se.id
}

func (se *SuscriptorEmail) Actualizar(mensaje string, datos interface{}) {
	if noticia, ok := datos.(Noticia); ok {
		// Verificar si está interesado en esta categoría
		for _, interes := range se.intereses {
			if interes == noticia.Categoria || interes == "todas" {
				notificacion := fmt.Sprintf("📧 Email a %s: %s - %s",
					se.email, noticia.Titulo, noticia.Categoria)
				se.notificaciones = append(se.notificaciones, notificacion)
				fmt.Printf("      %s\n", notificacion)
				return
			}
		}
	}
}

func (se *SuscriptorEmail) MostrarNotificaciones() {
	fmt.Printf("    Notificaciones de %s:\n", se.id)
	for i, notif := range se.notificaciones {
		fmt.Printf("      [%d] %s\n", i+1, notif)
	}
}

// Implementación de Observer - App Móvil
type AppMovil struct {
	id             string
	usuario        string
	configuracion  ConfiguracionApp
	notificaciones []NotificacionPush
}

type ConfiguracionApp struct {
	NotificacionesPush bool
	Sonido             bool
	Vibrar             bool
	CategoriasFiltro   []string
}

type NotificacionPush struct {
	Titulo  string
	Mensaje string
	Fecha   time.Time
	Leida   bool
}

func NuevaAppMovil(id, usuario string, config ConfiguracionApp) *AppMovil {
	return &AppMovil{
		id:             id,
		usuario:        usuario,
		configuracion:  config,
		notificaciones: make([]NotificacionPush, 0),
	}
}

func (am *AppMovil) ID() string {
	return am.id
}

func (am *AppMovil) Actualizar(mensaje string, datos interface{}) {
	if !am.configuracion.NotificacionesPush {
		return // Notificaciones deshabilitadas
	}

	if noticia, ok := datos.(Noticia); ok {
		// Verificar filtros de categoría
		if len(am.configuracion.CategoriasFiltro) > 0 {
			categoriaPermitida := false
			for _, categoria := range am.configuracion.CategoriasFiltro {
				if categoria == noticia.Categoria {
					categoriaPermitida = true
					break
				}
			}
			if !categoriaPermitida {
				return
			}
		}

		push := NotificacionPush{
			Titulo:  "Nueva Noticia",
			Mensaje: noticia.Titulo,
			Fecha:   time.Now(),
			Leida:   false,
		}

		am.notificaciones = append(am.notificaciones, push)

		efectos := ""
		if am.configuracion.Sonido {
			efectos += "🔊"
		}
		if am.configuracion.Vibrar {
			efectos += "📳"
		}

		fmt.Printf("      📱 Push a %s (%s): %s %s\n",
			am.usuario, am.id, push.Mensaje, efectos)
	}
}

func (am *AppMovil) MarcarComoLeida(indice int) {
	if indice >= 0 && indice < len(am.notificaciones) {
		am.notificaciones[indice].Leida = true
	}
}

func (am *AppMovil) NotificacionesNoLeidas() int {
	count := 0
	for _, notif := range am.notificaciones {
		if !notif.Leida {
			count++
		}
	}
	return count
}

// Implementación de Observer - Dashboard Web
type DashboardWeb struct {
	id           string
	admin        string
	estadisticas EstadisticasDashboard
}

type EstadisticasDashboard struct {
	TotalNoticias        int
	NoticiasPorCategoria map[string]int
	UltimaActualizacion  time.Time
}

func NuevoDashboardWeb(id, admin string) *DashboardWeb {
	return &DashboardWeb{
		id:    id,
		admin: admin,
		estadisticas: EstadisticasDashboard{
			NoticiasPorCategoria: make(map[string]int),
		},
	}
}

func (dw *DashboardWeb) ID() string {
	return dw.id
}

func (dw *DashboardWeb) Actualizar(mensaje string, datos interface{}) {
	if noticia, ok := datos.(Noticia); ok {
		dw.estadisticas.TotalNoticias++
		dw.estadisticas.NoticiasPorCategoria[noticia.Categoria]++
		dw.estadisticas.UltimaActualizacion = time.Now()

		fmt.Printf("      📊 Dashboard actualizado - Total: %d noticias\n",
			dw.estadisticas.TotalNoticias)
	}
}

func (dw *DashboardWeb) MostrarEstadisticas() {
	fmt.Printf("    Estadísticas del Dashboard (%s):\n", dw.admin)
	fmt.Printf("      Total de noticias: %d\n", dw.estadisticas.TotalNoticias)
	fmt.Printf("      Por categoría:\n")
	for categoria, count := range dw.estadisticas.NoticiasPorCategoria {
		fmt.Printf("        %s: %d\n", categoria, count)
	}
	fmt.Printf("      Última actualización: %s\n",
		dw.estadisticas.UltimaActualizacion.Format("15:04:05"))
}

// ============= SISTEMA DE MONITOREO DE TEMPERATURA =============

// Subject para temperatura
type SensorTemperatura struct {
	observers      []Observer
	temperatura    float64
	ubicacion      string
	alarmasActivas bool
}

func NuevoSensorTemperatura(ubicacion string) *SensorTemperatura {
	return &SensorTemperatura{
		observers:      make([]Observer, 0),
		ubicacion:      ubicacion,
		alarmasActivas: true,
	}
}

func (st *SensorTemperatura) Suscribir(observer Observer) {
	st.observers = append(st.observers, observer)
}

func (st *SensorTemperatura) Desuscribir(observer Observer) {
	for i, obs := range st.observers {
		if obs.ID() == observer.ID() {
			st.observers = append(st.observers[:i], st.observers[i+1:]...)
			return
		}
	}
}

func (st *SensorTemperatura) Notificar(mensaje string, datos interface{}) {
	if st.alarmasActivas {
		for _, observer := range st.observers {
			observer.Actualizar(mensaje, datos)
		}
	}
}

func (st *SensorTemperatura) ActualizarTemperatura(nuevaTemp float64) {
	tempAnterior := st.temperatura
	st.temperatura = nuevaTemp

	datosTemperatura := map[string]interface{}{
		"ubicacion":           st.ubicacion,
		"temperatura":         nuevaTemp,
		"temperaturaAnterior": tempAnterior,
		"fecha":               time.Now(),
	}

	mensaje := fmt.Sprintf("Temperatura en %s: %.1f°C", st.ubicacion, nuevaTemp)

	// Verificar si hay alertas
	if nuevaTemp > 35 {
		mensaje += " ⚠️ CALOR EXTREMO"
	} else if nuevaTemp < 0 {
		mensaje += " ❄️ TEMPERATURA BAJO CERO"
	}

	st.Notificar(mensaje, datosTemperatura)
}

func (st *SensorTemperatura) TemperaturaActual() float64 {
	return st.temperatura
}

// Observer para alertas de temperatura
type AlertaTemperatura struct {
	id              string
	umbralMax       float64
	umbralMin       float64
	alertasEnviadas []string
}

func NuevaAlertaTemperatura(id string, umbralMin, umbralMax float64) *AlertaTemperatura {
	return &AlertaTemperatura{
		id:              id,
		umbralMax:       umbralMax,
		umbralMin:       umbralMin,
		alertasEnviadas: make([]string, 0),
	}
}

func (at *AlertaTemperatura) ID() string {
	return at.id
}

func (at *AlertaTemperatura) Actualizar(mensaje string, datos interface{}) {
	if datosMap, ok := datos.(map[string]interface{}); ok {
		if temp, ok := datosMap["temperatura"].(float64); ok {
			ubicacion := datosMap["ubicacion"].(string)

			var alerta string
			if temp > at.umbralMax {
				alerta = fmt.Sprintf("🚨 ALERTA: Temperatura alta en %s: %.1f°C (máx: %.1f°C)",
					ubicacion, temp, at.umbralMax)
			} else if temp < at.umbralMin {
				alerta = fmt.Sprintf("🚨 ALERTA: Temperatura baja en %s: %.1f°C (mín: %.1f°C)",
					ubicacion, temp, at.umbralMin)
			}

			if alerta != "" {
				at.alertasEnviadas = append(at.alertasEnviadas, alerta)
				fmt.Printf("      %s\n", alerta)
			}
		}
	}
}

func (at *AlertaTemperatura) NumeroAlertas() int {
	return len(at.alertasEnviadas)
}

// Logger para temperatura
type LoggerTemperatura struct {
	id        string
	registros []RegistroTemperatura
}

type RegistroTemperatura struct {
	Ubicacion   string
	Temperatura float64
	Fecha       time.Time
}

func NuevoLoggerTemperatura(id string) *LoggerTemperatura {
	return &LoggerTemperatura{
		id:        id,
		registros: make([]RegistroTemperatura, 0),
	}
}

func (lt *LoggerTemperatura) ID() string {
	return lt.id
}

func (lt *LoggerTemperatura) Actualizar(mensaje string, datos interface{}) {
	if datosMap, ok := datos.(map[string]interface{}); ok {
		registro := RegistroTemperatura{
			Ubicacion:   datosMap["ubicacion"].(string),
			Temperatura: datosMap["temperatura"].(float64),
			Fecha:       datosMap["fecha"].(time.Time),
		}

		lt.registros = append(lt.registros, registro)
		fmt.Printf("      📝 Log: %s - %.1f°C registrado\n",
			registro.Ubicacion, registro.Temperatura)
	}
}

func (lt *LoggerTemperatura) UltimosRegistros(n int) []RegistroTemperatura {
	if n > len(lt.registros) {
		n = len(lt.registros)
	}

	start := len(lt.registros) - n
	return lt.registros[start:]
}

func ejemploPatronObserver() {
	fmt.Println("  === OBSERVER: SISTEMA DE NOTICIAS ===")

	// Crear centro de noticias
	centroNoticias := NuevoCentroNoticias()

	// Crear observers
	suscriptor1 := NuevoSuscriptorEmail("user1", "ana@email.com", []string{"tecnología", "ciencia"})
	suscriptor2 := NuevoSuscriptorEmail("user2", "carlos@email.com", []string{"todas"})

	appMovil := NuevaAppMovil("app_001", "María", ConfiguracionApp{
		NotificacionesPush: true,
		Sonido:             true,
		Vibrar:             false,
		CategoriasFiltro:   []string{"deportes", "tecnología"},
	})

	dashboard := NuevoDashboardWeb("dashboard_admin", "Administrador")

	// Suscribir observers
	centroNoticias.Suscribir(suscriptor1)
	centroNoticias.Suscribir(suscriptor2)
	centroNoticias.Suscribir(appMovil)
	centroNoticias.Suscribir(dashboard)

	fmt.Printf("    Centro de noticias con %d suscriptores\n\n", centroNoticias.NumeroObservers())

	// Publicar noticias
	noticias := []struct {
		titulo, contenido, categoria string
	}{
		{"Nueva versión de Go", "Go 1.23 incluye mejoras...", "tecnología"},
		{"Campeonato Mundial", "Los resultados del mundial...", "deportes"},
		{"Descubrimiento científico", "Nuevas investigaciones revelan...", "ciencia"},
		{"Actualización de seguridad", "Se recomienda actualizar...", "tecnología"},
	}

	for i, noticia := range noticias {
		fmt.Printf("  Publicando noticia %d:\n", i+1)
		centroNoticias.PublicarNoticia(noticia.titulo, noticia.contenido, noticia.categoria)
		fmt.Println()
	}

	// Mostrar estadísticas
	dashboard.MostrarEstadisticas()
	fmt.Printf("\n    Notificaciones no leídas en app: %d\n", appMovil.NotificacionesNoLeidas())

	// Desuscribir un observer
	fmt.Println("\n  Desuscribiendo user1:")
	centroNoticias.Desuscribir(suscriptor1)

	fmt.Println("\n  Publicando otra noticia:")
	centroNoticias.PublicarNoticia("Noticia final", "Esta es la última noticia...", "general")

	fmt.Println("\n  === OBSERVER: MONITOREO DE TEMPERATURA ===")

	// Crear sensor de temperatura
	sensor := NuevoSensorTemperatura("Sala de Servidores")

	// Crear observers para temperatura
	alerta := NuevaAlertaTemperatura("alerta_principal", 5.0, 30.0)
	logger := NuevoLoggerTemperatura("logger_temp")

	// Suscribir
	sensor.Suscribir(alerta)
	sensor.Suscribir(logger)

	fmt.Println("    Simulando cambios de temperatura:")

	temperaturas := []float64{20.0, 25.0, 32.0, 40.0, 15.0, -2.0, 22.0}

	for i, temp := range temperaturas {
		fmt.Printf("  Lectura %d:\n", i+1)
		sensor.ActualizarTemperatura(temp)
		fmt.Println()
	}

	fmt.Printf("    Total de alertas generadas: %d\n", alerta.NumeroAlertas())
	fmt.Println("    Últimas 3 lecturas:")

	ultimos := logger.UltimosRegistros(3)
	for i, registro := range ultimos {
		fmt.Printf("      [%d] %s: %.1f°C a las %s\n",
			i+1, registro.Ubicacion, registro.Temperatura, registro.Fecha.Format("15:04:05"))
	}
}
