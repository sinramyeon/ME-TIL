
# 서버 모니터링 설정하기
---

> 사용할 프로그램 :  Grafana, Telegraf, Influxdb


## 설치

### Grafana 설치

Grafana는 예쁜 메트릭 분석 및 시각화 오픈소스이다. InfluxDB부터 prometheus, glances까지 다양한 DB를 지원한다. 코드 한 줄 없이도 예쁜 패널을 만들 수 있다.

![](http://7ktqal.com1.z0.glb.clouddn.com/img/blog/grafana-demo2.png)
> 아주 예쁘고 뭔가 멋져보이는 Grafana 보드

Grafana 설치는 아주 간편한데, Docker라면 이미지만 받아서 설치하면 끝난다.

1) Docker 의 경우
`docker pull grafana/grafana`
`docker run -d --name=grafana -p 3000:3000 grafana/grafana`

Grafana 설정을 위해 아래와 같이 시작할 수도 있다.

```
docker run \
  -d \
  -p 3000:3000 \
  --name=grafana \
  -e "GF_SERVER_ROOT_URL=http://grafana.server.name" \
  -e "GF_SECURITY_ADMIN_PASSWORD=secret" \
  grafana/grafana
```

grafana([http://localhost:3000/](http://localhost:3000/))에는 root/root로 로그인하면 된다.

2) 윈도우 / 맥 / 리눅스의 경우

[윈도우](http://docs.grafana.org/installation/windows/), [맥](http://docs.grafana.org/installation/mac/) 및 [리눅스](http://docs.grafana.org/installation/debian/)는 공식 홈페이지를 참고하면 쉽게 설치된다.
 

### Telegraf 설치

Telegraf란 데이터 콜렉션을 위한 plugin-driven server agent 라고 홈페이지에 써져 있다.  다양한 input, output plugin을 지원해서 어떤 데이터를 이용할 때에도 유용하다.

Telegraf 설치는 더 쉽다.

1) Docker의 경우
`docker pull telegraf`

2) Linux의 경우

```
sudo yum install telegraf

serivce telegraf start
```

3) 윈도우 / 맥

윈도우와 맥 역시 [문서](https://docs.influxdata.com/telegraf/v1.5/introduction/installation/)를 참조해 설치하기 쉽다.


### InfluxDB 설치

InfluxDB는 시간별로 데이터를 저장하는 시계열 데이터베이스 중 하나이다. 보통 저 Grafana, telegraf, influxDB가 하나의 세트로 쓰이곤 한다. 여담으로 InfluxDB와 Telegraf 두개 다 Go언어로 개발되었다고 한다.

1) Docker에 설치
`docker pull influxdb`
사실 이것들 전부를 하나로 묶은 도커 이미지도 많기 때문에 그걸 하나 받아 써도 편리하다.
그럴 경우에는 다른 이미지를 사용해도 좋다.

2) 리눅스에 설치

```
bash
curl -sL https://repos.influxdata.com/influxdb.key | sudo apt-key add -
source /etc/lsb-release
echo "deb https://repos.influxdata.com/${DISTRIB_ID,,} ${DISTRIB_CODENAME} stable" | sudo tee /etc/apt/sources.list.d/influxdb.list

sudo apt-get update && sudo apt-get install influxdb

sudo service influxdb start
```

3) 윈도우, 맥에 설치

[공식 문서](https://docs.influxdata.com/influxdb/v1.5/introduction/installation/)가 잘 되어 있다.

---

## 설정

설정을 할 때는 순서대로 해야 한다. InfluxDB에서 받아올 정보를 저장하는 곳을 만든 후, Telegraf에서 그 저장소로 값을 보내고, Grafana에서 InfluxDB에 있는 값을 얻어오면 완성이다.

### InfluxDB설정

InfluxDB에서 Telegraf를 통해 값을 받아올 DB를 만들어야 한다.

```
$ influx

Connected to http://localhost:8086 version 0.9

InfluxDB shell 0.9

>

\> CREATE DATABASE 데이터베이스 이름
```

보통 InfluxDB 쉘은 8086 포트로, 또 InfluxDB web 은 8083 포트로 열려있다.


![](https://1.bp.blogspot.com/-P4WC86MJ8Ms/VzXol_XuFlI/AAAAAAAANNs/GOg3II4reLsm5MYcQXWQ5B8xAhfjjSk1ACLcB/s1600/11.png)
`127.0.0.1:8083` 으로 접근하면 위 창을 확인할 수 있다.


### Telegraf 설정

Telegraf에서 설정한 InfluxDB로 값을 보내야 합니다.

1.  /etc/telegraf 안에 telegraf.conf파일을 찾습니다.
2.  이 부분을 찾아 아까 내 아이피로 변경합니다.
   ```
    \[\[outputs.influxdb\]\]  
    #urls = \["[http://아까 influxdb에 설정한 내 ip:8086](http://ip-or-name-of-influxdb:8086/)"\] # required  
    urls  =  \["[http://10.10.10.20:8086](http://10.10.10.20:8086/)"\]  \# required  
    database  =  "telegraf"  \# required
```


### Grafana 설정

Grafana에 `root/root` 로 로그인합니다.  
`localhost:3003` 으로 접속하면, grafana datasource를 설정할 수 있습니다.

![](http://docs.grafana.org/v3.1/img/v2/add_KairosDB.jpg)

데이터소스를 추가합니다. InfluxDB와 Telegraf에 설정한 내용대로 입력하고 더하면 됩니다.
Telegraf의 경우 Type는 동일하게 Influxdb로 하고 밑에 database 이름을 telegraf로 해 주시면 됩니다.

---



