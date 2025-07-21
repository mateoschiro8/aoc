import { readdirSync } from "fs";
import { join } from "path";
import { spawnSync } from "child_process";

const entries = readdirSync(__dirname, {withFileTypes: true});

entries.forEach(entry => {
    if(entry.isDirectory() && entry.name.length <= 2) {
        console.log(`Corriendo ${entry.name}:`);
    
        const mainPath = join(__dirname, entry.name, "main.ts");    
        spawnSync("npx", ["ts-node", mainPath], { stdio: "inherit"});
    }
})

