package modules

import (
	"github.com/thanhpk/randstr"

    "time"
    "strconv"
)

func GenRndName() (string) {
	now := time.Now().Unix()
	return strconv.FormatInt(now, 10) + "_" + randstr.Hex(16)
}

func GenRndPDFName() (string) {
	return GenRndName() + ".pdf"
}