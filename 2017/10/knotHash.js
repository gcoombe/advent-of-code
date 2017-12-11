module.exports.knotHash1 = (lengths, size = 256) => {
    let skipSize = 0;
    let currentPos = 0;

    const list = Array.from(new Array(size), (val, index) => index);

    lengths.forEach(length => {
        let subList = list.slice(currentPos, currentPos + length);

        if (currentPos + length > list.length) {
            subList = subList.concat(list.slice(0, (currentPos + length) % list.length));
        }

        subList.reverse();

        subList.forEach((val, i) => {
            list[(currentPos + i) % list.length] = val;
        });

        currentPos = (currentPos + length + skipSize) % list.length;
        skipSize++;
    });

    return list[0] * list[1];
};

module.exports.knotHash2 = (input, size = 256) => {
    let skipSize = 0;
    let currentPos = 0;

    const sparseHash = Array.from(new Array(size), (val, index) => index);
    const lengths = input.split("").map(x => x.charCodeAt(0)).concat([17, 31, 73, 47, 23]);

    for (let r = 0; r < 64; r++) {
        lengths.forEach(length => {
            let subList = sparseHash.slice(currentPos, currentPos + length);

            if (currentPos + length > sparseHash.length) {
                subList = subList.concat(sparseHash.slice(0, (currentPos + length) % sparseHash.length));
            }

            subList.reverse();

            subList.forEach((val, i) => {
                sparseHash[(currentPos + i) % sparseHash.length] = val;
            });

            currentPos = (currentPos + length + skipSize) % sparseHash.length;
            skipSize++;
        });
    }


    const denseHash = [];
    for (let block = 0; block < 16; block++) {
        const subList = sparseHash.slice(block * 16, block * 16 + 16);
        denseHash.push(subList.reduce((accumulator, val) => {
            if (!accumulator) {
                return val;
            }
            return accumulator ^ val;
        }, null));
    }

    return denseHash.reduce((currString, val) => {
        return currString + val.toString(16).padStart(2, "0");
    },"");
};

