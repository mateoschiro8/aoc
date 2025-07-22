import { readFileSync } from "fs";
import { join } from "path";
const input: string = readFileSync(join(__dirname, "input.txt"), { encoding: "utf-8" })

const a: number = performance.now();

let score1 = 0, score2 = 0;
for(let i = 0; i < input.length - 14; i++) {
    let set1 = new Set(), set2 = new Set();

    for(let k = 0; k < 4; k++)
        set1.add(input[i + k]);
    for(let k = 0; k < 14; k++)
        set2.add(input[i + k]);
    
    if(set1.size == 4 && score1 == 0) 
        score1 = i + 4; 
    if(set2.size == 14 && score2 == 0) 
        score2 = i + 14; 
}

const b: number = performance.now();

console.log("Parte 1: " + score1);
console.log("Parte 2: " + score2);
console.log(`Tomó ${b - a} ms`)