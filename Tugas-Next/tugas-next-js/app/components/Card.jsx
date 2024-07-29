import Link from 'next/link';

const Card = () => {
  return (
    <div className="relative border border-gray-100" style={{ width: '300px' }}>
      <div className="relative object-cover w-full">
        <img src="https://fitinline.com/data/article/20210909/Foto-Produk-Baju-001.jpg" alt="Flowbite Logo" />
      </div>
      <div className="p-6">
        <small>
          <span className="bg-green-100 text-green-800 text-sm font-medium mr-2 px-2.5 py-0.5 rounded-r-lg dark:bg-green-200 dark:text-green-900">
            Kategori Produk
          </span>
        </small>
        <h5 className="mt-4">Nama Produk</h5>
        <ul className="mt-5 text-sm font-thin text-gray-500">
          <li>Stock: stok produk</li>
          <li className="text-lg font-bold">Harga: Rp Produk Harga</li>
        </ul>
        <div className="flex items-center justify-between mt-4 border">
          <button className="h-full px-2 text-black bg-gray-200">-</button>
          <input className="inline-block w-full h-full text-center focus:outline-none" placeholder="1" />
          <button className="h-full px-2 text-black bg-gray-200">+</button>
        </div>
        <button className="block w-full p-4 mt-5 text-sm font-medium text-white border rounded-sm bg-rose-400" type="button">
          Add to Cart
        </button>
        <Link href="" legacyBehavior>
          <a className="block w-full p-4 mt-2 text-sm font-medium text-center bg-white border rounded-sm text-rose-400 border-rose-400">
            Detail Product
          </a>
        </Link>
      </div>
    </div>
  );
};

export default Card;
