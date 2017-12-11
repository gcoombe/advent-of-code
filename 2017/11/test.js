const assert = require("assert");
const hexGrid = require("./hexGrid");

describe("Part 1", function () {
    it("Calculates test 1", function () {
        assert.equal(hexGrid.shortest("ne,ne,ne"), 3);
    });

    it("Calculates test 2", function () {
        assert.equal(hexGrid.shortest("ne,ne,sw,sw"), 0);
    });

    it("Calculates test 3", function () {
        assert.equal(hexGrid.shortest("ne,ne,s,s"), 2);
    });

    it("Calculates test 4", function () {
        assert.equal(hexGrid.shortest("se,sw,se,sw,sw"), 3);
    });
});