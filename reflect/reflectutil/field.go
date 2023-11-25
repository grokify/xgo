package reflectutil

import (
	"strings"

	reflections "github.com/oleiade/reflections"
)

func GetField(i any, fieldPath ...string) (any, error) {
	if len(fieldPath) == 0 {
		return i, nil
	}
	nextItem, err := reflections.GetField(i, strings.TrimSpace(fieldPath[0]))
	if err != nil || len(fieldPath) == 1 {
		return nextItem, err
	}
	return GetField(nextItem, fieldPath[1:]...)
}
