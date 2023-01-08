package hook

import (
	"regexp"
	"strconv"
)

const MAX_UUID_LENGHT = 36

type HookAdapter struct {
	//TODO: add filter like country code ..etc
}

var regexpZipCode = regexp.MustCompile(`^\d{5}$`)

func (h *HookAdapter) CheckId(id string) (uint64, error) {
	return strconv.ParseUint(id, 10, 64)
}

// Return true if given zip code match the regexp
func (h *HookAdapter) CheckZipCode(zip string) bool {
	return regexpZipCode.MatchString(zip)
}
