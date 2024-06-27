import React from "react";
import SimpleBox from "./SimpleBox";

export default function Column({handleClick, columnKey}: { handleClick: Function, columnKey: number }) {


    return (
        <div onClick={() => {handleClick(columnKey)}} style={{cursor: "pointer"}}>
            <SimpleBox key={0}/>
            <SimpleBox key={1}/>
            <SimpleBox key={2}/>
        </div>
    );
}
