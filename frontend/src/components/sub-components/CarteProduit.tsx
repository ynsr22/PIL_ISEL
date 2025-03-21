import { Link } from 'react-router-dom';
import { ImageWithFallback } from './ImageWithFallback';

interface ProductCardProps {
  id: number;
  nom: string;
  image: string;
  taille: string;
  prix: number;
}

// Composant pour les cartes de produits
export const ProductCard = ({ item }: { item: ProductCardProps }) => (
  <Link
    to={`/produit/${item.id}`}
    className="group bg-white rounded-lg shadow border border-gray-200 p-4 flex flex-col hover:shadow-lg transition-all duration-200 hover:border-yellow-400 h-full"
  >
    <div className="flex-1 flex items-center justify-center mb-4">
      <ImageWithFallback
        src={item.image}
        alt={`Image de ${item.nom}`}
      />
    </div>
    <div className="flex flex-col gap-2 mt-auto">
      <h3 className="text-lg font-semibold text-center text-gray-800 group-hover:text-yellow-600 transition-colors">
        {item.nom}
      </h3>
      <div className="flex justify-between items-center">
        <span className="text-sm text-gray-500">Taille {item.taille}</span>
        <span className="text-sm font-medium bg-yellow-50 text-yellow-600 px-2 py-1 rounded">
          {typeof item.prix === "number" ? item.prix.toFixed(2) : item.prix} €
        </span>
      </div>
    </div>
  </Link>
);
