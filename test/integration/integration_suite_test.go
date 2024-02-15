package integration

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/uptrace/bun"
	"github.com/vctaragao/assertdb/internal"
	"github.com/vctaragao/assertdb/internal/infra/database"
	"github.com/vctaragao/assertdb/internal/web"
	"github.com/vctaragao/assertdb/pkg"
)

type IntegrationSuite struct {
	suite.Suite
	dbConn pkg.DBConnection
	server *http.Server
}

func (s *IntegrationSuite) SetupSuite() {
	s.initDB()
	s.initApp()
}

func (s *IntegrationSuite) TearDownSuite() {
	s.closeApp()
}

func (s *IntegrationSuite) TearDownTest() {
	if err := s.dbConn.DbTx.Rollback(); err != nil {
		log.Fatal(err)
	}
}

func (s *IntegrationSuite) initDB() {
	dbConn := database.Connect()
	if err := database.Migrate(dbConn); err != nil {
		log.Fatal(err)
	}

	s.dbConn = s.NewTestSuiteDbConn(context.Background(), dbConn)
}

func (s *IntegrationSuite) initApp() {
	bananaAdapter := database.NewBananaAdapter(s.dbConn.DbTx)
	createService := internal.NewCreateBananaService(bananaAdapter)

	mux := http.NewServeMux()
	web.RegisterRoutes(mux, createService)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	s.server = server

	go func() {
		log.Println("Server running on port 8080...")
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	serverRunning := make(chan bool)
	defer close(serverRunning)

	go func() {
		retries := 5
		for retries > 0 {
			resp, err := http.Get("http://localhost:8080/healthz")
			if err != nil {
				log.Fatal(err)
			}

			if resp.StatusCode == http.StatusOK {
				serverRunning <- true
				break
			}

			retries--
			time.Sleep(500 * time.Millisecond)
		}
	}()

	select {
	case <-serverRunning:
		log.Println("Server started")
	case <-time.After(5 * time.Second):
		log.Fatal("Server did not start")
	}
}

func (s *IntegrationSuite) closeApp() {
	if err := s.dbConn.Close(); err != nil {
		log.Println(err)
	}

	if err := s.server.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}

	log.Println("App closed successfully")
}

func (s *IntegrationSuite) NewTestSuiteDbConn(ctx context.Context, dbConn *bun.DB) pkg.DBConnection {
	conn := pkg.DBConnection{DbConn: dbConn}
	conn.BeginTx(ctx)

	return conn
}

func TestIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test suite")
	}

	suite.Run(t, new(IntegrationSuite))
}
