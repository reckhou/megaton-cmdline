package api

import (
  "bytes"
  js "github.com/bitly/go-simplejson"
  "github.com/reckhou/megaton-cmdline/src/file"
  "github.com/reckhou/megaton-cmdline/src/globalVal"
  "io/ioutil"
  "log"
  "net/http"
)

func PushToOSS(project string) bool {
  if project == "" {
    return false
  }

  uri := "/api/pushToOSS?project=" + project
  resp := getURL(uri)
  return CheckResponse(uri, resp)
}

func PushToCDN(project string) bool {
  if project == "" {
    return false
  }

  uri := "/api/pushToCDN?project=" + project
  resp := getURL(uri)
  return CheckResponse(uri, resp)
}

func PublishAll(project, version string) bool {
  if project == "" || version == "" {
    return false
  }

  uri := "/api/publishAll?project=" + project + "&version=" + version
  resp := getURL(uri)
  return CheckResponse(uri, resp)
}

func NotifyPublish(project, version string) bool {
  if project == "" || version == "" {
    return false
  }

  uri := "/api/notifyPublish?project=" + project + "&version=" + version
  resp := getURL(uri)
  return CheckResponse(uri, resp)
}

func SetOnlineVersionID(project, versionID string) bool {
  if project == "" || versionID == "" {
    return false
  }

  uri := "/api/setOnlineVersionID?project=" + project + "&versionID=" + versionID
  resp := getURL(uri)
  return CheckResponse(uri, resp)
}

func CheckResponse(request string, response []byte) bool {
  if response == nil {
    log.Println("Megaton returns nil response.")
    return false
  }

  respJson, err := js.NewJson(response)
  if err != nil {
    log.Println("Response from Megaton:", string(response))
    return true
  }

  errCode := respJson.Get("error").MustInt()
  if errCode != 0 {
    log.Println("Megaton returns error:", errCode)
    return false
  }

  log.Println("Request", request, "succeed.")
  return true
}

func AutoPublish(project, version, versionID string) bool {
  if project == "" || version == "" || versionID == "" {
    return false
  }

  if !PublishAll(project, version) {
    return false
  }

  if !NotifyPublish(project, version) {
    return false
  }

  if !SetOnlineVersionID(project, versionID) {
    return false
  }

  return true
}

func UploadFile(localPath, project, fileName, relativePath, fileType string) bool {
  if localPath == "" || project == "" || fileName == "" || relativePath == "" || (fileType != "raw" && fileType != "flat") {
    return false
  }

  content := file.ReadFile(localPath)
  if content == nil {
    return false
  }

  transport := http.Transport{
    Dial: dialTimeout,
  }

  client := http.Client{
    Transport: &transport,
  }

  url := "http://" + globalVal.Args["MTAddr"] + "/api/uploadFile?project=" + project + "&relativePath=" + relativePath + "&name=" + fileName + "&fileType=" + fileType
  req, err := http.NewRequest("POST", url, bytes.NewReader(content))
  if err != nil {
    log.Println(err)
    return false
  }

  response, err := client.Do(req)
  if err != nil {
    log.Println(err)
    return false
  } else {
    defer response.Body.Close()
    responseContent, err := ioutil.ReadAll(response.Body)
    if err != nil {
      log.Println(err)
    }

    log.Println(string(responseContent))
  }

  return true
}
