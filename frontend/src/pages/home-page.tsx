import ProductCard from "@/components/ProductCard.tsx";

const HomePage = () => {
    return (
        <div className="container mt-10">
            <h1 className="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight">Type name</h1>
            <div className="grid grid-cols-3 place-items-center gap-10">
                <ProductCard/>
                <ProductCard/>
                <ProductCard/>
                <ProductCard/>
                <ProductCard/>
                <ProductCard/>
                <ProductCard/>
            </div>
        </div>
    );
};

export default HomePage;