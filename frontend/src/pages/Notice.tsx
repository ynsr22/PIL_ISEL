import { FC, memo } from "react";

// Composant Section réutilisable
interface SectionProps {
  title: string;
  content: React.ReactNode;
}

const Section: FC<SectionProps> = memo(({ title, content }) => (
  <section className="mb-8" aria-labelledby={title.replace(/\s+/g, "-").toLowerCase()}>
    <h2 id={title.replace(/\s+/g, "-").toLowerCase()} className="text-xl font-semibold mb-4">
      {title}
    </h2>
    {content}
  </section>
));

Section.displayName = "Section";

// Composant Notice principal
const Notice: FC = memo(() => {
  const sections = [
    {
      title: "1. Introduction",
      content: (
        <p className="text-gray-700">
          Cette notice est un exemple temporaire pour illustrer l'utilisation d'un site web en cours de développement.
          Le contenu sera mis à jour à mesure que le projet avance et que les fonctionnalités finales seront définies.
        </p>
      ),
    },
    {
      title: "2. Objectifs du site",
      content: (
        <ul className="list-disc pl-6 text-gray-700">
          <li>Faciliter la commande des moyens logistiques non motorisés.</li>
          <li>Centraliser les informations pour une meilleure gestion.</li>
          <li>Offrir une interface intuitive inspirée des boutiques en ligne.</li>
        </ul>
      ),
    },
    {
      title: "3. Fonctionnalités futures",
      content: (
        <p className="text-gray-700">
          Les fonctionnalités incluront la consultation des produits, la configuration personnalisée et la gestion des commandes.
          Des mises à jour régulières garantiront l'évolution et l'amélioration de l'expérience utilisateur.
        </p>
      ),
    },
  ];

  return (
    <div className="p-8 max-w-3xl mx-auto">
      <h1 className="text-3xl font-bold mb-6">Notice Exemple</h1>
      {sections.map((section, index) => (
        <Section key={index} title={section.title} content={section.content} />
      ))}
    </div>
  );
});

Notice.displayName = "Notice";

// Composant de la page
const Page: FC = () => (
  <main>
    <Notice />
  </main>
);

export default Page;
