#  [HTML 데이터 추출 방법(python)]


-  필요한 모듈

```
import html2text
from bs4 import BeautifulSoup

h = html2text.HTML2Text()
h.ignore_links = True

# 특정길이를 넘어갈 경우 줄바꿈이 들어가서 이를 방지하기 위해 넣음
h.body_width = 0
```

- 코드 활용 방법

```
try :
    # b -> html_text
    text = h.handle(b)

except Exception : # error 코드가 여러가지 나오므로
    # html문서중 약식으로 작성된 경우 오류 발생의 윟머이 있으므로 이를 처리하기 위해선 bs를 쓴다.
    soup = BeautifulSoup(b, "html.parser")
    sv = soup.get_text("\n", strip=True).encode("utf-8")
```
    
위 코드를 이용하면 html 데이터를 추출할 수 있다.

BeautifulSoup을 이용하면 오류 없이 처리가 가능하지만 줄바꿈이 원본과 다르게 될 수 있다.

예를들면 <span> tag가 있는 경우에 줄바꿈을 하게 되어 문제가 생길 수 있다.

따라서 위와 같이 처리를 하면 대부분의 html에 대해서 원하는 형태의 문장을 추출할 수 있다.
