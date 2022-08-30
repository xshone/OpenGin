package utils

import (
	"github.com/google/uuid"
)

func GetUUID() string {
	guid := uuid.New()
	guidStr := guid.String()
	return guidStr
}

func GetTraceId() string {
	return GetUUID()[:6]
}
