# 동시성

## 고루틴

고 버전 스레드

`go f(x, y, z)`

### 병렬성, 병행성

병렬성 Parallelism 두사람이 동시에 각각 업무를 보는 것
동시성(병행성) Concurrency 커피를 마시려고 신문 잠깐 보다 말고 커피 한모금 마시고 신문 보는것

```
func main() {
	go func() {
		fmt.Println("In goroutine")
	}()

	fmt.Println("main routine")
}
```

메인 루틴이 끝나기 전 go루틴이 시작이 안 될수도 있다.

### 고루틴 기다리기

```
// 이미지 파일들을 동시 다운로드받고 zip으로 압축하는 코드
var urls = []string{
	"http://image.com/img01.jpg",
	"....",
}

for _, url := range urls {

	go func(url string) {
	if _, err := download(url); err != nil {
		log.Fatal(err)
	}
	}(url)
}

filenames, err := filepath.Glob("*.jpg")
if err != nil {
	log.Fatal(err)
}

err = writeZip("images.zip", filenames)

if err!= nil {
	log.Fatal(err)
}

func download(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil{
		return "", err
	}
	filename, err := urlToFilename(url)
	if err != nil {
		return "", err
	}
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
defer f.Close()
_, err = io.Copy(f, resp.Body)
return filename, err
}

func urlToFilename(rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	return filepath.Base(url.Path), nil
}

func writeZip(outFilename string, filenames []string) error {
	outf, err := os.Create(outFilename)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(outf)
	for _, filename := range filenames {
		w, err := zw.Create(filename)
		if err != nil {
		return err
	}
	f, err:= os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(w, f)
	if err != nil {
		return err
	}
	}
return zw.Close()
}
```


이미지 파일이 모두 다운된 후 압축하려면 어떻게 해야 할까? syncd.WaitGroup을 사용한다.

```
var wg sync.WaitGroup
wg.AdD(len(urls))

for _, url := range urls {
	go func(url string) {
		defer wg.Done()
		if _, err := download(url); err != nil {
			log.Fatal(err)
		}(url)
	}
wg.Wait()
}
```

url 개수만큼의 고루틴이 끝나면 카운터가 0이 되어서 멈추게 된다.

## 채널

채널을 통해 서로 다른 고루틴 사이 통신을 할 수 있다.

```
c1 := make(chan int)
var chan int c2 = c1 //	 c2 = c1
var <- chan int c3 = c1 // 받기 전용
var chan<- int c4 = c1 // 보내기 전용
```

> c<-100 c에 100 보내기 / <-c c에서 자료 받기

### 일대일 단방향

```
funcExample_simpleChannel() {
	
}
```
