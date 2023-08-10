package jwt

import (
	"os"
	"strconv"
	"time"
)

func GetAccessExp() time.Duration {
	exp, _ := strconv.Atoi(os.Getenv("ACCESS_EXP"))
	return time.Duration(exp) * time.Minute
}

func GetRefreshExp() time.Duration {
	exp, _ := strconv.Atoi(os.Getenv("REFRESH_EXP"))
	return time.Duration(exp) * time.Minute
}
