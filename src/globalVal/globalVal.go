package globalVal

import (
  config "github.com/reckhou/goCfgMgr"
  "log"
  "os"
)

var (
  Verbose = false
  Args    map[string]string
)

func init() {
  defer func() {
    if err := recover(); err != nil {
      log.Println("Invalid config file.")
      os.Exit(1)
    }
  }()

  Args = make(map[string]string)
  Args["MTAddr"] = config.Get("address", nil).(string)
  if Args["MTAddr"] == "" {
    log.Fatal("Megaton server address not specified. Check config file first!")
  }
}
