package cli

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	internalHTTP "gitlab.com/EndowTheGreat/spark/http"
	"gitlab.com/EndowTheGreat/spark/markdown"

	"github.com/spf13/cobra"
)

var port string
var logger bool

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "The port to host web server on.")
	serveCmd.Flags().StringVarP(&markdown.Output, "dir", "d", "8080", "The output directory containing your converted files to serve.")
	serveCmd.Flags().BoolVarP(&logger, "log", "l", true, "Whether or not to log requests made to the web server.")
	serveCmd.MarkFlagRequired("dir")
}

var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Serve the converted files in your output directory.",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		handler := internalHTTP.NewRouter()
		handler.SetupRoutes(logger)
		server := &http.Server{
			Addr:    ":" + port,
			Handler: handler.Router,
		}
		log.Printf("Starting web server on :%v\n", port)
		serverCtx, serverStopCtx := context.WithCancel(context.Background())
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		go func() {
			<-sig
			shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
			defer cancel()
			go func() {
				<-shutdownCtx.Done()
				if shutdownCtx.Err() == context.DeadlineExceeded {
					log.Fatal("Graceful shutdown timed out, forcefully shutting down...")
				}
			}()
			if err := server.Shutdown(shutdownCtx); err != nil {
				log.Fatal(err)
			}
			serverStopCtx()
		}()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to instantiate HTTP server: ", err)
		}
		<-serverCtx.Done()
	},
}
