### 파일 삭제하기

---
파일, 디렉ㅌ리 삭제
#### rm


`rm 옵션 파일명`

삭제 확인 없이 삭제
`rm -f`

디렉토리를 삭제
`rm -r`

삭제 확인 없이 죄다 날려
`rm -rf`

---

날짜 기준으로 파일 삭제하기

`find 폴더 -name 파일명 -mtime + 일수 -delete`
`find 폴더 -name 파일명 -mtime + 일수 -exec rm -f {} \;`
`find 폴더 -name 파일명 -mtime + 일수 | xargs rm -f`

Ex) /backup/에서 3일 지난 파일 지우기
`find /backup/ -name "*.tgz" -mtime +2 -delete`

Ex2) 수

특정 파일만 골라 삭제하기

`find .-name "이름" -exec rm -f {} \;`
파일명엔 정규표현식 사용 가능
