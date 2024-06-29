import React from "react";
import Column from "./Column";
import Box from "@mui/material/Box";

export default function Grid({websocket, connected, handleColumnClick}: {websocket?: WebSocket, connected?: boolean, handleColumnClick: Function}) {


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
