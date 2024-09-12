package app

import (
	"log"
	customlog "pm/customLog"

    "database/sql"
    _ "github.com/go-sql-driver/mysql"
	"github.com/charmbracelet/bubbles/table"
)

func SqlConnect() (*sql.DB, error){
    // Custom logger
    customlog.Log()

    // Open Mysql connection
    db, err := sql.Open("mysql", "root:ZSe45rdx##@tcp(127.0.0.1:3306)/Logins")
    if err != nil {
        log.Fatalf("impossible to create the connection: %s", err)
    }

    // Test Mysql connection
    pingErr := db.Ping()
    if err != nil {
        log.Fatalf("impossible to ping the connection: %s", pingErr)
    }

    return db, err
}

func SqlList() []table.Row{
    // Custom logger
    customlog.Log()

    // Open and test connection
    db, err := SqlConnect()

    var (
        website     string
        username    string
        password    string
        itemsRow    []table.Row
    )

    // Conduct query

    // Select all from todo
    rows, err := db.Query("SELECT * FROM login")
    if err != nil {
        log.Fatalf("not able to select conduct query: %s", err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&website, &username, &password)
        if err != nil {
            log.Fatal(err)
        }
        // output rows here
        itemsRow = append(itemsRow, table.Row{website, username, password})
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    db.Close()

    return itemsRow
}

func SqlAdd(website, username, password string) {
    // Custom logger
    customlog.Log()

    // Open and test connection
    db, err := SqlConnect()

    // Conduct insert
    insertQuery := "INSERT INTO login (website, username, password) VALUES (?, ?, ?)"
    stmt, err := db.Prepare(insertQuery)
    if err != nil {
        log.Fatalf("not able to prepare insert query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(website, username, password)
    if err != nil {
        log.Fatalf("not able to execute insert query: %s", err)
    }

    db.Close()
}

func SqlRemove(website string) {
    // Open and test connection
    db, err := SqlConnect()

    // Conduct delete
    removeQuery := "DELETE FROM login WHERE website = ?"
    stmt, err := db.Prepare(removeQuery)
    if err != nil {
        log.Fatalf("not able to prepare remove query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(website)
    if err != nil {
        log.Fatalf("not able to execute remove query: %s", err)
    }

    db.Close()
}

func SqlUpdate(website, username, password string) {
    // Custom logger
    customlog.Log()

    // Open and test connection
    db, err := SqlConnect()

    // Conduct update
    deleteQuery := "DELETE FROM login WHERE website = ?"
    stmt, err := db.Prepare(deleteQuery)
    if err != nil {
        log.Fatalf("not able to prepare delete (update) query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(website)
    if err != nil {
        log.Fatalf("not able to execute delete (update) query: %s", err)
    }

    insertQuery := "INSERT INTO login (website, username, password) VALUES (?,?, ?)"
    stmt, err = db.Prepare(insertQuery)
    if err != nil {
        log.Fatalf("not able to prepare insert (update) query: %s", err)
    }
    defer stmt.Close()

    _, err = stmt.Exec(website, username, password)
    if err != nil {
        log.Fatalf("not abel to execute update query: %s", err)
    }

    db.Close()
}
