package api

import (
  js "github.com/bitly/go-simplejson"
  "log"
)

func PushToOSS(project string) bool {
  if project == "" {
    return false
  }

  uri := "/api/pushToOSS?project=" + project
  resp := getURL(uri)
  return CheckResponse(resp)
}

func PushToCDN(project string) bool {
  if project == "" {
    return false
  }

  uri := "/api/pushToCDN?project=" + project
  resp := getURL(uri)
  return CheckResponse(resp)
}

func PublishAll(project, version string) bool {
  if project == "" || version == "" {
    return false
  }

  uri := "/api/publishAll?project=" + project + "&version=" + version
  resp := getURL(uri)
  return CheckResponse(resp)
}

func NotifyPublish(project string) bool {
  if project == "" {
    return false
  }

  uri := "/api/notifyPublish?project=" + project
  resp := getURL(uri)
  return CheckResponse(resp)
}

func SetOnlineVersionID(project, versionID string) bool {
  if project == "" || versionID == "" {
    return false
  }

  uri := "/api/setOnlineVersionID?project=" + project + "&versionID=" + versionID
  resp := getURL(uri)
  return CheckResponse(resp)
}

func CheckResponse(response []byte) bool {
  if response == nil {
    log.Println("Megaton returns nil response.")
    return false
  }

  respJson, err := js.NewJson(response)
  if err != nil {
    log.Fatal("Invalid JSON response from Megaton: ", string(response))
  }

  errCode := respJson.Get("error").MustInt()
  if errCode != 0 {
    log.Println("Megaton returns error:", errCode)
    return false
  }

  return true
}

func AutoPublish(project, version, versionID string) bool {
  if project == "" || version == "" || versionID == "" {
    return false
  }

  if !PublishAll(project, version) {
    return false
  }

  if !NotifyPublish(project) {
    return false
  }

  if !SetOnlineVersionID(project, versionID) {
    return false
  }

  return true
}
