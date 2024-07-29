import Link from 'next/link';

const Navbar = () => {
  return (
    <header className="p-4 text-black bg-white">
      <nav className="flex items-center justify-between max-w-screen-xl px-4 mx-auto md:px-12">
        <Link href="/" legacyBehavior>
          <a className="text-2xl font-bold">LOGO</a>
        </Link>
        <div className="flex space-x-4">
          <Link href="/" legacyBehavior>
            <a>Home</a>
          </Link>
          <Link href="/products" legacyBehavior>
            <a>Product</a>
          </Link>
        </div>
      </nav>
    </header>
  );
};

export default Navbar;
