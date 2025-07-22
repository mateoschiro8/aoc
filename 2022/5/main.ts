import { readFileSync } from "fs";
import { join } from "path";
const input: string = readFileSync(join(__dirname, "input.txt"), { encoding: "utf-8" })

const a: number = performance.now();

const [crates, proc]: string[] = input.split("\n\n");

let score1 = "", score2 = "";
let stacks1: string[][] = [], stacks2: string[][] = [];

const stacksNumber: number = Number(crates[crates.length - 2]);
for (let i = 0; i <= stacksNumber; i++) {
    stacks1.push([]);
    stacks2.push([]);
}

crates.split("\n").slice(0, -1).forEach(l => {
    for(let i = 1; i < l.length; i+=4) {
        if(l[i] != " ") {
            stacks1[i / 4 + 3/4].unshift(l[i]);
            stacks2[i / 4 + 3/4].unshift(l[i]);
        }
    }
});

proc.split("\n").forEach(p => {
    const [w, x, y]: string[] = p.match(/\d+/g) as string[];
    const [a, b, c]: number[] = [Number(w), Number(x), Number(y)];

    stacks1[c].push(...stacks1[b].splice(stacks1[b].length - a, a).reverse())
    stacks2[c].push(...stacks2[b].splice(stacks2[b].length - a, a))
});

for(let i = 1; i <= stacksNumber; i++) {
    const [s1, s2] = [stacks1[i], stacks2[i]];
    if(s1.length > 0)
        score1 += s1[s1.length - 1]
    if(s2.length > 0)
        score2 += s2[s2.length - 1]
}

const b: number = performance.now();

console.log("Parte 1: " + score1);
console.log("Parte 2: " + score2);
console.log(`Tomó ${b - a} ms`)