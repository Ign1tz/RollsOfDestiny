import React, {useState} from "react";
import SimpleBox from "./SimpleBox";
import {column} from "../types/gameTypes";


export default function Column({handleClick, columnKey, column}: { handleClick: Function, columnKey: number, column: column | undefined}) {



    return (
        <div onClick={() => {handleClick(columnKey)}} style={{cursor: "pointer"}}>
            <SimpleBox key={0}  diceValue={column ? parseInt(column?.First) : null}/>
            <SimpleBox key={1}  diceValue={column ? parseInt(column?.Second) : null}/>
            <SimpleBox key={2}  diceValue={column ? parseInt(column?.Third) : null}/>
        </div>
    );
}