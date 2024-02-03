package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func gooseLogger(status string, errorNum int, errorMsg string, fileName string, err error) string {
	//~ Get current time and date
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")

	//* Log message
	logMessage := fmt.Sprintf("[%v] - Status: %v] - [Num: %v] - [Msg: %v] - [File: %v] - (System Error: %v)", timeString, status, errorNum, errorMsg, fileName, err)

	//^ Print out GOOSE_LOGGER message
	fmt.Println(logMessage)

	//& Write GOOSE_LOGGER message .log file | This will create if does not exist, appends to file
	logFile, err := os.OpenFile("Server_Logs.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "Error opening file:"
	}
	defer logFile.Close()

	//& Append to the log file
	writer := bufio.NewWriter(logFile)
	_, err = writer.WriteString(logMessage + "\n")

	if err != nil {
		fmt.Println("Logger Error: Error opening file:", err)
		return "Logger Error: Error opening file:"
	}

	//^ Flush the writer to ensure all content is written to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return "Error flushing writer:"
	}

	//^ Return string
	return logMessage
}
