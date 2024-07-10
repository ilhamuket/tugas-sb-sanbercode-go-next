var filterCarsPromise = require('./promise3.js');

filterCarsPromise("black", 2019)
    .then(cars => {
        console.log("Cars found:", cars);
    })
    .catch(error => {
        console.error(error.message);
    });

filterCarsPromise("silver", 2017)
    .then(cars => {
        console.log("Cars found:", cars);
    })
    .catch(error => {
        console.error(error.message);
    });

(async () => {
    try {
        let cars = await filterCarsPromise("grey", 2019);
        console.log("Cars found:", cars);
    } catch (error) {
        console.error(error.message);
    }

    try {
        let cars = await filterCarsPromise("grey", 2018);
        console.log("Cars found:", cars);
    } catch (error) {
        console.error(error.message);
    }

    try {
        let cars = await filterCarsPromise("black", 2020);
        console.log("Cars found:", cars);
    } catch (error) {
        console.error(error.message);
    }
})();
