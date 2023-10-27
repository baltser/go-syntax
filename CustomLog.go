package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//вызов os.OpenFile() сооздает файл журнала для записи, если он еще не
	//не существует, или же открывет его для записи, путем добавления новых
	// новых данных в конце (os.O_APPEND)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	iLog := log.New(f, "iLog", log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Mastering Go 3rd edition!")
}
