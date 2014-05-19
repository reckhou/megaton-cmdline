package api

import (
  gval "github.com/reckhou/megaton-cmdline/src/globalVal"
)

func PushToOSS() {
  uri := "/api/pushToOSS?project=" + gval.Args["project"]
  getURL(uri)
}
