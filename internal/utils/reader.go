package main

import (
	"context"
	"io"
	"net/http"
	"strconv"

	"github.com/celestix/gotgproto"
	"github.com/gotd/td/tg"
	"go.uber.org/zap"

	"github.com/celestix/gotgproto/utils" // Ruta original de telegramReader
)

var logger *zap.Logger

func main() {
	logger, _ = zap.NewDevelopment()
	defer logger.Sync()

	// Inicializa el cliente de Telegram
	client := gotgproto.NewClient() // Aquí deberías conectar y autenticar tu cliente antes de usarlo

	http.HandleFunc("/video", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		// TODO: Inicializa correctamente estos valores desde tu lógica
		var location tg.InputFileLocationClass = /* ubicación del archivo en Telegram */ nil
		start := int64(0)
		end := int64(1024 * 1024 * 100) // por ejemplo, 100 MB
		contentLength := end - start

		reader, err := utils.NewTelegramReader(ctx, client, location, start, end, contentLength)
		if err != nil {
			http.Error(w, "failed to create reader", http.StatusInternalServerError)
			return
		}
		defer reader.Close()

		w.Header().Set("Content-Type", "video/mp4")
		w.Header().Set("Content-Length", strconv.FormatInt(contentLength, 10))
		w.WriteHeader(http.StatusOK)
		_, _ = io.Copy(w, reader)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	logger.Sugar().Info("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
