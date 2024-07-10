var filterBooksPromise = require('./promise2.js');

filterBooksPromise(true, 50)
    .then(books => {
        console.log("Books found:", books);
    })
    .catch(error => {
        console.error(error.message);
    });

(async () => {
    try {
        let books = await filterBooksPromise(false, 250);
        console.log("Books found:", books);
    } catch (error) {
        console.error(error.message);
    }

    try {
        let books = await filterBooksPromise(true, 30);
        console.log("Books found:", books);
    } catch (error) {
        console.error(error.message);
    }
})();
