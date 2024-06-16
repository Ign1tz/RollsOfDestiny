import React from "react";
import Column from "./Column";
import Box from "@mui/material/Box";

export default function Grid() {
    const handleColumnClick = (key: number) => {
        console.log(`Grid received click from column ${key}`);
    };

    return (
        <Box
            display="flex"
            flexDirection="row"
            justifyContent="center"
            alignItems="center"
        >
            <Box><Column key={0} onClick={handleColumnClick} columnKey={0} /></Box>
            <Box><Column key={1} onClick={handleColumnClick} columnKey={1} /></Box>
            <Box><Column key={2} onClick={handleColumnClick} columnKey={2} /></Box>
        </Box>
    );
}
