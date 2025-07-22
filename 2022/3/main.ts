import { readFileSync } from "fs";
import { join } from "path";
const input: string = readFileSync(join(__dirname, "input.txt"), { encoding: "utf-8" })

const a: number = performance.now();

const priority = (c: string): number => {
    const code = c.charCodeAt(0);
    return code <= 90 ? code - 38 : code - 96;
}

let score1 = 0, score2 = 0;
const sacks = input.split("\n")

const checkSack = (rs: string): void => {
    const firstHalf: string = rs.slice(0, rs.length/2);
    const secondHalf: string = rs.slice(rs.length/2, rs.length);
    for(let i = 0; i < firstHalf.length; i++) {
        if(secondHalf.includes(firstHalf[i])) {
            score1 += priority(firstHalf[i]);
            break;
        }
    }
}

for(let i = 0; i < sacks.length; i += 3) {
    const [firstSack, secondSack, thirdSack] = [sacks[i], sacks[i + 1], sacks[i + 2]];
    
    checkSack(firstSack); checkSack(secondSack); checkSack(thirdSack);

    for(let j = 0; j < firstSack.length; j++) {
        if(secondSack.includes(firstSack[j]) && thirdSack.includes(firstSack[j])) {
            score2 += priority(firstSack[j]);
            break;
        }
    }
}

const b: number = performance.now();

console.log("Parte 1: " + score1);
console.log("Parte 2: " + score2);
console.log(`Tomó ${b - a} ms`)