package data

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	data *Store
}

// listen for 'go test' command --> run test methods
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

// run once, before test suite methods
func (suite *TestSuite) SetupSuite() {
	ds, err := NewStore("pg-test", os.Getenv("DB_DEV_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	if err != nil {
		log.Panic(err)
	}

	suite.data = ds
}

// run once, after test suite methods
func (ts *TestSuite) TearDownSuite() {
	log.Println("TearDownSuite()")
}

// run before each test
func (ts *TestSuite) SetupTest() {
	log.Println("SetupTest()")
}

// run after each test
func (ts *TestSuite) TearDownTest() {
	log.Println("TearDownTest()")
}

// run before each test
func (ts *TestSuite) BeforeTest(suiteName, testName string) {
	log.Println("BeforeTest()", suiteName, testName)
}

// run after each test
func (ts *TestSuite) AfterTest(suiteName, testName string) {
	log.Println("AfterTest()", suiteName, testName)
}

// Ther required tables are present
func (ts *TestSuite) TestDefault() {
	var exists bool

	// 'users' table must be present

	ts.data.db.QueryRow(`
		SELECT EXISTS (
    	SELECT FROM pg_tables
    	WHERE 
        schemaname = 'public' AND 
        tablename  = 'users'
    	);
		`).Scan(&exists)

	ts.True(exists, "The 'users' table should exist")

	// 'tokens' table must be present

	ts.data.db.QueryRow(`
		SELECT EXISTS (
			SELECT FROM pg_tables
			WHERE 
				schemaname = 'public' AND 
				tablename  = 'tokens'
			);
		`).Scan(&exists)

	ts.True(exists, "The 'tokens' table should exist")

	// 'pending_reset_passwords' table must be present

	ts.data.db.QueryRow(`
		SELECT EXISTS (
			SELECT FROM pg_tables
			WHERE 
				schemaname = 'public' AND 
				tablename  = 'pending_reset_passwords'
			);
		`).Scan(&exists)

	ts.True(exists, "The 'pending_reset_passwords' table should exist")
}
