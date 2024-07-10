var readBooksPromise = require('./promise.js');

var books = [
    { name: 'LOTR', timeSpent: 3000 },
    { name: 'Fidas', timeSpent: 2000 },
    { name: 'Kalkulus', timeSpent: 4000 }
];

let time = 10000;

const readBooksSequentially = (time, books) => {
    if (books.length === 0) {
        console.log("Semua buku telah dibaca atau waktu telah habis.");
        return;
    }

    readBooksPromise(time, books[0])
        .then(sisaWaktu => {
            readBooksSequentially(sisaWaktu, books.slice(1));
        })
        .catch(sisaWaktu => {
            console.log("Tidak cukup waktu untuk membaca buku selanjutnya.");
        });
};

readBooksSequentially(time, books);
