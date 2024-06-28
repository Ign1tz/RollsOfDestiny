import React from "react";
import Column from "./Column";
import Box from "@mui/material/Box";

export default function Grid({canPlace, setCanPlace, diceRoll, websocket, connected}: {canPlace: boolean, setCanPlace: Function, diceRoll: number | null, websocket?: WebSocket, connected?: boolean}) {
    const handleColumnClick = (key: number) => {
        console.log(connected)
        if (websocket && connected){
            websocket.send("test " + key)
        }
    };

    return (
        <Box
            display="flex"
            flexDirection="row"
            justifyContent="center"
            alignItems="center"
        >
            <Box><Column key={0} handleClick={handleColumnClick} canPlace={canPlace} setCanPlace={setCanPlace} columnKey={0} diceRoll={diceRoll}/></Box>
            <Box><Column key={1} handleClick={handleColumnClick} canPlace={canPlace} setCanPlace={setCanPlace} columnKey={1} diceRoll={diceRoll}/></Box>
            <Box><Column key={2} handleClick={handleColumnClick} canPlace={canPlace} setCanPlace={setCanPlace} columnKey={2} diceRoll={diceRoll}/></Box>
        </Box>
    );
}
