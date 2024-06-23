import React from "react";
import Column from "./Column";
import Box from "@mui/material/Box";

export default function Grid({canPlace, setCanPlace, diceRoll}: { canPlace: boolean, setCanPlace: Function, diceRoll: number | null }) {
    return (
        <Box
            display="flex"
            flexDirection="row"
            justifyContent="center"
            alignItems="center"
        >
            <Box><Column key={0} canPlace={canPlace} setCanPlace={setCanPlace} columnKey={0} diceRoll={diceRoll}/></Box>
            <Box><Column key={1} canPlace={canPlace} setCanPlace={setCanPlace} columnKey={1} diceRoll={diceRoll}/></Box>
            <Box><Column key={2} canPlace={canPlace} setCanPlace={setCanPlace} columnKey={2} diceRoll={diceRoll}/></Box>
        </Box>
    );
}
