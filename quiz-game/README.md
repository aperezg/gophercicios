# Ejercicio #1: Quiz Game

Este ejercicio consiste en crear un juego rápido de preguntas y respuestas.

## Detalles

El ejercicio se dividirá en dos partes para que sea más sencillo de explicar y de resolver. La segunda parte
será algo más complicada que la primera, así que si ves que te estancas ahí siéntete libre de hacer otro ejercicio
y volver sobre la segunda parte en otro momento.

### Parte 1

Crear un programa que leerá ciertas preguntas y respuestas proporcionadas mediante un fichero CSV, más detalles abajo, al usuario
le mostraremos solo las preguntas, y le informaremos de cuantas ha acertado y cuantas ha fallado. Aunque se haya equivocado la siguiente
pregunta debe ser mostrada indmediatamente.

El fichero CSV por defecto será `problem.csv` (con los datos de ejemplo mostrados abajo), pero el usuario puede pasar otro fichero
vía `flag`.

El fichero en cuestión tendrá el siguiente formato, separado por comas, donde la primera columna será la pregunta y la segunda la respuesta.

```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

Podemos asumir que será un test relativamente corto (menos de 100 preguntas) y con una sóla palabra/número como respuesta.

Al final del programa el test debe mostrar el total de preguntas correctas, las preguntas inválidas o no respondidas serán consideradas incrrectas.

**¡OJO CUIDADO!:** *El fichero CSV va separado por comas por defecto, así que sí escribimos una pregunta cómo, `"¿cuánto es 2+2, amigo?",4` nos lo dividirá en 3 parámetros en vez de 2. Os sugerimos echar un ojo al [paquete CSV](https://golang.org/pkg/encoding/csv/) y no escribir vuestro propio parser ;).*

### Parte 2

Vamos a poner las cosas un poco más díficiles, para ello adaptaremos nuestro código hecho en la parte 1, para añadir un temporizador.
El tiempo límite por defecto será de **30 segundos**, pero debería ser personalizable vía `flag`. 

Tu test debería parar en el mismo momento que ese tiempo haya sido alcanzado. Es decir, no debes esperar a que el usuario introduzca su respuesta, debe de pararse justo se alcance el tiempo y mostrar los resultados.

Esta vez como utilizaremos un tiempo, el usuario deberá de presionar enter (o cualquier otra tecla) antes de poner en marcha el contador.

Al final del test el programa debería de mostrar el total de preguntas correctas y cuantas preguntas hay en total, las preguntas no respondidas o invalidas serán consideradas incorrrectas.

## Bonus

Si aún no te has quedado satisfecho te proponemos un bonus...

1. Añadir un trimming de string y limpiar las entradas del usuario para asegurar que la pregunta sigue siendo válida a pesar de espacios extras, mayúsculas, etc. *Pista, echar un ojo al paquete [strings](https://golang.org/pkg/strings/).*
2. Añadir un `flag` nuevo que permita mostrar las preguntas desordenadas en cada ejecución.

## Ejercicio Original

Este ejercicio esta basado en la idea original de [Gophercises](https://gophercises.com/) creada por [Jon Calhoun](https://twitter.com/joncalhoun). El ejercicio original lo podéis encontrar [aquí](https://github.com/gophercises/quiz).