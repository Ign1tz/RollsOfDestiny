import React from "react";
import SimpleBox from "./SimpleBox";

export default function Column({onClick, columnKey}: { onClick: Function, columnKey: number }) {
    const handleClick = async () => {
        console.log(`Column ${columnKey} clicked`);
        const response = await fetch("http://localhost:8080/left", {
            method: "POST",
            headers: {
                'Accept': 'application/json, text/plain',
                'Content-Type': 'application/json;charset=UTF-8'
            },
            body: JSON.stringify({gameid: 'test', columnKey: columnKey.toString()})
        });
        if (!response.ok) {

        }
        if (onClick) {
            onClick(columnKey);
        }
    };

    return (
        <div onClick={handleClick} style={{cursor: "pointer"}}>
            <SimpleBox key={0}/>
            <SimpleBox key={1}/>
            <SimpleBox key={2}/>
        </div>
    );
}
