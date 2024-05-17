import "../css/Login.css"
import Button from "@mui/material/Button";

export default function LoginAndSignUpScreen({loggedIn, setLoggedIn}: {loggedIn: boolean, setLoggedIn: Function}) {

    const doWork = () => {
        setLoggedIn(!loggedIn)
        console.log(loggedIn)
    }
    // to be done, it wont globally switch the loggedIn state
    return (
        <div className = "loginSignUpDivision">
            <h2>Login / SignUp</h2>
            <h4> To be implemented ... </h4>
            <Button variant ="contained" onClick={() => window.location.href = "/"}> Login </Button>
            <Button variant ="contained" onClick={() => window.location.href = "/"}> SignUp </Button>
            <Button onClick={doWork}>TestButton to set LoggedIn</Button>
        </div>
    )
}