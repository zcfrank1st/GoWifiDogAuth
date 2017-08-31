package util

import (
    "strings"
)

func ParseRedirectUrl (html string) string {
    trans := strings.Replace(html, "<html><body>Please <a href='", "", -1)
    return strings.Replace(trans, "'>click here</a>.</body></html>", "", -1)
}
