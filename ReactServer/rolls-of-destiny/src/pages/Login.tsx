import "../css/Home.css"
import Button from "@mui/material/Button";
import SignUp from "./SignUp"

export default function LoginAndSignUpScreen({loggedIn, setLoggedIn}: {loggedIn: boolean, setLoggedIn: Function}) {

    const doWork = () => {
        setLoggedIn(!loggedIn)
        console.log(loggedIn)
    }

    return (
        <div className = "loginSignUpDivision">
            <h2>Login / SignUp</h2>
            <h4> To be implemented ... </h4>
            <SignUp/>
            <Button variant ="contained" onClick={() => window.location.href = "/"}> Login </Button>
            <Button variant ="contained" onClick={() => window.location.href = "/"}> SignUp </Button>
            <Button onClick={doWork}>TestButton to set LoggedIn</Button>
        </div>
    )
}