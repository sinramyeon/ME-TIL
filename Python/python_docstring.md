### docstring

---

설명 문자열(documentation string)을 알아보자!

```

```이것이다(쌍따옴표 세개)```
help(어쩌구)

```help()함수로 docstring 내용이 확인 가능하다.```


```

Ex)

```
def approximate_size(size, a_kilobyte_is_1024_bytes=True):
    '''Convert a file size to human-readable form.

    Keyword arguments:
    size -- file size in bytes
    a_kilobyte_is_1024_bytes -- if True (default), use multiples of 1024
                                if False, use multiples of 1000

    Returns: string

    '''
    
# 이렇게 하면 help(approximate_size) 시 저 내용이 출력된다.


```