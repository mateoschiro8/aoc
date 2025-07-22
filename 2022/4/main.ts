import { readFileSync } from "fs";
import { join } from "path";
const input: string[] = readFileSync(join(__dirname, "input.txt"), { encoding: "utf-8" }).split("\n");

const a: number = performance.now();

let score1 = 0, score2 = 0;

input.forEach(l => {
    const [w, x, y, z]: string[] = l.match(/\d+/g) as string[];
    const [a, b, c, d]: number[] = [Number(w), Number(x), Number(y), Number(z)];
    if((a <= c && b >= d) || (a >= c && b <= d)) {
        score1++;
        score2++;
    } else if((a <= c && b >= c) || (a <= d && b >= d))
        score2++;
})

const b: number = performance.now();

console.log("Parte 1: " + score1);
console.log("Parte 2: " + score2);
console.log(`Tomó ${b - a} ms`)