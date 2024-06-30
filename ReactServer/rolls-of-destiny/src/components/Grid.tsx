import React from "react";
import Column from "./Column";
import Box from "@mui/material/Box";

export default function Grid({websocket, connected, handleColumnClick, active, diceRoll}: {websocket?: WebSocket, connected?: boolean, handleColumnClick?: Function, active?: boolean, diceRoll: number | null}) {


    return (
        <Box
            display="flex"
            flexDirection="row"
            justifyContent="center"
            alignItems="center"
        >
            <Box><Column key={0} handleClick={active && handleColumnClick ? handleColumnClick : () => {}} columnKey={0} diceRoll={diceRoll} /></Box>
            <Box><Column key={1} handleClick={active && handleColumnClick ? handleColumnClick : () => {}} columnKey={1} diceRoll={diceRoll} /></Box>
            <Box><Column key={2} handleClick={active && handleColumnClick ? handleColumnClick : () => {}} columnKey={2} diceRoll={diceRoll} /></Box>
        </Box>
    );
}
