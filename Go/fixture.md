# fixture

fixture? fancy way to say sample data

```
var update = flag.Bool("update", false, "update golen files")

func TestFixture(t *testing T) {
	// table 로 쓸 자료 예

	for _, tc := range cases {
		actual := doSomething(tc)
		golen := filepath.Join("text-fixtureS", tc.Name+".golen")
		if *update {
			ioutil.WriteFile(golen, actual, 0644)
		}
	
		expected, _ := ioutil.ReadFile(golden)
		if !bytes.Equal(actual, expected) {
		// FAIL
		}
	}

}
```

- `go test` 는 현재 폴더를 패키지 디렉토리로 설정한다.
- `test-fixtures`를 상대경로로 테스트 데이터를 저장할 때 사용한다.
- config, model, binary 데이터를 불러올 때 유용하다.

즉 `f, err := os.Open("testdata/somefixture.json")` 과 같이 미리 저장되어있는 데이터를 불러오는 방식이 fixture라고 이해하면 편하다.

#### golden files

[golden files]([https://golang.org/src/cmd/gofmt/testdata/](https://golang.org/src/cmd/gofmt/testdata/))는 고언어 스탠다드 라이브러리에서도 사용되는 트릭인데, [https://www.youtube.com/watch?v=yszygk1cpEc](https://www.youtube.com/watch?v=yszygk1cpEc) 에서 배웠다.

루비에서 [비슷한 방식]([https://guides.rubyonrails.org/testing.html#the-test-database](https://guides.rubyonrails.org/testing.html#the-test-database))을 사용한다.
