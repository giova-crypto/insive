# Desencriptador XOR en Go
Este programa en Go fue diseñado para desencriptar un mensaje que fue cifrado utilizando la puerta lógica XOR. El mensaje encriptado se proporciona junto con algunas restricciones sobre la llave utilizada para el cifrado y el formato del mensaje encriptado.

## Requisitos
Go instalado en tu sistema. Puedes descargar e instalar Go desde su sitio oficial.
## Ejecución
Clona este repositorio [gitrepo](https://github.com/giova-crypto/insive) en tu máquina local:

`git clone https://github.com/giova-crypto/insive.git`
Navega al directorio del proyecto

## Ejecuta el programa:
`go run main.go`

El programa generará todas las posibles llaves, realizará la operación XOR entre cada llave y el mensaje encriptado utilizando múltiples procesos para aumentar el rendimiento, validará si los resultados corresponden a caracteres que cumplen con la expresión regular dada y finalmente guardará las llaves coincidentes y el resultado descifrado en un diccionario.

## Resultados
Una vez que el programa haya terminado de ejecutarse, verás los resultados en la salida estándar. Cada resultado mostrará la llave encontrada y el mensaje descifrado correspondiente.

# Proceso de pensamiento
En el camino a resolver este ejercicio conté con algunos pensamientos iniciales y algunos que siguieron en respuesta a inquietudes o problemas encontrados.

1. Buscar la llave
Dado que para desencriptar el mensaje era necesario encontrar la llave intenté algunas cosas como:
* Usar algo de ingeniería social para tratar de encontrar la llave en mensajes o conversaciones mantenidas. 
    Contras:
    * Resultaba difícil pensar en conversaciones mantenidas por palabras que hicieran match con el criterio además de poco probable que la llave fuese alguna palabra normal.
* Generar todas las llaves posibles
    * Pros:
        * No había posibilidad de no encontrar la llave.
    * Contras:
        * Las combinaciones de caracteres resultantes serían más de 400.000 siendo así que tendría que probar cada una en busca de la correcta y esto consumiría mucho tiempo.
2. Implementar una solución
Al saber como generar las llaves y conocer el proceso para desencriptar el mensaje sólo debía buscar código para interar sobre cada llave, operar entre esta y el mensaje, guardar el resultado y así hasta probar todas las llaves.

Sin embargo, esto conllevaba a algunos problemas:
* El tiempo sería algo alto en un proceso lineal de ejecución.
* Tendría que validar las más de 400.000 posibles soluciones para encontrar las que hiciesen sentido y no había una forma en cuanto a código de validar que una sentencia fuese gramaticalmente correcta (al menos no por lo que pude pensar y consultar).

Las soluciones a los problemas:
* Podría mejorar el rendimiento del programa ejecutando cada operación de desencripción del mensaje en un proceso diferente. Para esto elegí usar go y el package sync.
* Para elegir los posibles resultados correctos, al investigar, miré que lo que podría hacer en lugar de mirar el mensaje como un todo podría ser tomar cada caracter del mensaje y validar si este hacía match con la expresión regular dada en el ejercicio. Así entonces, guardaría todos los posibles mensajes coincidentes y su llave para una última evaluación manual.
