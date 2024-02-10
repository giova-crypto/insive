package main

import (
    "fmt"
    "regexp"
    "sync"
)

// Función para generar todas las posibles claves
func generarClaves(chars []rune, length int) []string {
    var claves []string
    generarClavesRecursivo(chars, length, "", &claves)
    return claves
}

func generarClavesRecursivo(chars []rune, length int, clave string, claves *[]string) {
    if length == 0 {
        *claves = append(*claves, clave)
        return
    }
    for _, char := range chars {
        nuevaClave := clave + string(char)
        generarClavesRecursivo(chars, length-1, nuevaClave, claves)
    }
}

// Función para realizar la operación XOR entre una clave y el texto encriptado
func operacionXOR(clave string, mensajeEncriptado []int) []int {
    resultado := make([]int, len(mensajeEncriptado))
    for i, caracter := range mensajeEncriptado {
        resultado[i] = caracter ^ int(clave[i%len(clave)])
    }
    return resultado
}

// Función para validar si los resultados corresponden a caracteres que cumplen con la expresión regular
func validarResultado(resultado []int) bool {
    regex := regexp.MustCompile(`[a-zA-Z0-9\s.,@\-_\/]`)
    for _, caracter := range resultado {
        if !regex.MatchString(string(rune(caracter))) {
            return false
        }
    }
    return true
}

func main() {
    // Mensaje encriptado
    mensajeEncriptado := []int{32, 2, 31, 8, 5, 14, 23, 0, 2, 2, 0, 65, 32, 2, 1, 13, 3, 30, 83, 38, 15, 8, 5, 0, 8, 9, 10, 65, 53, 6, 29, 2, 14, 2, 9, 77, 70, 15, 18, 18, 70, 11, 28, 6, 20, 6, 23, 14, 70, 3, 22, 18, 5, 14, 21, 19, 7, 21, 83, 4, 10, 71, 30, 4, 8, 20, 18, 11, 3, 73, 83, 65, 39, 15, 28, 19, 7, 75, 83, 17, 7, 21, 18, 65, 23, 18, 22, 65, 47, 9, 0, 8, 16, 2, 83, 19, 3, 4, 28, 15, 9, 29, 16, 0, 70, 19, 6, 65, 10, 8, 20, 19, 9, 75, 83, 18, 19, 5, 22, 65, 3, 11, 83, 2, 9, 3, 22, 65, 5, 8, 29, 65, 3, 11, 83, 16, 19, 2, 83, 19, 3, 20, 28, 13, 16, 14, 0, 21, 3, 71, 22, 18, 18, 2, 83, 4, 12, 2, 1, 2, 15, 4, 26, 14, 70, 2, 29, 65, 33, 14, 7, 41, 19, 5, 92, 38, 15, 19, 63, 0, 4, 71, 10, 65, 5, 8, 30, 17, 7, 21, 7, 4, 70, 2, 31, 65, 3, 9, 31, 0, 5, 2, 83, 0, 70, 20, 28, 17, 9, 21, 7, 4, 38, 14, 29, 18, 15, 17, 22, 79, 5, 11, 93}

    // Caracteres para generar las claves
    caracteres := []rune("abcdefghijklmnopqrstuvwxyz")

    // Generar todas las posibles claves
    claves := generarClaves(caracteres, 4)

    // Multiples procesos realizando las operaciones XOR para aumentar el rendimiento, uno para cada clave
    var wg sync.WaitGroup
    resultadoCanal := make(chan map[string]string, len(claves))

    for _, clave := range claves {
        wg.Add(1)
        go func(clave string) {
            defer wg.Done()
            resultado := operacionXOR(clave, mensajeEncriptado)
            if validarResultado(resultado) {
                resultadoDescifrado := make([]rune, len(resultado))
                for i := 0; i < len(resultado); i++ {
                    resultadoDescifrado[i] = rune(resultado[i])
                }
                resultadoCanal <- map[string]string{clave: string(resultadoDescifrado)}
            }
        }(clave)
    }

    // Esperar a que todos los procesos terminen
    go func() {
        wg.Wait()
        close(resultadoCanal)
    }()

    // Guardar las llaves coincidentes y el resultado descifrado en un diccionario
    resultados := make(map[string]string)
    for resultado := range resultadoCanal {
        for clave, descifrado := range resultado {
            resultados[clave] = descifrado
        }
    }

    // Imprimir los resultados
    for clave, descifrado := range resultados {
        fmt.Printf("Llave encontrada: %s, Mensaje descifrado: %s\n", clave, descifrado)
    }
}