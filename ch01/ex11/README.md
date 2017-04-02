# runall.bash

出力は、次の通りです。

```
0.20s     3154  http://t.co
0.24s    10672  http://google.com
0.26s    10596  http://google.co.jp
0.30s    13236  http://google.co.in
0.31s    10428  http://google.co.uk
0.28s    10490  http://google.it
0.29s    11009  http://google.com.hk
0.31s    11703  http://google.es
0.33s    10477  http://google.de
0.29s    10886  http://google.ca
0.37s     1204  http://pornhub.com
0.34s    91926  http://bing.com
0.35s     1204  http://livejasmin.com
0.42s    10498  http://google.fr
0.41s   138058  http://imgur.com
0.51s       81  http://baidu.com
0.63s   226015  http://stackoverflow.com
0.67s    19270  http://yahoo.co.jp
0.75s   105101  http://jd.com
0.77s    48424  http://msn.com
0.91s     1064  http://reddit.com
0.91s    61943  http://aliexpress.com
0.96s    12253  http://google.com.ru
0.97s    26028  http://wordpress.com
1.01s   320969  http://twitter.com
1.05s   254516  http://qq.com
1.16s   224073  http://tmall.com
1.21s    19456  http://instagram.com
1.33s   130998  http://weibo.com
1.43s   178967  http://ebay.com
1.45s   424235  http://yahoo.com
1.46s    81795  http://wikipedia.org
1.52s   619553  http://hao123.com
1.56s   116510  http://facebook.com
1.67s   600944  http://sina.com.cn
1.70s    71644  http://tumblr.com
1.79s   235260  http://amazon.com
1.82s   268353  http://netflix.com
2.04s    15643  http://live.com
2.40s   144196  http://ok.ru
2.41s    92982  http://microsoft.com
Get https://detail.tmall.com/?tbpm=3: stopped after 10 redirects
2.72s    56420  http://yandex.ru
3.11s     6563  http://vk.com
3.21s   431347  http://sohu.com
3.31s    10991  http://google.ru
3.52s    50316  http://taobao.com
4.19s    55772  http://linkedin.com
4.50s   549170  http://youtube.com
15.53s   278202  http://360.cn
15.53s elapsed
```

# あるウェブサイトが応答しない場合

あるウェブサイト (ここでは http://this-does-not-respond.example.com とします) が応答しない場合の出力は、次の通りです。

```
$ go run main.go http://google.com http://this-does-not-respond.example.com
0.31s    10664  http://google.com
Get http://this-does-not-respond.example.com: dial tcp this-does-not-respond.example.com:80: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
21.00s elapsed
```
