import {z} from 'zod'
import {useForm} from "react-hook-form";
import {zodResolver} from "@hookform/resolvers/zod";
import {
    Form,
    FormControl, FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage
} from "@/components/ui/form.tsx";
import {Input} from "@/components/ui/input.tsx";
import {Button} from "@/components/ui/button.tsx";
import {useNavigate} from "react-router-dom";

const formSchema = z.object({
    Email: z.string().min(5, {
        message: "Email must be 5 chars"
    }),
    Name: z.string(),
    Password: z.string(),
    RepeatPassword: z.string()
})

export const RegistrationPage = () => {
    const navigate = useNavigate()

    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            Email: "",
            Name: "",
            Password: "",
            RepeatPassword: ""
        }
    })

    const Login = (value: z.infer<typeof formSchema>) => {
        console.log(value)
    }

    return(
        <div className="flex justify-center items-center h-lvh">
            <div className="p-10 border-2 rounded-xl bg-zinc-950 drop-shadow-[0_10px_10px_rgba(100,100,100,0.1)]">
                <Form {...form}>
                    <form onSubmit={form.handleSubmit(Login)} className="space-y-5">
                        <FormField control={form.control} name="Email" render={({field}) => (
                            <FormItem>
                                <div className="flex flex-row items-center space-x-2">
                                    <FormLabel>Email</FormLabel>
                                    <FormControl>
                                        <Input placeholder="bob@bob.com" {...field}/>
                                    </FormControl>
                                </div>
                                <FormMessage/>
                            </FormItem>
                        )}/>
                        <FormField control={form.control} name="Name" render={({field}) => (
                            <FormItem>
                                <div className="flex flex-row items-center space-x-2">
                                    <FormLabel>Name</FormLabel>
                                    <FormControl>
                                        <Input placeholder="Bob" {...field}/>
                                    </FormControl>
                                </div>
                                <FormDescription>
                                    What your name?
                                </FormDescription>
                                <FormMessage/>
                            </FormItem>
                        )}/>
                        <FormField control={form.control} name="Password" render={({field}) => (
                            <FormItem>
                                <div className="flex flex-row items-center space-x-2">
                                    <FormLabel>Password</FormLabel>
                                    <FormControl>
                                        <Input placeholder="qwerty12345" {...field} type="password"/>
                                    </FormControl>
                                </div>
                                <FormMessage/>
                            </FormItem>
                        )}/>
                        <FormField control={form.control} name="RepeatPassword" render={({field}) => (
                            <FormItem>
                                <div className="flex flex-row items-center space-x-2">
                                    <FormLabel>Repeat password</FormLabel>
                                    <FormControl>
                                        <Input placeholder="qwerty12345" {...field} type="password"/>
                                    </FormControl>
                                </div>
                                <FormMessage/>
                            </FormItem>
                        )}/>
                        <Button className="w-full" type="submit">Sign up</Button>
                    </form>
                </Form>
                <div className="font-sans">
                    Have account?
                    <Button variant="link" onClick={() => navigate('/login')}>Sign in</Button>
                </div>
            </div>
        </div>
    )
}