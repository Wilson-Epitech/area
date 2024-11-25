package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct {
	id string
	email string
	password string
	createdAt string
}

type Service struct {
	id string
	name string
}

type Trigger struct {
	id string
	name string
	description string
	serviceId string
}

type Action struct {
	id string
	name string
	description string
	serviceId string
}

type Workflow struct {
	id string
	name string
	ownerId string
	actionId string
	triggerId string
	isActivated bool
	createdAt string
}

func retrieveDatabaseInfos() (string, int, string, string, string, error) {
	port, err := strconv.Atoi(os.Getenv("PGPORT"))
	if err != nil {
		panic(err)
	}
	return os.Getenv("PGHOST"), port, os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"), err
}

func connectToDatabase() (*sql.DB, error) {
	godotenv.Load()

	host, port, user, password, dbname, err := retrieveDatabaseInfos()
	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	  "password=%s dbname=%s sslmode=require",
	  host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
	  panic(err)
	}

	err = db.Ping()
	if err != nil {
	  panic(err)
	}

	fmt.Println("Successfully connected!")
	return db, err
}

func createUser(db *sql.DB, email string, password string) error {
	sqlStatement := `
    INSERT INTO users (email, password)
    VALUES ($1, $2);
    `
    _, err := db.Exec(sqlStatement, email, password)
    if err != nil {
      return err
    }
	return nil
}

func createService(db *sql.DB, name string) error {
	sqlStatement := `
    INSERT INTO services (name)
    VALUES ($1);
    `
    _, err := db.Exec(sqlStatement, name)
    if err != nil {
		return err
    }
	return nil
}

func createTrigger(db *sql.DB, name string, description string, serviceName string) error {
	retrievedService, err := findServiceByName(db, serviceName)
	if err != nil {
		return err
	}

	sqlStatement := `INSERT INTO triggers (name, description, serviceid) VALUES ($1, $2, $3);`
    _, err = db.Exec(sqlStatement, name, description, retrievedService.id)
    if err != nil {
    	return err
    }
	return nil
}

func createAction(db *sql.DB, name string, description string, serviceName string) error {
	retrievedService, err := findServiceByName(db, serviceName)
	if err != nil {
		return err
	}

	sqlStatement := `INSERT INTO actions (name, description, serviceid) VALUES ($1, $2, $3);`
    _, err = db.Exec(sqlStatement, name, description, retrievedService.id)
    if err != nil {
    	return err
    }
	return nil
}

func createWorkflow(db *sql.DB, name string, ownerEmail string, triggerName string, actionName string) error {
	retrievedUser, err := findUserByEmail(db, ownerEmail)
	if err != nil {
		return err
	}

	retrievedTrigger, err := findTriggerByName(db, triggerName)
	if err != nil {
		return err
	}

	retrievedAction, err := findActionByName(db, actionName)
	if err != nil {
		return err
	}

	sqlStatement := `INSERT INTO workflows (name, ownerid, triggerid, actionid, isactivated) VALUES ($1, $2, $3, $4, $5);`
    _, err = db.Exec(sqlStatement, name, retrievedUser.id, retrievedTrigger.id, retrievedAction.id, true)
    if err != nil {
    	return err
    }
	return nil
}

func findUserByEmail(db *sql.DB, email string) (User, error) {
	var retrievedUser User
	sqlStatement := `SELECT * FROM users where email = ($1)`

    row := db.QueryRow(sqlStatement, email)
	err := row.Scan(&retrievedUser.email, &retrievedUser.password, &retrievedUser.id, &retrievedUser.createdAt)
	if err != nil {
		return retrievedUser, err
	}
	return retrievedUser, nil
}

func findServiceByName(db *sql.DB, name string) (Service, error) {
	var retrievedService Service
	sqlStatement := `SELECT * FROM services where name = ($1)`

    row := db.QueryRow(sqlStatement, name)
	err := row.Scan(&retrievedService.id, &retrievedService.name)
	if err != nil {
		return retrievedService, err
	}
	return retrievedService, nil
}

func findTriggerByName(db *sql.DB, name string) (Trigger, error) {
	var retrievedTrigger Trigger
	sqlStatement := `SELECT * FROM triggers where name = ($1)`

    row := db.QueryRow(sqlStatement, name)
	err := row.Scan(&retrievedTrigger.id, &retrievedTrigger.serviceId, &retrievedTrigger.name, &retrievedTrigger.description)
	if err != nil {
		return retrievedTrigger, err
	}
	return retrievedTrigger, nil
}

func findActionByName(db *sql.DB, name string) (Action, error) {
	var retrievedAction Action
	sqlStatement := `SELECT * FROM actions where name = ($1)`

    row := db.QueryRow(sqlStatement, name)
	err := row.Scan(&retrievedAction.id, &retrievedAction.serviceId, &retrievedAction.name, &retrievedAction.description)
	if err != nil {
		return retrievedAction, err
	}
	return retrievedAction, nil
}

func findWorkflowByName(db *sql.DB, name string) (Workflow, error) {
	var retrievedWorkflow Workflow
	sqlStatement := `SELECT * FROM workflows where name = ($1)`

    row := db.QueryRow(sqlStatement, name)
	err := row.Scan(&retrievedWorkflow.id, &retrievedWorkflow.name, &retrievedWorkflow.ownerId, &retrievedWorkflow.actionId,
		&retrievedWorkflow.triggerId, &retrievedWorkflow.isActivated, &retrievedWorkflow.createdAt)
	if err != nil {
		return retrievedWorkflow, err
	}
	return retrievedWorkflow, nil
}

func main() {
	db, err := connectToDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
