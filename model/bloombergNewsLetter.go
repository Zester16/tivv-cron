package model

import "os"

/*
* all blm keys
 */
const BLM_FIVE = "blm_five"
const BLM_ECO = "blm_eco"
const BLM_EUROPE = "blm_europe"
const BLM_TECH = "blm_tech"

/*
* all nyt keys
 */

const NYT_DEALBOOK = "nyt_dealbook"
const NYT_MORNING_AUS = "nyt_morning_aus"
const NYT_MORNING_APAC = "nyt_morning_apac"
const NYT_MORNING_EUROPE = "nyt_morning_europe"
const NYT_MORINIG_US = "nyt_morning_us"
const NYT_EVENING_US = "nyt_evening_us"

var BloombergUrls = map[string]string{}

type test struct {
	blmUrls map[string]string
	nytUrls map[string]string
}

var BlmTest = test{}

func (t *test) SetUrls() {
	blmUrls := map[string]string{BLM_FIVE: os.Getenv(BLM_FIVE),
		BLM_ECO: os.Getenv(BLM_ECO),
	}
	nytUrls := map[string]string{
		NYT_DEALBOOK:       os.Getenv(NYT_DEALBOOK),
		NYT_MORNING_AUS:    os.Getenv(NYT_MORNING_AUS),
		NYT_MORNING_APAC:   os.Getenv(NYT_MORNING_APAC),
		NYT_MORNING_EUROPE: os.Getenv(NYT_MORNING_EUROPE),
		NYT_MORINIG_US:     os.Getenv(NYT_MORINIG_US),
		NYT_EVENING_US:     os.Getenv(NYT_EVENING_US),
	}
	t.blmUrls = blmUrls
	t.nytUrls = nytUrls

}

func (t *test) GetBLMUrls() map[string]string {
	return t.blmUrls
}

func (t *test) GetNYTUrls() map[string]string {
	return t.nytUrls
}
