"use client";

import { memo } from "react";
import Image from "next/image";

interface accessoires {
    id: number;
    nom: string;
    prix: number;
    image?: string;
}

interface CarteAccessoireProps {
    accessoire: accessoires;
    isSelected: boolean;
    isSelectable: boolean;
    onClick: () => void;
}

const CarteAccessoire = memo(({ accessoire, isSelected, isSelectable, onClick }: CarteAccessoireProps) => (
    <div 
        className={`p-2 border rounded transition-colors relative
            ${isSelectable
                ? (isSelected ? 'border-yellow-500 bg-yellow-50 cursor-pointer' : 'border-gray-200 hover:border-yellow-300 cursor-pointer')
                : 'border-gray-200 bg-gray-50'}`}
        onClick={isSelectable ? onClick : undefined}
        role={isSelectable ? 'button' : undefined}
        tabIndex={isSelectable ? 0 : undefined}
        aria-label={accessoire.nom}
        aria-pressed={isSelectable ? isSelected : undefined}
        aria-disabled={isSelectable ? undefined : true}
        onKeyDown={isSelectable ? (e: React.KeyboardEvent<HTMLDivElement>) => e.key === "Enter" && onClick() : undefined}
    >
        <div className="relative w-full h-40">
        <Image
            src={accessoire.image || '/accessoires/defaut.png'}
            alt={accessoire.nom}
            layout="fill"
            objectFit="contain"
            className="rounded"
            loading="lazy"
        /></div>
        <h3 className="text-sm font-medium">{accessoire.nom}</h3>
        <p className="text-xs text-gray-600">{accessoire.prix.toLocaleString()}€</p>
        {isSelected && isSelectable && (
            <div className="bg-yellow-500 text-white text-xs rounded-full w-5 h-5 flex items-center justify-center absolute top-1 right-1">
                ✓
            </div>
        )}
    </div>
));

CarteAccessoire.displayName = 'CarteAccessoire';

export default CarteAccessoire;
