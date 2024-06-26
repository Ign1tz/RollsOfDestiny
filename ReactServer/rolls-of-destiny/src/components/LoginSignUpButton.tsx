import Button from "@mui/material/Button";
import Box from "@mui/material/Box";

export default function LoginSignUpButton() {
    return (
        <Box color = "black">
            <Button style={{marginInlineEnd: "20px"}} variant="contained" color = "inherit" onClick={() => window.location.href="/login"}>Login</Button>
            <Button  variant="contained" color = "inherit" onClick={() => window.location.href="/signup"}>Signup</Button>
        </Box>
    )
}