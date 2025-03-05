import { createContext, useState, ReactNode } from 'react';

interface RechercheContextProps {
    recherche: string;
    setRecherche: (recherche: string) => void;
}

export const RechercheContext = createContext<RechercheContextProps | null>(null);

export const RechercheProvider = ({ children }: { children: ReactNode}) => {
    const [recherche, setRecherche] = useState<string>('');
    return (
        <RechercheContext.Provider value={{ recherche, setRecherche }}>
            {children}
        </RechercheContext.Provider>
    )
}