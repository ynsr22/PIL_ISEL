from sqlalchemy import Column, Integer, String, ForeignKey
from sqlalchemy.orm import relationship
from database import Base

# --- Table des Catégories ---
class Categorie(Base):
    __tablename__ = "categories"

    id = Column(Integer, primary_key=True, index=True)
    nom = Column(String, nullable=False)

    moyens_roulants = relationship("MoyenRoulant", back_populates="categorie")
    accessoire_defauts = relationship("AccessoireDefaut", back_populates="categorie")
    compatibilites = relationship("Compatibilite", back_populates="categorie")

# --- Table des Moyens Roulants ---
class MoyenRoulant(Base):
    __tablename__ = "moyens_roulants"

    id = Column(Integer, primary_key=True, index=True)
    nom = Column(String, nullable=False)
    categorie_id = Column(Integer, ForeignKey("categories.id"), nullable=False)
    roues = Column(Integer, nullable=False)
    emplacement = Column(String)
    type_base = Column(String)
    taille = Column(String)
    departement = Column(String)
    image = Column(String, nullable=True)
    prix = Column(Integer, nullable=False)

    categorie = relationship("Categorie", back_populates="moyens_roulants")

# --- Table des Accessoires ---
class Accessoire(Base):
    __tablename__ = "liste_accessoire"

    id = Column(Integer, primary_key=True, index=True)
    nom = Column(String, nullable=False)
    image = Column(String, nullable=True)
    prix = Column(Integer, nullable=False)

    compatibilites = relationship("Compatibilite", back_populates="accessoire")
    accessoire_defauts = relationship("AccessoireDefaut", back_populates="accessoire")

# --- Table des Compatibilités entre Catégories et Accessoires ---
class Compatibilite(Base):
    __tablename__ = "compatibilites"

    id = Column(Integer, primary_key=True, index=True)
    categorie_id = Column(Integer, ForeignKey("categories.id"), nullable=False)
    accessoire_id = Column(Integer, ForeignKey("liste_accessoire.id"), nullable=False)

    categorie = relationship("Categorie", back_populates="compatibilites")
    accessoire = relationship("Accessoire", back_populates="compatibilites")

# --- Table des Accessoires par Défaut pour une Catégorie ---
class AccessoireDefaut(Base):
    __tablename__ = "accessoire_defaut"

    id = Column(Integer, primary_key=True, index=True)
    categorie_id = Column(Integer, ForeignKey("categories.id"), nullable=False)
    accessoire_id = Column(Integer, ForeignKey("liste_accessoire.id"), nullable=False)

    categorie = relationship("Categorie", back_populates="accessoire_defauts")
    accessoire = relationship("Accessoire", back_populates="accessoire_defauts")
