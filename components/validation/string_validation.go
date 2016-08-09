package validation

import (
	"regexp"
	"net/url"
	"strings"
	"encoding/base64"
	"github.com/DeepForestTeam/mobiussign/components/log"
)

func init() {

}

func IsSha512(str string) bool {
	if len(str) != 128 && len(str) != 64 {
		return false
	}
	for _, v := range str {
		if ('F' < v || v < 'A') && ('f' < v || v < 'a') && ('9' < v || v < '0') {
			return false
		}
	}
	return true
}

func IsUrl(str string) bool {
	if str == "" || len(str) >= 2083 || len(str) <= 3 {
		return false
	}
	u, err := url.Parse(str)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}
	url_pattern := `^((ftp|https?):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(([a-zA-Z0-9]([a-zA-Z0-9-]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|((www\.)?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`
	check_url := regexp.MustCompile(url_pattern)
	match := check_url.MatchString(str)
	return match
}

func IsBase64(str string) bool {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\t", "", -1)

	if len(str) < 3 {
		return false
	}
	for _, v := range str {
		if ('Z' < v || v < 'A') && ('z' < v || v < 'a') && ('9' < v || v < '0') && (v != '+' && v != '/' && v != '=') {
			return false
		}
	}
	_, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}