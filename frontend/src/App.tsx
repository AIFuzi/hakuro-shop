import {ThemeProvider} from "@/components/ui/theme-provider.tsx";
import {BrowserRouter, Route, Routes} from "react-router-dom";
import HomePage from "@/pages/home-page.tsx";
import NotFoundPage from "@/pages/not-found-page.tsx";
import Navbar from "@/components/navbar.tsx";
import {LoginPage} from "@/pages/login-page.tsx";
import {RegistrationPage} from "@/pages/register-page.tsx";

function App() {
  return (
    <ThemeProvider defaultTheme="dark">
        <BrowserRouter>
            <Navbar/>
            <Routes>
                <Route path="/signup" element={<RegistrationPage/>}/>
                <Route path="/login" element={<LoginPage/>}/>
                <Route path="/" element={<HomePage/>}/>
                <Route path="/*" element={<NotFoundPage/>}/>
            </Routes>
        </BrowserRouter>
    </ThemeProvider>
  )
}

export default App
