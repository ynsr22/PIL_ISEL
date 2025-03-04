import { useContext } from "react";
import Image from "next/image";
import Link from "next/link";
import { SearchContext } from "../components/search";
import { Search, ShoppingCart, BookOpen } from "lucide-react";
import logo from "../../public/logo_renault.png";
import logo_responsive from "../../public/logo_renault_responsive.png";

const SearchBar: React.FC<{ searchQuery: string; setSearchQuery: (query: string) => void }> = ({ searchQuery, setSearchQuery }) => (
  <form
    className="flex items-center w-full max-w-xs h-9 bg-white rounded-full overflow-hidden border border-gray-200 focus-within:ring-2 focus-within:ring-black focus-within:ring-opacity-20"
    onSubmit={(e) => e.preventDefault()}
    role="search"
  >
    <input
      type="search"
      value={searchQuery}
      onChange={(e) => setSearchQuery(e.target.value)}
      placeholder="Recherche..."
      className="w-full min-w-0 px-3 py-1 text-gray-700 bg-transparent focus:outline-none"
      aria-label="Rechercher des produits"
    />
    <button
      type="submit"
      className="flex-shrink-0 flex items-center justify-center w-10 h-full bg-black text-white hover:bg-gray-800 transition-colors"
      aria-label="Lancer la recherche"
    >
      <Search className="w-5 h-5" />
    </button>
  </form>
);

const CartIcon: React.FC = () => (
  <Link href="/panier" aria-label="Accéder au panier">
    <div className="relative flex-shrink-0 cursor-pointer hover:bg-white rounded-full p-2 transition-colors group">
      <ShoppingCart className="w-6 h-6 group-hover:scale-105 transition-transform" />
    </div>
  </Link>
);

const Header: React.FC = () => {
  const { searchQuery, setSearchQuery } = useContext(SearchContext);

  return (
    <header className="w-full h-14 bg-[#FFEA2F] flex items-center justify-between px-4 relative shadow-sm">
      <div className="flex items-center absolute left-4 z-10">
        <Link href="/" className="flex items-center" aria-label="Accueil Renault">
          <Image src={logo} alt="Logo Renault" width={110} height={40} priority className="hidden sm:block" />
          <Image src={logo_responsive} alt="Logo Renault" width={35} height={35} priority className="block sm:hidden" />
        </Link>
      </div>
      <div className="flex justify-center items-center w-full">
        <SearchBar searchQuery={searchQuery} setSearchQuery={setSearchQuery} />
      </div>
      <div className="flex items-center absolute right-4 z-10">
        <CartIcon />
      </div>
    </header>
  );
};

const Navigation: React.FC = () => (
  <nav className="w-full h-10 bg-black shadow-md" aria-label="Navigation principale">
    <div className="max-w-6xl mx-auto h-full flex justify-center items-center">
      <Link href="/notice" className="h-full flex justify-center items-center gap-1 group py-2 px-4 text-white hover:text-[#FFEA2F] transition-colors duration-200" aria-label="Accéder aux notices d'utilisation">
        <BookOpen className="w-5 h-5 stroke-white group-hover:stroke-[#FFEA2F] transition-colors duration-200" />
        <span className="font-medium">Notice d&#39;utilisation</span>
      </Link>
    </div>
  </nav>
);

export { Header, Navigation };
