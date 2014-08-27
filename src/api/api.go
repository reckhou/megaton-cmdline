package api

import (
  js "github.com/bitly/go-simplejson"
  "github.com/reckhou/megaton-cmdline/src/file"
  gval "github.com/reckhou/megaton-cmdline/src/globalVal"
  "log"
)

func GetSavePath() string {
  if gval.Args["save"] != "" {
    return gval.Args["save"]
  }

  return "./tmp.mtsav"
}

func SaveResponse(content []byte) bool {
  savePath := GetSavePath()
  return file.WriteFile(savePath, content)
}

func PushToOSS(project string) bool {
  if project == "" {
    log.Fatal("Illegal param!")
  }

  uri := "/api/pushToOSS?project=" + project
  resp := getURL(uri)
  return CheckResponse(uri, resp)
}

func PushToCDN(project string) bool {
  if project == "" {
    log.Fatal("Illegal param!")
  }

  uri := "/api/pushToCDN?project=" + project
  resp := getURL(uri)
  return CheckResponse(uri, resp)
}

func PublishAll(project, version, assetVersion string) bool {
  if project == "" || version == "" {
    log.Fatal("Illegal param!")
  }

  uri := "/api/publishAll?project=" + project + "&version=" + version
  if assetVersion != "" {
    uri += "&assetVersion=" + assetVersion
  }
  resp := getURL(uri)
  return CheckResponse(uri, resp)
}

func NotifyPublish(project, version string) bool {
  if project == "" || version == "" {
    log.Fatal("Illegal param!")
  }

  uri := "/api/notifyPublish?project=" + project + "&version=" + version
  resp := getURL(uri)
  return CheckResponse(uri, resp)
}

func SetOnlineVersionID(project, versionID string) bool {
  if project == "" || versionID == "" {
    log.Fatal("Illegal param!")
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

  if request != "" {
    log.Println("Request", request, "succeed.")
  }

  return true
}

func AutoPublish(project, version, versionID, assetVersion string) bool {
  if project == "" || version == "" || versionID == "" {
    log.Fatal("Illegal param!")
  }

  if !PublishAll(project, version, assetVersion) {
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
    log.Fatal("Illegal param!")
  }

  content := file.ReadFile(localPath)
  if content == nil {
    log.Fatal("Read file at", localPath, "failed.")
  }

  uri := "/api/uploadFile?project=" + project + "&relativePath=" + relativePath + "&fileName=" + fileName + "&fileType=" + fileType
  responseContent := postURL(uri, content)

  return CheckResponse(uri, responseContent)
}

func RemoveFile(project, fileName string) bool {
  if project == "" || fileName == "" {
    log.Fatal("Illegal param!")
  }

  uri := "/api/removeFile?project=" + project + "&fileName=" + fileName
  resp := getURL(uri)
  return CheckResponse(uri, resp)
}

func GetTag(project, fileName string) bool {
  if project == "" || fileName == "" {
    log.Fatal("Illegal param!")
  }

  uri := "/api/getTag?project=" + project + "&fileName=" + fileName
  resp := getURL(uri)

  if !CheckResponse(uri, resp) {
    return false
  }

  log.Println("Tag of", fileName, string(resp))
  return SaveResponse(resp)
}

func SetTag(project, fileName, tags string) bool {
  if project == "" || fileName == "" || tags == "" {
    log.Fatal("Illegal param!")
  }

  uri := "/api/setTag?project=" + project + "&fileName=" + fileName + "&tags=" + tags
  resp := getURL(uri)

  return CheckResponse(uri, resp)
}

func SetProfile(project, version, name, contentPath string) bool {
  if project == "" || version == "" || name == "" || contentPath == "" {
    log.Fatal("Illegal param!")
  }

  content := file.ReadFile(contentPath)
  if content == nil {
    log.Fatal("Read file at", contentPath, "failed.")
  }

  uri := "/api/setProfile?project=" + project + "&version=" + version + "&name=" + name
  responseContent := postURL(uri, content)

  return CheckResponse(uri, responseContent)
}

func GetProfile(project, version string) bool {
  if project == "" || version == "" {
    log.Fatal("Illegal param!")
  }

  uri := "/api/getProfile?project=" + project + "&version=" + version
  responseContent := getURL(uri)

  if !CheckResponse(uri, responseContent) {
    return false
  }

  log.Println(string(responseContent))
  return SaveResponse(responseContent)
}

func RemoveProfile(project, version, name string) bool {
  if project == "" || version == "" || name == "" {
    log.Fatal("Illegal param!")
  }

  uri := "/api/removeProfile?project=" + project + "&version=" + version + "&name=" + name
  responseContent := getURL(uri)

  return CheckResponse(uri, responseContent)
}

func CopyProfile(project, fromVersion, fromName, toVersion, toName string) bool {
  if project == "" || fromVersion == "" || fromName == "" || toVersion == "" || toName == "" {
    log.Fatal("Illegal param!")
  }

  uri := "/api/copyProfile?project=" + project + "&fromVersion=" + fromVersion + "&fromName=" + fromName + "&toVersion=" + toVersion + "&toName=" + toName
  responseContent := getURL(uri)

  return CheckResponse(uri, responseContent)
}

func AddPatchFile(project, version, file, profile string) bool {
  if project == "" || version == "" || file == "" || profile == "" {
    log.Fatal("Illegal param!")
  }

  uri := "/addPatchFile?project=" + project + "&version=" + version + "&file=" + file + "&profile=" + profile
  responseContent := getURL(uri)

  return CheckResponse(uri, responseContent)
}

func RemovePatchFile(project, version, file, profile string) bool {
  if project == "" || version == "" || file == "" || profile == "" {
    log.Fatal("Illegal param!")
  }

  uri := "/removePatchFile?project=" + project + "&version=" + version + "&file=" + file + "&profile=" + profile
  responseContent := getURL(uri)

  return CheckResponse(uri, responseContent)
}
