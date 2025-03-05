"use client";

import React, { useCallback, memo, useEffect, useState } from "react";
import RangeSlider from "./RangeSlider";

// Types
interface FilterSectionProps {
  title: string;
  isOpen: boolean;
  onToggle: () => void;
  children: React.ReactNode;
}

interface Category {
  id: number;
  nom: string;
}

interface FilterComponentProps {
  priceRange: [number, number];
  setPriceRange: React.Dispatch<React.SetStateAction<[number, number]>>;
  selectedDepartments: string[];
  setSelectedDepartments: React.Dispatch<React.SetStateAction<string[]>>;
  selectedCategories: number[];
  setSelectedCategories: React.Dispatch<React.SetStateAction<number[]>>;
}

interface OpenSections {
  department: boolean;
  category: boolean;
  price: boolean;
}

interface PriceRange {
  min: number;
  max: number;
}

// Constants
const DEPARTMENTS = ["Tôlerie", "Montage"] as const;

// API Services
const API_URL = "http://localhost:8000";

// UI Components
const FilterSection: React.FC<FilterSectionProps> = memo(
  ({ title, isOpen, onToggle, children }) => (
    <section className="border-b border-gray-200 last:border-b-0">
      <button
        onClick={onToggle}
        className="w-full flex justify-between items-center py-3 px-4 text-left text-lg font-bold focus:outline-none cursor-pointer hover:bg-gray-50 transition-colors"
        aria-expanded={isOpen}
      >
        {title}
        <span className="text-xl">{isOpen ? "−" : "+"}</span>
      </button>
      {isOpen && (
        <div className="p-4 space-y-2 animate-fadeIn">
          {children}
        </div>
      )}
    </section>
  )
);

FilterSection.displayName = "FilterSection";

// Helper function for selection toggle
const toggleSelection = <T,>(currentSelections: T[], itemToToggle: T): T[] =>
  currentSelections.includes(itemToToggle)
    ? currentSelections.filter((item) => item !== itemToToggle)
    : [...currentSelections, itemToToggle];

// Main Component
const FilterComponent: React.FC<FilterComponentProps> = ({
  priceRange,
  setPriceRange,
  selectedDepartments,
  setSelectedDepartments,
  selectedCategories,
  setSelectedCategories,
}) => {
  const [openSections, setOpenSections] = useState<OpenSections>({
    department: false,
    category: false,
    price: false,
  });

  const [categories, setCategories] = useState<Category[]>([]);
  const [categoriesLoading, setCategoriesLoading] = useState(true);

  const [priceRangeData, setPriceRangeData] = useState<PriceRange>({ min: 0, max: 1000 });
  const [priceRangeLoading, setPriceRangeLoading] = useState(true);

  // Fetch categories
  useEffect(() => {
    const fetchCategories = async () => {
      setCategoriesLoading(true);
      try {
        const response = await fetch(`${API_URL}/categories`);
        if (!response.ok) throw new Error(`Erreur serveur: ${response.status}`);
        const data = await response.json();
        if (Array.isArray(data)) setCategories(data);
      } catch (error) {
        console.error("Erreur lors du chargement des catégories:", error);
      } finally {
        setCategoriesLoading(false);
      }
    };

    fetchCategories();
  }, []);

  // Fetch price range
  useEffect(() => {
    const fetchPriceRange = async () => {
      setPriceRangeLoading(true);
      try {
        const response = await fetch(`${API_URL}/moyens`);
        if (!response.ok) throw new Error(`Erreur serveur: ${response.status}`);
        const data = await response.json();

        if (Array.isArray(data) && data.length > 0) {
          const prices = data.map((item) => parseFloat(item.prix));
          const min = Math.floor(Math.min(...prices));
          const max = Math.ceil(Math.max(...prices));

          setPriceRangeData({ min, max });
          setPriceRange([min, max]);
        }
      } catch (error) {
        console.error("Erreur lors du chargement des prix:", error);
      } finally {
        setPriceRangeLoading(false);
      }
    };

    fetchPriceRange();
  }, [setPriceRange]);

  const toggleSection = useCallback((section: keyof OpenSections) => {
    setOpenSections((prev) => ({ ...prev, [section]: !prev[section] }));
  }, []);

  const handleDepartmentChange = useCallback(
    (department: string) => {
      setSelectedDepartments((prev) => toggleSelection(prev, department));
    },
    [setSelectedDepartments]
  );

  const handleCategoryChange = useCallback(
    (categoryId: number) => {
      setSelectedCategories((prev) => toggleSelection(prev, categoryId));
    },
    [setSelectedCategories]
  );

  return (
    <div className="w-full border border-gray-200 rounded-lg shadow-md bg-white overflow-hidden">
      {/* Départements Section */}
      <FilterSection
        title="Départements"
        isOpen={openSections.department}
        onToggle={() => toggleSection("department")}
      >
        <div>
          {DEPARTMENTS.map((dept) => (
            <label
              key={dept}
              className="flex items-center gap-2 cursor-pointer py-1 px-1 rounded hover:bg-gray-50 transition-colors"
            >
              <input
                type="checkbox"
                checked={selectedDepartments.includes(dept)}
                onChange={() => handleDepartmentChange(dept)}
                className="h-4 w-4 rounded border-gray-300 text-black focus:ring-black"
              />
              <span>{dept}</span>
            </label>
          ))}
        </div>
      </FilterSection>

      {/* Catégories Section */}
      <FilterSection
        title="Catégories"
        isOpen={openSections.category}
        onToggle={() => toggleSection("category")}
      >
        {categoriesLoading ? (
          <div className="flex justify-center py-2">
            <div className="animate-pulse h-5 w-24 bg-gray-200 rounded"></div>
          </div>
        ) : categories.length === 0 ? (
          <p className="text-sm text-gray-500">Aucune catégorie disponible</p>
        ) : (
          <div className="space-y-2 max-h-60 overflow-y-auto pr-1">
            {categories.map((category) => (
              <label
                key={category.id}
                className="flex items-center gap-2 cursor-pointer py-1 px-1 rounded hover:bg-gray-50 transition-colors"
              >
                <input
                  type="checkbox"
                  checked={selectedCategories.includes(category.id)}
                  onChange={() => handleCategoryChange(category.id)}
                  className="h-4 w-4 rounded border-gray-300 text-black focus:ring-black"
                />
                <span>{category.nom}</span>
              </label>
            ))}
          </div>
        )}
      </FilterSection>

      {/* Prix Section */}
      <FilterSection
        title="Prix"
        isOpen={openSections.price}
        onToggle={() => toggleSection("price")}
      >
        {priceRangeLoading ? (
          <div className="flex flex-col space-y-4 py-2">
            <div className="animate-pulse h-2 w-full bg-gray-200 rounded"></div>
            <div className="flex justify-between">
              <div className="animate-pulse h-6 w-16 bg-gray-200 rounded-full"></div>
              <div className="animate-pulse h-6 w-16 bg-gray-200 rounded-full"></div>
            </div>
          </div>
        ) : (
          <>
            <RangeSlider
              min={priceRangeData.min}
              max={priceRangeData.max}
              value={priceRange}
              onChange={setPriceRange}
            />
          </>
        )}
      </FilterSection>
    </div>
  );
};

export default memo(FilterComponent);