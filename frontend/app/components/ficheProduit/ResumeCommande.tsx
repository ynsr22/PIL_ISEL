import { memo } from 'react';

interface Accessoires {
    id: number;
    nom: string;
    prix: number;
    image?: string;
}

interface ResumeCommandeProps {
    quantite: number;
    accessoiresParDefaut: Accessoires[];
    selectedAccessoiresOptionnels: Accessoires[];
    totalPrice: number;
    prixProduit: number;
}

const ResumeCommande = memo(({ quantite, accessoiresParDefaut, selectedAccessoiresOptionnels, totalPrice, prixProduit }: ResumeCommandeProps) => {
    const totalPrixAccessoiresParDefaut = accessoiresParDefaut.reduce((sum, acc) => sum + acc.prix, 0);
    const totalPrixAccessoiresOptionnels = selectedAccessoiresOptionnels.reduce((sum, acc) => sum + acc.prix, 0);

    return (
        <div className="bg-gray-50 p-2 rounded h-auto">
            <h3 className="text-sm font-semibold mb-1">Résumé de la commande</h3>
            <div className="text-xs">
                <p>Produit: {prixProduit}€ x {quantite}</p>
                <p className="mt-1">Total accessoires par défaut : {totalPrixAccessoiresParDefaut}€ x {quantite} = {totalPrixAccessoiresParDefaut * quantite}€</p>
                <p className="mt-1">Total accessoires optionnels : {totalPrixAccessoiresOptionnels}€ x {quantite} = {totalPrixAccessoiresOptionnels * quantite}€</p>

                <div className="border-t pt-1 mt-1" />
                <p className="text-base font-bold">Total : {totalPrice.toLocaleString()}€</p>
            </div>
        </div>
    );
});

ResumeCommande.displayName = 'ResumeCommande';

export default ResumeCommande;