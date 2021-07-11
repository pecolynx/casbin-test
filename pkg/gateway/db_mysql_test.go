package gateway

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	gormMySQL "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	testDBURL string
)

func XXinit() {
	fmt.Println("init")

	testDBHost := os.Getenv("TEST_DB_HOST")
	if testDBHost == "" {
		testDBHost = "127.0.0.1"
	}

	testDBPort := os.Getenv("TEST_DB_PORT")
	if testDBPort == "" {
		testDBPort = "3307"
	}

	testDBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Asia%%2FTokyo", "user", "password", testDBHost, testDBPort, "testdb")

	db, err := gorm.Open(gormMySQL.Open(testDBURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	pos := strings.Index(wd, "pkg")
	dir := wd[0:pos] + "sqls"
	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+dir, "mysql", driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(fmt.Errorf("Failed to up. err: %w", err))
		}
	}
}

func TestA(t *testing.T) {

}
