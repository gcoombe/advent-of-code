//Used http://keekerdc.com/2011/03/hexagon-grids-coordinate-systems-and-distance-calculations/

module.exports.shortest = (input) => {
    const moves = input.split(",");

    const childPos = [0,0,0];
    moves.forEach(move => {
        if (move === "n") {
            childPos[0] = childPos[0] - 1;
            childPos[1] = childPos[1] + 1;
        }

        if (move === "s") {
            childPos[0] = childPos[0] + 1;
            childPos[1] = childPos[1] - 1;
        }

        if (move === "ne") {
            childPos[1] = childPos[1] + 1;
            childPos[2] = childPos[2] - 1;
        }

        if (move === "sw") {
            childPos[1] = childPos[1] - 1;
            childPos[2] = childPos[2] + 1;
        }

        if (move === "se") {
            childPos[0] = childPos[0] + 1;
            childPos[2] = childPos[2] - 1;
        }

        if (move === "nw") {
            childPos[0] = childPos[0] - 1;
            childPos[2] = childPos[2] + 1;
        }
    });

    return calculateDistance(childPos);
};

const calculateDistance = pos => {
    return (Math.abs(pos[0]) + Math.abs(pos[1]) + Math.abs(pos[2])) / 2;
};

module.exports.farthest = (input) => {
    const moves = input.split(",");

    let maxDistance = 0;
    const childPos = [0,0,0];
    moves.forEach(move => {
        if (move === "n") {
            childPos[0] = childPos[0] - 1;
            childPos[1] = childPos[1] + 1;
        }

        if (move === "s") {
            childPos[0] = childPos[0] + 1;
            childPos[1] = childPos[1] - 1;
        }

        if (move === "ne") {
            childPos[1] = childPos[1] + 1;
            childPos[2] = childPos[2] - 1;
        }

        if (move === "sw") {
            childPos[1] = childPos[1] - 1;
            childPos[2] = childPos[2] + 1;
        }

        if (move === "se") {
            childPos[0] = childPos[0] + 1;
            childPos[2] = childPos[2] - 1;
        }

        if (move === "nw") {
            childPos[0] = childPos[0] - 1;
            childPos[2] = childPos[2] + 1;
        }

        maxDistance = Math.max(maxDistance, calculateDistance(childPos));
    });
    return maxDistance;
};