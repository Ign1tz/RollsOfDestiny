import React from "react";
import Box from "@mui/material/Box";

const diceImages = [
    "../images/1.png",
    "../images/2.png",
    "../images/3.png",
    "../images/4.png",
    "../images/5.png",
    "../images/6.png"
];

export default function SimpleBox({ diceValue }: { diceValue: number | null }) {
    return (
        <Box
            height={50}
            width={50}
            my={0}
            mx={0}
            display="flex"
            alignItems="center"
            justifyContent="center"
            p={2}
            sx={{ border: "2px solid grey" }}
        >
            {diceValue !== null && (
                <img src={diceImages[diceValue - 1]} alt={`Dice ${diceValue}`} height={50} width={50} />
            )}
        </Box>
    );
}
