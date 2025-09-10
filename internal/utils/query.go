package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

type DateRange struct {
	From *time.Time
	To   *time.Time
}

func ParseDateRange(c *gin.Context) (DateRange, error) {
	layout := "2006-01-02"
	var dr DateRange
	if from := c.Query("from"); from != "" {
		t, err := time.Parse(layout, from); if err != nil { return dr, err }
		dr.From = &t
	}
	if to := c.Query("to"); to != "" {
		t, err := time.Parse(layout, to); if err != nil { return dr, err }
		// include end of day
		tt := t.Add(24*time.Hour - time.Nanosecond)
		dr.To = &tt
	}
	return dr, nil
}
