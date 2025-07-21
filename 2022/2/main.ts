import { readFileSync } from "fs";
import { join } from "path";
const input: string = readFileSync(join(__dirname, "input.txt"), { encoding: "utf-8" })

const at: number = performance.now();

const points: Map<string, number> = new Map([["X", 1], ["Y", 2], ["Z", 3]])
const duels: Map<string, Map<string, number>> = new Map([
  ["A", new Map([["X", 3],["Y", 6],["Z", 0]])],
  ["B", new Map([["X", 0],["Y", 3],["Z", 6]])],
  ["C", new Map([["X", 6],["Y", 0],["Z", 3]])]
]);
const strategy: Map<string, Map<string, string>> = new Map([
  ["A", new Map([["X", "Z"],["Y", "X"],["Z", "Y"]])],
  ["B", new Map([["X", "X"],["Y", "Y"],["Z", "Z"]])],
  ["C", new Map([["X", "Y"],["Y", "Z"],["Z", "X"]])]
]);

let score1 = 0, score2 = 0;
input.split("\n").forEach(d => {
    const plays: string[] = d.split(" ");
    const he = plays[0]; 
    let me = plays[1];
    score1 += (points.get(me) ?? 0) + (duels.get(he)?.get(me) ?? 0);
    me = strategy.get(he)?.get(me) ?? "";
    score2 += (points.get(me) ?? 0) + (duels.get(he)?.get(me) ?? 0);
})

const bt: number = performance.now();

console.log("Parte 1: " + score1);
console.log("Parte 2: " + score2);
console.log(`Tomó ${bt - at} ms`)