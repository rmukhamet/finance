package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/rmukhamet/finance/config"
	"github.com/rmukhamet/finance/rest"
	"go.uber.org/zap"
)

func run(ctx context.Context, w io.Writer, args []string) error {
	// Prepare config
	// create app
	// create connections
	// create controllers
	// init app logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()
	// logger := zap.NewExample()

	// defer logger.Sync()

	// stdLogger := zap.NewStdLog(logger)
	// stdLogger.Info("sdfsdfds")

	cfg := config.Config{}

	srv := rest.NewServer(
		sugar,
		&cfg,
		// tenantsStore,
		// slackLinkStore,
		// msteamsLinkStore,
		// proxy,
	)
	httpServer := &http.Server{
		Addr:    net.JoinHostPort("0.0.0.0" /* config.Host */, "80" /* config.Port */),
		Handler: srv,
	}
	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		if err := httpServer.Shutdown(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	ctx := context.Background()

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)

	defer cancel()

	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

/*
os.Args []string Аргументы, передаваемые при исполнении вашей программы. Также используется для флагов парсинга.
os.Stdin io.Reader Для считывания ввода
os.Stdout io.Writer Для записи вывода
os.Stderr io.Writer Для записи логов ошибок
os.Getenv func(string) string Для чтения переменных окружения
os.Getwd func() (string, error) Получение рабочей папки
*/
