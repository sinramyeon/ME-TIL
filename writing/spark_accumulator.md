[스파크 기초 시리즈]스파크 어큐뮬레이터

스파크를 이미 사용하고 계시다면, 특히 인터랙티브 쉘(쥬피터, 제플린)에서 사용하고 있으신 분들이라면, 이게 불편하지 않으실까 싶다.

분석내용을 변수에 담아놨는데, 같은 연산을 계속 반복해야 하는 경우가 분명 생기고 있을 것이다. 10초, 20초만 기다려도 분노 조절 장애가 오거늘... 성능 최적화를 안 한 내 스파크 코드는 무려 10분간 돌아가고 있다. 그런데 이걸 또 계산하라고? 그런 짓은 절대로 빅데이터를 다룰 때 하지 말아야 할 짓거리 중 하나이다.



스파크 map(), filter()연산시 함수를 넘길 때 변수를 쓰는데, 변수 복사본을 쓰는 것에 지나지 않으므로 변경사항이 날아가버린다. 

이럴 때 공유 변수인 어큐뮬레이터와 브로드캐스트 변수를 사용해야 한다.



accumulator
스파크의 공유 변수 중 하나이며, sc.accumulator()를 이용하여 불러온다.

어큐물레이터 메서드 첫 번째 매개변수에는 초기값을, 두 번째 매개변수에는 이름을 설정할 수 있다.

어큐뮬레이터로 자주 하는 작업의 예를 들어보도록 하겠다.



작업 수행 중 발생하는 일 세기
file = sc.textFile(inputFile)
blankLines = sc.accumulator(0)
# 0으로 초기화한 인트형 어큐뮬레이터 선언

def extractCallSigns(line) : # 콜 사인 리스트를 로그에서 읽어온다
    global blankLines
    if (line == "") : # 공백 체크
        blankLines += 1
    return line.split(" ")

callSigns = file.flatMap(extractCallSigns)
callSigns.saveAsTextFile(outputDir + "/callsigns")
print(blankLines.value)
# 공백라인 찾기
이 코드에서 유의할 점은, 올바른 값을 saveAsTextFile()이후에나 알 수 있다는 것이다. 어큐뮬레이터 값 증가는 map()이 실제로 실행되는 저장단계에서 일어나기 때문이다. 지연 평가를 항상 잊지 말자.



어큐뮬레이터의 동작 단계는 아래와 같다.



1. 드라이버에서 SparkContext.accumulator(initialValue)를 호출해서 초기값을 가진 어큐뮬레이터를 만든다.

반환타입은 org.apache.spark.Accumulator[T]객체고 T는 선언한 초기값 타입이다.

2. 스파크 클로져 작업 노드에서 어큐뮬레이터에 값을 더했다.

3. 드라이버 프로그램에서 value속성으로 어큐뮬레이터 값에 접근한다.



다른 예제를 하나 더, 어큐뮬레이터로 에러를 세는 예제를 봐보겠다.



validSignCount = sc.accumulator(0)
invalidSignCount = sc.accumulator(0)
# 어큐뮬레이터를 일종의 쓰기전용 변수라고 생각하자.

def validateSign(sign) :
    global validSignCount, invalidSignCount
    if re.match("정규식", sign) :
        validSignCount += 1
        return True
    else :
        invalidSignCount += 1
        return False

validSigns = callSigns.filter(validateSign)
contactCount = validSigns.map(lambda sign : (sign, 1)).reduceByKey(lambda (x,y) : x+y)

# 연산 시도
# 어큐뮬레이터는 이때서야 가동 된다.
contactCount.count()


스파크는 느려진거나 장애가 발생한 애들은 다시 재실행해서 고치려고 한다. 아무 이상이 없어도 메모리가 꽉 차가서 다시 재실행 할 수도 있다.



그 와중 각 태스크 업데이트는 어큐뮬레이터에 한번씩만 반영된다.

 즉 장애나 반복연산이 일어났을 때 전 값을 갖고 있을 수도 있으므로, 이럴때 꼭 foreach()등을 사용해서 제대로된 값을 얻어야한다.



accumulator 이용시의 주의점을 알아보자.

1. 가산 시

변환하려다 만든 rdd를 여러곳에서 참조하다가 여기저기서 더하는 바람에 중복 가산이 생겼다. 이럴때 꼭 foreach로 안전하게 가산하자.



2. 참조 시

어큐뮬레이터 값을 함수 내에서 참조하다가, 로컬값만 참조해서 이상한 결과가 나왔다. 참조할 때 내가 지금 어디값을 갖다쓰는건지 두번 세번 확인하자.



