# Metrics-go: a minimal tool for memory metrics collection

Metrics-go is a lab project to investigate on the **runtime** package
features.
The project idea started from this article:
<https://scene-si.org/2018/08/06/basic-monitoring-of-go-apps-with-the-runtime-package/>

The initial basic idea was extended with the usage of separate goroutines to
get and print the output, an interface for future expand of the package and a demo
main funcion.

The main program just runs a dummy function that prints a simple message on 
standard output along with the metrics in a simple log format.

```
Starting memory stats collector.
Doing some crazy stuff to entertain you...
2018/08/11 17:53:19 {"Alloc":82968,"TotalAlloc":82968,"Sys":1740800,"Mallocs":202,"Frees":2,"LiveObjects":200,"PauseTotalNs":0,"NumGC":0,"NumGoroutine":3}
Doing some crazy stuff to entertain you...
2018/08/11 17:53:20 {"Alloc":105568,"TotalAlloc":105568,"Sys":1740800,"Mallocs":342,"Frees":24,"LiveObjects":318,"PauseTotalNs":0,"NumGC":0,"NumGoroutine":3}
Doing some crazy stuff to entertain you...
2018/08/11 17:53:21 {"Alloc":106480,"TotalAlloc":106480,"Sys":1740800,"Mallocs":348,"Frees":24,"LiveObjects":324,"PauseTotalNs":0,"NumGC":0,"NumGoroutine":3}
Doing some crazy stuff to entertain you...
2018/08/11 17:53:22 {"Alloc":107392,"TotalAlloc":107392,"Sys":1740800,"Mallocs":354,"Frees":24,"LiveObjects":330,"PauseTotalNs":0,"NumGC":0,"NumGoroutine":3}
Doing some crazy stuff to entertain you...
2018/08/11 17:53:23 {"Alloc":108304,"TotalAlloc":108304,"Sys":1740800,"Mallocs":360,"Frees":24,"LiveObjects":336,"PauseTotalNs":0,"NumGC":0,"NumGoroutine":3}
Stopping stats printer.
Stopping memory stats collector.
```

The package can be easily included in custom code for more exciting tests.

### TODO
- Adopt the **logrus** package
- Let the user print the metrics on a separate file
- Add tests

### Author
Gianni Salinetti <gbsalinetti@extraordy.com>

