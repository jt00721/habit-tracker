package testutils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jt00721/habit-tracker/internal/repository"
	"github.com/jt00721/habit-tracker/internal/routes"
	"github.com/jt00721/habit-tracker/internal/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TestServer struct {
	*httptest.Server
}

func NewTestDB(t *testing.T) (*gorm.DB, func()) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Warning: .env file not found, check Load path")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_TEST_DB"), os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Reduce test output noise
	})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	setupSQL, err := os.ReadFile("../../test/testdata/setup.sql")
	if err != nil {
		t.Fatalf("Failed to read setup SQL: %v", err)
	}
	if err := db.Exec(string(setupSQL)).Error; err != nil {
		t.Fatalf("Failed to execute setup SQL: %v", err)
	}

	return db, func() {
		teardownSQL, err := os.ReadFile("../../test/testdata/teardown.sql")
		if err != nil {
			t.Fatalf("Failed to read teardown SQL: %v", err)
		}
		if err := db.Exec(string(teardownSQL)).Error; err != nil {
			t.Fatalf("Failed to execute teardown SQL: %v", err)
		}
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}
}

func NewTestServer(t *testing.T) (*TestServer, *usecase.HabitUsecase, func()) {
	db, teardownDB := NewTestDB(t)

	repo := &repository.HabitRepository{DB: db}
	habitUc := &usecase.HabitUsecase{HabitRepo: repo}

	router := gin.Default()
	router.Use(gin.Recovery())
	routes.SetupRoutes(router, habitUc)

	server := httptest.NewServer(router)

	return &TestServer{server}, habitUc, func() {
		server.Close() // close test server
		teardownDB()   // teardown DB
	}
}

func (ts *TestServer) Get(t *testing.T, path string) (*http.Response, []byte) {
	res, err := http.Get(ts.URL + path)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	return res, body
}
