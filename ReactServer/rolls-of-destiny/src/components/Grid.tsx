import React from "react";
import Column from "./Column";
import Box from "@mui/material/Box";

export default function Grid() {
    const handleColumnClick = (key: number) => {
        console.log(`Grid received click from column ${key}`);
    };

    const columns = [
        <Column key={0} onClick={handleColumnClick} columnKey={0} />,
        <Column key={1} onClick={handleColumnClick} columnKey={1} />,
        <Column key={2} onClick={handleColumnClick} columnKey={2} />
    ];

    return (
        <Box
            display="flex"
            flexDirection="row"
            justifyContent="center"
            alignItems="center"
            minHeight="40vh"
        >
            {columns.map((column, index) => (
                <Box key={index} flexGrow={1}> {}
                    {column}
                </Box>
            ))}
        </Box>
    );
}
