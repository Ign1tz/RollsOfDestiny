import {ButtonGroup} from "@mui/material";
import Button from "@mui/material/Button";

export default function homeScreenButtonGroup() {
    return (
        <ButtonGroup>
            <Button onClick = {() => window.location.href = "/game"}>
                Play Against Bot
            </Button>
            <Button onClick = {() => window.location.href = "/game"}>
                Play Against Real Enemy
            </Button>
        </ButtonGroup>
    )
}