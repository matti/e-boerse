package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/go-resty/resty/v2"
)

type priceResponse struct {
	Price float64
}

func main() {
	var hours int
	var mode string
	flag.IntVar(&hours, "hours", 1, "")
	flag.StringVar(&mode, "mode", "float", "")
	flag.Parse()

	var times []time.Time
	now := time.Now()
	for i := 0; i < hours; i++ {
		times = append(times, now.Add(time.Hour*time.Duration(i)))
	}

	client := resty.New()

	var total float64
	for _, t := range times {
		var priceResponse *priceResponse
		u := fmt.Sprintf("https://api.porssisahko.net/v1/price.json?date=%s&hour=%d", t.Format("2006-01-02"), t.Hour())
		if resp, err := client.R().SetResult(&priceResponse).Get(u); err != nil {
			log.Println(resp.StatusCode(), err)
		} else {
			total = total + priceResponse.Price
		}
	}

	switch mode {
	case "float":
		fmt.Println(total)
	case "int":
		fmt.Println(int(math.Round(total)))
	default:
		log.Fatalln("unknown mode: ", mode)
	}
}
