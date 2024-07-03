import React from "react";
import Column from "./Column";
import Box from "@mui/material/Box";
import {column, enemyInfo, yourInfo} from "../types/gameTypes";

export default function Grid({handleColumnClick, active, grid}: {handleColumnClick?: Function, active?: boolean, grid: yourInfo | null}) {

    return (
        <Box
            display="flex"
            flexDirection="row"
            justifyContent="center"
            alignItems="center"
        >
            <Box><Column key={0} handleClick={active && handleColumnClick && !grid?.LeftColumn.IsFull ? handleColumnClick : () => {console.log("inactive")}} columnKey={0} column={grid?.LeftColumn} /></Box>
            <Box><Column key={1} handleClick={active && handleColumnClick && !grid?.MiddleColumn.IsFull ? handleColumnClick : () => {}} columnKey={1} column={grid?.MiddleColumn} /></Box>
            <Box><Column key={2} handleClick={active && handleColumnClick && !grid?.RightColumn.IsFull ? handleColumnClick : () => {}} columnKey={2} column={grid?.RightColumn} /></Box>
        </Box>
    );
}
