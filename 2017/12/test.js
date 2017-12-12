const search = require("./search");
const fs = require("fs");
const assert = require("assert");

describe("Part 1", function () {
    it("Solves correctly", function () {
        const input = fs.readFileSync(__dirname + "/testInput.txt", "utf-8");
        assert.equal(search.solve(input, "0"), 6);
    });
});