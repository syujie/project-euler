const multiplesOf3and5 = (limit) => {
    let total = 0;

    for (let i = 0; i < limit; i++) {
        if (i % 3 === 0 || i % 5 === 0) {
            total += i;
        }
    }

    console.log(total);
};

multiplesOf3and5(1000);
