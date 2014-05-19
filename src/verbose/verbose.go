package verbose

import (
  gval "github.com/reckhou/megaton-cmdline/src/globalVal"
  "log"
)

func Print(v ...interface{}) {
  if gval.Verbose {
    log.Println(v)
  }
}
