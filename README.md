# HTTP/2 Timeout Test

This minimal application comes to a complete halt as the vegeta report shows.

```sh
$ echo "GET https://localhost:3000/" | vegeta attack -insecure -duration=5s | vegeta report
Requests      [total, rate]            250, 50.20
Duration      [total, attack, wait]    34.980297934s, 4.979999591s, 30.000298343s
Latencies     [mean, 50, 95, 99, max]  17.880625607s, 30.000422117s, 30.000777718s, 30.001158261s, 30.002248396s
Bytes In      [total, mean]            1414, 5.66
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  40.40%
Status Codes  [code:count]             200:101  0:149  
Error Set:
Get https://localhost:3000/: http2: timeout awaiting response headers
```

This is also experienced in browsers, try running and refresh
https://localhost:3000/ a few times and it will just sit there.

Tested on Ubuntu and Debian linux/amd64 1.8beta2 and stable 1.7.4.
