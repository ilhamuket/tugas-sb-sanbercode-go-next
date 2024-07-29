import Layout from './components/Layout';
import Card from './components/Card';
import Image from 'next/image';
import bannerImage from './assets/shoping_hero.svg'; 
import Link from 'next/link';

export default function Home() {
  return (
    <Layout>
      <section style={{ backgroundColor: '#ffe4e4' }} className="px-8 py-20 text-black md:px-16">
        <div className="flex flex-col items-center max-w-screen-lg mx-auto md:flex-row">
          <div className="flex-1 text-center md:text-left">
            <h1 className="mb-4 text-5xl font-bold">Temukan produk kamu disini!</h1>
            <p className="mb-4 text-lg">Nikmati diskon hingga 100% setiap pembelian yang kamu lakukan</p>
            <button className="px-6 py-3 text-white bg-red-500 rounded-xl">
              Find Product
            </button>
          </div>
          <div className="flex justify-center flex-1 mt-10 md:mt-0">
            <Image src={bannerImage} alt="Banner Image" className="max-w-xs md:max-w-md" />
          </div>
        </div>
      </section>
      
      {/* New Search Bar Section */}
      <section className="px-8 py-4 bg-white md:px-16">
        <div className="flex flex-col items-center max-w-screen-lg mx-auto md:flex-row md:items-start">
          <div className="flex items-center w-full mb-4 md:mb-0 md:flex-1">
            <div className="flex items-center justify-center w-10 h-10 bg-pink-500 rounded-full">
              <svg className="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" /></svg>
            </div>
            <input
              type="text"
              placeholder="Search Product..."
              className="w-full p-2 ml-2 border border-gray-300 rounded-lg md:w-3/4 lg:w-4/5"
            />
          </div>
          <div className="flex flex-wrap items-center justify-center w-full gap-4 md:justify-end md:flex-1">
            <Link href="/category/shirt" className="text-blue-500 hover:underline">Shirt</Link>
            <Link href="/category/electronics" className="text-blue-500 hover:underline">Electronics</Link>
            <Link href="/category/games" className="text-blue-500 hover:underline">Games</Link>
            <Link href="/category/hijab" className="text-blue-500 hover:underline">Hijab</Link>
            <Link href="/category/shoes" className="text-blue-500 hover:underline">Shoes</Link>
            <Link href="/category/laptops" className="text-blue-500 hover:underline">Laptops</Link>
          </div>
        </div>
      </section>
      
      <section className="px-8 mt-10 bg-white md:px-16">
        <div className="max-w-screen-lg mx-auto">
          <div className="mb-6">
            <div className="flex justify-between w-full">
              <h2 className="text-2xl font-semibold text-black">This is recommended for you!</h2>
            </div>
          </div>
          <div className="flex flex-wrap justify-center gap-4">
            <Card />
            <Card />
            <Card />
          </div>
        </div>
      </section>
    </Layout>
  );
}
