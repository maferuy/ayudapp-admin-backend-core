package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/maferuy/ayudapp-admin-backend-core/handlers"
	"github.com/maferuy/ayudapp-admin-backend-core/utils"

	_ "github.com/maferuy/ayudapp-admin-backend-core/docs"
)

//@title API de Administración de Ayud App
//@version 1.0
//@description API para administradores de Ayud App
//
//@contact.name Martin Fernandez
//@contact.url https://mafer.dev
//@contact.email maramal@outlook.com
//
//@host localhost:8080
//@BasePath /api/
//@query.collection-format multi
//
//@securityDefinitions.apikey ApiKeyAuth
//@in header
//@name Authorization
func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("No se pudo cargar la configuración: ", err)
	}

	utils.SetupLogOutput()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	server, err := handlers.NewServer(config)
	if err != nil {
		log.Fatal("No se pudo crear el servidor: ", err)
	}

	defer func() {
		if err := server.Client.Disconnect(context.TODO()); err != nil {
			log.Fatal("No fue posible desconectar la base de datos: ", err)
		}
	}()

	httpServer := &http.Server{
		Addr:    "localhost:" + config.Port,
		Handler: server.Router,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error al iniciar el servidor: %s", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("Deteniendo servidor. Presiona Ctrl+C para forzar.")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Servidor detenido forzosamente: %s", err)
	}

	log.Println("Servidor detenido.")
}
