# 京东商城搜索框联想

爬取京东商城搜索框输入的字的联想匹配

## 1.运行方式

```
go run ./main.go -kw gtx -proxy 127.0.0.1:1234
kw为搜索的关键字
proxy为代理(如果被反爬的话，走clash或者ssr的代理端口)
```

实例

```
>>go run ./main.go -kw gtx -proxy 127.0.0.1:1234
设置代理为: 127.0.0.1:7890
jQuery24400509([{"key":"gtx1650super"},{"key":"gtx1660super"},{"key":"gtx1060"},{"key":"gtx780"},{"key":"gtx750"},{"key":"gtx1650"},{"k
ey":"gtx1660"},{"key":"gtx1660ti微星"},{"key":"gtx1660ti"},{"key":"gtx1050"},{"key":"gtx1060ti"},{"key":"gtx1060 6g显卡"},{"key":"gtx76
0"},{"version":"SAK7|MIXTAG_SAK7R,SAK7_M_AM_L5369,SAK7_M_COL_U17677,SAK7_S_AM_R,SAK7_SC_PD_R,SAK7_SM_PB_L16675,SAK7_SS_PM_R|"}])  
```

