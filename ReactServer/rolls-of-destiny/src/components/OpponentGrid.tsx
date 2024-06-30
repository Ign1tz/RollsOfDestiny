import React from "react";
import Column from "./Column";
import Box from "@mui/material/Box";

export default function OpponentGrid({diceRoll}: { diceRoll: number | null }) {
    return (
        <Box
            display="flex"
            flexDirection="row"
            justifyContent="center"
            alignItems="center"
        >
            <Box><Column key={0} columnKey={0} diceRoll={diceRoll}/></Box>
            <Box><Column key={1} columnKey={1} diceRoll={diceRoll}/></Box>
            <Box><Column key={2} columnKey={2} diceRoll={diceRoll}/></Box>
        </Box>
    );
}
