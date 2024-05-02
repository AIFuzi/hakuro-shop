import ProductCard from "@/components/ProductCard.tsx";
import {
    Dialog,
    DialogContent,
    DialogDescription, DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger
} from "@/components/ui/dialog.tsx";
import {Button} from "@/components/ui/button.tsx";
import {useState} from "react";

const HomePage = () => {
    const [open, setOpen] = useState(false)

    return (
        <div className="container mt-10">
            <h1 className="scroll-m-20 pb-2 text-3xl font-semibold tracking-tight">Type name</h1>
            <div className="grid grid-cols-3 place-items-center gap-10">
                <ProductCard/>
                <ProductCard/>
                <ProductCard/>
                <ProductCard/>
                <ProductCard/>
                <ProductCard/>
                <ProductCard/>
                <Button onClick={() => setOpen(true)}>Open</Button>
                <Dialog open={open}>
                    <DialogTrigger></DialogTrigger>
                    <DialogContent>
                        <DialogHeader>
                            <DialogTitle>Bob</DialogTitle>
                            <DialogDescription>
                                Test bob ac
                            </DialogDescription>
                        </DialogHeader>
                        <DialogFooter>
                            <Button onClick={() => setOpen(false)}>Close</Button>
                        </DialogFooter>
                    </DialogContent>
                </Dialog>
            </div>
        </div>
    );
};

export default HomePage;