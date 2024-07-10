var readBooks = require('./callback.js');

var books = [
    { name: 'LOTR', timeSpent: 3000 },
    { name: 'Fidas', timeSpent: 2000 },
    { name: 'Kalkulus', timeSpent: 4000 },
    { name: 'komik', timeSpent: 1000 }
];

function readBooksRecursive(time, index) {
    if (index < books.length) {
        readBooks(time, books[index], (sisaWaktu) => {
            if (sisaWaktu > 0) {
                readBooksRecursive(sisaWaktu, index + 1);
            }
        });
    }
}

readBooksRecursive(10000, 0);
