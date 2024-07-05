import React from "react";
import Column from "./Column";
import Box from "@mui/material/Box";
import {enemyInfo, yourInfo} from "../types/gameTypes";
import OpponentColumn from "./OpponentColumn";

export default function OpponentGrid({grid}: { grid: enemyInfo | null }) {
    return (
        <Box
            display="flex"
            flexDirection="row"
            justifyContent="center"
            alignItems="center"
            style={{width: "100%", minHeight: "100%"}}
        >
            <Box style={{height: "100%", aspectRatio: "1/3"}}><OpponentColumn key={0} columnKey={0} column={grid?.LeftColumn} handleClick={() => {}}/></Box>
            <Box style={{height: "100%", aspectRatio: "1/3"}}><OpponentColumn key={1} columnKey={1} column={grid?.MiddleColumn} handleClick={() => {}}/></Box>
            <Box style={{height: "100%", aspectRatio: "1/3"}}><OpponentColumn key={2} columnKey={2} column={grid?.RightColumn} handleClick={() => {}}/></Box>
        </Box>
    );
}
