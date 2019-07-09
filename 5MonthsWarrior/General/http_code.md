# http status code

 비고 | 클래스 | 설명
 --- | ---- | ---
 1xx | Informational | 리퀘스트를 받아들여 처리중
 2xx | Success | 리퀘스트 정상 처리
 3xx | Redirection | 리퀘스트를 완료하기 위해 추가적 동작이 필요
 4xx | Client Error | 서버는 리퀘스트 이해 불가능
 5xx | Server Error | 서버는 리퀘스트 처리 실패
 
 ___
 
 2xx 성공
 200 OK
 : 리퀘스트 정상 처리
 204 No Content
 : 정상 처리했으나 돌려줄 리소스가 없음
 206 Partial Content
 : 정상 처리했으며 범위가 지정된 리퀘스트에 의해 부분적 GET 리퀘스트
 
 ---
 
 3xx 리다이렉트
 
 301 Moved Permanently
 : 새 URI가 부여되었으므로 그 URI 이용 필요
 302 Found
 : 일시적 이동
 303 See other
 : 다른 URI로 GET
 304 Not Modifed, 307 Temporary Redirect 등
 
 ---
 
 4xx 클라이언트 에러
 
 400 Bad Request
 : 리퀘스트 구문 에러
 
 401 Unauthorized
 : HTTP 인증정보 필요
 
 403 Forbidden
 : 리퀘스트된 리소스의 액세스 거부
 
 404 Not Found
 : 서버에 그딴주소가 없음
 
 ---
 
 5xx 서버 에러
 
 500 Internal Server Error
 : 리퀘스트 처리 도중 에러
 
 503 Service Unavaliable
 : 일시적 과부하...
 
