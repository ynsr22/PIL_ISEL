"use client";

import React, {
    useState,
    useEffect,
    useCallback,
    useMemo,
    ChangeEvent,
} from "react";

interface RangeSliderProps {
    min: number;
    max: number;
    value: [number, number];
    onChange: (value: [number, number]) => void;
    step?: number;
    currency?: string;
    formatValue?: (value: number) => string;
}

const RangeSlider: React.FC<RangeSliderProps> = ({
    min,
    max,
    value: externalValue,
    onChange,
    step = 10,
    currency = "€",
    formatValue = (value) => `${value}${currency}`,
}) => {
    // Valeur interne du slider, synchronisée avec la valeur externe
    const [localValue, setLocalValue] = useState<[number, number]>(externalValue);
    const [activeThumb, setActiveThumb] = useState<"min" | "max" | null>(null);
    const [isDragging, setIsDragging] = useState(false);

    // Synchronisation lorsque la prop `value` est modifiée
    useEffect(() => {
        setLocalValue(externalValue);
    }, [externalValue]);

    const minPercent = useMemo(
        () => ((localValue[0] - min) / (max - min)) * 100,
        [localValue, min, max]
    );
    const maxPercent = useMemo(
        () => ((localValue[1] - min) / (max - min)) * 100,
        [localValue, min, max]
    );

    const handleMinChange = useCallback(
        (e: ChangeEvent<HTMLInputElement>) => {
            setActiveThumb("min");

            const newMin = Math.min(Number(e.target.value), localValue[1] - step);
            const finalMin = Math.max(newMin, min);

            const newValue: [number, number] = [finalMin, localValue[1]];
            setLocalValue(newValue);
            onChange(newValue);
        },
        [localValue, onChange, min, step]
    );

    const handleMaxChange = useCallback(
        (e: ChangeEvent<HTMLInputElement>) => {
            setActiveThumb("max");

            const newMax = Math.max(Number(e.target.value), localValue[0] + step);
            const finalMax = Math.min(newMax, max);

            const newValue: [number, number] = [localValue[0], finalMax];
            setLocalValue(newValue);
            onChange(newValue);
        },
        [localValue, onChange, max, step]
    );

    const handleThumbMouseDown = useCallback((thumb: "min" | "max") => {
        setActiveThumb(thumb);
        setIsDragging(true);
    }, []);

    const handleMouseUp = useCallback(() => {
        setIsDragging(false);
        setActiveThumb(null);
    }, []);

    // Register global mouse up event listener
    useEffect(() => {
        document.addEventListener("mouseup", handleMouseUp);
        document.addEventListener("touchend", handleMouseUp);
        return () => {
            document.removeEventListener("mouseup", handleMouseUp);
            document.removeEventListener("touchend", handleMouseUp);
        };
    }, [handleMouseUp]);

    return (
        <div className="relative w-full pt-6 pb-8" role="group" aria-label="Price range slider">
            {/* Affichage des valeurs sélectionnées */}
            <div className="flex justify-between mb-4 text-sm font-medium text-gray-700">
                <span>{formatValue(localValue[0])}</span>
                <span>{formatValue(localValue[1])}</span>
            </div>

            {/* Piste du slider */}
            <div className="relative h-2 bg-gray-200 rounded-full">
                <div
                    className="absolute h-full bg-blue-500 rounded-full transition-all duration-150"
                    style={{
                        left: `${minPercent}%`,
                        width: `${maxPercent - minPercent}%`,
                    }}
                />
            </div>

            {/* Input de la valeur min */}
            <input
                type="range"
                min={min}
                max={max}
                step={step}
                value={localValue[0]}
                onChange={handleMinChange}
                onMouseDown={() => handleThumbMouseDown("min")}
                onTouchStart={() => handleThumbMouseDown("min")}
                className="absolute top-1 left-0 w-full h-2 appearance-none bg-transparent cursor-pointer focus:outline-none"
                style={{
                    zIndex: activeThumb === "min" ? 5 : 3,
                    pointerEvents: activeThumb === "max" ? "none" : "auto",
                }}
                aria-label="Minimum price"
            />

            {/* Input de la valeur max */}
            <input
                type="range"
                min={min}
                max={max}
                step={step}
                value={localValue[1]}
                onChange={handleMaxChange}
                onMouseDown={() => handleThumbMouseDown("max")}
                onTouchStart={() => handleThumbMouseDown("max")}
                className="absolute top-1 left-0 w-full h-2 appearance-none bg-transparent cursor-pointer focus:outline-none"
                style={{
                    zIndex: activeThumb === "max" ? 5 : 4,
                    pointerEvents: activeThumb === "min" ? "none" : "auto",
                }}
                aria-label="Maximum price"
            />

            {/* Thumbs personnalisés */}
            <div
                className={`absolute top-1 -translate-y-1/2 w-5 h-5 bg-white border-2 border-blue-500 rounded-full shadow-md
                    ${isDragging && activeThumb === "min" ? "scale-110 cursor-grabbing" : "cursor-grab"}
                    transition-all duration-150`}
                style={{ left: `calc(${minPercent}% - 10px)` }}
                aria-hidden="true"
            />
            <div
                className={`absolute top-1 -translate-y-1/2 w-5 h-5 bg-white border-2 border-blue-500 rounded-full shadow-md
                    ${isDragging && activeThumb === "max" ? "scale-110 cursor-grabbing" : "cursor-grab"}
                    transition-all duration-150`}
                style={{ left: `calc(${maxPercent}% - 10px)` }}
                aria-hidden="true"
            />
        </div>
    );
};

export default React.memo(RangeSlider);