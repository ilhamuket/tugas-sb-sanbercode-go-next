// Introduce
console.log("---------- Soal 1 ----------");
function introduce(name, gender, job, age) {
    if (gender === "laki-laki") {
      return `Pak ${name} adalah seorang ${job} yang berusia ${age} tahun`;
    } else {
      return `Bu ${name} adalah seorang ${job} yang berusia ${age} tahun`;
    }
}
var john = introduce("John", "laki-laki", "penulis", "30");
console.log(john); // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
var sarah = introduce("Sarah", "perempuan", "model", "28");
console.log(sarah); // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

// Karakter Unik
console.log("---------- Soal 2 ----------");
function uniqueCharacters(str) {
    let result = '';
    let normalizedStr = str.toLowerCase().replace(/\s/g, ''); 
    for (let i = 0; i < normalizedStr.length; i++) {
      let char = normalizedStr.charAt(i);
      if (normalizedStr.indexOf(char) === normalizedStr.lastIndexOf(char)) {
        result += char;
      }
    }
    return result;
}
var text = "Super Bootcamp Fullstack Dev 2022";
console.log(uniqueCharacters(text)); // expected output: rbmfkdv0


// Terbesar dan Terkecil
console.log("---------- Soal 3 ----------");
function findMinMax(arr) {
    let min = Math.min(...arr);
    let max = Math.max(...arr);
    return `angka terbesar adalah ${max} dan angka terkecil adalah ${min}`;
}
var angka = [2, 3, 1, 9, 12, 8, 9, 7];
console.log(findMinMax(angka)); // expected output: angka terbesar adalah 12 dan angka terkecil adalah 1

// Nama Besar
console.log("---------- Soal 4 ----------");
function longestName(arr) {
    let longest = arr.reduce((a, b) => (a.length > b.length ? a : b));
    return longest;
}
var names = ["Andrew Gillett", "Chris Sawyer", "David Walsh", "John D Rockefeller"];
console.log(longestName(names)); // expected output: John D Rockefeller

// Arrange String
console.log("---------- Soal 5 ----------");
function arrangeString(str) {
    return str.split('').sort().join('');
}
console.log(arrangeString("bahasa")); // Output : aaabhs
console.log(arrangeString("similikiti")); // Output : iiiiiklmst
console.log(arrangeString("sanbercode")); // Output : abcdeenors
console.log(arrangeString("")); // Output : ""

console.log("---------- Soal 6 ----------");

function compressString(str) {
    let sortedStr = str.split('').sort().join(''); 
    let compressed = '';
    let count = 1;
    
    for (let i = 0; i < sortedStr.length; i++) {
        if (sortedStr[i] === sortedStr[i + 1]) {
            count++;
        } else {
            compressed += sortedStr[i] + count;
            count = 1;
        }
    }
    
    return compressed.length < str.length ? compressed : sortedStr;
}

console.log(compressString("abrakadabra")); // a5b2d1k1r2
console.log(compressString("aabcccccaaa")); // a5b1c5
console.log(compressString("abdul")); // abdlu
console.log(compressString("maman")); // aamn



// Palindrome
console.log("---------- Soal 7 ----------");
function palindrome(kata) {
    return kata === kata.split('').reverse().join('');
}
console.log(palindrome('katak')); // true
console.log(palindrome('blanket')); // false
console.log(palindrome('nababan')); // true
console.log(palindrome('haji ijah')); // true
console.log(palindrome('mister')); // false

// Palindrome Angka
console.log("---------- Soal 8 ----------");
function angkaPalindrome(num) {
    function isPalindrome(n) {
      return n.toString() === n.toString().split('').reverse().join('');
    }
    let nextNum = num + 1;
    while (!isPalindrome(nextNum)) {
      nextNum++;
    }
    return nextNum;
}
console.log(angkaPalindrome(8)); // 9
console.log(angkaPalindrome(10)); // 11
console.log(angkaPalindrome(117)); // 121
console.log(angkaPalindrome(175)); // 181
console.log(angkaPalindrome(1000)); // 1001

// Pasangan Angka Terbesar
console.log("---------- Soal 9 ----------");
function pasanganTerbesar(num) {
    let str = num.toString();
    let max = 0;
    for (let i = 0; i < str.length - 1; i++) {
      let pair = parseInt(str.substring(i, i + 2));
      if (pair > max) {
        max = pair;
      }
    }
    return max;
}
console.log(pasanganTerbesar(641573)); // 73
console.log(pasanganTerbesar(12783456)); // 83
console.log(pasanganTerbesar(910233)); // 91
console.log(pasanganTerbesar(71856421)); // 85
console.log(pasanganTerbesar(79918293)); // 99

// Cek Permutasi
console.log("---------- Soal 10 ----------");
function cekPermutasi(str1, str2) {
    if (str1.length !== str2.length) return false;
    let sortedStr1 = str1.split('').sort().join('');
    let sortedStr2 = str2.split('').sort().join('');
    return sortedStr1 === sortedStr2;
}
console.log(cekPermutasi("abah", "baha")); // true
console.log(cekPermutasi("ondel", "delon")); // true
console.log(cekPermutasi("paul sernine", "arsene lupin")); // true
console.log(cekPermutasi("taco", "taca")); // false

// URLify
console.log("---------- Soal 11 ----------");
function urlify(str, length) {
    let trimmedStr = str.substring(0, length);
    return trimmedStr.split(' ').join('%20');
}
console.log(urlify("Mr John Smith    ", 13)); // Mr%20John%20Smith
console.log(urlify("Bizzare world of Javascript     ", 27)); // Bizzare%20world%20of%20Javascript

// Mengubah Array menjadi object
console.log("---------- Soal 12 ----------");
var arrayDaftarPeserta = ["John Doe", "laki-laki", "baca buku", 1992];
var objDaftarPeserta = {
    nama: arrayDaftarPeserta[0],
    jenis_kelamin: arrayDaftarPeserta[1],
    hobi: arrayDaftarPeserta[2],
    tahun_lahir: arrayDaftarPeserta[3],
};
console.log(objDaftarPeserta);

// Menampilkan Huruf Vokal Saja
console.log("---------- Soal 13 ----------");
function tampilHurufVokal(sentence) {
    let vokal = 'aiueoAIUEO';
    let result = '';
    for (let char of sentence) {
      if (vokal.includes(char) || !isNaN(char)) {
        result += char;
      }
    }
    return result;
}
var sentence = "Super Bootcamp Golang Nextjs 2024";
console.log(tampilHurufVokal(sentence)); // "ueooaoae2024"

console.log("---------- Soal 14 ----------");

// Data buah yang diberikan
var buahData = [
    { nama: "Nanas", warna: "Kuning", "ada bijinya": "tidak", harga: 9000 },
    { nama: "Jeruk", warna: "Oranye", "ada bijinya": "ada", harga: 8000 },
    { nama: "Semangka", warna: "Hijau & Merah", "ada bijinya": "ada", harga: 10000 },
    { nama: "Pisang", warna: "Kuning", "ada bijinya": "tidak", harga: 5000 }
];

// Membuat array yang berisi objek buah yang tidak memiliki biji
var buahTanpaBiji = buahData.filter(buah => buah["ada bijinya"] === "tidak");

// Menampilkan hasil dalam bentuk array of objects
console.log(buahTanpaBiji);



// Mengelompokkan data
console.log("---------- Soal 15 ----------");
var people = [
    { name: "John", job: "Programmer", gender: "male", age: 30 },
    { name: "Sarah", job: "Model", gender: "female", age: 27 },
    { name: "Jack", job: "Engineer", gender: "male", age: 25 },
    { name: "Ellie", job: "Designer", gender: "female", age: 35 },
    { name: "Danny", job: "Footballer", gender: "male", age: 30 }
];
var malesAbove29 = people.filter(p => p.gender === "male" && p.age > 29);
console.log(malesAbove29);

// Rata-rata Usia
console.log("---------- Soal 16 ----------");
function rataRataUsia(people) {
    let totalAge = people.reduce((sum, person) => sum + person.age, 0);
    return totalAge / people.length;
}
console.log(rataRataUsia(people)); // expected output: rata-rata usia

// Urutkan data berdasarkan umur
console.log("---------- Soal 17 ----------");
people.sort((a, b) => a.age - b.age);
people.forEach((p, index) => {
    console.log(`${index + 1}. ${p.name}`);
});

// Menambahkan data ke property object
console.log("---------- Soal 18 ----------");
var phone = {
    name: "Samsung Galaxy Note 20",
    brand: "Samsung",
    colors: ["Black"],
    release: 2020
};
function addColors(color) {
    phone.colors.push(color);
}
addColors("Gold");
addColors("Silver");
addColors("Brown");
console.log(phone);

// Filter dan Munculkan Data
console.log("---------- Soal 19 ----------");
var phones = [
    { name: "Samsung Galaxy A52", brand: "Samsung", year: 2021, colors: ["black", "white"] },
    { name: "Redmi Note 10 Pro", brand: "Xiaomi", year: 2021, colors: ["white", "blue"] },
    { name: "Redmi Note 9 Pro", brand: "Xiaomi", year: 2020, colors: ["white", "blue", "black"] },
    { name: "Iphone 12", brand: "Apple", year: 2020, colors: ["silver", "gold"] },
    { name: "Iphone 11", brand: "Apple", year: 2019, colors: ["gold", "black", "silver"] },
];
var blackPhones = phones.filter(phone => phone.colors.includes("black")).sort((a, b) => a.year - b.year);
blackPhones.forEach((phone, index) => console.log(`${index + 1}. ${phone.name}, colors available: ${phone.colors.join(', ')}`));
