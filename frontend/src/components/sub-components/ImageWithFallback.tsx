import { useState } from "react";

interface ImageWithFallbackProps extends React.ImgHTMLAttributes<HTMLImageElement> {
  src: string;
  alt: string;
}

export const ImageWithFallback: React.FC<ImageWithFallbackProps> = ({ src, alt, ...props }) => {
  const [imgSrc, setImgSrc] = useState(src);
  const [isLoading, setIsLoading] = useState(true);

  return (
    <div className="relative w-full h-48 bg-gray-50">
      {/* Loader si l'image charge */}
      {isLoading && imgSrc !== "/placeholder.jpg" && (
        <div className="absolute inset-0 flex items-center justify-center">
          <div className="w-10 h-10 border-4 border-gray-200 border-t-yellow-500 rounded-full animate-spin" />
        </div>
      )}

      {/* Image avec gestion du fallback */}
      <img
        src={imgSrc}
        alt={alt}
        className={`transition-opacity duration-300 ${isLoading ? "opacity-0" : "opacity-100"} object-contain w-full h-full`}
        onLoad={() => setIsLoading(false)}
        onError={() => {
          if (imgSrc !== "/placeholder.jpg") {
            setImgSrc("/placeholder.jpg");
          }
          setIsLoading(false);
        }}
        {...props}
      />
    </div>
  );
};
