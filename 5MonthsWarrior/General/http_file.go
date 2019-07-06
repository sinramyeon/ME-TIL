// curl -F "name=Michael Jackson" -F "thumbnail=@photo.jpg" http://localhost:8080

package main

import (
  "bytes"
  "io"
  "log"
  "mime/multipart"
  "net/http"
  "os"
)

func main() {

  var buffer bytes.buffer
  writer := multipart.NewWriter(&buffer)
  writer.WriteField("name", "Michael Jackson")
  
  fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
  if err != nil {
    panic(err)
  }
  
  readFile, err := os.Open("photo.jpg")
  if err != nil {
    panic(err)
  }
  
  defer readFile.Close()
  io.Copy(fileWriter, readFile)
  writer.Close()
  
  resp, err := http.Post("http://localhost:8080", writer.FormDataContentType(), &buffer)
   if err != nil {
    panic(err)
  }
  
  log.Println("Status", resp.Status)
}
