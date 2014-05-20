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
    return
  }

  parseArgs(os.Args)
  if gval.Args["project"] == "" {
    log.Fatal("Project not given.")
  }

  for i := 1; i < argCnt; i++ {
    arg := os.Args[i]

    if arg == "-pa" || arg == "--publish-all" {
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
    }

    existArgs[arg] = ""
  }

  return true
}

func usage() {
  log.Println("-p  | --project <project>                      The project you are working on. This option is a MUST.")
  log.Println("-pa | --publish-all <version>                  Publish all dat files with different profiles.")
  log.Println("-np | --notify-publish <version>               Notify Kiloton to update dat files' info to latest publish through Megaton.")
  log.Println("-ov | --online-versionID <versionID>           Set online versionID of client. The client with higher versionID will be treat as test environment by Kiloton.")
  log.Println("-ap | --auto-publish <version> <versionID>     A combination of -pa, -np and -ov. Used for fully automatic publish process.")
  log.Println("-po | --push-to-oss                            Push asset to Aliyun OSS.")
  log.Println("-pc | --push-to-cdn                            Push asset to CDN.")
  log.Println("-v  | --verbose                                Print debug log.")
  log.Println("-a  | --address <address>                      Appoint Megaton's address instead of address in config.")
  log.Println("--usage                                        This usage help.")
}
