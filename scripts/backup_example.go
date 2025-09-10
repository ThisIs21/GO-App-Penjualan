// DISCLAIMER: jalan kalau server punya mysqldump & izin jalankan proses.
package main

import (
	"os/exec"
	"time"
	"log"
	"os"
)

func main() {
	filename := "backup-" + time.Now().Format("20060102-150405") + ".sql"
	cmd := exec.Command(os.Getenv("MYSQLDUMP_PATH"),
		"-h", os.Getenv("DB_HOST"), "-P", os.Getenv("DB_PORT"),
		"-u", os.Getenv("DB_USER"), "-p"+os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
	out, err := os.Create(filename); if err != nil { log.Fatal(err) }
	defer out.Close()
	cmd.Stdout = out
	log.Println("running mysqldump to", filename)
	if err := cmd.Run(); err != nil { log.Fatal(err) }
	log.Println("done")
}
