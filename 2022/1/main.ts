import { readFileSync } from "fs";
import { join } from "path";
const input: string = readFileSync(join(__dirname, "input.txt"), { encoding: "utf-8" })

const a: number = performance.now();

const values: string[] = input.split("\n\n")
const sumas: number[] = values.map(value => {
    return value.split("\n").map(Number).reduce((p, c) => p + c, 0);
})

sumas.sort((a, b) => b - a);

const b: number = performance.now();

console.log("Parte 1: " + sumas[0]);
console.log("Parte 2: " + (sumas[0] + sumas[1] + sumas[2]));
console.log(`Tomó ${b - a} ms`)