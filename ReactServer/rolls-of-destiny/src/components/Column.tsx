// Column.jsx
import React from "react";
import SimpleBox from "./SimpleBox";

export default function Column({ onClick, columnKey }: { onClick: Function, columnKey: number}) {
    const handleClick = () => {
        console.log(`Column ${columnKey} clicked`);
        if (onClick) {
            onClick(columnKey);
        }
    };

    return (
        <div onClick={handleClick} style={{ cursor: "pointer" }}>
            <SimpleBox key={0} />
            <SimpleBox key={1} />
            <SimpleBox key={2} />
        </div>
    );
}
