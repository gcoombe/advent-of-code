#!/usr/bin/env node
const knotHash  = require("./knotHash");

const input = "225,171,131,2,35,5,0,13,1,246,54,97,255,98,254,110";

console.log("Part 1: " + knotHash.knotHash1(input.split(",")));
console.log("Part 2: " + knotHash.knotHash2(input));