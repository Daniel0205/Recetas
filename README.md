# Requisitos:
+ CockroachDB  v19.1.0 <br />
+ Manejador de paquetes npm v6.9.0 <br />
+ go version go1.11.9 <br />

# Configuracion de CockroachDB 
+ Ejecutar el siguiente comando en la consola:<br />
>`cockroach start --insecure --listen-addr=localhost:26257`

+ Despues en una nueva consola ejecutar:<br />
>`cockroach start 
--insecure 
--store=node2 
--listen-addr=localhost:26258 
--http-addr=localhost:8081 
--join=localhost:26257`
  
+ Estos comandos empezaran a correr los nodos donde se encontrar la base de datos, por lo tanto deben estar corriendo en simultaneo con al aplicacion.


# Configuracion de la base de datos
+ En una nueva consola ejecutar los siguientes comandos:<br />
> `cockroach sql --insecure`<br />
> `CREATE DATABASE IF NOT EXISTS recetas;`<br />
> `set database = recetas;`<br />

+ Y finalmente ejecutar los scripts para crear la base de datos, que se encuentran en el archivo [scripts.sql](https://github.com/Daniel0205/Recetas/blob/master/scripts.sql) del repositorio
	
  
# Ejecutar API-REST
+ Ubicarse en la carpeta '/api REST' del proyecto y ejecutar los comandos:
> `go build main.go`<br />
> `./main`<br />

+ Este ultimo comando dejara el api-rest esperando  en el puerto http://localhost:3000/ por peticiones.
+ El api rest debe estar corriendo en simultaneo con la aplicacion para poder responder las peticiones del cliente.

# Ejecutar la aplicación web
-En una nueva consola ubicarse en la carpeta '/app' y ejecutar los siguientes comandos:
> `npm install`<br />
> `npm run dev`<br />

Finalmente este ultimo comando, iniciara la aplicacion en el navegador.

