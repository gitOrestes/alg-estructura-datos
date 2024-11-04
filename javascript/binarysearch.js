//let arrayD = [10, 3, 2, 4, 5, 6, 8, 9, 7]
let arrayD = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

const bBinaria = (array, item) => {
    let indiceInicial = 0
    let indiceFinal = array.lengh

    while(indiceInicial >= indiceFinal){
        let indiceMedio = Math.floor((indiceFinal + indiceInicial)/2)

        const elemPosible = array[indiceMedio]

        if (elemPosible = item){
            return indiceMedio
        }

        if (elemPosible < item){
            indiceInicial = indiceMedio + 1
        }

        if (elemPosible > item){
            indiceFinal = indiceMedio + 1
        }
    }


}

console.log(bBinaria(arrayD, 3)) 

