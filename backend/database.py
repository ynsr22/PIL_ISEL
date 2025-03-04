# database.py
import os
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, declarative_base

# Exemple: "postgresql://postgres:postgres@localhost:5433/postgres"
DATABASE_URL = os.getenv("DATABASE_URL", "postgresql://postgres:postgres@192.168.0.26:5433/postgres")

engine = create_engine(DATABASE_URL, echo=False)

SessionLocal = sessionmaker(bind=engine, autocommit=False, autoflush=False)

Base = declarative_base()
