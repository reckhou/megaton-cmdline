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
  log.Println(os.Args)
  argCnt := len(os.Args)
  if argCnt < 2 {
    usage()
    return
  }

  chkArgs(os.Args)
  if gval.Args["project"] == "" {
    log.Fatal("Project not given.")
  }

  for i := 1; i < argCnt; i++ {
    arg := os.Args[i]

    if arg == "-po" || arg == "--push-to-oss" {
      api.PushToOSS()
    } else if arg == "-v" || arg == "--verbose" {
      gval.Verbose = true
    }
  }
}

func chkArgs(args []string) bool {
  argCnt := len(args)
  for i := 0; i < argCnt; i++ {
    arg := args[i]
    if _, isExist := existArgs[arg]; isExist {
      log.Fatal("Option ", arg, " already exists.")
    }

    if (arg == "-p" || arg == "--project") && (i+1 < argCnt) {
      gval.Args["project"] = os.Args[i+1]
      i++
    }

    existArgs[arg] = ""
  }

  return true
}

func usage() {
  log.Println("-p | --project <project>  The project you are working on.")
  log.Println("-po | --push-to-oss    Push asset to Aliyun OSS.")
  log.Println("-pc | --push-to-cdn    Push asset to CDN.")
  log.Println("-v | --verbose    Print debug log.")
  log.Println("-a | --address    Set megaton server's address.")
  log.Println("--usage    This usage help.")
}
