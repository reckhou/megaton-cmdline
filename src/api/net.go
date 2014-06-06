package api

import (
  "bytes"
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

func postURL(uri string, content []byte) (resp []byte) {
  if uri == "" || content == nil {
    return
  }

  url := URIToURL(uri)
  verbose.Print(url)

  resp = make([]byte, 0)
  transport := http.Transport{
    Dial: dialTimeout,
  }

  client := http.Client{
    Transport: &transport,
  }

  req, err := http.NewRequest("POST", url, bytes.NewReader(content))
  if err != nil {
    log.Println(err)
    return nil
  }

  response, err := client.Do(req)
  if err != nil {
    log.Println(err)
    return nil
  }
  defer response.Body.Close()
  resp, err = ioutil.ReadAll(response.Body)
  if err != nil {
    log.Println(err)
    return nil
  }

  return resp
}
