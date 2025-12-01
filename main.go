package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Inicializa el router de Gin con la configuración por defecto (incluye logger y middleware de recuperación).
	r := gin.Default()

	// 2. Define el único endpoint en la ruta raíz ("/").
	// Cuando alguien haga una petición GET a "/", se ejecutará esta función.
	r.GET("/", func(c *gin.Context) {
		// Envía una respuesta de texto plano con el código de estado 200 OK.
		c.String(http.StatusOK, "Bienvenido a mi deploy")
	})

	// 3. Lógica para el puerto (¡Esto es CLAVE para el despliegue!)
	// Render (y la mayoría de plataformas) te asignan un puerto a través de la variable de entorno PORT.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Si no se define la variable, usa el puerto 8080 para desarrollo local.
	}

	// 4. Inicia el servidor en el puerto correcto.
	// r.Run() escuchará en una dirección como ":8080".
	r.Run(":" + port)
}
