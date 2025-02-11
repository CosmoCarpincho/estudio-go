# OJO ver nomenclatura
https://www.reddit.com/r/golang/comments/rowff0/module_naming_conventions/?tl=es-es

- Idealmente una palabra. -> módulo/paquete SINO recomienda usar alias para nombre del paquete
- Los módulos pueden contener guiones -
- PERO los paquetes deben tener identificador válido : letra, dígito, guion bajo

Ej: 

github.com/user/my-module  el nombre de su paquete raísz podria ser my

``` go
import my "github.com/user/my-module"
```

https://go.dev/blog/package-names

https://go.dev/doc/effective_go#names

READ -> [go modules reference](https://gophers-latam.github.io/posts/2021/04/golang-referencia-esencial-go-modules/)