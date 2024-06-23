import React from "react";
import Column from "./Column";
import Box from "@mui/material/Box";

export default function Grid({ onColumnClick, diceRoll }: { onColumnClick: Function, diceRoll: number | null }) {
    return (
        <Box
            display="flex"
            flexDirection="row"
            justifyContent="center"
            alignItems="center"
        >
            <Box><Column key={0} onClick={onColumnClick} columnKey={0} diceRoll={diceRoll} /></Box>
            <Box><Column key={1} onClick={onColumnClick} columnKey={1} diceRoll={diceRoll} /></Box>
            <Box><Column key={2} onClick={onColumnClick} columnKey={2} diceRoll={diceRoll} /></Box>
        </Box>
    );
}
