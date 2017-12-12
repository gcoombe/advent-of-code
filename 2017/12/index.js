const fs = require("fs");
const search = require("./search");

const input = fs.readFileSync(__dirname + "/part1.txt", "utf-8");
console.log("Part 1", search.part1(input, "0"));
console.log("Part 2", search.part2(input));