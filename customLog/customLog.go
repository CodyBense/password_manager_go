package customlog

import (
	"log"
	"os"
)

func Log() {
    // Log to custom file
    LOG_FILE := "/tmp/pm_log"

    // Open log file
    logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        log.Panic(err)
    }

    // Set log output
    log.SetOutput(logFile)

    // Set log flags
    log.SetFlags(log.Lshortfile | log.LstdFlags)
}
