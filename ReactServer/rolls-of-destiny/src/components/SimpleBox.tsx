import React from "react";
import Box from "@mui/material/Box";
import dice1 from "../images/1.png";
import dice2 from "../images/2.png";
import dice3 from "../images/3.png";
import dice4 from "../images/4.png";
import dice5 from "../images/5.png";
import dice6 from "../images/6.png";


const diceImages = [dice1, dice2, dice3, dice4, dice5, dice6];

export default function SimpleBox({diceValue}: { diceValue: number | null }) {
    return (
        <Box
            height={80}
            width={80}
            my={0}
            mx={0}
            display="flex"
            alignItems="center"
            justifyContent="center"
            p={2}
            sx={{
                border: "2px solid black",
                backgroundColor: "darkgray",
            }}
        >
            {diceValue !== null && (
                <img src={diceImages[diceValue - 1]} alt={`Dice ${diceValue}`} height={80} width={80} style={{borderRadius: "10%" }}/>
            )}
        </Box>
    );
}
