```
curl   http://6ec69753.ngrok.io/offers                                                                     

[{"Name":"dsasas","Details":"sasas0","ImageURL":"https://via.placeholder.com/150","ID":"0"},{"Name":"dsasas","Details":"sasas1","ImageURL":"https://via.placeholder.com/150","ID":"1"},{"Name":"dsasas","Details":"sasas2","ImageURL":"https://via.placeholder.com/150","ID":"2"},{"Name":"dsasas","Details":"sasas3","ImageURL":"https://via.placeholder.com/150","ID":"3"},{"Name":"dsasas","Details":"sasas4","ImageURL":"https://via.placeholder.com/150","ID":"4"},{"Name":"dsasas","Details":"sasas5","ImageURL":"https://via.placeholder.com/150","ID":"5"},{"Name":"dsasas","Details":"sasas6","ImageURL":"https://via.placeholder.com/150","ID":"6"},{"Name":"dsasas","Details":"sasas7","ImageURL":"https://via.placeholder.com/150","ID":"7"},{"Name":"dsasas","Details":"sasas8","ImageURL":"https://via.placeholder.com/150","ID":"8"},{"Name":"dsasas","Details":"sasas9","ImageURL":"https://via.placeholder.com/150","ID":"9"}]
```

```
curl -XPOST http://localhost:3333/offers -d'{"Name":"dsasas","Details":"sasas0","ImageURL":"https://via.placeholder.com/150","ID":"0"}'
{"Name":"dsasas","Details":"sasas0","ImageURL":"https://via.placeholder.com/150","ID":"0"}
```


```
curl  -XDELETE   http://6ec69753.ngrok.io/offers/0                                                  
{"success":true}
```