import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Header, Navigation } from "./components/Header";
import Catalogue from "./pages/Catalogue";
import PageProduit from "./pages/PageProduit";
import Panier from "./pages/Panier";
import Notice from "./pages/Notice";
import { RechercheProvider } from "./components/sub-components/Recherche";

function App() {
  return (
    <Router>
      <RechercheProvider>
      <Header />
      <Navigation />  
      <main className="container mx-auto p-4">
        <Routes>
          <Route path="/" element={<Catalogue />} />
          <Route path="/produit/:id" element={<PageProduit />} />
          <Route path="/panier" element={<Panier />} />
          <Route path="/notice" element={<Notice />} />
          <Route path="*" element={<h1>404 - Page non trouvée</h1>} />
        </Routes>
      </main>
      </RechercheProvider>
    </Router>
  );
}

export default App;
