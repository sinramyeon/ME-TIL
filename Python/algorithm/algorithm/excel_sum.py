# excel_sum.py

"""
32,232
48,329
848,438
식의 엑셀 셀들을 더해주는 함수
엑셀 쓸때 sum함수 값을 다시한번 확인할때 유용
인쇄돼는 값을 셀에 다시한번 쳐주면 됨
"""
def excel_sum(num) :
    num = num.replace("\n", "+").replace(",", "").strip()
    print(num)