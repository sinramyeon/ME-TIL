// MIME타입으로 image/jpeg를 설정
...

import "net/textproto"

part := make(textproto.MIMEheader)
part.Set("Content-Type", "image/jpeg")
part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.jpg"`)

fileWriter, err := writer.CreatePart(part)
if err != nil {
  panic(err)
}

readFile, err := os.Open("photo.jpg")
if err != nil {
  panic(err)
}

io.Copy(fileWriter, readFile)

...
