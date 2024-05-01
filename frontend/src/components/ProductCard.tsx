const ProductCard = () => {
    return (
        <div className="flex flex-col cursor-pointer">
            <img alt="" className="object-contain w-full rounded-2xl" src="http://localhost:4000/6d87337d-d03b-444e-aa47-e2b16aa5b45a.png"/>
            <h2 className="scroll-m-20 text-2xl font-semibold tracking-tight">Product name</h2>
            <h3 className="text-red-500">$400</h3>
        </div>
    );
};

export default ProductCard;