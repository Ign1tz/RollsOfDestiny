import React from "react";
import Box from "@mui/material/Box";

export default function SimpleBox() {
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
        ></Box>
    );
}
