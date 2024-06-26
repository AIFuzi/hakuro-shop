import {observer} from "mobx-react-lite";
import {Button} from "@/components/ui/button.tsx";
import {useNavigate} from "react-router-dom";
import logo from '../assets/shop-logo.svg'

const Navbar = observer(() => {
    const navigate = useNavigate()

    return (
        <div className="sticky top-0 border-b-2 border-zinc-900">
            <div className="container">
                <div className="h-14 flex items-center justify-between backdrop-blur-sm bg-zinc-950/70">
                    <div>
                        <Button variant="link" onClick={() => navigate('/')}>
                            <img alt="" src={logo} width="200" className="hover:contrast-0"/>
                        </Button>
                    </div>
                    <div>
                        <ul className="flex flex-row justify-between">
                            <li><Button variant="link">Shirt</Button></li>
                            <li><Button variant="link">Shirt</Button></li>
                            <li><Button variant="link">Shirt</Button></li>
                            <li><Button variant="link">Shirt</Button></li>
                            <li><Button variant="link">Shirt</Button></li>
                        </ul>
                    </div>
                    <div className="mr-5">
                        <Button variant="ghost" onClick={() => navigate('/signup')} className="mr-2">Sign Up</Button>
                        <Button variant="ghost" onClick={() => navigate('/login')}>Login</Button>
                    </div>
                </div>
            </div>
        </div>
    );
});

export default Navbar;