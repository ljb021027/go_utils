import requests
import time

url="https://www.baidu.com"

payload=[{'mesg': 'one'},{'mesg':'two'}]

s = requests.Session()
r = s.post(url, data=payload)
print(r.headers)
print(r.status_code)
print(r.content)