package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("https://yandex.ru/pogoda/moscow")

	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	re, _ := regexp.Compile(`current-weather__thermometer_type_now[\s\w"']*>([^<]+)`)

	match := re.FindStringSubmatch(string(body))

	fmt.Println(match[1])
}
