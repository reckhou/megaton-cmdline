package api

import (
  gval "github.com/reckhou/megaton-cmdline/src/globalVal"
  "github.com/reckhou/megaton-cmdline/src/verbose"
  "io/ioutil"
  "log"
  "net"
  "net/http"
  "time"
)

var (
  gTransport http.Transport
  gClient    http.Client
)

var timeout = time.Duration(2 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
  return net.DialTimeout(network, addr, timeout)
}

func init() {
  gTransport = http.Transport{
    Dial: dialTimeout,
  }

  gClient = http.Client{
    Transport: &gTransport,
  }
}

func URIToURL(uri string) string {
  return "http://" + gval.Args["MTAddr"] + uri
}

func getURL(uri string) (content []byte) {
  url := URIToURL(uri)
  verbose.Print("Get URL ", url)

  timeStart := time.Now()
  req, err := http.NewRequest("GET", url, nil)
  response, err := gClient.Do(req)
  timeEnd := time.Now()
  if err != nil {
    log.Fatal(err)
  } else {
    defer response.Body.Close()
    responseContent, err := ioutil.ReadAll(response.Body)
    if err != nil {
      log.Fatal(err)
    }
    delta := timeEnd.Sub(timeStart)
    verbose.Print("Query using: ", delta.String())
    content = responseContent
    verbose.Print("The calculated length is:", len(string(responseContent)), "for the url:", url)
    verbose.Print("   ", response.StatusCode)
    hdr := response.Header
    for key, value := range hdr {
      verbose.Print("   ", key, ":", value)
    }
    verbose.Print(string(content))
  }

  return content
}
