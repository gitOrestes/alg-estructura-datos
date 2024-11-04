//* FOR EACH

const numeros = [5, 2, 4, 6, 1, 25, 6]

numeros.forEach( e => console.log(e))

numeros.forEach( (e, i) => console.log(e + 'tiene el indice ' + i))

numeros.forEach( (e, i) => e % 2 === 0 ? console.log("par") : console.log(impar))

//* MAP

let dobles = numeros.map(e => (e * 2))

console.log(dobles)

//* FILTER

const impares = numeros.filter(e => e % 2 !== 0)
console.log(impares)

//* REDUCE

sum = numeros.reduce((acc, elem) => acc + elem, 100)
console.log(sum)

//* FIND

let encontrado = numeros.find(e => e % 2 === 0)
console.log(sum)