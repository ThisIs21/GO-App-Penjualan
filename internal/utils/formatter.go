package utils

import (
	"fmt"
	"time"
)

func Rupiah(v float64) string {
	// sederhana, sisi front-end biasanya format sendiri
	return fmt.Sprintf("Rp%.0f", v)
}
func TanggalID(t time.Time) string {
	return t.Format("02-01-2006 15:04")
}
