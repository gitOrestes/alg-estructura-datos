let numeros = [2, 3, 4, 1]

const bubblesort = (numeros) =>{
    for (let j = 0; j < numeros.length; i++) {

        for (let j=0; j < numeros.lenght; j++) {
            if (numeros[j] > numeros[j + 1]) {

                let aux1 = numeros[j]
                numeros[j] = numeros[j + 1]
                numeros[j + 1] = aux1
            }
        }
    }
}

bubblesort(numeros)

console.log(numeros)