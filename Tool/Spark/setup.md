# spark 막간 : 로컬에 스파크 설치하기

---

물론 docker로 많은 잘 감싸져있는 이미지들이 많다. 그렇지만 윈도우에서 docker가 잘 안된다(나의 컴퓨터에서만 그럴수도 있다). 그러므로 vm 위에 스파크 및 쥬피터 노트북, 파이썬까지 설치해 보겠다.

#### 버추얼박스 설치하기

1. [https://www.virtualbox.org/](https://www.virtualbox.org/)에서 버추얼박스를 깐다. user manual 이 있기 때문에 혹시 문제가 있다면 참고하면 된다.

2. [https://www.ubuntu.com](https://www.ubuntu.com)에서 우분투를 설치한다. 쓰기 쉽게 Ubuntu Desktop으로 설치한다. iso 파일로 받아질 것이다.

3. 버추얼박스를 실행 후, New로 새 os를 만든다. Type은 Linux, Version은 Ubuntu64bit(컴퓨터가 32bit라면 32로) 를 설정한 후, 메모리 사이즈는 자기 컴퓨터를 믿고 설정한다. 하드디스크는 정해진 대로 설정해도 상관없는데 우리는 Fixed size로 하드디스크를 설정했다.

4. 프로세스가 완료되면 가상으로 만들어진 자신의 os를 확인해볼 수 있는데, settings탭에서 프로세서 등을 설정할 수 있다. 데이터 분석을 하기 위해서라면 좀 늘려서 적용해주는게 정신건강에 이롭다.

5.  start를 누르고, select start-up disk가 뜨면 아까 다운받아둔 ubuntu desktop을 선택하고 start를 누른다. 이렇게 하면 ubuntu가 내가 만들어 둔 버추얼박스에 설치된다.
6.  ubuntu 기본 설정을 한 후, (만약 ubuntu설치가 처음이라면 검색하면 많이 나온다) 파이썬, 쥬피터, 스파크를 설치하면 된다.


#### 파이썬과 쥬피터, 스파크 설치하기

스파크 2.2는 파이썬 3.6을 지원하지만... 미래에 까실 분들은, 꼭 파이썬 버전과 스파크 버전이 혼용되는지부터 확인하셔야 한다.

1. ubuntu에는 python3이 깔려있으니 깔필요가 없다.
2. 그렇지만 쥬피터를 깔아야 한다.
3. `pip3 install jupyter`로 깔려고 하면 아마 pip이 없을것이다.
4. `sudo install python3-pip`로 pip부터 설치한다.
5. 다시 jupyter를 설치한다.
6. `jupyter notebook`으로 쥬피터 노트북 가동을 확인한다. 쥬피터 가동 시 로그인 창이 뜨는 경우, 커맨드 창에 뜨는 토큰을 사용하면 된다.
7. 다음에는 자바를 설치해야 한다. `sudo apt-get update` 적용을 한 후, `sudo apt-get install default-jre`로 자바를 받는다.
8. `java-version`이 성공하면 자바가 설치된 것이다.
9. 이번엔 스칼라도 깔아야만 한다.... `sudo apt-get install scala`로 설치한 후, `scala-version`으로 설치를 확인한다.
10. 자바, 스칼라, 파이썬을 연결하는 `pip3 install py4j` 도 설치한다
11. 이제 하둡과 스파크만 설치하면 된다. 우분투 데스크탑 내에서 인터넷을 켜다가 스파크, 하둡 최근 릴리즈를 받아온다.
12. `sudo tar -zxvf spark파일하고 hadoop파일` 로 압축을 푼다.
13. `export SPARK_HOME="아까 압축 푼 주소"`로 꼭 패스 설정을 해 준다.
14. `export PATH=$SPARK_HOME:$PATH` 로 패스에 추가한다.
15. `export PYTHONPATH=$SPARK_HOME/python:$PYTHONPATH`로 파이썬 설정 또한 한다.
16. `export PYSPARK_DRIVER_PYTHON="jupyter"` 로 쥬피터 노트북을 파이스파크 드라이버로 잡고, `export PYSPARK_DRIVER_PYTHON_OPTS="notebook"` 과 `export PYSPARK_PYTHON=python3`으로 파이썬3을 파이스파크에 쓰겠다고 명시한다.                                                                                                                                                                                                                                                            
17.  혹시 모를 퍼미션 에러가 있을 수도 있으니 관련 파일과 폴더는 모두 `chmod 777`로 권한을 변경해 둔다.
18. 해당 폴더의 내에서 파이썬을 구동 시, `import pyspark`가 가능하다면 성공한 것이다.


(도커가 잘 돌아간다면 그냥 도커로 묶인 이미지 하나를 받아서 심신건강을 찾는것이 좋다.)

### 파이스파크 셋업 오류 해결하기

파이스파크, 제플린, 기타등등을 연동하면서 큰 고통에 부착할 때가 있다.

#### 파이썬 버전 에러 발생 시

```
Exception: Python in worker has different version 2.7 than that in driver 3.6, PySpark cannot run with different minor versions.Please check environment variables PYSPARK_PYTHON and PYSPARK_DRIVER_PYTHON are correctly set.
```

보통 파이썬 2, 3이 공존하는 환경일 때 이러한 에러가 발생한다. 우리의 경우에는 파이썬 2.대에서 3.대로 업데이트하면서 이런 에러가 발생했는데, 파이썬 2. 대의 경로와 환경설정을 3.대로 옮겨왔는데도 계속 쓰고 있었기 때문이다.

이럴 때는 `/svc/zeppelin/conf/zeppelin-env.sh ` 를 확인하자.

그리고 `export PYSPARK_PYTHON=python3` `export PYSPARK_DRIVER_PYTHON=python3` 로 파이썬 경로를 명시해 준다.

또 `/spark/conf/spark-env.sh` 에도 동일 설정을 넣어준다.

*만약, 이렇게 해도 안 될 시?*

python3대신 python3의 절대 경로를 넣는다.(관련 [참고문서](https://wwija.com/computer-internet-technology/3864536_using-pyspark-in-zeppelin-with-python3-on-spark-2-1-0.html))

제플린에서 기동을 확인하고 싶다면, 버전을 제플린에서 확인해보자.

```
%pyspark

import sys
sys.version_info
```

제대로 자신이 선택한 파이썬 버전인 `sys.version_info(major=3, minor=6, micro=3, releaselevel='final', serial=0)` 가 나오면 정상이다.

#### py4j 에러 혹은 파이썬 자체를 못찾을 때

```
Py4JJavaError: An error occurred while calling o203.showString. : org.apache.spark.SparkException: Job aborted due to stage failure: Task 1 in stage 47.0 failed 4 times, most recent failure: Lost task 1.3 in stage 47.0 (TID 6719, data06.skylake, executor 59): java.io.IOException: Cannot run program "python3": error=2, No such file or directory at java.lang.ProcessBuilder.start(ProcessBuilder.java:1048)
```

우리는 이러한 에러가 났고, 다른 파이스파크 쉘에서는 잘 작동하는데 몇 쉘에서만 되지 않는 오류가 발생했다. 특히 RDD연산을 하는 행에서 제대로 실행이 되지 않았다.

이런 경우에는
1. 파이썬 설치 폴더 권한을 777로 준다.
2. PYSPARK_PYTHON 경로를 다시 잘 본다
3. 만약 클러스터 모드로 돌리고 있다면, 해당 노드 내에 모두 파이썬을 깐다.

이 해결방법은 [여기](https://stackoverflow.com/questions/45498049/encountered-a-error-that-cant-run-program-on-pyspark)에서 나왔으며, 위와 같은 행위를 다 해도 못 찾는다고 할 때가 있다.

그럴땐 혹시 zeppelin자체의 문제가 아닌지 확인한다. 구형 zeppelin에서 이와 비슷한 [문제](https://community.hortonworks.com/questions/16436/cants-get-pyspark-interpreter-to-work-on-zeppelin.html)가 발생한적이 있었기 때문이다.

1. 제플린 우측 상단 바 내 `Interpreter` 설정에 들어가, Spark 설정을 찾는다.
2. Spark Interpreter 환경 설정 중, `zeppelin.pyspark.python` 부문이 혹시 `python`혹은 `python3`으로 되어 있다면, 파이썬이 깔린 절대 경로의 위치로 변경한다.



#### 결론

보통 어느 에러가 발생하던 경로나 환경설정 문제이기 때문에, 첫째로는 PATH, 둘째로는 권한, 셋째로는 env나 sh파일, 넷째로는 웹 UI의 설정을 확인하면 대다수는 해결된다.
