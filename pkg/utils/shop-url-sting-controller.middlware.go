package utils

import (
	"fmt"
	"strings"
)

func ShopURLStringController(url string) (string, bool) {
	if strings.Contains(url, "/shop/") && url != "/shop" {
		fmt.Printf("발동 %v\n", url)

		return url[len("/shop/"):], true
	}

	return "", false
}
