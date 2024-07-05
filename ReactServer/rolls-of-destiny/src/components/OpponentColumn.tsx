import React, {useState} from "react";
import SimpleBox from "./SimpleBox";
import {column} from "../types/gameTypes";
import Box from "@mui/material/Box";


export default function OpponentColumn({handleClick, columnKey, column}: { handleClick: Function, columnKey: number, column: column | undefined}) {



    return (
        <Box onClick={() => {handleClick(columnKey)}} >
            <SimpleBox key={2}  diceValue={column ? parseInt(column?.Third) : null}/>
            <SimpleBox key={1}  diceValue={column ? parseInt(column?.Second) : null}/>
            <SimpleBox key={0}  diceValue={column ? parseInt(column?.First) : null}/>
        </Box>
    );
}