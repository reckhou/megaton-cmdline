package main

import (
  "github.com/reckhou/megaton-cmdline/src/api"
  gval "github.com/reckhou/megaton-cmdline/src/globalVal"
  "log"
  "os"
)

var (
  usageMsg  = "Use --usage for help."
  existArgs = make(map[string]string)
)

func main() {
  log.Println("Megaton command line tool V1.0")
  log.Println("Megaton host:", gval.Args["MTAddr"])

  argCnt := len(os.Args)
  if argCnt < 2 {
    usage()
  }

  parseArgs(os.Args)
  if gval.Args["project"] == "" {
    log.Println("Project not given.")
    usage()
  }

  for i := 1; i < argCnt; i++ {
    arg := os.Args[i]
    if arg == "-h" || arg == "--usage" {
      usage()
    } else if arg == "-pa" || arg == "--publish-all" {
      api.PublishAll(gval.Args["project"], gval.Args["version"])
    } else if arg == "-np" || arg == "--notify-publish" {
      api.NotifyPublish(gval.Args["project"], gval.Args["version"])
    } else if arg == "-ov" || arg == "--online-version" {
      api.SetOnlineVersionID(gval.Args["project"], gval.Args["versionID"])
    } else if arg == "-ap" || arg == "--auto-publish" {
      api.AutoPublish(gval.Args["project"], gval.Args["version"], gval.Args["versionID"])
    } else if arg == "-po" || arg == "--push-to-oss" {
      api.PushToOSS(gval.Args["project"])
    } else if arg == "-pc" || arg == "--push-to-cdn" {
      api.PushToCDN(gval.Args["project"])
    } else if arg == "-uf" || arg == "--upload-file" {
      api.UploadFile(gval.Args["localPath"], gval.Args["project"], gval.Args["fileName"], gval.Args["relativePath"], gval.Args["fileType"])
    } else if arg == "-rf" || arg == "--remove-file" {
      api.RemoveFile(gval.Args["project"], gval.Args["fileName"])
    } else if (arg == "-gt" || arg == "--get-tag") && (i+1 < argCnt) {
      api.GetTag(gval.Args["project"], gval.Args["fileName"])
    } else if (arg == "-st" || arg == "--set-tag") && (i+2 < argCnt) {
      api.SetTag(gval.Args["project"], gval.Args["fileName"], gval.Args["tags"])
    } else if (arg == "-sp" || arg == "--set-profile") && (i+3 < argCnt) {
      api.SetProfile(gval.Args["project"], gval.Args["version"], gval.Args["name"], gval.Args["contentPath"])
    } else if (arg == "-gp" || arg == "--get-profile") && (i+1 < argCnt) {
      api.GetProfile(gval.Args["project"], gval.Args["version"])
    } else if (arg == "-rp" || arg == "--remove-profile") && (i+1 < argCnt) {
      api.RemoveProfile(gval.Args["project"], gval.Args["version"], gval.Args["name"])
    } else if (arg == "-cp" || arg == "--copy-profile") && (i+4 < argCnt) {
      api.CopyProfile(gval.Args["project"], gval.Args["fromVersion"], gval.Args["fromName"], gval.Args["toVersion"], gval.Args["toName"])
    }
  }
}

func parseArgs(args []string) bool {
  argCnt := len(args)
  for i := 0; i < argCnt; i++ {
    arg := args[i]
    if _, isExist := existArgs[arg]; isExist {
      log.Fatal("Option ", arg, " already exists.")
    }

    if (arg == "-p" || arg == "--project") && (i+1 < argCnt) {
      gval.Args["project"] = os.Args[i+1]
      i++
    } else if (arg == "-s" || arg == "--save") && (i+1 < argCnt) {
      gval.Args["save"] = os.Args[i+1]
      i++
    } else if (arg == "-pa" || arg == "--publish-all") && (i+1 < argCnt) {
      gval.Args["version"] = os.Args[i+1]
      i++
    } else if (arg == "-np" || arg == "--notify-publish") && (i+1 < argCnt) {
      gval.Args["version"] = os.Args[i+1]
      i++
    } else if (arg == "-ov" || arg == "--online-version") && (i+1 < argCnt) {
      gval.Args["versionID"] = os.Args[i+1]
      i++
    } else if (arg == "-ap" || arg == "--auto-publish") && (i+2 < argCnt) {
      gval.Args["version"] = os.Args[i+1]
      gval.Args["versionID"] = os.Args[i+2]
      i += 2
    } else if (arg == "-a" || arg == "--address") && (i+1 < argCnt) {
      gval.Args["MTAddr"] = os.Args[i+1]
      i++
    } else if arg == "-v" || arg == "--verbose" {
      gval.Verbose = true
    } else if (arg == "-uf" || arg == "--upload-file") && (i+4 < argCnt) {
      gval.Args["localPath"] = os.Args[i+1]
      gval.Args["fileName"] = os.Args[i+2]
      gval.Args["relativePath"] = os.Args[i+3]
      gval.Args["fileType"] = os.Args[i+4]
      i += 4
    } else if (arg == "-rf" || arg == "--remove-file") && (i+1 < argCnt) {
      gval.Args["fileName"] = os.Args[i+1]
      i++
    } else if (arg == "-gt" || arg == "--get-tag") && (i+1 < argCnt) {
      gval.Args["fileName"] = os.Args[i+1]
      i++
    } else if (arg == "-st" || arg == "--set-tag") && (i+2 < argCnt) {
      gval.Args["fileName"] = os.Args[i+1]
      gval.Args["tags"] = os.Args[i+2]
      i += 2
    } else if (arg == "-sp" || arg == "--set-profile") && (i+3 < argCnt) {
      gval.Args["version"] = os.Args[i+1]
      gval.Args["name"] = os.Args[i+2]
      gval.Args["contentPath"] = os.Args[i+3]
      i += 3
    } else if (arg == "-gp" || arg == "--get-profile") && (i+1 < argCnt) {
      gval.Args["version"] = os.Args[i+1]
      i++
    } else if (arg == "-rp" || arg == "--remove-profile") && (i+2 < argCnt) {
      gval.Args["version"] = os.Args[i+1]
      gval.Args["name"] = os.Args[i+2]
      i += 2
    } else if (arg == "-cp" || arg == "--copy-profile") && (i+4 < argCnt) {
      gval.Args["fromVersion"] = os.Args[i+1]
      gval.Args["fromName"] = os.Args[i+2]
      gval.Args["toVersion"] = os.Args[i+3]
      gval.Args["toName"] = os.Args[i+4]
      i += 4
    }

    existArgs[arg] = ""
  }

  return true
}

func usage() {
  log.Println("-p  | --project <project>                      The project you are working on. This option is a MUST.")
  log.Println("-s  | --save <path>                            Appoint result save path of all \"get\" APIs(optional). Cmdline tool will save to \"tmp.mtsav\" under current directory if save path not appointed.")
  log.Println("-v  | --verbose                                Print debug log.")
  log.Println("-a  | --address <address>                      Appoint Megaton's address instead of address in config.")
  log.Println("-pa | --publish-all <version>                  Publish all dat files with different profiles.")
  log.Println("-np | --notify-publish <version>               Notify Kiloton to update dat files' info to latest publish through Megaton.")
  log.Println("-ov | --online-versionID <versionID>           Set online versionID of client. The client with higher versionID will be treat as test environment by Kiloton.")
  log.Println("-ap | --auto-publish <version> <versionID>     A combination of -pa, -np and -ov. Used for fully automatic publish process.")
  log.Println("-po | --push-to-oss                            Push asset to Aliyun OSS.")
  log.Println("-pc | --push-to-cdn                            Push asset to CDN.")
  log.Println("-uf | --upload-file <localPath> <fileName> <relativePath> <fileType> Upload file to megaton, fileType must be \"raw\" or \"flat\".")
  log.Println("-rf | --remove-file <fileName>                 Remove a file, this will not delete file in file system, only mark it as removed.")
  log.Println("-gt | --get-tag <fileName>                     Get file's tag, tags are seperated by \",\" .")
  log.Println("-st | --set-tag <fileName> <tags>              Set file's tag, tags are seperated by \",\" .")
  log.Println("-sp | --set-profile <version> <name> <contentPath> Set a profile, profile has same name will be overwritten.")
  log.Println("-gp | --get-profile <version>                  Get all profiles under specific version.")
  log.Println("-rp | --remove-profile <version> <name>        Remove profile under specific version.")
  log.Println("-cp | --copy-profile <fromVersion> <fromName> <toVersion> <toName> Copy profile under specific version.")
  log.Println("-h  | --usage                                  This help.")
  os.Exit(0)
}
