package util

import (
    "os/exec"
    "log"
    "strings"
)

func UUID() string {
    // uuid for linux
    out, err := exec.Command("uuidgen").Output()
    if err != nil {
        log.Fatal(err)
    }

    return strings.ToLower(strings.Replace(string(out), "-", "", -1))
}