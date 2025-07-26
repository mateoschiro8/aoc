import { readFileSync } from "fs";
import { join } from "path";
const input: string = readFileSync(join(__dirname, "input.txt"), { encoding: "utf-8" });

const a: number = performance.now();

interface INode {
    name: string;
    parent: Directory | null;

    size(score: { value: number }): number;
    checkToDelete(smax: number, score: { value: number }): void;
    isDirectory(): boolean;
}

type Node = Filee | Directory;

class Filee implements INode {
    name: string;
    _size: number;
    parent: Directory | null;

    constructor(name: string, size: number, parent: Directory | null) {
        this.name = name;
        this._size = size;
        this.parent = parent;
    }

    size(score: { value: number }): number {
        return this._size;
    }

    checkToDelete(smax: number, score: { value: number }) { }

    isDirectory(): boolean {
        return false;
    }
}

class Directory implements INode {
    name: string;
    things: Node[];
    parent: Directory | null;

    constructor(name: string, parent: Directory | null) {
        this.name = name;
        this.things = [];
        this.parent = parent;
    }
    
    addNode(node: Node) {
        this.things.push(node);
    }

    size(score: { value: number }): number {
        const s = this.things.reduce((total, node) => total + node.size(score), 0);
        if(s <= 100000)
            score.value += s;
        return s;
    }

    searchSubDir(name: string): Directory | null {
        for(let i = 0; i < this.things.length; i++) {
            if(this.things[i].name == name)
                return this.things[i] as Directory
        }
        this.things.push(new Directory(name, this));
        return this.things[this.things.length - 1] as Directory;
    }

    checkToDelete(smax: number, score: { value: number }) {
        for(let i = 0; i < this.things.length; i++) {
            if(this.things[i].isDirectory() && (70000000 - smax + this.things[i].size({value: 0})) > 30000000) {
                score.value = Math.min(this.things[i].size({value: 0}), score.value)
            }
            this.things[i].checkToDelete(smax, score);
        }

    }

    isDirectory(): boolean {
        return true;
    }
}

let cwd: Directory | null = null;
let root: Directory | null = null;

input.split("\n").forEach(l => {
    const cmd: string[] = l.split(" ");
    
    if(cmd[1] === "ls")
        return;

    if(cmd[1] === "cd") {
        if(cwd == null) {
            cwd = new Directory(cmd[2], null)
            root = cwd;
        }
        else  
            cwd = (cmd[2] === "..") ? (cwd?.parent ?? cwd) : (cwd?.searchSubDir(cmd[2]) ?? null);
    }
    else if(cmd[0] == "dir") {
        cwd?.addNode(new Directory(cmd[1], cwd));
    } else {
        const [s, n]: (number | string)[] = [Number(cmd[0]), cmd[1]];
        cwd?.addNode(new Filee(n as string, s as number, cwd));
    }
});

let score1 = { value: 0 }; 
const smax = root!.size(score1);

let score2 = { value: smax };
root!.checkToDelete(smax, score2);

const b: number = performance.now();

console.log("Parte 1: " + score1.value);
console.log("Parte 2: " + score2.value);
console.log(`Tomó ${b - a} ms`)