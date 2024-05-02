const ProductCard = () => {
    return (
        <div className="flex flex-col cursor-pointer">
            <img alt="" className="object-contain w-full rounded-2xl" src={import.meta.env.VITE_API_PRODUCT_IMAGE}/>
            <h2 className="scroll-m-20 text-2xl font-semibold tracking-tight">Product name</h2>
            <h3 className="text-red-500">$400</h3>
        </div>
    );
};

export default ProductCard;