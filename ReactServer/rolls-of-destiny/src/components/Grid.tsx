import React from "react";
import Column from "./Column";
import Box from "@mui/material/Box";

export default function Grid({websocket, connected}: {websocket?: WebSocket, connected?: boolean}) {
    const handleColumnClick = (key: number) => {
        console.log(connected)
        if (websocket && connected){
            console.log(websocket)
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
            <Box><Column key={0} handleClick={handleColumnClick} columnKey={0} /></Box>
            <Box><Column key={1} handleClick={handleColumnClick} columnKey={1} /></Box>
            <Box><Column key={2} handleClick={handleColumnClick} columnKey={2} /></Box>
        </Box>
    );
}
