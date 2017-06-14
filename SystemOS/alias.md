##### alias 설정

---

1. 설정파일 열기

`vim ~/.bashrc`

2. alias 설정하기

```
# .bashrc

# User specific aliases and functions

alias rm='rm -i'
alias cp='cp -i'
alias mv='mv -i'
alias mstart='/opt/apache-tomcat-8.0.44/bin/startup.sh'
alias mstop='/opt/apache-tomcat-8.0.44/bin/shutdown.sh'
alias mlog='tail -f /opt/apache-tomcat-8.0.44/logs/catalina.out'

# Source global definitions
if [ -f /etc/bashrc ]; then
        . /etc/bashrc
fi
```
파일에서 alias를 추가합니다.

`alias 추가할 단어='실행할 파일'`

3. 설정파일 적용하기

`source ~/.bashrc`

