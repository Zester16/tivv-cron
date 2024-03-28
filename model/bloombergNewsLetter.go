package model

import "os"

const BLM_FIVE = "blm_five"
const BLM_ECO = "blm_eco"
const BLM_EUROPE = "blm_europe"
const BLM_TECH = "blm_tech"

var BloombergUrls = map[string]string{}

type test struct {
	blmUrls map[string]string
}

var BlmTest = test{}

func (t *test) SetUrls() {
	bloomberg := map[string]string{BLM_FIVE: os.Getenv(BLM_FIVE),
		BLM_ECO: os.Getenv(BLM_ECO),
	}
	t.blmUrls = bloomberg

}
