package common

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/comail/colog"
)

// CustomLogger output log file.
func CustomLogger() error {
	file, err := os.OpenFile(GetLogFilePath(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Println(fmt.Sprintf("%v: %v", "error", err.Error()))
		return err
	}
	colog.Register()
	colog.SetOutput(io.MultiWriter(file, os.Stdout))
	colog.SetFormatter(&colog.StdFormatter{
		Flag: log.Ldate | log.Ltime | log.Lshortfile,
	})
	return nil
}
