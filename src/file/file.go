package file

import (
  "log"
  "os"
)

func ReadFile(fullPath string) []byte {
  file, err := os.Open(fullPath)
  if err != nil {
    log.Println(err)
    return nil
  }

  fileLen, _ := file.Seek(0, 2)
  data := make([]byte, fileLen)
  file.Seek(0, 0)
  file.Read(data)
  //log.Printf("read %d bytes from %s", readLen, fullPath)

  file.Close()
  return data
}
