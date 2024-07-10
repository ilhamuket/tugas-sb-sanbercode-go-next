console.log("---------- Soal 1 ----------");
// Fungsi untuk menghitung luas lingkaran
const luasLingkaran = (radius) => {
    return Math.PI * radius * radius;
};

// Fungsi untuk menghitung keliling lingkaran
const kelilingLingkaran = (radius) => {
    return 2 * Math.PI * radius;
};

console.log(luasLingkaran(10)); // Output: 314.1592653589793
console.log(kelilingLingkaran(10)); // Output: 62.83185307179586


console.log("---------- Soal 2 ----------");
const introduce = (...args) => {
    const [name, age, gender, job] = args;
    const title = gender === "Laki-Laki" ? "Pak" : "Bu";
    return `${title} ${name} adalah seorang ${job} yang berusia ${age} tahun`;
};

const perkenalanJohn = introduce("john", "30", "Laki-Laki", "penulis");
console.log(perkenalanJohn); // Output: "Pak john adalah seorang penulis yang berusia 30 tahun"

const perkenalanSarah = introduce("sarah", "28", "Perempuan", "guru");
console.log(perkenalanSarah); // Output: "Bu sarah adalah seorang guru yang berusia 28 tahun"


console.log("---------- Soal 3 ----------");
const newFunction = (firstName, lastName) => ({
    firstName,
    lastName,
    fullName() {
        console.log(`${firstName} ${lastName}`);
    }
});

console.log(newFunction("John", "Doe").firstName); // Output: John
console.log(newFunction("Richard", "Roe").lastName); // Output: Roe
newFunction("William", "Imoh").fullName(); // Output: William Imoh


console.log("---------- Soal 4 ----------");
let phone = {
    name: "Galaxy Note 20",
    brand: "Samsung",
    year: 2020,
    colors: ["Mystic Bronze", "Mystic White", "Mystic Black"]
};

const { name: phoneName, brand: phoneBrand, year, colors: [colorBronze, , colorBlack] } = phone;

console.log(phoneBrand, phoneName, year, colorBlack, colorBronze); 
// Output: Samsung Galaxy Note 20 2020 Mystic Black Mystic Bronze


console.log("---------- Soal 5 ----------");
let warna = ["biru", "merah", "kuning", "hijau"];
let dataBukuTambahan = {
    penulis: "john doe",
    tahunTerbit: 2020
};
let buku = {
    nama: "pemograman dasar",
    jumlahHalaman: 172,
    warnaSampul: ["hitam"]
};

buku = {
    ...buku,
    warnaSampul: [...buku.warnaSampul, ...warna],
    ...dataBukuTambahan
};

console.log(buku);
/* Output:
{
    nama: "pemograman dasar",
    jumlahHalaman: 172,
    warnaSampul: ["hitam", "biru", "merah", "kuning", "hijau"],
    penulis: "john doe",
    tahunTerbit: 2020
}
*/


console.log("---------- Soal 6 ----------");
const addProducts = (samsung, newProducts) => ({
    ...samsung,
    products: [...samsung.products, ...newProducts]
});

let samsung = {
    name: "Samsung",
    products: [
        { name: "Samsung Galaxy Note 10", colors: ["black", "gold", "silver"] },
        { name: "Samsung Galaxy Note 10s", colors: ["blue", "silver"] },
        { name: "Samsung Galaxy Note 20s", colors: ["white", "black"] }
    ]
};

let newProducts = [
    { name: "Samsung Galaxy A52", colors: ["white", "black"] },
    { name: "Samsung Galaxy M52", colors: ["blue", "grey", "white"] }
];

samsung = addProducts(samsung, newProducts);

console.log(samsung);



console.log("---------- Soal 7 ----------");
const createObject = (nama, domisili, umur) => ({ nama, domisili, umur });

let data = ["Bondra", "Medan", 25];
console.log(createObject(...data)); 
// Output: { nama: "Bondra", domisili: "Medan", umur: 25 }


console.log("---------- Soal 8 ----------");
const graduate = (...data) => {
    return data.reduce((result, item) => {
        result[item.class] = result[item.class] || [];
        result[item.class].push(item.name);
        return result;
    }, {});
};

const data1 = [
    { name: "Ahmad", class: "adonis" },
    { name: "Regi", class: "laravel" },
    { name: "Bondra", class: "adonis" },
    { name: "Iqbal", class: "vuejs" },
    { name: "Putri", class: "laravel" }
];

const data2 = [
    { name: "Yogi", class: "react" },
    { name: "Fikri", class: "agile" },
    { name: "Arief", class: "agile" }
];

console.log(graduate(...data1));
// Output:
// {
//     adonis: ["Ahmad", "Bondra"],
//     laravel: ["Regi", "Putri"],
//     vuejs: ["Iqbal"]
// }

console.log(graduate(...data2));
// Output:
// {
//     react: ["Yogi"],
//     agile: ["Fikri", "Arief"]
// }


