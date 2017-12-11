const assert = require('assert');
const knotHash = require("./knotHash");

describe("KnotHash", function () {

    describe("Part 1", function () {

        it("Calculates knotHash correctly", function () {
            assert.equal(knotHash.knotHash1([3, 4, 1, 5], 5), 12);
        });
    });

    describe("Part 2", function () {
        it("Calculates ''", function () {
            assert.equal(knotHash.knotHash2(""), "a2582a3a0e66e6e86e3812dcb672a272");
        });

        it("Calculates AoC 2017", function () {
            assert.equal(knotHash.knotHash2("AoC 2017"), "33efeb34ea91902bb2f59c9920caa6cd");
        });

        it("Calculates 1 2 3", function () {
            assert.equal(knotHash.knotHash2("1,2,3"), "3efbe78a8d82f29979031a4aa0b16a9d");
        });
    });
});