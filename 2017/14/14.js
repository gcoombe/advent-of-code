const knotHash = require("../10/knotHash");


const buildGrid = (input) => {
    return [...Array(128).keys()].map(i => input + "-" + i)
        .map(row => knotHash.knotHash2(row))
        .map(str => str.match(/.{1,2}/g))
        .map(row => {
            return row.reduce((accumulator, hexVal) => {
                const bits = parseInt(hexVal, 16).toString(2).padStart(8, "0");
                return accumulator + bits;
            }, "");
        });
};

module.exports.part1 = (input) => {
    const grid = buildGrid(input);

    return grid.map(row => {
            return row.split("").reduce((accumulator, bit) => {
                return accumulator + parseInt(bit, 10);
            }, 0);
        }).reduce((accumulator, row) => {
           return accumulator + row
        }, 0);

};

const dfs = (grid, i, j, seen) => {
    const key = i + "-" + j;
    if (seen.has(key)) {
        return
    } else if (grid[i][j] !== "1") {
        return
    } else {
        seen.add(key);
    }

    if (i > 0) {
        dfs(grid, i-1, j, seen)
    }
    if (j > 0 ){
        dfs(grid, i, j -1, seen);
    }
    if (i < 127) {
        dfs(grid, i+1, j, seen);
    }
    if (j < 127) {
        dfs (grid, i, j + 1, seen);
    }
};

module.exports.part2 = (input) => {
    const grid =  buildGrid(input);
    const seen = new Set();
    let n = 0;
    for (let i = 0; i <= 127; i++) {
        for (let j = 0; j <= 127; j++) {
            const key = i + "-" + j;
            if (!seen.has(key) && grid[i][j] === "1") {
                n++;
                dfs(grid, i, j, seen);
            }
        }
    }
    return n;
};