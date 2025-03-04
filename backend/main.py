from fastapi import FastAPI, Depends, HTTPException
from sqlalchemy.orm import Session
from database import SessionLocal
from fastapi.middleware.cors import CORSMiddleware
import models

app = FastAPI(title="API Moyens Logistiques - FastAPI + SQLAlchemy + PostgreSQL")

# Liste des origines autorisées (ajuste selon ton déploiement)
origins = ["http://localhost:3000"]

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Dépendance pour récupérer une session DB
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@app.get("/")
def root():
    return {"message": "Bienvenue sur l'API Moyens Logistiques !"}

# --- Catégories ---
@app.get("/categories")
def get_categories(db: Session = Depends(get_db)):
    """Retourne toutes les catégories de moyens roulants."""
    return db.query(models.Categorie).all()

# --- Moyens roulants ---
@app.get("/moyens")
def get_moyens(db: Session = Depends(get_db)):
    """Retourne tous les moyens roulants."""
    return db.query(models.MoyenRoulant).all()

@app.get("/moyens/{moyen_id}")
def get_moyen(moyen_id: int, db: Session = Depends(get_db)):
    """Retourne un moyen roulant spécifique."""
    moyen = db.query(models.MoyenRoulant).filter(models.MoyenRoulant.id == moyen_id).first()
    if not moyen:
        raise HTTPException(status_code=404, detail="Moyen roulant non trouvé")
    return moyen

# --- Accessoires ---
@app.get("/accessoires")
def get_accessoires(db: Session = Depends(get_db)):
    """Retourne tous les accessoires."""
    return db.query(models.Accessoire).all()

@app.get("/accessoires/{accessoire_id}")
def get_accessoire(accessoire_id: int, db: Session = Depends(get_db)):
    """Retourne un accessoire spécifique."""
    accessoire = db.query(models.Accessoire).filter(models.Accessoire.id == accessoire_id).first()
    if not accessoire:
        raise HTTPException(status_code=404, detail="Accessoire non trouvé")
    return accessoire

# --- Accessoires par défaut pour une catégorie ---
@app.get("/categories/{categorie_id}/accessoires_defauts")
def get_accessoires_defauts(categorie_id: int, db: Session = Depends(get_db)):
    """Retourne les accessoires par défaut pour une catégorie."""
    accessoires_defauts = db.query(models.AccessoireDefaut).filter(models.AccessoireDefaut.categorie_id == categorie_id).all()
    accessoire_ids = [ad.accessoire_id for ad in accessoires_defauts]
    accessoires = db.query(models.Accessoire).filter(models.Accessoire.id.in_(accessoire_ids)).all()
    return accessoires

# --- Compatibilités (catégories <-> accessoires) ---
@app.get("/categories/{categorie_id}/accessoires")
def get_accessoires_for_categorie(categorie_id: int, db: Session = Depends(get_db)):
    """Retourne les accessoires compatibles avec une catégorie donnée."""
    compatibilites = db.query(models.Compatibilite).filter(models.Compatibilite.categorie_id == categorie_id).all()
    accessoire_ids = [c.accessoire_id for c in compatibilites]
    accessoires = db.query(models.Accessoire).filter(models.Accessoire.id.in_(accessoire_ids)).all()
    return accessoires

# --- Compatibilités (moyens roulants <-> accessoires) ---
@app.get("/moyens/{moyen_id}/accessoires")
def get_accessoires_for_moyen(moyen_id: int, db: Session = Depends(get_db)):
    """Retourne les accessoires compatibles avec un moyen roulant donné."""
    moyen = db.query(models.MoyenRoulant).filter(models.MoyenRoulant.id == moyen_id).first()
    if not moyen:
        raise HTTPException(status_code=404, detail="Moyen roulant non trouvé")

    compatibilites = db.query(models.Compatibilite).filter(models.Compatibilite.categorie_id == moyen.categorie_id).all()
    accessoire_ids = [c.accessoire_id for c in compatibilites]
    accessoires = db.query(models.Accessoire).filter(models.Accessoire.id.in_(accessoire_ids)).all()
    return accessoires
